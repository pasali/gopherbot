package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// CardAction An action on a card
// swagger:model CardAction
type CardAction struct {

	// URL Picture which will appear on the button, next to text label.
	Image string `json:"image,omitempty"`

	// Text description which appear on the button.
	Title string `json:"title,omitempty"`

	// Defines the type of action implemented by this button.
	Type string `json:"type,omitempty"`

	// Supplementary parameter for action. Content of this property depends on the ActionType
	Value string `json:"value,omitempty"`
}

// Validate validates this card action
func (m *CardAction) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
