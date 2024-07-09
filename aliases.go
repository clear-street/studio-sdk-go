// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"github.com/clear-street/studio-sdk-go/internal/apierror"
	"github.com/clear-street/studio-sdk-go/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type LocateOrder = shared.LocateOrder

// The status of the locate order.
//
// This is an alias to an internal type.
type LocateOrderStatus = shared.LocateOrderStatus

// This is an alias to an internal value.
const LocateOrderStatusPending = shared.LocateOrderStatusPending

// This is an alias to an internal value.
const LocateOrderStatusOffered = shared.LocateOrderStatusOffered

// This is an alias to an internal value.
const LocateOrderStatusFilled = shared.LocateOrderStatusFilled

// This is an alias to an internal value.
const LocateOrderStatusRejected = shared.LocateOrderStatusRejected

// This is an alias to an internal value.
const LocateOrderStatusDeclined = shared.LocateOrderStatusDeclined

// This is an alias to an internal value.
const LocateOrderStatusExpired = shared.LocateOrderStatusExpired

// This is an alias to an internal value.
const LocateOrderStatusCancelled = shared.LocateOrderStatusCancelled

// This is an alias to an internal type.
type Order = shared.Order

// The type of order, can be one of the following:
//
//   - `limit`: A limit order will execute at-or-better than the limit price you
//     specify
//   - `market`: An order that will execute at the prevailing market prices
//
// This is an alias to an internal type.
type OrderOrderType = shared.OrderOrderType

// This is an alias to an internal value.
const OrderOrderTypeLimit = shared.OrderOrderTypeLimit

// This is an alias to an internal value.
const OrderOrderTypeMarket = shared.OrderOrderTypeMarket

// Buy, sell, sell-short indicator.
//
// This is an alias to an internal type.
type OrderSide = shared.OrderSide

// This is an alias to an internal value.
const OrderSideBuy = shared.OrderSideBuy

// This is an alias to an internal value.
const OrderSideSell = shared.OrderSideSell

// This is an alias to an internal value.
const OrderSideSellShort = shared.OrderSideSellShort

// Simplified order state, which is inferred from `OrderStatus`. Makes it easier to
// determine whether an order can be executed against.
//
//   - `open`: Order _can_ potentially be executed against.
//   - `rejected`: Order _cannot_ be executed against because it was rejected. This
//     is a terminal state.
//   - `closed`: Order _cannot_ be executed against. This is a terminal state.
//
// This is an alias to an internal type.
type OrderState = shared.OrderState

// This is an alias to an internal value.
const OrderStateOpen = shared.OrderStateOpen

// This is an alias to an internal value.
const OrderStateRejected = shared.OrderStateRejected

// This is an alias to an internal value.
const OrderStateClosed = shared.OrderStateClosed

// Granular order status using
// [standard values come FIX tag 39](https://www.fixtrading.org/online-specification/order-state-changes).
//
// This is an alias to an internal type.
type OrderStatus = shared.OrderStatus

// This is an alias to an internal value.
const OrderStatusNew = shared.OrderStatusNew

// This is an alias to an internal value.
const OrderStatusPartiallyFilled = shared.OrderStatusPartiallyFilled

// This is an alias to an internal value.
const OrderStatusFilled = shared.OrderStatusFilled

// This is an alias to an internal value.
const OrderStatusCanceled = shared.OrderStatusCanceled

// This is an alias to an internal value.
const OrderStatusReplaced = shared.OrderStatusReplaced

// This is an alias to an internal value.
const OrderStatusPendingCancel = shared.OrderStatusPendingCancel

// This is an alias to an internal value.
const OrderStatusStopped = shared.OrderStatusStopped

// This is an alias to an internal value.
const OrderStatusRejected = shared.OrderStatusRejected

// This is an alias to an internal value.
const OrderStatusSuspended = shared.OrderStatusSuspended

// This is an alias to an internal value.
const OrderStatusPendingNew = shared.OrderStatusPendingNew

// This is an alias to an internal value.
const OrderStatusCalculated = shared.OrderStatusCalculated

// This is an alias to an internal value.
const OrderStatusExpired = shared.OrderStatusExpired

// This is an alias to an internal value.
const OrderStatusAcceptedForBidding = shared.OrderStatusAcceptedForBidding

// This is an alias to an internal value.
const OrderStatusPendingReplace = shared.OrderStatusPendingReplace

// This is an alias to an internal value.
const OrderStatusDoneForDay = shared.OrderStatusDoneForDay

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
//
// This is an alias to an internal type.
type OrderStrategyType = shared.OrderStrategyType

// This is an alias to an internal value.
const OrderStrategyTypeSor = shared.OrderStrategyTypeSor

// This is an alias to an internal value.
const OrderStrategyTypeDark = shared.OrderStrategyTypeDark

// This is an alias to an internal value.
const OrderStrategyTypeAp = shared.OrderStrategyTypeAp

// This is an alias to an internal value.
const OrderStrategyTypePov = shared.OrderStrategyTypePov

// This is an alias to an internal value.
const OrderStrategyTypeTwap = shared.OrderStrategyTypeTwap

// This is an alias to an internal value.
const OrderStrategyTypeVwap = shared.OrderStrategyTypeVwap

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
//
// This is an alias to an internal type.
type OrderTimeInForce = shared.OrderTimeInForce

// This is an alias to an internal value.
const OrderTimeInForceDay = shared.OrderTimeInForceDay

// This is an alias to an internal value.
const OrderTimeInForceIoc = shared.OrderTimeInForceIoc

// This is an alias to an internal value.
const OrderTimeInForceDayPlus = shared.OrderTimeInForceDayPlus

// This is an alias to an internal value.
const OrderTimeInForceAtOpen = shared.OrderTimeInForceAtOpen

// This is an alias to an internal value.
const OrderTimeInForceAtClose = shared.OrderTimeInForceAtClose

// The last reason why this order was updated
//
// This is an alias to an internal type.
type OrderOrderUpdateReason = shared.OrderOrderUpdateReason

// This is an alias to an internal value.
const OrderOrderUpdateReasonPlace = shared.OrderOrderUpdateReasonPlace

// This is an alias to an internal value.
const OrderOrderUpdateReasonModify = shared.OrderOrderUpdateReasonModify

// This is an alias to an internal value.
const OrderOrderUpdateReasonCancel = shared.OrderOrderUpdateReasonCancel

// This is an alias to an internal value.
const OrderOrderUpdateReasonExecutionReport = shared.OrderOrderUpdateReasonExecutionReport

// This is an alias to an internal value.
const OrderOrderUpdateReasonCancelReject = shared.OrderOrderUpdateReasonCancelReject

// This is an alias to an internal type.
type PnlSummaryForAccount = shared.PnlSummaryForAccount

// This is an alias to an internal type.
type Position = shared.Position

// This is an alias to an internal type.
type RegtMarginSimulation = shared.RegtMarginSimulation

// This is an alias to an internal type.
type Trade = shared.Trade

// The side this trade occurred on.
//
// This is an alias to an internal type.
type TradeSide = shared.TradeSide

// This is an alias to an internal value.
const TradeSideBuy = shared.TradeSideBuy

// This is an alias to an internal value.
const TradeSideSell = shared.TradeSideSell

// This is an alias to an internal value.
const TradeSideSellShort = shared.TradeSideSellShort
