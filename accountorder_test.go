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

func TestAccountOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Orders.New(
		context.TODO(),
		"x",
		studiosdk.AccountOrderNewParams{
			OrderType:    studiosdk.F(studiosdk.AccountOrderNewParamsOrderTypeLimit),
			Quantity:     studiosdk.F("x"),
			Side:         studiosdk.F(studiosdk.AccountOrderNewParamsSideBuy),
			StrategyType: studiosdk.F(studiosdk.AccountOrderNewParamsStrategyTypeSor),
			Symbol:       studiosdk.F("AAPL"),
			TimeInForce:  studiosdk.F(studiosdk.AccountOrderNewParamsTimeInForceDay),
			LocateBroker: studiosdk.F("x"),
			Price:        studiosdk.F("x"),
			ReferenceID:  studiosdk.F("my-order-id-123"),
			SymbolFormat: studiosdk.F(studiosdk.AccountOrderNewParamsSymbolFormatCms),
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

func TestAccountOrderGet(t *testing.T) {
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
	_, err := client.Accounts.Orders.Get(
		context.TODO(),
		"x",
		"x",
	)
	if err != nil {
		var apierr *studiosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Orders.List(
		context.TODO(),
		"x",
		studiosdk.AccountOrderListParams{
			From:      studiosdk.F(int64(1710613560668)),
			PageSize:  studiosdk.F(int64(1)),
			PageToken: studiosdk.F("string"),
			To:        studiosdk.F(int64(1710613560668)),
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

func TestAccountOrderDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Orders.Delete(
		context.TODO(),
		"x",
		studiosdk.AccountOrderDeleteParams{
			Symbol:       studiosdk.F("AAPL"),
			SymbolFormat: studiosdk.F(studiosdk.AccountOrderDeleteParamsSymbolFormatCms),
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

func TestAccountOrderCancel(t *testing.T) {
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
	err := client.Accounts.Orders.Cancel(
		context.TODO(),
		"x",
		"x",
	)
	if err != nil {
		var apierr *studiosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
