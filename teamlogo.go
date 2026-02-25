// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cjavdev/believe-go/internal/apiform"
	"github.com/cjavdev/believe-go/internal/apijson"
	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
	"github.com/cjavdev/believe-go/packages/respjson"
)

// TeamLogoService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamLogoService] method instead.
type TeamLogoService struct {
	Options []option.RequestOption
}

// NewTeamLogoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTeamLogoService(opts ...option.RequestOption) (r TeamLogoService) {
	r = TeamLogoService{}
	r.Options = opts
	return
}

// Delete a team's logo.
func (r *TeamLogoService) Delete(ctx context.Context, fileID string, body TeamLogoDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.TeamID == "" {
		err = errors.New("missing required team_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("teams/%s/logo/%s", url.PathEscape(body.TeamID), fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Download a team's logo by file ID.
func (r *TeamLogoService) Download(ctx context.Context, fileID string, query TeamLogoDownloadParams, opts ...option.RequestOption) (res *TeamLogoDownloadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TeamID == "" {
		err = errors.New("missing required team_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("teams/%s/logo/%s", url.PathEscape(query.TeamID), fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a logo image for a team. Accepts image files (jpg, png, gif, webp).
func (r *TeamLogoService) Upload(ctx context.Context, teamID string, body TeamLogoUploadParams, opts ...option.RequestOption) (res *FileUpload, err error) {
	opts = slices.Concat(r.Options, opts)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return
	}
	path := fmt.Sprintf("teams/%s/logo", url.PathEscape(teamID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response model for file uploads.
type FileUpload struct {
	ChecksumSha256 string    `json:"checksum_sha256" api:"required"`
	ContentType    string    `json:"content_type" api:"required"`
	FileID         string    `json:"file_id" api:"required" format:"uuid"`
	Filename       string    `json:"filename" api:"required"`
	SizeBytes      int64     `json:"size_bytes" api:"required"`
	UploadedAt     time.Time `json:"uploaded_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChecksumSha256 respjson.Field
		ContentType    respjson.Field
		FileID         respjson.Field
		Filename       respjson.Field
		SizeBytes      respjson.Field
		UploadedAt     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileUpload) RawJSON() string { return r.JSON.raw }
func (r *FileUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamLogoDownloadResponse = any

type TeamLogoDeleteParams struct {
	TeamID string `path:"team_id" api:"required" json:"-"`
	paramObj
}

type TeamLogoDownloadParams struct {
	TeamID string `path:"team_id" api:"required" json:"-"`
	paramObj
}

type TeamLogoUploadParams struct {
	// Logo image file
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r TeamLogoUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
