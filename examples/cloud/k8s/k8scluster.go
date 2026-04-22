package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func main() {
	// GCORE_API_KEY, GCORE_BASE_URL, GCORE_CLOUD_PROJECT_ID, and
	// GCORE_CLOUD_REGION_ID are picked up automatically by gcore.NewClient
	// when set in the environment — no need to read them explicitly.
	for _, v := range []string{"GCORE_CLOUD_PROJECT_ID", "GCORE_CLOUD_REGION_ID"} {
		if os.Getenv(v) == "" {
			log.Fatalf("%s environment variable is required", v)
		}
	}

	// TODO set a pre-existing SSH keypair name before running. The K8s cluster
	// create endpoint requires a keypair to bootstrap nodes.
	keypairName := os.Getenv("GCORE_CLOUD_K8S_KEYPAIR_NAME")
	if keypairName == "" {
		log.Fatalf("GCORE_CLOUD_K8S_KEYPAIR_NAME environment variable is required")
	}

	client := gcore.NewClient()

	// Look up a small flavor and the latest version to use for cluster creation.
	flavors := listFlavors(&client)
	poolFlavorID := pickPoolFlavor(flavors)

	versions := listVersions(&client)
	clusterVersion := pickCreateVersion(versions)

	// Cluster lifecycle.
	clusterName := createCluster(&client, poolFlavorID, clusterVersion, keypairName)

	getCluster(&client, clusterName)
	listClusters(&client)

	updateCluster(&client, clusterName)

	// Upgrade the cluster to the newest available version, if one exists.
	upgradeVersions := listVersionsForUpgrade(&client, clusterName)
	if len(upgradeVersions) > 0 {
		upgradeCluster(&client, clusterName, upgradeVersions[len(upgradeVersions)-1].Version)
		// The upgrade task finishes before the cluster itself settles back to
		// Provisioned, so subsequent mutations (e.g. creating a new pool) can
		// be rejected with "Cluster or pool is not ready". Wait it out.
		waitForClusterReady(&client, clusterName)
	} else {
		fmt.Printf("No upgrade versions available for cluster %s; skipping upgrade\n", clusterName)
	}

	getCertificate(&client, clusterName)
	getKubeconfig(&client, clusterName)

	// Pool operations.
	listPools(&client, clusterName)
	existingPool := getFirstPool(&client, clusterName)

	checkPoolQuota(&client, poolFlavorID)
	secondPool := createPool(&client, clusterName, poolFlavorID)
	updatePool(&client, clusterName, secondPool)
	resizePool(&client, clusterName, secondPool)
	listPoolNodes(&client, clusterName, secondPool)
	deletePool(&client, clusterName, secondPool)

	// Ensure the initial pool is still reachable after the second pool is gone.
	getPool(&client, clusterName, existingPool)

	listClusterNodes(&client, clusterName)

	// Final cleanup.
	deleteCluster(&client, clusterName)
}

func createCluster(client *gcore.Client, poolFlavorID, version, keypairName string) string {
	fmt.Println("\n=== CREATE K8S CLUSTER ===")

	params := cloud.K8SClusterNewParams{
		Name:    "gcore-go-example-k8s",
		Keypair: keypairName,
		Version: version,
		Pools: []cloud.K8SClusterNewParamsPool{{
			Name:               "gcore-go-pool",
			FlavorID:           poolFlavorID,
			MinNodeCount:       1,
			MaxNodeCount:       gcore.Int(2),
			BootVolumeSize:     gcore.Int(50),
			BootVolumeType:     "standard",
			AutoHealingEnabled: gcore.Bool(true),
			IsPublicIpv4:       gcore.Bool(true),
			ServergroupPolicy:  "soft-anti-affinity",
			Labels: map[string]string{
				"pool": "gcore-go-example",
			},
		}},
	}

	cluster, err := client.Cloud.K8S.Clusters.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating k8s cluster: %v", err)
	}

	fmt.Printf("Created K8s cluster: Name=%s, Version=%s, Status=%s\n",
		cluster.Name, cluster.Version, cluster.Status)
	fmt.Println("==========================")

	return cluster.Name
}

func getCluster(client *gcore.Client, clusterName string) {
	fmt.Println("\n=== GET K8S CLUSTER ===")

	cluster, err := client.Cloud.K8S.Clusters.Get(context.Background(), clusterName, cloud.K8SClusterGetParams{})
	if err != nil {
		fmt.Printf("Error getting k8s cluster: %v\n", err)
		return
	}

	fmt.Printf("K8s cluster: Name=%s, Version=%s, Status=%s, Pools=%d\n",
		cluster.Name, cluster.Version, cluster.Status, len(cluster.Pools))
	fmt.Println("=======================")
}

