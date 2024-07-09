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
	"github.com/stainless-sdks/studio-sdk-go/shared"
)

// AccountOrderService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountOrderService] method instead.
type AccountOrderService struct {
	Options []option.RequestOption
}

// NewAccountOrderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountOrderService(opts ...option.RequestOption) (r *AccountOrderService) {
	r = &AccountOrderService{}
	r.Options = opts
	return
}

// Creates a new order and sends to our internal systems for execution. Note that a
// successful call to this endpoint does not necessarily mean your order has been
// accepted, e.g. a downstream venue might reject your order. You should therefore
// utilize our WebSocket APIs to listen for changes in order lifecycle events.
func (r *AccountOrderService) New(ctx context.Context, accountID string, body AccountOrderNewParams, opts ...option.RequestOption) (res *AccountOrderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an order that was previously created.
func (r *AccountOrderService) Get(ctx context.Context, accountID string, orderID string, opts ...option.RequestOption) (res *AccountOrderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/orders/%s", accountID, orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List orders for a given account for the current trading day, filtered on the
// given query parameters.
func (r *AccountOrderService) List(ctx context.Context, accountID string, query AccountOrderListParams, opts ...option.RequestOption) (res *AccountOrderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Attempts to cancel all open orders for a given account. Cancelling an order
// cannot be guaranteed as there might be in-flight executions.
func (r *AccountOrderService) Delete(ctx context.Context, accountID string, body AccountOrderDeleteParams, opts ...option.RequestOption) (res *AccountOrderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Attempts to cancel an existing order. Cancelling an order cannot be guaranteed
// as there might be in-flight executions.
func (r *AccountOrderService) Cancel(ctx context.Context, accountID string, orderID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/orders/%s", accountID, orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AccountOrderNewResponse struct {
	// An internally generated unique ID for this order.
	OrderID string                      `json:"order_id,required"`
	JSON    accountOrderNewResponseJSON `json:"-"`
}

// accountOrderNewResponseJSON contains the JSON metadata for the struct
// [AccountOrderNewResponse]
type accountOrderNewResponseJSON struct {
	OrderID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountOrderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOrderNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOrderGetResponse struct {
	Order shared.Order                `json:"order,required"`
	JSON  accountOrderGetResponseJSON `json:"-"`
}

// accountOrderGetResponseJSON contains the JSON metadata for the struct
// [AccountOrderGetResponse]
type accountOrderGetResponseJSON struct {
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountOrderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOrderGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOrderListResponse struct {
	Data []shared.Order `json:"data,required"`
	// Cursor for the next page of results.
	NextPageToken string                       `json:"next_page_token"`
	JSON          accountOrderListResponseJSON `json:"-"`
}

// accountOrderListResponseJSON contains the JSON metadata for the struct
// [AccountOrderListResponse]
type accountOrderListResponseJSON struct {
	Data          apijson.Field
	NextPageToken apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountOrderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOrderListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOrderDeleteResponse struct {
	// Array of order IDs that were attempted to be cancelled.
	Data []string                       `json:"data,required"`
	JSON accountOrderDeleteResponseJSON `json:"-"`
}

// accountOrderDeleteResponseJSON contains the JSON metadata for the struct
// [AccountOrderDeleteResponse]
type accountOrderDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountOrderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOrderDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOrderNewParams struct {
	// The type of order, can be one of the following:
	//
	//   - `limit`: A limit order will execute at-or-better than the limit price you
	//     specify
	//   - `market`: An order that will execute at the prevailing market prices
	OrderType param.Field[AccountOrderNewParamsOrderType] `json:"order_type,required"`
	// The maximum quantity to be executed.
	Quantity param.Field[string] `json:"quantity,required"`
	// Buy, sell, sell-short indicator.
	Side param.Field[AccountOrderNewParamsSide] `json:"side,required"`
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
	StrategyType param.Field[AccountOrderNewParamsStrategyType] `json:"strategy_type,required"`
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
	TimeInForce param.Field[AccountOrderNewParamsTimeInForce] `json:"time_in_force,required"`
	// Name of the broker that provided you inventory for a short-sale. Required if
	// `side` is `sell-short`. If you procured inventory through us, you can use
	// `CLST`.
	LocateBroker param.Field[string] `json:"locate_broker"`
	// The price to execute at-or-better.
	Price param.Field[string] `json:"price"`
	// An ID that you provide.
	ReferenceID param.Field[string] `json:"reference_id"`
	// Denotes the format of the provided `symbol` field.
	SymbolFormat param.Field[AccountOrderNewParamsSymbolFormat] `json:"symbol_format"`
}

func (r AccountOrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of order, can be one of the following:
//
//   - `limit`: A limit order will execute at-or-better than the limit price you
//     specify
//   - `market`: An order that will execute at the prevailing market prices
type AccountOrderNewParamsOrderType string

const (
	AccountOrderNewParamsOrderTypeLimit  AccountOrderNewParamsOrderType = "limit"
	AccountOrderNewParamsOrderTypeMarket AccountOrderNewParamsOrderType = "market"
)

func (r AccountOrderNewParamsOrderType) IsKnown() bool {
	switch r {
	case AccountOrderNewParamsOrderTypeLimit, AccountOrderNewParamsOrderTypeMarket:
		return true
	}
	return false
}

// Buy, sell, sell-short indicator.
type AccountOrderNewParamsSide string

const (
	AccountOrderNewParamsSideBuy       AccountOrderNewParamsSide = "buy"
	AccountOrderNewParamsSideSell      AccountOrderNewParamsSide = "sell"
	AccountOrderNewParamsSideSellShort AccountOrderNewParamsSide = "sell-short"
)

func (r AccountOrderNewParamsSide) IsKnown() bool {
	switch r {
	case AccountOrderNewParamsSideBuy, AccountOrderNewParamsSideSell, AccountOrderNewParamsSideSellShort:
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
type AccountOrderNewParamsStrategyType string

const (
	AccountOrderNewParamsStrategyTypeSor  AccountOrderNewParamsStrategyType = "sor"
	AccountOrderNewParamsStrategyTypeDark AccountOrderNewParamsStrategyType = "dark"
	AccountOrderNewParamsStrategyTypeAp   AccountOrderNewParamsStrategyType = "ap"
	AccountOrderNewParamsStrategyTypePov  AccountOrderNewParamsStrategyType = "pov"
	AccountOrderNewParamsStrategyTypeTwap AccountOrderNewParamsStrategyType = "twap"
	AccountOrderNewParamsStrategyTypeVwap AccountOrderNewParamsStrategyType = "vwap"
)

func (r AccountOrderNewParamsStrategyType) IsKnown() bool {
	switch r {
	case AccountOrderNewParamsStrategyTypeSor, AccountOrderNewParamsStrategyTypeDark, AccountOrderNewParamsStrategyTypeAp, AccountOrderNewParamsStrategyTypePov, AccountOrderNewParamsStrategyTypeTwap, AccountOrderNewParamsStrategyTypeVwap:
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
type AccountOrderNewParamsTimeInForce string

const (
	AccountOrderNewParamsTimeInForceDay     AccountOrderNewParamsTimeInForce = "day"
	AccountOrderNewParamsTimeInForceIoc     AccountOrderNewParamsTimeInForce = "ioc"
	AccountOrderNewParamsTimeInForceDayPlus AccountOrderNewParamsTimeInForce = "day-plus"
	AccountOrderNewParamsTimeInForceAtOpen  AccountOrderNewParamsTimeInForce = "at-open"
	AccountOrderNewParamsTimeInForceAtClose AccountOrderNewParamsTimeInForce = "at-close"
)

func (r AccountOrderNewParamsTimeInForce) IsKnown() bool {
	switch r {
	case AccountOrderNewParamsTimeInForceDay, AccountOrderNewParamsTimeInForceIoc, AccountOrderNewParamsTimeInForceDayPlus, AccountOrderNewParamsTimeInForceAtOpen, AccountOrderNewParamsTimeInForceAtClose:
		return true
	}
	return false
}

// Denotes the format of the provided `symbol` field.
type AccountOrderNewParamsSymbolFormat string

const (
	AccountOrderNewParamsSymbolFormatCms AccountOrderNewParamsSymbolFormat = "cms"
	AccountOrderNewParamsSymbolFormatOsi AccountOrderNewParamsSymbolFormat = "osi"
)

func (r AccountOrderNewParamsSymbolFormat) IsKnown() bool {
	switch r {
	case AccountOrderNewParamsSymbolFormatCms, AccountOrderNewParamsSymbolFormatOsi:
		return true
	}
	return false
}

type AccountOrderListParams struct {
	// Milliseconds since epoch timestamp. This will constrain the search for orders
	// created after this timestamp, inclusively. Timestamps for orders prior the
	// current trading day will be ignored.
	From param.Field[int64] `query:"from"`
	// Number of orders to return per page.
	PageSize param.Field[int64] `query:"page_size"`
	// Cursor for the page to return.
	PageToken param.Field[string] `query:"page_token"`
	// Milliseconds since epoch timestamp. This will constrain the search for orders
	// created before this timestamp, inclusively. Timestamps for orders beyond the
	// current trading day will be ignored.
	To param.Field[int64] `query:"to"`
}

// URLQuery serializes [AccountOrderListParams]'s query parameters as `url.Values`.
func (r AccountOrderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountOrderDeleteParams struct {
	// Cancel orders only for this specific symbol. If this is omitted, all open orders
	// will be cancelled.
	Symbol param.Field[string] `query:"symbol"`
	// Format of the provided symbol.
	SymbolFormat param.Field[AccountOrderDeleteParamsSymbolFormat] `query:"symbol_format"`
}

// URLQuery serializes [AccountOrderDeleteParams]'s query parameters as
// `url.Values`.
func (r AccountOrderDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format of the provided symbol.
type AccountOrderDeleteParamsSymbolFormat string

const (
	AccountOrderDeleteParamsSymbolFormatCms AccountOrderDeleteParamsSymbolFormat = "cms"
	AccountOrderDeleteParamsSymbolFormatOsi AccountOrderDeleteParamsSymbolFormat = "osi"
)

func (r AccountOrderDeleteParamsSymbolFormat) IsKnown() bool {
	switch r {
	case AccountOrderDeleteParamsSymbolFormatCms, AccountOrderDeleteParamsSymbolFormatOsi:
		return true
	}
	return false
}
