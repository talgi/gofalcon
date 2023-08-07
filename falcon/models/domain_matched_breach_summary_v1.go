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

// DomainMatchedBreachSummaryV1 domain matched breach summary v1
//
// swagger:model domain.MatchedBreachSummaryV1
type DomainMatchedBreachSummaryV1 struct {

	// Community/colloquial exposed data event name.
	CommunityName string `json:"community_name,omitempty"`

	// The level of confidence regarding data veridicality. Options for likely authentic, confirmed authentic (default: unverified).
	ConfidenceLevel string `json:"confidence_level,omitempty"`

	// credentials domains
	CredentialsDomains []string `json:"credentials_domains"`

	// credentials ips
	CredentialsIps []string `json:"credentials_ips"`

	// The description of the breach
	// Required: true
	Description *string `json:"description"`

	// The date the exposed data event occurred.
	EventDate string `json:"event_date,omitempty"`

	// CrowdStrike-generated unique exposed data event identifier.
	EventID string `json:"event_id,omitempty"`

	// The date when the data was leaked online
	// Format: date-time
	ExposureDate strfmt.DateTime `json:"exposure_date,omitempty"`

	// The set of fields which were breached: 'email', 'password', 'login_id', 'phone', etc.
	// Required: true
	Fields []string `json:"fields"`

	// Metadata regarding the file(s) where exposed data records where found.
	Files []*DomainFileDetailsV1 `json:"files"`

	// idp send date
	// Format: date-time
	IdpSendDate strfmt.DateTime `json:"idp_send_date,omitempty"`

	// idp send status
	IdpSendStatus string `json:"idp_send_status,omitempty"`

	// The name of the breach
	// Required: true
	Name *string `json:"name"`

	// Exposed Data Event Threat Actor/Group: Moniker(s) or real name(s) of the individual/group who unveiled confidential data.
	ObtainedBy string `json:"obtained_by,omitempty"`

	// Where the leak was found.
	URL string `json:"url,omitempty"`
}

// Validate validates this domain matched breach summary v1
func (m *DomainMatchedBreachSummaryV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExposureDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdpSendDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainMatchedBreachSummaryV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *DomainMatchedBreachSummaryV1) validateExposureDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExposureDate) { // not required
		return nil
	}

	if err := validate.FormatOf("exposure_date", "body", "date-time", m.ExposureDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainMatchedBreachSummaryV1) validateFields(formats strfmt.Registry) error {

	if err := validate.Required("fields", "body", m.Fields); err != nil {
		return err
	}

	return nil
}

func (m *DomainMatchedBreachSummaryV1) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainMatchedBreachSummaryV1) validateIdpSendDate(formats strfmt.Registry) error {
	if swag.IsZero(m.IdpSendDate) { // not required
		return nil
	}

	if err := validate.FormatOf("idp_send_date", "body", "date-time", m.IdpSendDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainMatchedBreachSummaryV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain matched breach summary v1 based on the context it is used
func (m *DomainMatchedBreachSummaryV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainMatchedBreachSummaryV1) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if m.Files[i] != nil {

			if swag.IsZero(m.Files[i]) { // not required
				return nil
			}

			if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainMatchedBreachSummaryV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainMatchedBreachSummaryV1) UnmarshalBinary(b []byte) error {
	var res DomainMatchedBreachSummaryV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