func listClusters(client *gcore.Client) {
	fmt.Println("\n=== LIST K8S CLUSTERS ===")

	clusters, err := client.Cloud.K8S.Clusters.List(context.Background(), cloud.K8SClusterListParams{})
	if err != nil {
		fmt.Printf("Error listing k8s clusters: %v\n", err)
		return
	}

	for i, c := range clusters.Results {
		fmt.Printf("  %d. K8s cluster: Name=%s, Version=%s, Status=%s\n",
			i+1, c.Name, c.Version, c.Status)
	}

	if len(clusters.Results) == 0 {
		fmt.Println("  No k8s clusters found.")
	}

	fmt.Println("=========================")
}

func updateCluster(client *gcore.Client, clusterName string) {
	fmt.Println("\n=== UPDATE K8S CLUSTER ===")

	// Override a cluster-autoscaler parameter — this is a safe, idempotent
	// change that exercises UpdateAndPoll.
	params := cloud.K8SClusterUpdateParams{
		AutoscalerConfig: map[string]string{
			"scale-down-unneeded-time": "5m",
		},
	}

	cluster, err := client.Cloud.K8S.Clusters.UpdateAndPoll(context.Background(), clusterName, params)
	if err != nil {
		log.Fatalf("Error updating k8s cluster: %v", err)
	}

	fmt.Printf("Updated K8s cluster: Name=%s, AutoscalerConfig=%v\n",
		cluster.Name, cluster.AutoscalerConfig)
	fmt.Println("==========================")
}

func listVersionsForUpgrade(client *gcore.Client, clusterName string) []cloud.K8SClusterVersion {
	fmt.Println("\n=== LIST K8S CLUSTER UPGRADE VERSIONS ===")

	versions, err := client.Cloud.K8S.Clusters.ListVersionsForUpgrade(
		context.Background(), clusterName, cloud.K8SClusterListVersionsForUpgradeParams{})
	if err != nil {
		fmt.Printf("Error listing k8s upgrade versions: %v\n", err)
		return nil
	}

	for i, v := range versions.Results {
		fmt.Printf("  %d. Upgrade version: %s\n", i+1, v.Version)
	}

	if len(versions.Results) == 0 {
		fmt.Println("  Cluster is already on the latest available version.")
	}

	fmt.Println("==========================================")
	return versions.Results
}

func upgradeCluster(client *gcore.Client, clusterName, targetVersion string) {
	fmt.Println("\n=== UPGRADE K8S CLUSTER ===")

	params := cloud.K8SClusterUpgradeParams{
		Version: targetVersion,
	}

	cluster, err := client.Cloud.K8S.Clusters.UpgradeAndPoll(context.Background(), clusterName, params)
	if err != nil {
		log.Fatalf("Error upgrading k8s cluster: %v", err)
	}

	fmt.Printf("Upgraded K8s cluster: Name=%s, Version=%s\n",
		cluster.Name, cluster.Version)
	fmt.Println("===========================")
}

func getCertificate(client *gcore.Client, clusterName string) {
	fmt.Println("\n=== GET K8S CLUSTER CERTIFICATE ===")

	cert, err := client.Cloud.K8S.Clusters.GetCertificate(
		context.Background(), clusterName, cloud.K8SClusterGetCertificateParams{})
	if err != nil {
		fmt.Printf("Error getting k8s cluster certificate: %v\n", err)
		return
	}

	fmt.Printf("Certificate bytes=%d, Key bytes=%d\n",
		len(cert.Certificate), len(cert.Key))
	fmt.Println("===================================")
}

func waitForClusterReady(client *gcore.Client, clusterName string) {
	fmt.Println("\n=== WAIT FOR K8S CLUSTER READY ===")

	const (
		pollInterval = 10 * time.Second
		maxAttempts  = 60
	)
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		cluster, err := client.Cloud.K8S.Clusters.Get(
			context.Background(), clusterName, cloud.K8SClusterGetParams{})
		if err != nil {
			log.Fatalf("Error polling cluster status: %v", err)
		}

		pending := ""
		if cluster.Status != cloud.K8SClusterStatusProvisioned {
			pending = fmt.Sprintf("cluster=%s", cluster.Status)
		}
		for _, p := range cluster.Pools {
			if p.Status != "Running" {
				pending += fmt.Sprintf(" pool[%s]=%s", p.Name, p.Status)
			}
		}

		if pending == "" {
			fmt.Printf("Cluster %s is ready (all pools Running)\n", clusterName)
			fmt.Println("==================================")
			return
		}
		fmt.Printf("  attempt %d: waiting on%s\n", attempt, pending)
		time.Sleep(pollInterval)
	}
	log.Fatalf("Cluster %s did not become ready after %d attempts", clusterName, maxAttempts)
}

func deleteCluster(client *gcore.Client, clusterName string) {
	fmt.Println("\n=== DELETE K8S CLUSTER ===")

	err := client.Cloud.K8S.Clusters.DeleteAndPoll(
		context.Background(), clusterName, cloud.K8SClusterDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting k8s cluster: %v", err)
	}

	fmt.Printf("Deleted K8s cluster: Name=%s\n", clusterName)
	fmt.Println("==========================")
}
