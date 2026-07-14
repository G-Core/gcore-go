package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cdn"
	"github.com/G-Core/gcore-go/packages/param"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read
	// from the GCORE_API_KEY environment variable if omitted.

	// Origin group ID is required to create the CDN resource that presets are
	// applied to.
	originGroupID, err := strconv.ParseInt(os.Getenv("GCORE_CDN_ORIGIN_GROUP_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CDN_ORIGIN_GROUP_ID environment variable is required and must be a valid integer")
	}

	client := gcore.NewClient()

	// Pick a preset for each object type from the catalog.
	resourcePreset, rulePreset := findPresets(&client)
	if resourcePreset == nil {
		log.Fatalf("No CDN resource preset found to apply in this example")
	}

	// Create a CDN resource to apply presets to.
	resource := createCDNResource(&client, originGroupID)
	// Always clean the resource up, even if a later step fails.
	defer deactivateAndDeleteCDNResource(&client, resource.ID)

	// --- Applied preset lifecycle for a CDN resource ---

	// Apply the resource preset to the CDN resource.
	applyPreset(&client, resourcePreset.ID, resource.ID)

	// List the objects the preset is currently applied to.
	getAppliedObjects(&client, resourcePreset.ID)

	// Inspect the preset applied to the CDN resource and the fields it manages.
	getResourcePreset(&client, resource.ID)

	// Unapply the preset from the CDN resource.
	unapplyPreset(&client, resourcePreset.ID, resource.ID)

	// --- Applied preset lifecycle for a rule ---

	// A Rule-type preset is required to demonstrate the rule flow. It may not be
	// available on every account, so skip gracefully when it is missing.
	if rulePreset == nil {
		fmt.Println("\nNo Rule preset available, skipping the rule-preset flow")
		return
	}

	// Create a rule on the CDN resource to apply the rule preset to.
	rule := createRule(&client, resource.ID)

	// Apply the rule preset to the rule.
	applyPreset(&client, rulePreset.ID, rule.ID)

	// Inspect the preset applied to the rule and the fields it manages.
	getRulePreset(&client, resource.ID, rule.ID)

	// Unapply the preset from the rule.
	unapplyPreset(&client, rulePreset.ID, rule.ID)
}

func findPresets(client *gcore.Client) (resourcePreset, rulePreset *cdn.PresetDetail) {
	fmt.Println("\n=== LIST PRESETS ===")

	result, err := client.CDN.Presets.List(context.Background(), cdn.PresetListParams{})
	if err != nil {
		log.Fatalf("Error listing presets: %v", err)
	}

	for i, preset := range result.Results {
		fmt.Printf("  %d. Preset: ID=%d, Name=%s, ObjectType=%s\n",
			i+1, preset.ID, preset.Name, preset.ObjectType)

		// Remember the first preset found for each object type.
		switch preset.ObjectType {
		case cdn.PresetDetailObjectTypeCDNResource:
			if resourcePreset == nil {
				resourcePreset = &result.Results[i]
			}
		case cdn.PresetDetailObjectTypeRule:
			if rulePreset == nil {
				rulePreset = &result.Results[i]
			}
		}
	}
	fmt.Println("====================")

	return resourcePreset, rulePreset
}

func createCDNResource(client *gcore.Client, originGroupID int64) *cdn.CDNResource {
	fmt.Println("\n=== CREATE CDN RESOURCE ===")

	// Use a unique CNAME to avoid conflicts.
	cname := fmt.Sprintf("cdn-applied-preset-example-%d.example.com", time.Now().Unix())

	resource, err := client.CDN.CDNResources.New(context.Background(), cdn.CDNResourceNewParams{
		Cname:       cname,
		OriginGroup: param.NewOpt(originGroupID),
		Active:      gcore.Bool(true),
	})
	if err != nil {
		log.Fatalf("Error creating CDN resource: %v", err)
	}

	fmt.Printf("Created CDN Resource: ID=%d, Cname=%s, Active=%v\n",
		resource.ID, resource.Cname, resource.Active)
	fmt.Println("===========================")

	return resource
}

