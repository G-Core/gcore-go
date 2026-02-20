// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// PlaylistVideoService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlaylistVideoService] method instead.
type PlaylistVideoService struct {
	Options []option.RequestOption
}

// NewPlaylistVideoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlaylistVideoService(opts ...option.RequestOption) (r PlaylistVideoService) {
	r = PlaylistVideoService{}
	r.Options = opts
	return
}

// Shows ordered array of playlist videos
func (r *PlaylistVideoService) List(ctx context.Context, playlistID int64, opts ...option.RequestOption) (res *[]PlaylistVideo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/playlists/%v/videos", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
