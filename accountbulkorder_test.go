// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/studio-sdk-go"
	"github.com/stainless-sdks/studio-sdk-go/internal/testutil"
	"github.com/stainless-sdks/studio-sdk-go/option"
)

func TestAccountBulkOrderNew(t *testing.T) {
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
	_, err := client.Accounts.BulkOrders.New(
		context.TODO(),
		"x",
		studiosdk.AccountBulkOrderNewParams{
			Orders: studiosdk.F([]studiosdk.AccountBulkOrderNewParamsOrder{{
				ReferenceID:  studiosdk.F("my-order-id-123"),
				OrderType:    studiosdk.F(studiosdk.AccountBulkOrderNewParamsOrdersOrderTypeLimit),
				Side:         studiosdk.F(studiosdk.AccountBulkOrderNewParamsOrdersSideBuy),
				Quantity:     studiosdk.F("x"),
				Price:        studiosdk.F("x"),
				TimeInForce:  studiosdk.F(studiosdk.AccountBulkOrderNewParamsOrdersTimeInForceDay),
				StrategyType: studiosdk.F(studiosdk.AccountBulkOrderNewParamsOrdersStrategyTypeSor),
				LocateBroker: studiosdk.F("x"),
				Symbol:       studiosdk.F("AAPL"),
				SymbolFormat: studiosdk.F(studiosdk.AccountBulkOrderNewParamsOrdersSymbolFormatCms),
			}}),
		},
	)
	if err != nil {
		var apierr *studiosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
