// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/studio-sdk-go/internal/apijson"
	"github.com/stainless-sdks/studio-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/studio-sdk-go/internal/param"
	"github.com/stainless-sdks/studio-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/studio-sdk-go/option"
)

// InstrumentService contains methods and other services that help with interacting
// with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstrumentService] method instead.
type InstrumentService struct {
	Options []option.RequestOption
}

// NewInstrumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstrumentService(opts ...option.RequestOption) (r *InstrumentService) {
	r = &InstrumentService{}
	r.Options = opts
	return
}

// Get an instrument by the given symbol
func (r *InstrumentService) Get(ctx context.Context, symbol string, query InstrumentGetParams, opts ...option.RequestOption) (res *Instrument, err error) {
	opts = append(r.Options[:], opts...)
	if symbol == "" {
		err = errors.New("missing required symbol parameter")
		return
	}
	path := fmt.Sprintf("instruments/%s", symbol)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Instrument struct {
	// The asset class of the symbol.
	AssetClass InstrumentAssetClass `json:"asset_class,required"`
	// A description of the instrument.
	Description string `json:"description,required"`
	// The primary exchange for the instrument.
	PrimaryExchange string             `json:"primary_exchange,required"`
	Symbols         []InstrumentSymbol `json:"symbols,required"`
	JSON            instrumentJSON     `json:"-"`
}

// instrumentJSON contains the JSON metadata for the struct [Instrument]
type instrumentJSON struct {
	AssetClass      apijson.Field
	Description     apijson.Field
	PrimaryExchange apijson.Field
	Symbols         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Instrument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instrumentJSON) RawJSON() string {
	return r.raw
}

// The asset class of the symbol.
type InstrumentAssetClass string

const (
	InstrumentAssetClassOther  InstrumentAssetClass = "other"
	InstrumentAssetClassEquity InstrumentAssetClass = "equity"
	InstrumentAssetClassOption InstrumentAssetClass = "option"
	InstrumentAssetClassDebt   InstrumentAssetClass = "debt"
)

func (r InstrumentAssetClass) IsKnown() bool {
	switch r {
	case InstrumentAssetClassOther, InstrumentAssetClassEquity, InstrumentAssetClassOption, InstrumentAssetClassDebt:
		return true
	}
	return false
}

type InstrumentSymbol struct {
	Symbol string `json:"symbol"`
	// Denotes the format of the provided `symbol` field.
	SymbolFormat InstrumentSymbolsSymbolFormat `json:"symbol_format"`
	JSON         instrumentSymbolJSON          `json:"-"`
}

// instrumentSymbolJSON contains the JSON metadata for the struct
// [InstrumentSymbol]
type instrumentSymbolJSON struct {
	Symbol       apijson.Field
	SymbolFormat apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InstrumentSymbol) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instrumentSymbolJSON) RawJSON() string {
	return r.raw
}

// Denotes the format of the provided `symbol` field.
type InstrumentSymbolsSymbolFormat string

const (
	InstrumentSymbolsSymbolFormatCms InstrumentSymbolsSymbolFormat = "cms"
	InstrumentSymbolsSymbolFormatOsi InstrumentSymbolsSymbolFormat = "osi"
)

func (r InstrumentSymbolsSymbolFormat) IsKnown() bool {
	switch r {
	case InstrumentSymbolsSymbolFormatCms, InstrumentSymbolsSymbolFormatOsi:
		return true
	}
	return false
}

type InstrumentGetParams struct {
	// The format of the provided symbol.
	SymbolFormat param.Field[InstrumentGetParamsSymbolFormat] `query:"symbol_format"`
}

// URLQuery serializes [InstrumentGetParams]'s query parameters as `url.Values`.
func (r InstrumentGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The format of the provided symbol.
type InstrumentGetParamsSymbolFormat string

const (
	InstrumentGetParamsSymbolFormatCms InstrumentGetParamsSymbolFormat = "cms"
	InstrumentGetParamsSymbolFormatOsi InstrumentGetParamsSymbolFormat = "osi"
)

func (r InstrumentGetParamsSymbolFormat) IsKnown() bool {
	switch r {
	case InstrumentGetParamsSymbolFormatCms, InstrumentGetParamsSymbolFormatOsi:
		return true
	}
	return false
}
