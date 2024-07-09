// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/studio-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/studio-sdk-go/option"
)

// EntityRegtMarginService contains methods and other services that help with
// interacting with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityRegtMarginService] method instead.
type EntityRegtMarginService struct {
	Options []option.RequestOption
}

// NewEntityRegtMarginService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEntityRegtMarginService(opts ...option.RequestOption) (r *EntityRegtMarginService) {
	r = &EntityRegtMarginService{}
	r.Options = opts
	return
}

// Get the latest Reg-T margin calculation for the given entity
func (r *EntityRegtMarginService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *RegtMargin, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/regt-margin", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
