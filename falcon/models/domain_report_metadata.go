// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainReportMetadata domain report metadata
//
// swagger:model domain.ReportMetadata
type DomainReportMetadata struct {

	// created by user id
	// Required: true
	CreatedByUserID *string `json:"created_by_user_id"`

	// created by uuid
	// Required: true
	CreatedByUUID *string `json:"created_by_uuid"`

	// discover params
	DiscoverParams *DomainDiscoverParams `json:"discover_params,omitempty"`

	// last scheduled execution
	LastScheduledExecution *DomainLastScheduledExecution `json:"last_scheduled_execution,omitempty"`

	// last unscheduled execution
	// Required: true
	LastUnscheduledExecution *DomainLastUnscheduledExecution `json:"last_unscheduled_execution"`

	// subtype
	// Required: true
	Subtype *string `json:"subtype"`

	// xdr params
	XdrParams *DomainXDRParams `json:"xdr_params,omitempty"`
}

// Validate validates this domain report metadata
func (m *DomainReportMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedByUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedByUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscoverParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastScheduledExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUnscheduledExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubtype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXdrParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainReportMetadata) validateCreatedByUserID(formats strfmt.Registry) error {

	if err := validate.Required("created_by_user_id", "body", m.CreatedByUserID); err != nil {
		return err
	}

	return nil
}

func (m *DomainReportMetadata) validateCreatedByUUID(formats strfmt.Registry) error {

	if err := validate.Required("created_by_uuid", "body", m.CreatedByUUID); err != nil {
		return err
	}

	return nil
}

func (m *DomainReportMetadata) validateDiscoverParams(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscoverParams) { // not required
		return nil
	}

	if m.DiscoverParams != nil {
		if err := m.DiscoverParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discover_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discover_params")
			}
			return err
		}
	}

	return nil
}

func (m *DomainReportMetadata) validateLastScheduledExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.LastScheduledExecution) { // not required
		return nil
	}

	if m.LastScheduledExecution != nil {
		if err := m.LastScheduledExecution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_scheduled_execution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_scheduled_execution")
			}
			return err
		}
	}

	return nil
}

func (m *DomainReportMetadata) validateLastUnscheduledExecution(formats strfmt.Registry) error {

	if err := validate.Required("last_unscheduled_execution", "body", m.LastUnscheduledExecution); err != nil {
		return err
	}

	if m.LastUnscheduledExecution != nil {
		if err := m.LastUnscheduledExecution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_unscheduled_execution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_unscheduled_execution")
			}
			return err
		}
	}

	return nil
}

func (m *DomainReportMetadata) validateSubtype(formats strfmt.Registry) error {

	if err := validate.Required("subtype", "body", m.Subtype); err != nil {
		return err
	}

	return nil
}

func (m *DomainReportMetadata) validateXdrParams(formats strfmt.Registry) error {
	if swag.IsZero(m.XdrParams) { // not required
		return nil
	}

	if m.XdrParams != nil {
		if err := m.XdrParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("xdr_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("xdr_params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain report metadata based on the context it is used
func (m *DomainReportMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiscoverParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastScheduledExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUnscheduledExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateXdrParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainReportMetadata) contextValidateDiscoverParams(ctx context.Context, formats strfmt.Registry) error {

	if m.DiscoverParams != nil {

		if swag.IsZero(m.DiscoverParams) { // not required
			return nil
		}

		if err := m.DiscoverParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discover_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discover_params")
			}
			return err
		}
	}

	return nil
}

func (m *DomainReportMetadata) contextValidateLastScheduledExecution(ctx context.Context, formats strfmt.Registry) error {

	if m.LastScheduledExecution != nil {

		if swag.IsZero(m.LastScheduledExecution) { // not required
			return nil
		}

		if err := m.LastScheduledExecution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_scheduled_execution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_scheduled_execution")
			}
			return err
		}
	}

	return nil
}

func (m *DomainReportMetadata) contextValidateLastUnscheduledExecution(ctx context.Context, formats strfmt.Registry) error {

	if m.LastUnscheduledExecution != nil {

		if err := m.LastUnscheduledExecution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_unscheduled_execution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_unscheduled_execution")
			}
			return err
		}
	}

	return nil
}

func (m *DomainReportMetadata) contextValidateXdrParams(ctx context.Context, formats strfmt.Registry) error {

	if m.XdrParams != nil {

		if swag.IsZero(m.XdrParams) { // not required
			return nil
		}

		if err := m.XdrParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("xdr_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("xdr_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainReportMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainReportMetadata) UnmarshalBinary(b []byte) error {
	var res DomainReportMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
