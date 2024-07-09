// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/clear-street/studio-sdk-go"
	"github.com/clear-street/studio-sdk-go/internal/testutil"
	"github.com/clear-street/studio-sdk-go/option"
)

func TestEntityPnlSummaryGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := studiosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Entities.PnlSummary.Get(context.TODO(), "x")
	if err != nil {
		var apierr *studiosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
