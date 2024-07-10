// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/clear-street/studio-sdk-go"
	"github.com/clear-street/studio-sdk-go/internal/testutil"
	"github.com/clear-street/studio-sdk-go/option"
)

func TestUsage(t *testing.T) {
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
	entity, err := client.Entities.Get(context.TODO(), "<your_entity_id>")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", entity.EntityID)
}
