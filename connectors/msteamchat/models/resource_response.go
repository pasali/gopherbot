package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ResourceResponse A response containing a resource ID
// swagger:model ResourceResponse
type ResourceResponse struct {

	// Id of the resource
	ID string `json:"id,omitempty"`
}

// Validate validates this resource response
func (m *ResourceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
