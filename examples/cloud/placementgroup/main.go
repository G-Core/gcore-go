package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/packages/param"
)

func main() {
	// set GCORE_CLOUD_PROJECT_ID and GCORE_CLOUD_REGION_ID before running
	// set GCORE_INSTANCE_ID to an active test VM in that project and region
	region, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil || region <= 0 {
		log.Fatal("GCORE_CLOUD_REGION_ID must be a positive integer")
	}
	instanceID := os.Getenv("GCORE_INSTANCE_ID")
	if instanceID == "" {
		log.Fatal("GCORE_INSTANCE_ID must identify an active test VM without a placement group")
	}
	client := gcore.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	if err := run(ctx, &client, region, instanceID); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, client *gcore.Client, region int64, instanceID string) (err error) {
	instance, err := client.Cloud.Instances.Get(ctx, instanceID, cloud.InstanceGetParams{})
	if err != nil {
		return fmt.Errorf("read test VM: %w", err)
	}
	if instance.RegionID != region || instance.Status != "active" || instance.PlacementGroupID != "" {
		return errors.New("choose an active test VM in the target region with no placement group")
	}

	groups := client.Cloud.PlacementGroups
	name := fmt.Sprintf("sdk-placement-%d", time.Now().UnixNano())
	group, err := groups.NewAndPoll(ctx, cloud.PlacementGroupNewParams{
		Name: name, RegionID: region, Policy: cloud.PlacementGroupNewParamsPolicyAntiAffinity,
	})
	if err != nil {
		return fmt.Errorf("create placement group: %w", err)
	}
	// use a fresh context so a failed or timed-out step cannot prevent cleanup
	defer func() {
		cleanup, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		if deleteErr := groups.DeleteAndPoll(cleanup, group.ID, cloud.PlacementGroupDeleteParams{}); deleteErr != nil {
			err = errors.Join(err, fmt.Errorf("delete placement group %s: %w", group.ID, deleteErr))
			return
		}
		fmt.Printf("Deleted group %s\n", group.ID)
	}()

	pager := groups.ListAutoPaging(ctx, cloud.PlacementGroupListParams{
		RegionIDs: []int64{region}, Name: cloud.PlacementGroupListParamsName{Exact: []string{name}},
	})
	for pager.Next() {
		fmt.Printf("Found group %s\n", pager.Current().ID)
	}
	if err := pager.Err(); err != nil {
		return fmt.Errorf("list placement groups: %w", err)
	}
	details, err := groups.Get(ctx, group.ID, cloud.PlacementGroupGetParams{})
	if err != nil {
		return fmt.Errorf("get placement group: %w", err)
	}
	fmt.Printf("Group %s uses policy %s\n", details.Name, details.Policy)

	if _, err := client.Cloud.Instances.UpdateAndPoll(ctx, instanceID, cloud.InstanceUpdateParams{
		PlacementGroupID: gcore.String(group.ID),
	}); err != nil {
		return fmt.Errorf("add test VM to group: %w", err)
	}
	details, err = groups.Get(ctx, group.ID, cloud.PlacementGroupGetParams{})
	if err != nil {
		return fmt.Errorf("read group members: %w", err)
	}
	fmt.Printf("Member IDs: %v\n", details.MemberIDs)

	if _, err := client.Cloud.Instances.UpdateAndPoll(ctx, instanceID, cloud.InstanceUpdateParams{
		PlacementGroupID: param.Null[string](),
	}); err != nil {
		return fmt.Errorf("remove test VM from group: %w", err)
	}
	return nil
}
