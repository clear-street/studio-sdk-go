// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/clear-street/studio-sdk-go"
	"github.com/clear-street/studio-sdk-go/internal/apijson"
)

type LocateOrder struct {
	// Account ID for the account.
	AccountID string `json:"account_id,required"`
	// Unique locate ID assigned by us.
	LocateOrderID string `json:"locate_order_id,required"`
	// Unique MPID assigned by us.
	Mpid string `json:"mpid,required"`
	// The timestamp indicating when the locate order was requested.
	RequestedAt int64 `json:"requested_at,required"`
	// String representation of quantity.
	RequestedQuantity string `json:"requested_quantity,required"`
	// The status of the locate order.
	Status LocateOrderStatus `json:"status,required"`
	Symbol string            `json:"symbol,required"`
	// The timestamp indicating when the locate order was last updated.
	UpdatedAt int64 `json:"updated_at,required"`
	// The rate charged if the instrument is held overnight.
	BorrowRate string `json:"borrow_rate"`
	// Comment from the desk.
	DeskComment string `json:"desk_comment"`
	// The timestamp indicating when the locate-order will expire.
	ExpiresAt int64 `json:"expires_at"`
	// The locate ID, available once the locate order has been offered
	LocateID string `json:"locate_id"`
	// The timestamp indicating when the locate-order was located.
	LocatedAt int64 `json:"located_at"`
	// The quantity that has been located.
	LocatedQuantity string `json:"located_quantity"`
	// The reference ID provided by you.
	ReferenceID string `json:"reference_id"`
	// The total cost of the locate.
	TotalCost string `json:"total_cost"`
	// Comment from the trader.
	TraderComment string          `json:"trader_comment"`
	JSON          locateOrderJSON `json:"-"`
}

// locateOrderJSON contains the JSON metadata for the struct [LocateOrder]
type locateOrderJSON struct {
	AccountID         apijson.Field
	LocateOrderID     apijson.Field
	Mpid              apijson.Field
	RequestedAt       apijson.Field
	RequestedQuantity apijson.Field
	Status            apijson.Field
	Symbol            apijson.Field
	UpdatedAt         apijson.Field
	BorrowRate        apijson.Field
	DeskComment       apijson.Field
	ExpiresAt         apijson.Field
	LocateID          apijson.Field
	LocatedAt         apijson.Field
	LocatedQuantity   apijson.Field
	ReferenceID       apijson.Field
	TotalCost         apijson.Field
	TraderComment     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LocateOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locateOrderJSON) RawJSON() string {
	return r.raw
}

// The status of the locate order.
type LocateOrderStatus string

const (
	LocateOrderStatusPending   LocateOrderStatus = "pending"
	LocateOrderStatusOffered   LocateOrderStatus = "offered"
	LocateOrderStatusFilled    LocateOrderStatus = "filled"
	LocateOrderStatusRejected  LocateOrderStatus = "rejected"
	LocateOrderStatusDeclined  LocateOrderStatus = "declined"
	LocateOrderStatusExpired   LocateOrderStatus = "expired"
	LocateOrderStatusCancelled LocateOrderStatus = "cancelled"
)

func (r LocateOrderStatus) IsKnown() bool {
	switch r {
	case LocateOrderStatusPending, LocateOrderStatusOffered, LocateOrderStatusFilled, LocateOrderStatusRejected, LocateOrderStatusDeclined, LocateOrderStatusExpired, LocateOrderStatusCancelled:
		return true
	}
	return false
}

