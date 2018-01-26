package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ResolvedByUserWrapper resolved by user wrapper
// swagger:model ResolvedByUserWrapper
type ResolvedByUserWrapper struct {

	// is set
	IsSet bool `json:"isSet,omitempty"`

	// value
	Value bool `json:"value,omitempty"`
}

// Validate validates this resolved by user wrapper
func (m *ResolvedByUserWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}