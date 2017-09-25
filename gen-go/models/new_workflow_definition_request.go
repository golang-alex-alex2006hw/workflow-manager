package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// NewWorkflowDefinitionRequest new workflow definition request
// swagger:model NewWorkflowDefinitionRequest
type NewWorkflowDefinitionRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// manager
	Manager Manager `json:"manager,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// start at
	StartAt string `json:"startAt,omitempty"`

	// states
	States []*State `json:"states"`
}

// Validate validates this new workflow definition request
func (m *NewWorkflowDefinitionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManager(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStates(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewWorkflowDefinitionRequest) validateManager(formats strfmt.Registry) error {

	if swag.IsZero(m.Manager) { // not required
		return nil
	}

	if err := m.Manager.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *NewWorkflowDefinitionRequest) validateStates(formats strfmt.Registry) error {

	if swag.IsZero(m.States) { // not required
		return nil
	}

	for i := 0; i < len(m.States); i++ {

		if swag.IsZero(m.States[i]) { // not required
			continue
		}

		if m.States[i] != nil {

			if err := m.States[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
