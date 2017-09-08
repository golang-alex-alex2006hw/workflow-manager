package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Workflow workflow
// swagger:model Workflow
type Workflow struct {

	// created at
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last updated
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// tasks
	Tasks []*Job `json:"tasks"`

	// workflow definition
	WorkflowDefinition *WorkflowDefinition `json:"workflowDefinition,omitempty"`
}

// Validate validates this workflow
func (m *Workflow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTasks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkflowDefinition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workflow) validateTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {

		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {

			if err := m.Tasks[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Workflow) validateWorkflowDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowDefinition) { // not required
		return nil
	}

	if m.WorkflowDefinition != nil {

		if err := m.WorkflowDefinition.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