type Order struct {
	// Account ID for the account.
	AccountID string `json:"account_id,required"`
	// When the order was created in milliseconds since epoch.
	CreatedAt int64 `json:"created_at,required"`
	// The quantity that has been filled.
	FilledQuantity string `json:"filled_quantity,required"`
	// An internally generated unique ID for this order.
	OrderID string `json:"order_id,required"`
	// The type of order, can be one of the following:
	//
	//   - `limit`: A limit order will execute at-or-better than the limit price you
	//     specify
	//   - `market`: An order that will execute at the prevailing market prices
	OrderType OrderOrderType `json:"order_type,required"`
	// The requested quantity on this order.
	Quantity string `json:"quantity,required"`
	// Buy, sell, sell-short indicator.
	Side OrderSide `json:"side,required"`
	// Simplified order state, which is inferred from `OrderStatus`. Makes it easier to
	// determine whether an order can be executed against.
	//
	//   - `open`: Order _can_ potentially be executed against.
	//   - `rejected`: Order _cannot_ be executed against because it was rejected. This
	//     is a terminal state.
	//   - `closed`: Order _cannot_ be executed against. This is a terminal state.
	State OrderState `json:"state,required"`
	// Granular order status using
	// [standard values come FIX tag 39](https://www.fixtrading.org/online-specification/order-state-changes).
	Status OrderStatus `json:"status,required"`
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
	StrategyType OrderStrategyType `json:"strategy_type,required"`
	Symbol       string            `json:"symbol,required"`
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
	TimeInForce OrderTimeInForce `json:"time_in_force,required"`
	// When the order was updated in milliseconds since epoch.
	UpdatedAt int64 `json:"updated_at,required"`
	// A monotonically increasing number indicating the version of this order. A higher
	// number indicates a more recent version of the order.
	Version int64 `json:"version,required"`
	// Calculated average price of all fills on this order.
	AveragePrice float64 `json:"average_price"`
	// The last reason why this order was updated
	OrderUpdateReason OrderOrderUpdateReason `json:"order_update_reason"`
	// The requsted limit price on this order.
	Price string `json:"price"`
	// The ID you provided when creating this order.
	ReferenceID string `json:"reference_id"`
	// Free form text typically contains reasons for a reject.
	Text string    `json:"text"`
	JSON orderJSON `json:"-"`
}

// orderJSON contains the JSON metadata for the struct [Order]
type orderJSON struct {
	AccountID         apijson.Field
	CreatedAt         apijson.Field
	FilledQuantity    apijson.Field
	OrderID           apijson.Field
	OrderType         apijson.Field
	Quantity          apijson.Field
	Side              apijson.Field
	State             apijson.Field
	Status            apijson.Field
	StrategyType      apijson.Field
	Symbol            apijson.Field
	TimeInForce       apijson.Field
	UpdatedAt         apijson.Field
	Version           apijson.Field
	AveragePrice      apijson.Field
	OrderUpdateReason apijson.Field
	Price             apijson.Field
	ReferenceID       apijson.Field
	Text              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Order) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderJSON) RawJSON() string {
	return r.raw
}

// The type of order, can be one of the following:
//
//   - `limit`: A limit order will execute at-or-better than the limit price you
//     specify
//   - `market`: An order that will execute at the prevailing market prices
type OrderOrderType string

const (
	OrderOrderTypeLimit  OrderOrderType = "limit"
	OrderOrderTypeMarket OrderOrderType = "market"
)

func (r OrderOrderType) IsKnown() bool {
	switch r {
	case OrderOrderTypeLimit, OrderOrderTypeMarket:
		return true
	}
	return false
}

// Buy, sell, sell-short indicator.
type OrderSide string

const (
	OrderSideBuy       OrderSide = "buy"
	OrderSideSell      OrderSide = "sell"
	OrderSideSellShort OrderSide = "sell-short"
)

func (r OrderSide) IsKnown() bool {
	switch r {
	case OrderSideBuy, OrderSideSell, OrderSideSellShort:
		return true
	}
	return false
}

// Simplified order state, which is inferred from `OrderStatus`. Makes it easier to
// determine whether an order can be executed against.
//
//   - `open`: Order _can_ potentially be executed against.
//   - `rejected`: Order _cannot_ be executed against because it was rejected. This
//     is a terminal state.
//   - `closed`: Order _cannot_ be executed against. This is a terminal state.
type OrderState string

