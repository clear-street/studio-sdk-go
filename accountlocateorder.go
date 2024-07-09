// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/clear-street/studio-sdk-go/internal/apijson"
	"github.com/clear-street/studio-sdk-go/internal/param"
	"github.com/clear-street/studio-sdk-go/internal/requestconfig"
	"github.com/clear-street/studio-sdk-go/option"
	"github.com/clear-street/studio-sdk-go/shared"
)

// AccountLocateOrderService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLocateOrderService] method instead.
type AccountLocateOrderService struct {
	Options []option.RequestOption
}

// NewAccountLocateOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLocateOrderService(opts ...option.RequestOption) (r *AccountLocateOrderService) {
	r = &AccountLocateOrderService{}
	r.Options = opts
	return
}

// Create locate order to borrow inventory for short-selling.
func (r *AccountLocateOrderService) New(ctx context.Context, accountID string, body AccountLocateOrderNewParams, opts ...option.RequestOption) (res *shared.LocateOrder, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/locate-orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get locate order by its unique locate order ID.
func (r *AccountLocateOrderService) Get(ctx context.Context, accountID string, locateOrderID string, opts ...option.RequestOption) (res *shared.LocateOrder, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if locateOrderID == "" {
		err = errors.New("missing required locate_order_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/locate-orders/%s", accountID, locateOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Accept or decline locate order that has been offered.
func (r *AccountLocateOrderService) Update(ctx context.Context, accountID string, locateOrderID string, body AccountLocateOrderUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if locateOrderID == "" {
		err = errors.New("missing required locate_order_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/locate-orders/%s", accountID, locateOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// List all locate orders
func (r *AccountLocateOrderService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountLocateOrderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/locate-orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountLocateOrderListResponse struct {
	Data []shared.LocateOrder               `json:"data,required"`
	JSON accountLocateOrderListResponseJSON `json:"-"`
}

// accountLocateOrderListResponseJSON contains the JSON metadata for the struct
// [AccountLocateOrderListResponse]
type accountLocateOrderListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLocateOrderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLocateOrderListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountLocateOrderNewParams struct {
	// The market participant where the locate will be sent.
	Mpid param.Field[string] `json:"mpid,required"`
	// String representation of quantity.
	Quantity param.Field[string] `json:"quantity,required"`
	// Your unique ID for this locate order.
	ReferenceID param.Field[string] `json:"reference_id,required"`
	Symbol      param.Field[string] `json:"symbol,required"`
	// Any additional comments for the locate request.
	Comments param.Field[string] `json:"comments"`
}

func (r AccountLocateOrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountLocateOrderUpdateParams struct {
	// Accept or decline the locate order.
	Accept param.Field[bool] `json:"accept,required"`
}

func (r AccountLocateOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
