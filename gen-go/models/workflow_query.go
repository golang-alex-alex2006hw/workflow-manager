package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// WorkflowQuery workflow query
// swagger:model WorkflowQuery
type WorkflowQuery struct {

	// limit
	// Maximum: 10000
	Limit int64 `json:"limit,omitempty"`

	// oldest first
	OldestFirst bool `json:"oldestFirst,omitempty"`

	// page token
	PageToken string `json:"pageToken,omitempty"`

	// status
	Status WorkflowStatus `json:"status,omitempty"`

	// summary only
	SummaryOnly *bool `json:"summaryOnly,omitempty"`

	// workflow definition name
	// Required: true
	WorkflowDefinitionName *string `json:"workflowDefinitionName"`
}

// Validate validates this workflow query
func (m *WorkflowQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkflowDefinitionName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowQuery) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MaximumInt("limit", "body", int64(m.Limit), 10000, false); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowQuery) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowQuery) validateWorkflowDefinitionName(formats strfmt.Registry) error {

	if err := validate.Required("workflowDefinitionName", "body", m.WorkflowDefinitionName); err != nil {
		return err
	}

	return nil
}