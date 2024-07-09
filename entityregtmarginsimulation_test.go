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

func TestEntityRegtMarginSimulationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.RegtMarginSimulations.New(
		context.TODO(),
		"x",
		studiosdk.EntityRegtMarginSimulationNewParams{
			Name:           studiosdk.F("string"),
			IgnoreExisting: studiosdk.F(true),
			Prices: studiosdk.F([]studiosdk.EntityRegtMarginSimulationNewParamsPrice{{
				Symbol:       studiosdk.F("AAPL"),
				SymbolFormat: studiosdk.F(studiosdk.EntityRegtMarginSimulationNewParamsPricesSymbolFormatCms),
				Price:        studiosdk.F("x"),
			}, {
				Symbol:       studiosdk.F("AAPL"),
				SymbolFormat: studiosdk.F(studiosdk.EntityRegtMarginSimulationNewParamsPricesSymbolFormatCms),
				Price:        studiosdk.F("x"),
			}, {
				Symbol:       studiosdk.F("AAPL"),
				SymbolFormat: studiosdk.F(studiosdk.EntityRegtMarginSimulationNewParamsPricesSymbolFormatCms),
				Price:        studiosdk.F("x"),
			}}),
			Trades: studiosdk.F([]studiosdk.EntityRegtMarginSimulationNewParamsTrade{{
				Symbol:       studiosdk.F("AAPL"),
				SymbolFormat: studiosdk.F(studiosdk.EntityRegtMarginSimulationNewParamsTradesSymbolFormatCms),
				Side:         studiosdk.F(studiosdk.EntityRegtMarginSimulationNewParamsTradesSideBuy),
				Quantity:     studiosdk.F("x"),
				Price:        studiosdk.F("x"),
			}, {
				Symbol:       studiosdk.F("AAPL"),
				SymbolFormat: studiosdk.F(studiosdk.EntityRegtMarginSimulationNewParamsTradesSymbolFormatCms),
				Side:         studiosdk.F(studiosdk.EntityRegtMarginSimulationNewParamsTradesSideBuy),
				Quantity:     studiosdk.F("x"),
				Price:        studiosdk.F("x"),
			}, {
				Symbol:       studiosdk.F("AAPL"),
				SymbolFormat: studiosdk.F(studiosdk.EntityRegtMarginSimulationNewParamsTradesSymbolFormatCms),
				Side:         studiosdk.F(studiosdk.EntityRegtMarginSimulationNewParamsTradesSideBuy),
				Quantity:     studiosdk.F("x"),
				Price:        studiosdk.F("x"),
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

func TestEntityRegtMarginSimulationGet(t *testing.T) {
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
	_, err := client.Entities.RegtMarginSimulations.Get(
		context.TODO(),
		"x",
		"6460030d-8ed4-19d3-818e-e87b36e90005",
	)
	if err != nil {
		var apierr *studiosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
