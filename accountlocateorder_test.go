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

func TestAccountLocateOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.LocateOrders.New(
		context.TODO(),
		"x",
		studiosdk.AccountLocateOrderNewParams{
			Mpid:        studiosdk.F("x"),
			Quantity:    studiosdk.F("x"),
			ReferenceID: studiosdk.F("my-order-id-123"),
			Symbol:      studiosdk.F("AAPL"),
			Comments:    studiosdk.F("comments"),
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

func TestAccountLocateOrderGet(t *testing.T) {
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
	_, err := client.Accounts.LocateOrders.Get(
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

func TestAccountLocateOrderUpdate(t *testing.T) {
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
	err := client.Accounts.LocateOrders.Update(
		context.TODO(),
		"x",
		"x",
		studiosdk.AccountLocateOrderUpdateParams{
			Accept: studiosdk.F(true),
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

func TestAccountLocateOrderList(t *testing.T) {
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
	_, err := client.Accounts.LocateOrders.List(context.TODO(), "x")
	if err != nil {
		var apierr *studiosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
