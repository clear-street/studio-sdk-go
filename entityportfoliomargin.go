// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/clear-street/studio-sdk-go/internal/requestconfig"
	"github.com/clear-street/studio-sdk-go/option"
)

// EntityPortfolioMarginService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityPortfolioMarginService] method instead.
type EntityPortfolioMarginService struct {
	Options []option.RequestOption
}

// NewEntityPortfolioMarginService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEntityPortfolioMarginService(opts ...option.RequestOption) (r *EntityPortfolioMarginService) {
	r = &EntityPortfolioMarginService{}
	r.Options = opts
	return
}

// Get latest portfolio margin calculation for the given entity
func (r *EntityPortfolioMarginService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *PortfolioMargin, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/portfolio-margin", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
