// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/clear-street/studio-sdk-go/internal/requestconfig"
	"github.com/clear-street/studio-sdk-go/option"
	"github.com/clear-street/studio-sdk-go/shared"
)

// AccountPnlSummaryService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPnlSummaryService] method instead.
type AccountPnlSummaryService struct {
	Options []option.RequestOption
}

// NewAccountPnlSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPnlSummaryService(opts ...option.RequestOption) (r *AccountPnlSummaryService) {
	r = &AccountPnlSummaryService{}
	r.Options = opts
	return
}

// Get PNL summary for a given account.
func (r *AccountPnlSummaryService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *shared.PnlSummaryForAccount, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pnl-summary", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
