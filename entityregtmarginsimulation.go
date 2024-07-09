// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sachnk/studio-sdk-go/internal/apijson"
	"github.com/sachnk/studio-sdk-go/internal/param"
	"github.com/sachnk/studio-sdk-go/internal/requestconfig"
	"github.com/sachnk/studio-sdk-go/option"
	"github.com/sachnk/studio-sdk-go/shared"
)

// EntityRegtMarginSimulationService contains methods and other services that help
// with interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityRegtMarginSimulationService] method instead.
type EntityRegtMarginSimulationService struct {
	Options []option.RequestOption
}

// NewEntityRegtMarginSimulationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEntityRegtMarginSimulationService(opts ...option.RequestOption) (r *EntityRegtMarginSimulationService) {
	r = &EntityRegtMarginSimulationService{}
	r.Options = opts
	return
}

// Simulate Reg-T margin calculation for a given hypothetical set of prices and/or
// trades. This is useful for understanding the impact of price fluctuations or
// trades on margin requirements. Once a simulation is created, it remains
// available for 48-hours, after which it will automatically be deleted.
//
// Simulations created through the API are visible in the Studio UI under the Risk
// & Margin section, after enabling the "Risk Simulations" toggle.
func (r *EntityRegtMarginSimulationService) New(ctx context.Context, entityID string, body EntityRegtMarginSimulationNewParams, opts ...option.RequestOption) (res *EntityRegtMarginSimulationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/regt-margin-simulations", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a Reg-T margin simluation that was previously created. Note, simulations are
// automatically deleted after 48-hours.
func (r *EntityRegtMarginSimulationService) Get(ctx context.Context, entityID string, simulationID SimulationIDParam, opts ...option.RequestOption) (res *shared.RegtMarginSimulation, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	if simulationID == "" {
		err = errors.New("missing required simulation_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/regt-margin-simulations/%s", entityID, simulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SimulationID = string

type SimulationIDParam = string

type EntityRegtMarginSimulationNewResponse struct {
	// Unique ID for a simulation.
	SimulationID SimulationID                              `json:"simulation_id,required" format:"uuid"`
	JSON         entityRegtMarginSimulationNewResponseJSON `json:"-"`
}

// entityRegtMarginSimulationNewResponseJSON contains the JSON metadata for the
// struct [EntityRegtMarginSimulationNewResponse]
type entityRegtMarginSimulationNewResponseJSON struct {
	SimulationID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntityRegtMarginSimulationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityRegtMarginSimulationNewResponseJSON) RawJSON() string {
	return r.raw
}

type EntityRegtMarginSimulationNewParams struct {
	// A name for this simulation for reference.
	Name param.Field[string] `json:"name,required"`
	// If true, the simulation will ignore any existing positions and balances in the
	// account. Set to true if you want to simulate from a clean slate, i.e. an empty
	// account.
	IgnoreExisting param.Field[bool] `json:"ignore_existing"`
	// List of prices to use in the simulation, i.e. fair-market-values you specify for
	// each symbol. If this is not provided, current market prices will be used, if
	// they are available.
	Prices param.Field[[]EntityRegtMarginSimulationNewParamsPrice] `json:"prices"`
	// List of hypothetical trades to include in the simulation, if any.
	Trades param.Field[[]EntityRegtMarginSimulationNewParamsTrade] `json:"trades"`
}

func (r EntityRegtMarginSimulationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityRegtMarginSimulationNewParamsPrice struct {
	// The price to use in the simulation.
	Price param.Field[string] `json:"price,required"`
	// The symbol for the instrument.
	Symbol param.Field[string] `json:"symbol,required"`
	// Denotes the format of the provided `symbol` field.
	SymbolFormat param.Field[EntityRegtMarginSimulationNewParamsPricesSymbolFormat] `json:"symbol_format"`
}

func (r EntityRegtMarginSimulationNewParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Denotes the format of the provided `symbol` field.
type EntityRegtMarginSimulationNewParamsPricesSymbolFormat string

const (
	EntityRegtMarginSimulationNewParamsPricesSymbolFormatCms EntityRegtMarginSimulationNewParamsPricesSymbolFormat = "cms"
	EntityRegtMarginSimulationNewParamsPricesSymbolFormatOsi EntityRegtMarginSimulationNewParamsPricesSymbolFormat = "osi"
)

func (r EntityRegtMarginSimulationNewParamsPricesSymbolFormat) IsKnown() bool {
	switch r {
	case EntityRegtMarginSimulationNewParamsPricesSymbolFormatCms, EntityRegtMarginSimulationNewParamsPricesSymbolFormatOsi:
		return true
	}
	return false
}

type EntityRegtMarginSimulationNewParamsTrade struct {
	// The price of the simulated trade.
	Price param.Field[string] `json:"price,required"`
	// The quantity of the simulated trade.
	Quantity param.Field[string] `json:"quantity,required"`
	// The side of the simulated trade.
	Side param.Field[EntityRegtMarginSimulationNewParamsTradesSide] `json:"side,required"`
	// The symbol for the instrument.
	Symbol param.Field[string] `json:"symbol,required"`
	// Denotes the format of the provided `symbol` field.
	SymbolFormat param.Field[EntityRegtMarginSimulationNewParamsTradesSymbolFormat] `json:"symbol_format"`
}

func (r EntityRegtMarginSimulationNewParamsTrade) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The side of the simulated trade.
type EntityRegtMarginSimulationNewParamsTradesSide string

const (
	EntityRegtMarginSimulationNewParamsTradesSideBuy  EntityRegtMarginSimulationNewParamsTradesSide = "buy"
	EntityRegtMarginSimulationNewParamsTradesSideSell EntityRegtMarginSimulationNewParamsTradesSide = "sell"
)

func (r EntityRegtMarginSimulationNewParamsTradesSide) IsKnown() bool {
	switch r {
	case EntityRegtMarginSimulationNewParamsTradesSideBuy, EntityRegtMarginSimulationNewParamsTradesSideSell:
		return true
	}
	return false
}

// Denotes the format of the provided `symbol` field.
type EntityRegtMarginSimulationNewParamsTradesSymbolFormat string

const (
	EntityRegtMarginSimulationNewParamsTradesSymbolFormatCms EntityRegtMarginSimulationNewParamsTradesSymbolFormat = "cms"
	EntityRegtMarginSimulationNewParamsTradesSymbolFormatOsi EntityRegtMarginSimulationNewParamsTradesSymbolFormat = "osi"
)

func (r EntityRegtMarginSimulationNewParamsTradesSymbolFormat) IsKnown() bool {
	switch r {
	case EntityRegtMarginSimulationNewParamsTradesSymbolFormatCms, EntityRegtMarginSimulationNewParamsTradesSymbolFormatOsi:
		return true
	}
	return false
}