const (
	OrderStateOpen     OrderState = "open"
	OrderStateRejected OrderState = "rejected"
	OrderStateClosed   OrderState = "closed"
)

func (r OrderState) IsKnown() bool {
	switch r {
	case OrderStateOpen, OrderStateRejected, OrderStateClosed:
		return true
	}
	return false
}

// Granular order status using
// [standard values come FIX tag 39](https://www.fixtrading.org/online-specification/order-state-changes).
type OrderStatus string

const (
	OrderStatusNew                OrderStatus = "new"
	OrderStatusPartiallyFilled    OrderStatus = "partially-filled"
	OrderStatusFilled             OrderStatus = "filled"
	OrderStatusCanceled           OrderStatus = "canceled"
	OrderStatusReplaced           OrderStatus = "replaced"
	OrderStatusPendingCancel      OrderStatus = "pending-cancel"
	OrderStatusStopped            OrderStatus = "stopped"
	OrderStatusRejected           OrderStatus = "rejected"
	OrderStatusSuspended          OrderStatus = "suspended"
	OrderStatusPendingNew         OrderStatus = "pending-new"
	OrderStatusCalculated         OrderStatus = "calculated"
	OrderStatusExpired            OrderStatus = "expired"
	OrderStatusAcceptedForBidding OrderStatus = "accepted-for-bidding"
	OrderStatusPendingReplace     OrderStatus = "pending-replace"
	OrderStatusDoneForDay         OrderStatus = "done-for-day"
)

