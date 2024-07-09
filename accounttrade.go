// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/sachnk/studio-sdk-go/internal/apijson"
	"github.com/sachnk/studio-sdk-go/internal/apiquery"
	"github.com/sachnk/studio-sdk-go/internal/param"
	"github.com/sachnk/studio-sdk-go/internal/requestconfig"
	"github.com/sachnk/studio-sdk-go/option"
	"github.com/sachnk/studio-sdk-go/shared"
)

// AccountTradeService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTradeService] method instead.
type AccountTradeService struct {
	Options []option.RequestOption
}

// NewAccountTradeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountTradeService(opts ...option.RequestOption) (r *AccountTradeService) {
	r = &AccountTradeService{}
	r.Options = opts
	return
}

// Get trade a trade by its unique trade ID.
func (r *AccountTradeService) Get(ctx context.Context, accountID string, tradeID string, opts ...option.RequestOption) (res *shared.Trade, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tradeID == "" {
		err = errors.New("missing required trade_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/trades/%s", accountID, tradeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List trades for a given account for the current trading day.
func (r *AccountTradeService) List(ctx context.Context, accountID string, query AccountTradeListParams, opts ...option.RequestOption) (res *AccountTradeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/trades", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountTradeListResponse struct {
	Data []shared.Trade `json:"data,required"`
	// Cursor for the next page of results.
	NextPageToken string                       `json:"next_page_token"`
	JSON          accountTradeListResponseJSON `json:"-"`
}

// accountTradeListResponseJSON contains the JSON metadata for the struct
// [AccountTradeListResponse]
type accountTradeListResponseJSON struct {
	Data          apijson.Field
	NextPageToken apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountTradeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTradeListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountTradeListParams struct {
	// Number of trades to return per page.
	PageSize param.Field[int64] `query:"page_size"`
	// Cursor for the page to return.
	PageToken param.Field[string] `query:"page_token"`
}

// URLQuery serializes [AccountTradeListParams]'s query parameters as `url.Values`.
func (r AccountTradeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
