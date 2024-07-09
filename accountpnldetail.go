// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/clear-street/studio-sdk-go/internal/apijson"
	"github.com/clear-street/studio-sdk-go/internal/requestconfig"
	"github.com/clear-street/studio-sdk-go/option"
)

// AccountPnlDetailService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPnlDetailService] method instead.
type AccountPnlDetailService struct {
	Options []option.RequestOption
}

// NewAccountPnlDetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPnlDetailService(opts ...option.RequestOption) (r *AccountPnlDetailService) {
	r = &AccountPnlDetailService{}
	r.Options = opts
	return
}

// List PNL details for a given account.
func (r *AccountPnlDetailService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountPnlDetailListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pnl-details", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountPnlDetailListResponse struct {
	Data []AccountPnlDetailListResponseData `json:"data,required"`
	JSON accountPnlDetailListResponseJSON   `json:"-"`
}

// accountPnlDetailListResponseJSON contains the JSON metadata for the struct
// [AccountPnlDetailListResponse]
type accountPnlDetailListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPnlDetailListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPnlDetailListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPnlDetailListResponseData struct {
	// Account ID for the account.
	AccountID string `json:"account_id,required"`
	// The asset class of the symbol.
	AssetClass AccountPnlDetailListResponseDataAssetClass `json:"asset_class,required"`
	// Quantity of a given instrument bought.
	BoughtQuantity string `json:"bought_quantity,required"`
	// Total buys of a given instrument.
	Buys int64 `json:"buys,required"`
	// Profit and loss from intraday trading activities.
	DayPnl float64 `json:"day_pnl,required"`
	// Name of the legal entity.
	EntityID string `json:"entity_id,required"`
	// Absolute market value of long and short market values.
	GrossMarketValue float64 `json:"gross_market_value,required"`
	// Market value net of long and short market values.
	NetMarketValue float64 `json:"net_market_value,required"`
	// Profit and loss from previous trading date.
	OvernightPnl float64 `json:"overnight_pnl,required"`
	// Price used in this pnl calculation.
	Price float64 `json:"price,required"`
	// String representation of quantity.
	Quantity string `json:"quantity,required"`
	// Profit and loss realized from position closing trading activity.
	RealizedPnl float64 `json:"realized_pnl,required"`
	// Total sells of a given instrument.
	Sells int64 `json:"sells,required"`
	// Market value of a given instrument a the start of a trading day.
	SodMarketValue float64 `json:"sod_market_value,required"`
	// Price at the start of a trading day.
	SodPrice float64 `json:"sod_price,required"`
	// Quantity of a given instrument at the start of a trading day.
	SodQuantity string `json:"sod_quantity,required"`
	// Quantity of a given instrument sold.
	SoldQuantity string `json:"sold_quantity,required"`
	Symbol       string `json:"symbol,required"`
	// Description of the symbol.
	SymbolDescription string `json:"symbol_description,required"`
	// Milliseconds since epoch.
	Timestamp int64 `json:"timestamp,required"`
	// Total fees incurred from trading activities.
	TotalFees float64 `json:"total_fees,required"`
	// `realized_pnl + unrealized_pnl`
	TotalPnl float64 `json:"total_pnl,required"`
	// The underlying instrument.
	Underlier string `json:"underlier,required"`
	// Profit and loss from market changes.
	UnrealizedPnl float64                              `json:"unrealized_pnl,required"`
	JSON          accountPnlDetailListResponseDataJSON `json:"-"`
}

// accountPnlDetailListResponseDataJSON contains the JSON metadata for the struct
// [AccountPnlDetailListResponseData]
type accountPnlDetailListResponseDataJSON struct {
	AccountID         apijson.Field
	AssetClass        apijson.Field
	BoughtQuantity    apijson.Field
	Buys              apijson.Field
	DayPnl            apijson.Field
	EntityID          apijson.Field
	GrossMarketValue  apijson.Field
	NetMarketValue    apijson.Field
	OvernightPnl      apijson.Field
	Price             apijson.Field
	Quantity          apijson.Field
	RealizedPnl       apijson.Field
	Sells             apijson.Field
	SodMarketValue    apijson.Field
	SodPrice          apijson.Field
	SodQuantity       apijson.Field
	SoldQuantity      apijson.Field
	Symbol            apijson.Field
	SymbolDescription apijson.Field
	Timestamp         apijson.Field
	TotalFees         apijson.Field
	TotalPnl          apijson.Field
	Underlier         apijson.Field
	UnrealizedPnl     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountPnlDetailListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPnlDetailListResponseDataJSON) RawJSON() string {
	return r.raw
}

// The asset class of the symbol.
type AccountPnlDetailListResponseDataAssetClass string

const (
	AccountPnlDetailListResponseDataAssetClassOther  AccountPnlDetailListResponseDataAssetClass = "other"
	AccountPnlDetailListResponseDataAssetClassEquity AccountPnlDetailListResponseDataAssetClass = "equity"
	AccountPnlDetailListResponseDataAssetClassOption AccountPnlDetailListResponseDataAssetClass = "option"
	AccountPnlDetailListResponseDataAssetClassDebt   AccountPnlDetailListResponseDataAssetClass = "debt"
)

func (r AccountPnlDetailListResponseDataAssetClass) IsKnown() bool {
	switch r {
	case AccountPnlDetailListResponseDataAssetClassOther, AccountPnlDetailListResponseDataAssetClassEquity, AccountPnlDetailListResponseDataAssetClassOption, AccountPnlDetailListResponseDataAssetClassDebt:
		return true
	}
	return false
}
