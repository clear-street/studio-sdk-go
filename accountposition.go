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

// AccountPositionService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPositionService] method instead.
type AccountPositionService struct {
	Options []option.RequestOption
}

// NewAccountPositionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountPositionService(opts ...option.RequestOption) (r *AccountPositionService) {
	r = &AccountPositionService{}
	r.Options = opts
	return
}

// Get current positions for a given account for a given symbol.
func (r *AccountPositionService) Get(ctx context.Context, accountID string, symbol string, opts ...option.RequestOption) (res *shared.Position, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if symbol == "" {
		err = errors.New("missing required symbol parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/positions/%s", accountID, symbol)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List current positions for a given account.
func (r *AccountPositionService) List(ctx context.Context, accountID string, query AccountPositionListParams, opts ...option.RequestOption) (res *AccountPositionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/positions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountPositionListResponse struct {
	Data []shared.Position `json:"data,required"`
	// Cursor for the next page of results.
	NextPageToken string                          `json:"next_page_token"`
	JSON          accountPositionListResponseJSON `json:"-"`
}

// accountPositionListResponseJSON contains the JSON metadata for the struct
// [AccountPositionListResponse]
type accountPositionListResponseJSON struct {
	Data          apijson.Field
	NextPageToken apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountPositionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPositionListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPositionListParams struct {
	// Number of positions to return per page.
	PageSize param.Field[int64] `query:"page_size"`
	// Cursor for the page to return.
	PageToken param.Field[string] `query:"page_token"`
}

// URLQuery serializes [AccountPositionListParams]'s query parameters as
// `url.Values`.
func (r AccountPositionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
