package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ReceiptItem An item on a receipt card
// swagger:model ReceiptItem
type ReceiptItem struct {

	// Image
	Image *CardImage `json:"image,omitempty"`

	// Amount with currency
	Price string `json:"price,omitempty"`

	// Number of items of given kind
	Quantity string `json:"quantity,omitempty"`

	// Subtitle appears just below Title field, differs from Title in font styling only
	Subtitle string `json:"subtitle,omitempty"`

	// This action will be activated when user taps on the Item bubble.
	Tap *CardAction `json:"tap,omitempty"`

	// Text field appears just below subtitle, differs from Subtitle in font styling only
	Text string `json:"text,omitempty"`

	// Title of the Card
	Title string `json:"title,omitempty"`
}

// Validate validates this receipt item
func (m *ReceiptItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTap(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReceiptItem) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {

		if err := m.Image.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ReceiptItem) validateTap(formats strfmt.Registry) error {

	if swag.IsZero(m.Tap) { // not required
		return nil
	}

	if m.Tap != nil {

		if err := m.Tap.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
