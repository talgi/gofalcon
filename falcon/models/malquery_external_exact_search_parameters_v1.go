// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MalqueryExternalExactSearchParametersV1 malquery external exact search parameters v1
//
// swagger:model malquery.ExternalExactSearchParametersV1
type MalqueryExternalExactSearchParametersV1 struct {

	// Additional search options
	Options *MalqueryExternalHuntOptions `json:"options,omitempty"`

	// Patterns to search for
	// Required: true
	Patterns []*MalquerySearchParameter `json:"patterns"`
}

// Validate validates this malquery external exact search parameters v1
func (m *MalqueryExternalExactSearchParametersV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatterns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MalqueryExternalExactSearchParametersV1) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {
		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *MalqueryExternalExactSearchParametersV1) validatePatterns(formats strfmt.Registry) error {

	if err := validate.Required("patterns", "body", m.Patterns); err != nil {
		return err
	}

	for i := 0; i < len(m.Patterns); i++ {
		if swag.IsZero(m.Patterns[i]) { // not required
			continue
		}

		if m.Patterns[i] != nil {
			if err := m.Patterns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patterns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patterns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this malquery external exact search parameters v1 based on the context it is used
func (m *MalqueryExternalExactSearchParametersV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePatterns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MalqueryExternalExactSearchParametersV1) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.Options != nil {

		if swag.IsZero(m.Options) { // not required
			return nil
		}

		if err := m.Options.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *MalqueryExternalExactSearchParametersV1) contextValidatePatterns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Patterns); i++ {

		if m.Patterns[i] != nil {

			if swag.IsZero(m.Patterns[i]) { // not required
				return nil
			}

			if err := m.Patterns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patterns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patterns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MalqueryExternalExactSearchParametersV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MalqueryExternalExactSearchParametersV1) UnmarshalBinary(b []byte) error {
	var res MalqueryExternalExactSearchParametersV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
