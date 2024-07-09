// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sachnk/studio-sdk-go/internal/requestconfig"
	"github.com/sachnk/studio-sdk-go/option"
)

// EntityPnlSummaryService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityPnlSummaryService] method instead.
type EntityPnlSummaryService struct {
	Options []option.RequestOption
}

// NewEntityPnlSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEntityPnlSummaryService(opts ...option.RequestOption) (r *EntityPnlSummaryService) {
	r = &EntityPnlSummaryService{}
	r.Options = opts
	return
}

// Get PNL summary for all accounts in an entity.
func (r *EntityPnlSummaryService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *PnlSummary, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/pnl-summary", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
