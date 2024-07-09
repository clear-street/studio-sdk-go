// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/studio-sdk-go/internal/apijson"
	"github.com/stainless-sdks/studio-sdk-go/internal/param"
	"github.com/stainless-sdks/studio-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/studio-sdk-go/option"
)

// AccountBulkOrderService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBulkOrderService] method instead.
type AccountBulkOrderService struct {
	Options []option.RequestOption
}

// NewAccountBulkOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountBulkOrderService(opts ...option.RequestOption) (r *AccountBulkOrderService) {
	r = &AccountBulkOrderService{}
	r.Options = opts
	return
}

// Creates multiple orders in a single request, up to 1000. Note that a successful
// call to this endpoint does not necessarily mean your orders have been accepted,
// e.g. a downstream venue might reject your order. You should therefore utilize
// our WebSocket APIs to listen for changes in order lifecycle events.
//
// The response will contain an array of objects, indicating whether your order was
// submitted. If the order was submitted, the `order_id` field will be populated
// with the order ID assigned to this order. If the order was rejected, the
// `reason` field will be populated with the reason for rejection. The data array
// returned in the response object is guaranteed to be ordered in the same order as
// the orders you provided in the request. Again, note that even if your order was
// submitted, that doesn't mean it was _accepted_, and may still be rejected by
// downstream venues.
func (r *AccountBulkOrderService) New(ctx context.Context, accountID string, body AccountBulkOrderNewParams, opts ...option.RequestOption) (res *AccountBulkOrderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/bulk-orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountBulkOrderNewResponse struct {
	// Array indicating whether each respective order was submitted or not. This array
	// is guaranteed to be sorted in the same order as the orders you provided in your
	// request.
	Data []AccountBulkOrderNewResponseData `json:"data,required"`
	// Total number of orders rejected
	Rejected int64 `json:"rejected,required"`
	// Total number of orders submitted
	Submitted int64                           `json:"submitted,required"`
	JSON      accountBulkOrderNewResponseJSON `json:"-"`
}

// accountBulkOrderNewResponseJSON contains the JSON metadata for the struct
// [AccountBulkOrderNewResponse]
type accountBulkOrderNewResponseJSON struct {
	Data        apijson.Field
	Rejected    apijson.Field
	Submitted   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBulkOrderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBulkOrderNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBulkOrderNewResponseData struct {
	// True if the order was submitted successfully, false otherwise.
	Submitted bool `json:"submitted,required"`
	// If the order was submitted, the order ID assigned to this order. Empty if the
	// order was rejected.
	OrderID string `json:"order_id"`
	// If the order rejected, the reason for rejection. Empty if the order was
	// accepted.
	Reason string                              `json:"reason"`
	JSON   accountBulkOrderNewResponseDataJSON `json:"-"`
}

// accountBulkOrderNewResponseDataJSON contains the JSON metadata for the struct
// [AccountBulkOrderNewResponseData]
type accountBulkOrderNewResponseDataJSON struct {
	Submitted   apijson.Field
	OrderID     apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBulkOrderNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBulkOrderNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountBulkOrderNewParams struct {
	// An array of orders to create.
	Orders param.Field[[]AccountBulkOrderNewParamsOrder] `json:"orders,required"`
}

func (r AccountBulkOrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBulkOrderNewParamsOrder struct {
	// The type of order, can be one of the following:
	//
	//   - `limit`: A limit order will execute at-or-better than the limit price you
	//     specify
	//   - `market`: An order that will execute at the prevailing market prices
	OrderType param.Field[AccountBulkOrderNewParamsOrdersOrderType] `json:"order_type,required"`
	// The maximum quantity to be executed.
	Quantity param.Field[string] `json:"quantity,required"`
	// Buy, sell, sell-short indicator.
	Side param.Field[AccountBulkOrderNewParamsOrdersSide] `json:"side,required"`
	// Strategy type used for execution, can be one of below. Note, we use sensible
	// defaults for strategy parameters at the moment. In future, we will provide a way
	// to provide specify these parameters.
	//
	// - `sor`: Smart order router
	// - `dark`: Dark pool
	// - `ap`: Arrival price
	// - `pov`: Percentage of volume
	// - `twap`: Time weighted average price
	// - `vwap`: Volume weighted average price
	//
	// For more information on these strategies, please refer to our
	// [documentation](https://docs.clearstreet.io/studio/docs/execution-strategies).
	StrategyType param.Field[AccountBulkOrderNewParamsOrdersStrategyType] `json:"strategy_type,required"`
	// The symbol this order is for. See `symbol_format` for supported symbol formats.
	Symbol param.Field[string] `json:"symbol,required"`
	// The lifecycle enforcement of this order.
	//
	//   - `day`: The order will exist for the duration of the current trading session
	//   - `ioc`: The order will immediately be executed or cancelled
	//   - `day-plus`: The order will exist only for the duration the current trading
	//     session plus extended hours, if applicable
	//   - `at-open`: The order will exist only for the opening auction of the next
	//     session
	//   - `at-close`: The order will exist only for the closing auction of the current
	//     session
	TimeInForce param.Field[AccountBulkOrderNewParamsOrdersTimeInForce] `json:"time_in_force,required"`
	// Name of the broker that provided you inventory for a short-sale. Required if
	// `side` is `sell-short`. If you procured inventory through us, you can use
	// `CLST`.
	LocateBroker param.Field[string] `json:"locate_broker"`
	// The price to execute at-or-better.
	Price param.Field[string] `json:"price"`
	// An ID that you provide.
	ReferenceID param.Field[string] `json:"reference_id"`
	// Denotes the format of the provided `symbol` field.
	SymbolFormat param.Field[AccountBulkOrderNewParamsOrdersSymbolFormat] `json:"symbol_format"`
}

func (r AccountBulkOrderNewParamsOrder) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of order, can be one of the following:
//
//   - `limit`: A limit order will execute at-or-better than the limit price you
//     specify
//   - `market`: An order that will execute at the prevailing market prices
type AccountBulkOrderNewParamsOrdersOrderType string

const (
	AccountBulkOrderNewParamsOrdersOrderTypeLimit  AccountBulkOrderNewParamsOrdersOrderType = "limit"
	AccountBulkOrderNewParamsOrdersOrderTypeMarket AccountBulkOrderNewParamsOrdersOrderType = "market"
)

func (r AccountBulkOrderNewParamsOrdersOrderType) IsKnown() bool {
	switch r {
	case AccountBulkOrderNewParamsOrdersOrderTypeLimit, AccountBulkOrderNewParamsOrdersOrderTypeMarket:
		return true
	}
	return false
}

// Buy, sell, sell-short indicator.
type AccountBulkOrderNewParamsOrdersSide string

const (
	AccountBulkOrderNewParamsOrdersSideBuy       AccountBulkOrderNewParamsOrdersSide = "buy"
	AccountBulkOrderNewParamsOrdersSideSell      AccountBulkOrderNewParamsOrdersSide = "sell"
	AccountBulkOrderNewParamsOrdersSideSellShort AccountBulkOrderNewParamsOrdersSide = "sell-short"
)

func (r AccountBulkOrderNewParamsOrdersSide) IsKnown() bool {
	switch r {
	case AccountBulkOrderNewParamsOrdersSideBuy, AccountBulkOrderNewParamsOrdersSideSell, AccountBulkOrderNewParamsOrdersSideSellShort:
		return true
	}
	return false
}

// Strategy type used for execution, can be one of below. Note, we use sensible
// defaults for strategy parameters at the moment. In future, we will provide a way
// to provide specify these parameters.
//
// - `sor`: Smart order router
// - `dark`: Dark pool
// - `ap`: Arrival price
// - `pov`: Percentage of volume
// - `twap`: Time weighted average price
// - `vwap`: Volume weighted average price
//
// For more information on these strategies, please refer to our
// [documentation](https://docs.clearstreet.io/studio/docs/execution-strategies).
type AccountBulkOrderNewParamsOrdersStrategyType string

const (
	AccountBulkOrderNewParamsOrdersStrategyTypeSor  AccountBulkOrderNewParamsOrdersStrategyType = "sor"
	AccountBulkOrderNewParamsOrdersStrategyTypeDark AccountBulkOrderNewParamsOrdersStrategyType = "dark"
	AccountBulkOrderNewParamsOrdersStrategyTypeAp   AccountBulkOrderNewParamsOrdersStrategyType = "ap"
	AccountBulkOrderNewParamsOrdersStrategyTypePov  AccountBulkOrderNewParamsOrdersStrategyType = "pov"
	AccountBulkOrderNewParamsOrdersStrategyTypeTwap AccountBulkOrderNewParamsOrdersStrategyType = "twap"
	AccountBulkOrderNewParamsOrdersStrategyTypeVwap AccountBulkOrderNewParamsOrdersStrategyType = "vwap"
)

func (r AccountBulkOrderNewParamsOrdersStrategyType) IsKnown() bool {
	switch r {
	case AccountBulkOrderNewParamsOrdersStrategyTypeSor, AccountBulkOrderNewParamsOrdersStrategyTypeDark, AccountBulkOrderNewParamsOrdersStrategyTypeAp, AccountBulkOrderNewParamsOrdersStrategyTypePov, AccountBulkOrderNewParamsOrdersStrategyTypeTwap, AccountBulkOrderNewParamsOrdersStrategyTypeVwap:
		return true
	}
	return false
}

// The lifecycle enforcement of this order.
//
//   - `day`: The order will exist for the duration of the current trading session
//   - `ioc`: The order will immediately be executed or cancelled
//   - `day-plus`: The order will exist only for the duration the current trading
//     session plus extended hours, if applicable
//   - `at-open`: The order will exist only for the opening auction of the next
//     session
//   - `at-close`: The order will exist only for the closing auction of the current
//     session
type AccountBulkOrderNewParamsOrdersTimeInForce string

const (
	AccountBulkOrderNewParamsOrdersTimeInForceDay     AccountBulkOrderNewParamsOrdersTimeInForce = "day"
	AccountBulkOrderNewParamsOrdersTimeInForceIoc     AccountBulkOrderNewParamsOrdersTimeInForce = "ioc"
	AccountBulkOrderNewParamsOrdersTimeInForceDayPlus AccountBulkOrderNewParamsOrdersTimeInForce = "day-plus"
	AccountBulkOrderNewParamsOrdersTimeInForceAtOpen  AccountBulkOrderNewParamsOrdersTimeInForce = "at-open"
	AccountBulkOrderNewParamsOrdersTimeInForceAtClose AccountBulkOrderNewParamsOrdersTimeInForce = "at-close"
)

func (r AccountBulkOrderNewParamsOrdersTimeInForce) IsKnown() bool {
	switch r {
	case AccountBulkOrderNewParamsOrdersTimeInForceDay, AccountBulkOrderNewParamsOrdersTimeInForceIoc, AccountBulkOrderNewParamsOrdersTimeInForceDayPlus, AccountBulkOrderNewParamsOrdersTimeInForceAtOpen, AccountBulkOrderNewParamsOrdersTimeInForceAtClose:
		return true
	}
	return false
}

// Denotes the format of the provided `symbol` field.
type AccountBulkOrderNewParamsOrdersSymbolFormat string

const (
	AccountBulkOrderNewParamsOrdersSymbolFormatCms AccountBulkOrderNewParamsOrdersSymbolFormat = "cms"
	AccountBulkOrderNewParamsOrdersSymbolFormatOsi AccountBulkOrderNewParamsOrdersSymbolFormat = "osi"
)

func (r AccountBulkOrderNewParamsOrdersSymbolFormat) IsKnown() bool {
	switch r {
	case AccountBulkOrderNewParamsOrdersSymbolFormatCms, AccountBulkOrderNewParamsOrdersSymbolFormatOsi:
		return true
	}
	return false
}
