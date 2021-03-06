package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// AttachmentData Attachment data
// swagger:model AttachmentData
type AttachmentData struct {

	// Name of the attachment
	Name string `json:"name,omitempty"`

	// original content
	OriginalBase64 strfmt.Base64 `json:"originalBase64,omitempty"`

	// Thumbnail
	ThumbnailBase64 strfmt.Base64 `json:"thumbnailBase64,omitempty"`

	// content type of the attachmnet
	Type string `json:"type,omitempty"`
}

// Validate validates this attachment data
func (m *AttachmentData) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
