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

// AccountEasyBorrowService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEasyBorrowService] method instead.
type AccountEasyBorrowService struct {
	Options []option.RequestOption
}

// NewAccountEasyBorrowService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountEasyBorrowService(opts ...option.RequestOption) (r *AccountEasyBorrowService) {
	r = &AccountEasyBorrowService{}
	r.Options = opts
	return
}

// List all current easy-to-borrow stock symbols. This list changes dynamically
// daily.
func (r *AccountEasyBorrowService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountEasyBorrowListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/easy-borrows", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountEasyBorrowListResponse struct {
	Data []string                          `json:"data,required"`
	JSON accountEasyBorrowListResponseJSON `json:"-"`
}

// accountEasyBorrowListResponseJSON contains the JSON metadata for the struct
// [AccountEasyBorrowListResponse]
type accountEasyBorrowListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEasyBorrowListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEasyBorrowListResponseJSON) RawJSON() string {
	return r.raw
}
