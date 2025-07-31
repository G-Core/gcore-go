// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/security"
)

func TestEventListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Security.Events.List(context.TODO(), security.EventListParams{
		AlertType: security.EventListParamsAlertTypeDDOSAlert,
		DateFrom: security.EventListParamsDateFromUnion{
			OfTime: gcore.Time(time.Now()),
		},
		DateTo: security.EventListParamsDateToUnion{
			OfTime: gcore.Time(time.Now()),
		},
		Limit:               gcore.Int(1),
		Offset:              gcore.Int(0),
		Ordering:            security.EventListParamsOrderingAttackStartTime,
		TargetedIPAddresses: gcore.String("targeted_ip_addresses"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
