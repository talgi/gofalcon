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

// MalqueryQueryMetaInfo malquery query meta info
//
// swagger:model malquery.QueryMetaInfo
type MalqueryQueryMetaInfo struct {

	// pagination
	Pagination *MsaPaging `json:"pagination,omitempty"`

	// powered by
	PoweredBy string `json:"powered_by,omitempty"`

	// Elapsed time since the request started in seconds
	QueryTime float64 `json:"query_time,omitempty"`

	// Request ID returned after creating a hunt or exact search
	Reqid string `json:"reqid,omitempty"`

	// Request status. Possible values: inprogress, failed, done
	Status string `json:"status,omitempty"`

	// trace id
	// Required: true
	TraceID *string `json:"trace_id"`

	// writes
	Writes *MsaspecWrites `json:"writes,omitempty"`
}

// Validate validates this malquery query meta info
func (m *MalqueryQueryMetaInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWrites(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MalqueryQueryMetaInfo) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *MalqueryQueryMetaInfo) validateTraceID(formats strfmt.Registry) error {

	if err := validate.Required("trace_id", "body", m.TraceID); err != nil {
		return err
	}

	return nil
}

func (m *MalqueryQueryMetaInfo) validateWrites(formats strfmt.Registry) error {
	if swag.IsZero(m.Writes) { // not required
		return nil
	}

	if m.Writes != nil {
		if err := m.Writes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("writes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this malquery query meta info based on the context it is used
func (m *MalqueryQueryMetaInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrites(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MalqueryQueryMetaInfo) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {

		if swag.IsZero(m.Pagination) { // not required
			return nil
		}

		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *MalqueryQueryMetaInfo) contextValidateWrites(ctx context.Context, formats strfmt.Registry) error {

	if m.Writes != nil {

		if swag.IsZero(m.Writes) { // not required
			return nil
		}

		if err := m.Writes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("writes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MalqueryQueryMetaInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MalqueryQueryMetaInfo) UnmarshalBinary(b []byte) error {
	var res MalqueryQueryMetaInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
