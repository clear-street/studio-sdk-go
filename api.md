# Shared Response Types

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#LocateOrder">LocateOrder</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#PnlSummaryForAccount">PnlSummaryForAccount</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#Position">Position</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#RegtMarginSimulation">RegtMarginSimulation</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#Trade">Trade</a>

# Entities

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#Entity">Entity</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#PnlSummary">PnlSummary</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#PortfolioMargin">PortfolioMargin</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#RegtMargin">RegtMargin</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityListResponse">EntityListResponse</a>

Methods:

- <code title="get /entities/{entity_id}">client.Entities.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities">client.Entities.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityListResponse">EntityListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PnlSummaries

Methods:

- <code title="get /entities/{entity_id}/pnl-summary">client.Entities.PnlSummaries.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityPnlSummaryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#PnlSummary">PnlSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RegtMargins

Methods:

- <code title="get /entities/{entity_id}/regt-margin">client.Entities.RegtMargins.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityRegtMarginService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#RegtMargin">RegtMargin</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PortfolioMargins

Methods:

- <code title="get /entities/{entity_id}/portfolio-margin">client.Entities.PortfolioMargins.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityPortfolioMarginService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#PortfolioMargin">PortfolioMargin</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RegtMarginSimulations

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#SimulationIDParam">SimulationIDParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#SimulationID">SimulationID</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityRegtMarginSimulationNewResponse">EntityRegtMarginSimulationNewResponse</a>

Methods:

- <code title="post /entities/{entity_id}/regt-margin-simulations">client.Entities.RegtMarginSimulations.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityRegtMarginSimulationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityRegtMarginSimulationNewParams">EntityRegtMarginSimulationNewParams</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityRegtMarginSimulationNewResponse">EntityRegtMarginSimulationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities/{entity_id}/regt-margin-simulations/{simulation_id}">client.Entities.RegtMarginSimulations.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#EntityRegtMarginSimulationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, simulationID <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#SimulationIDParam">SimulationIDParam</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#RegtMarginSimulation">RegtMarginSimulation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountListResponse">AccountListResponse</a>

Methods:

- <code title="get /accounts/{account_id}">client.Accounts.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BulkOrders

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountBulkOrderNewResponse">AccountBulkOrderNewResponse</a>

Methods:

- <code title="post /accounts/{account_id}/bulk-orders">client.Accounts.BulkOrders.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountBulkOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountBulkOrderNewParams">AccountBulkOrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountBulkOrderNewResponse">AccountBulkOrderNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orders

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderNewResponse">AccountOrderNewResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderGetResponse">AccountOrderGetResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderListResponse">AccountOrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderDeleteResponse">AccountOrderDeleteResponse</a>

Methods:

- <code title="post /accounts/{account_id}/orders">client.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderNewParams">AccountOrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderNewResponse">AccountOrderNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/orders/{order_id}">client.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderGetResponse">AccountOrderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/orders">client.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderListParams">AccountOrderListParams</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderListResponse">AccountOrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/orders">client.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderDeleteParams">AccountOrderDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderDeleteResponse">AccountOrderDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/orders/{order_id}">client.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountOrderService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Trades

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountTradeListResponse">AccountTradeListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/trades/{trade_id}">client.Accounts.Trades.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountTradeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, tradeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#Trade">Trade</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/trades">client.Accounts.Trades.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountTradeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountTradeListParams">AccountTradeListParams</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountTradeListResponse">AccountTradeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Positions

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountPositionListResponse">AccountPositionListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/positions/{symbol}">client.Accounts.Positions.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountPositionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, symbol <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#Position">Position</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/positions">client.Accounts.Positions.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountPositionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountPositionListParams">AccountPositionListParams</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountPositionListResponse">AccountPositionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LocateOrders

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountLocateOrderListResponse">AccountLocateOrderListResponse</a>

Methods:

- <code title="post /accounts/{account_id}/locate-orders">client.Accounts.LocateOrders.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountLocateOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountLocateOrderNewParams">AccountLocateOrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#LocateOrder">LocateOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/locate-orders/{locate_order_id}">client.Accounts.LocateOrders.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountLocateOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, locateOrderID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#LocateOrder">LocateOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}/locate-orders/{locate_order_id}">client.Accounts.LocateOrders.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountLocateOrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, locateOrderID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountLocateOrderUpdateParams">AccountLocateOrderUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /accounts/{account_id}/locate-orders">client.Accounts.LocateOrders.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountLocateOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountLocateOrderListResponse">AccountLocateOrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EasyBorrows

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountEasyBorrowListResponse">AccountEasyBorrowListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/easy-borrows">client.Accounts.EasyBorrows.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountEasyBorrowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountEasyBorrowListResponse">AccountEasyBorrowListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PnlSummary

Methods:

- <code title="get /accounts/{account_id}/pnl-summary">client.Accounts.PnlSummary.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountPnlSummaryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go/shared#PnlSummaryForAccount">PnlSummaryForAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PnlDetails

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountPnlDetailListResponse">AccountPnlDetailListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/pnl-details">client.Accounts.PnlDetails.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountPnlDetailService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#AccountPnlDetailListResponse">AccountPnlDetailListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Instruments

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#Instrument">Instrument</a>

Methods:

- <code title="get /instruments/{symbol}">client.Instruments.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#InstrumentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, symbol <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#InstrumentGetParams">InstrumentGetParams</a>) (<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go">studiosdk</a>.<a href="https://pkg.go.dev/github.com/clear-street/studio-sdk-go#Instrument">Instrument</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