func createRule(client *gcore.Client, resourceID int64) *cdn.CDNResourceRule {
	fmt.Println("\n=== CREATE RULE ===")

	rule, err := client.CDN.CDNResources.Rules.New(context.Background(), resourceID, cdn.CDNResourceRuleNewParams{
		Name:     "applied-preset-example",
		Rule:     "/static",
		RuleType: 0,
	})
	if err != nil {
		log.Fatalf("Error creating rule: %v", err)
	}

	fmt.Printf("Created Rule: ID=%d, Name=%s\n", rule.ID, rule.Name)
	fmt.Println("===================")

	return rule
}

func applyPreset(client *gcore.Client, presetID, objectID int64) {
	fmt.Println("\n=== APPLY PRESET ===")

	result, err := client.CDN.Presets.Applied.Apply(context.Background(), presetID, cdn.PresetAppliedApplyParams{
		ObjectID: objectID,
	})
	if err != nil {
		log.Fatalf("Error applying preset: %v", err)
	}

	fmt.Printf("Applied preset %d to object %d: %s\n", presetID, objectID, result.Message)
	fmt.Println("====================")
}

func getAppliedObjects(client *gcore.Client, presetID int64) {
	fmt.Println("\n=== GET APPLIED OBJECTS ===")

	result, err := client.CDN.Presets.Applied.GetObjects(context.Background(), presetID)
	if err != nil {
		log.Fatalf("Error getting applied objects: %v", err)
	}

	if applied := result.AsAppliedObjects(); len(applied.ObjectIDs) > 0 {
		fmt.Printf("Preset %d is applied to %s objects: %v\n",
			presetID, applied.ObjectType, applied.ObjectIDs)
	} else {
		fmt.Printf("Preset %d is not applied to any objects: %s\n",
			presetID, result.AsNoAppliedObjects().Message)
	}
	fmt.Println("===========================")
}

func getResourcePreset(client *gcore.Client, resourceID int64) {
	fmt.Println("\n=== GET RESOURCE PRESET ===")

	applied, err := client.CDN.Presets.Applied.GetResourcePreset(context.Background(), resourceID)
	if err != nil {
		log.Fatalf("Error getting resource preset: %v", err)
	}

	fmt.Printf("CDN resource %d has preset: ID=%d, Name=%s, Deletable=%v, ManagedFields=%v\n",
		resourceID, applied.ID, applied.Name, applied.Deletable, applied.Fields)
	fmt.Println("===========================")
}

func getRulePreset(client *gcore.Client, resourceID, ruleID int64) {
	fmt.Println("\n=== GET RULE PRESET ===")

	applied, err := client.CDN.Presets.Applied.GetRulePreset(context.Background(), ruleID, cdn.PresetAppliedGetRulePresetParams{
		ResourceID: resourceID,
	})
	if err != nil {
		log.Fatalf("Error getting rule preset: %v", err)
	}

	fmt.Printf("Rule %d has preset: ID=%d, Name=%s, Deletable=%v, ManagedFields=%v\n",
		ruleID, applied.ID, applied.Name, applied.Deletable, applied.Fields)
	fmt.Println("=======================")
}

func unapplyPreset(client *gcore.Client, presetID, objectID int64) {
	fmt.Println("\n=== UNAPPLY PRESET ===")

	err := client.CDN.Presets.Applied.Unapply(context.Background(), objectID, cdn.PresetAppliedUnapplyParams{
		PresetID: presetID,
	})
	if err != nil {
		log.Fatalf("Error unapplying preset: %v", err)
	}

	fmt.Printf("Unapplied preset %d from object %d\n", presetID, objectID)
	fmt.Println("======================")
}

func deactivateAndDeleteCDNResource(client *gcore.Client, resourceID int64) {
	fmt.Println("\n=== DEACTIVATE AND DELETE CDN RESOURCE ===")

	// DeactivateAndDelete first deactivates the CDN resource by setting
	// active=false, then deletes it — the Delete operation requires the resource
	// to be deactivated first.
	err := client.CDN.CDNResources.DeactivateAndDelete(context.Background(), resourceID)
	if err != nil {
		log.Fatalf("Error deactivating and deleting CDN resource: %v", err)
	}

	fmt.Printf("CDN Resource with ID %d successfully deactivated and deleted\n", resourceID)
	fmt.Println("==========================================")
}
