// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// JobAttempt job attempt
// swagger:model JobAttempt
type JobAttempt struct {

	// container instance a r n
	ContainerInstanceARN string `json:"containerInstanceARN,omitempty"`

	// created at
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// exit code
	ExitCode int64 `json:"exitCode,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// started at
	StartedAt strfmt.DateTime `json:"startedAt,omitempty"`

	// stopped at
	StoppedAt strfmt.DateTime `json:"stoppedAt,omitempty"`

	// task a r n
	TaskARN string `json:"taskARN,omitempty"`
}

// Validate validates this job attempt
func (m *JobAttempt) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *JobAttempt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobAttempt) UnmarshalBinary(b []byte) error {
	var res JobAttempt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
