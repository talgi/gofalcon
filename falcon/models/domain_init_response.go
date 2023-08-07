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

// DomainInitResponse domain init response
//
// swagger:model domain.InitResponse
type DomainInitResponse struct {

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// existing aid sessions
	// Required: true
	ExistingAidSessions *int32 `json:"existing_aid_sessions"`

	// offline queued
	// Required: true
	OfflineQueued *bool `json:"offline_queued"`

	// previous commands
	PreviousCommands []string `json:"previous_commands"`

	// pwd
	Pwd string `json:"pwd,omitempty"`

	// scripts
	// Required: true
	Scripts []*DomainScriptHelp `json:"scripts"`

	// session id
	// Required: true
	SessionID *string `json:"session_id"`
}

// Validate validates this domain init response
func (m *DomainInitResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExistingAidSessions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfflineQueued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScripts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainInitResponse) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainInitResponse) validateExistingAidSessions(formats strfmt.Registry) error {

	if err := validate.Required("existing_aid_sessions", "body", m.ExistingAidSessions); err != nil {
		return err
	}

	return nil
}

func (m *DomainInitResponse) validateOfflineQueued(formats strfmt.Registry) error {

	if err := validate.Required("offline_queued", "body", m.OfflineQueued); err != nil {
		return err
	}

	return nil
}

func (m *DomainInitResponse) validateScripts(formats strfmt.Registry) error {

	if err := validate.Required("scripts", "body", m.Scripts); err != nil {
		return err
	}

	for i := 0; i < len(m.Scripts); i++ {
		if swag.IsZero(m.Scripts[i]) { // not required
			continue
		}

		if m.Scripts[i] != nil {
			if err := m.Scripts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scripts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scripts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainInitResponse) validateSessionID(formats strfmt.Registry) error {

	if err := validate.Required("session_id", "body", m.SessionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain init response based on the context it is used
func (m *DomainInitResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScripts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainInitResponse) contextValidateScripts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Scripts); i++ {

		if m.Scripts[i] != nil {

			if swag.IsZero(m.Scripts[i]) { // not required
				return nil
			}

			if err := m.Scripts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scripts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scripts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainInitResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainInitResponse) UnmarshalBinary(b []byte) error {
	var res DomainInitResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
