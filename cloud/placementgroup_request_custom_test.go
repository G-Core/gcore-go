package cloud_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"slices"
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/option"
)

func TestPlacementGroupListGlobalFilters(t *testing.T) {
	const path = "/cloud/v1/projects/1/placement_groups"
	members := []string{"instance-1", "instance-2"}
	for _, regionIDs := range [][]int64{nil, {3, 5}} {
		t.Run("regions_"+strconv.Itoa(len(regionIDs)), func(t *testing.T) {
			var calls atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				page := calls.Add(1) - 1
				wantQuery := url.Values{
					"limit":        {"1"},
					"offset":       {strconv.Itoa(int(page))},
					"name[prefix]": {"prod-", "web-"},
					"policies":     {"affinity", "anti-affinity"},
					"member_ids":   members,
				}
				for _, id := range regionIDs {
					wantQuery.Add("region_ids", strconv.FormatInt(id, 10))
				}
				if r.Method != http.MethodGet || r.URL.Path != path {
					t.Errorf("request = %s %s, want GET %s", r.Method, r.URL.Path, path)
				}
				if got := r.URL.Query(); !reflect.DeepEqual(got, wantQuery) {
					t.Errorf("page %d: query = %v, want %v", page, got, wantQuery)
				}
				var managedBy any
				if page == 1 {
					managedBy = map[string]string{"type": "system"}
				}
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(map[string]any{
					"count": 2,
					"results": []map[string]any{{
						"id": "group-" + strconv.Itoa(int(page)), "name": "prod-app", "policy": "affinity",
						"region_id": 3 + 2*page, "managed_by": managedBy, "member_ids": members,
						"created_at": "2026-11-19T12:00:00Z", "updated_at": "2026-11-19T12:00:00Z",
					}},
				})
			}))
			defer server.Close()
			client := newTaskFlowClient(server)
			pager := client.Cloud.PlacementGroups.ListAutoPaging(context.Background(), cloud.PlacementGroupListParams{
				Limit: gcore.Int(1), Offset: gcore.Int(0), RegionIDs: regionIDs, MemberIDs: members,
				Name:     cloud.PlacementGroupListParamsName{Prefix: []string{"prod-", "web-"}},
				Policies: []string{"affinity", "anti-affinity"},
			})
			var results []cloud.PlacementGroup
			for pager.Next() {
				results = append(results, pager.Current())
			}
			if err := pager.Err(); err != nil {
				t.Fatal(err)
			}
			if len(results) != 2 || calls.Load() != 2 {
				t.Fatalf("got %d results from %d requests, want two pages", len(results), calls.Load())
			}
			if results[0].RegionID != 3 || results[1].RegionID != 5 {
				t.Errorf("regions = %d, %d, want 3 and 5", results[0].RegionID, results[1].RegionID)
			}
			if results[0].JSON.ManagedBy.Valid() || !results[1].JSON.ManagedBy.Valid() || results[1].ManagedBy.Type != "system" {
				t.Errorf("unexpected managed_by values: %s; %s", results[0].RawJSON(), results[1].RawJSON())
			}
			if !slices.Equal(results[0].MemberIDs, members) {
				t.Errorf("member_ids = %v, want %v", results[0].MemberIDs, members)
			}
		})
	}
}

func TestPlacementGroupCreateWithNoDefaultRegion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const path = "/cloud/v1/projects/1/placement_groups"
		if r.Method != http.MethodPost || r.URL.Path != path || r.URL.RawQuery != "" {
			t.Errorf("request = %s %s, want POST %s", r.Method, r.URL, path)
		}
		var body map[string]any
		want := map[string]any{"region_id": float64(5), "name": "web", "policy": "soft-affinity"}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || !reflect.DeepEqual(body, want) {
			t.Errorf("body = %v, error = %v, want %v", body, err, want)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"tasks": []string{"task-1"}})
	}))
	defer server.Close()
	service := cloud.NewPlacementGroupService(option.WithBaseURL(server.URL), option.WithAPIKey("test"), option.WithCloudProjectID(1))
	result, err := service.New(context.Background(), cloud.PlacementGroupNewParams{
		RegionID: 5, Name: "web", Policy: cloud.PlacementGroupNewParamsPolicySoftAffinity,
	})
	if err != nil || result == nil || !slices.Equal(result.Tasks, []string{"task-1"}) {
		t.Fatalf("create = %#v, error = %v, want task-1", result, err)
	}
}
