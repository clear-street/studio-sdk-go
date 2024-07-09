// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/studio-sdk-go/internal/apijson"
	"github.com/stainless-sdks/studio-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/studio-sdk-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options      []option.RequestOption
	BulkOrders   *AccountBulkOrderService
	Orders       *AccountOrderService
	Trades       *AccountTradeService
	Positions    *AccountPositionService
	LocateOrders *AccountLocateOrderService
	EasyBorrows  *AccountEasyBorrowService
	PnlSummary   *AccountPnlSummaryService
	PnlDetails   *AccountPnlDetailService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.BulkOrders = NewAccountBulkOrderService(opts...)
	r.Orders = NewAccountOrderService(opts...)
	r.Trades = NewAccountTradeService(opts...)
	r.Positions = NewAccountPositionService(opts...)
	r.LocateOrders = NewAccountLocateOrderService(opts...)
	r.EasyBorrows = NewAccountEasyBorrowService(opts...)
	r.PnlSummary = NewAccountPnlSummaryService(opts...)
	r.PnlDetails = NewAccountPnlDetailService(opts...)
	return
}

// Get an account by its ID.
func (r *AccountService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all available accounts.
func (r *AccountService) List(ctx context.Context, opts ...option.RequestOption) (res *AccountListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Account struct {
	// Account ID for the account.
	AccountID string `json:"account_id,required"`
	// Entity ID for the legal entity.
	EntityID string      `json:"entity_id,required"`
	Name     string      `json:"name,required"`
	JSON     accountJSON `json:"-"`
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	AccountID   apijson.Field
	EntityID    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountJSON) RawJSON() string {
	return r.raw
}

type AccountListResponse struct {
	Data []Account               `json:"data"`
	JSON accountListResponseJSON `json:"-"`
}

// accountListResponseJSON contains the JSON metadata for the struct
// [AccountListResponse]
type accountListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseJSON) RawJSON() string {
	return r.raw
}
