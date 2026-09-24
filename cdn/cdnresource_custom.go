// Custom code. This file is not generated and is preserved across codegen runs.
// It isolates the hand-written DeactivateAndDelete convenience method from
// generated code to eliminate merge conflicts.

package cdn

import (
	"context"

	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// DeactivateAndDelete is a utility method that first deactivates the CDN resource
// by setting the `active` attribute to `false`, then deletes the resource from
// the system permanently.
//
// This method is useful because the Delete operation requires the CDN resource to
// be deactivated first. DeactivateAndDelete handles both steps in a single call.
func (r *CDNResourceService) DeactivateAndDelete(ctx context.Context, resourceID int64, opts ...option.RequestOption) (err error) {
	params := CDNResourceUpdateParams{
		Active: param.NewOpt(false),
	}
	_, err = r.Update(ctx, resourceID, params, opts...)
	if err != nil {
		return err
	}

	return r.Delete(ctx, resourceID, opts...)
}