func (r OrderStatus) IsKnown() bool {
	switch r {
	case OrderStatusNew, OrderStatusPartiallyFilled, OrderStatusFilled, OrderStatusCanceled, OrderStatusReplaced, OrderStatusPendingCancel, OrderStatusStopped, OrderStatusRejected, OrderStatusSuspended, OrderStatusPendingNew, OrderStatusCalculated, OrderStatusExpired, OrderStatusAcceptedForBidding, OrderStatusPendingReplace, OrderStatusDoneForDay:
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
type OrderStrategyType string

const (
	OrderStrategyTypeSor  OrderStrategyType = "sor"
	OrderStrategyTypeDark OrderStrategyType = "dark"
	OrderStrategyTypeAp   OrderStrategyType = "ap"
	OrderStrategyTypePov  OrderStrategyType = "pov"
	OrderStrategyTypeTwap OrderStrategyType = "twap"
	OrderStrategyTypeVwap OrderStrategyType = "vwap"
)

func (r OrderStrategyType) IsKnown() bool {
	switch r {
	case OrderStrategyTypeSor, OrderStrategyTypeDark, OrderStrategyTypeAp, OrderStrategyTypePov, OrderStrategyTypeTwap, OrderStrategyTypeVwap:
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
type OrderTimeInForce string

const (
	OrderTimeInForceDay     OrderTimeInForce = "day"
	OrderTimeInForceIoc     OrderTimeInForce = "ioc"
	OrderTimeInForceDayPlus OrderTimeInForce = "day-plus"
	OrderTimeInForceAtOpen  OrderTimeInForce = "at-open"
	OrderTimeInForceAtClose OrderTimeInForce = "at-close"
)

func (r OrderTimeInForce) IsKnown() bool {
	switch r {
	case OrderTimeInForceDay, OrderTimeInForceIoc, OrderTimeInForceDayPlus, OrderTimeInForceAtOpen, OrderTimeInForceAtClose:
		return true
	}
	return false
}

// The last reason why this order was updated
type OrderOrderUpdateReason string

const (
	OrderOrderUpdateReasonPlace           OrderOrderUpdateReason = "place"
	OrderOrderUpdateReasonModify          OrderOrderUpdateReason = "modify"
	OrderOrderUpdateReasonCancel          OrderOrderUpdateReason = "cancel"
	OrderOrderUpdateReasonExecutionReport OrderOrderUpdateReason = "execution-report"
	OrderOrderUpdateReasonCancelReject    OrderOrderUpdateReason = "cancel-reject"
)

func (r OrderOrderUpdateReason) IsKnown() bool {
	switch r {
	case OrderOrderUpdateReasonPlace, OrderOrderUpdateReasonModify, OrderOrderUpdateReasonCancel, OrderOrderUpdateReasonExecutionReport, OrderOrderUpdateReasonCancelReject:
		return true
	}
	return false
}

type PnlSummaryForAccount struct {
	// Account ID
	AccountID string                   `json:"account_id,required"`
	JSON      pnlSummaryForAccountJSON `json:"-"`
	studiosdk.PnlSummary
}

// pnlSummaryForAccountJSON contains the JSON metadata for the struct
// [PnlSummaryForAccount]
type pnlSummaryForAccountJSON struct {
	AccountID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PnlSummaryForAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pnlSummaryForAccountJSON) RawJSON() string {
	return r.raw
}

type Position struct {
	// Account ID for the account.
	AccountID string `json:"account_id"`
	// String representation of quantity.
	Quantity string       `json:"quantity"`
	Symbol   string       `json:"symbol"`
	JSON     positionJSON `json:"-"`
}

// positionJSON contains the JSON metadata for the struct [Position]
type positionJSON struct {
	AccountID   apijson.Field
	Quantity    apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Position) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r positionJSON) RawJSON() string {
	return r.raw
}

type RegtMarginSimulation struct {
	// The margin calculation after applying simulated trades.
	After studiosdk.RegtMargin `json:"after,required"`
	// The margin calculation before applying simulated trades.
	Before studiosdk.RegtMargin `json:"before,required"`
	// Timestamp of when this simulation was created.
	CreatedAt int64 `json:"created_at,required"`
	// Name of this simulation that you provided when creating it.
	Name string `json:"name,required"`
	// Unique ID for a simulation.
	SimulationID studiosdk.SimulationID   `json:"simulation_id,required" format:"uuid"`
	JSON         regtMarginSimulationJSON `json:"-"`
}

// regtMarginSimulationJSON contains the JSON metadata for the struct
// [RegtMarginSimulation]
type regtMarginSimulationJSON struct {
	After        apijson.Field
	Before       apijson.Field
	CreatedAt    apijson.Field
	Name         apijson.Field
	SimulationID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RegtMarginSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regtMarginSimulationJSON) RawJSON() string {
	return r.raw
}

type Trade struct {
	// When this trade happened in milliseconds since epoch.
	CreatedAt int64 `json:"created_at,required"`
	// The order ID of the order this trade occurred on.
	OrderID string `json:"order_id,required"`
	// The traded price.
	Price string `json:"price,required"`
	// The amount that was traded.
	Quantity string `json:"quantity,required"`
	// The side this trade occurred on.
	Side TradeSide `json:"side,required"`
	// Unique trade ID assigned by us.
	TradeID string `json:"trade_id,required"`
	// Account ID for the account.
	AccountID string `json:"account_id"`
	// The symbol this trade was for.
	Symbol string    `json:"symbol"`
	JSON   tradeJSON `json:"-"`
}

// tradeJSON contains the JSON metadata for the struct [Trade]
type tradeJSON struct {
	CreatedAt   apijson.Field
	OrderID     apijson.Field
	Price       apijson.Field
	Quantity    apijson.Field
	Side        apijson.Field
	TradeID     apijson.Field
	AccountID   apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Trade) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tradeJSON) RawJSON() string {
	return r.raw
}

// The side this trade occurred on.
type TradeSide string

const (
	TradeSideBuy       TradeSide = "buy"
	TradeSideSell      TradeSide = "sell"
	TradeSideSellShort TradeSide = "sell-short"
)

func (r TradeSide) IsKnown() bool {
	switch r {
	case TradeSideBuy, TradeSideSell, TradeSideSellShort:
		return true
	}
	return false
}
