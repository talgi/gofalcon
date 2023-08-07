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

// K8sregAWSAccountResp k8sreg a w s account resp
//
// swagger:model k8sreg.AWSAccountResp
type K8sregAWSAccountResp struct {

	// account id
	// Required: true
	AccountID *string `json:"account_id"`

	// aws permissions status
	// Required: true
	AwsPermissionsStatus []*K8sregAccountPermissionsStatus `json:"aws_permissions_status"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// cloudformation url
	CloudformationURL string `json:"cloudformation_url,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// from cspm
	// Required: true
	FromCspm *bool `json:"from_cspm"`

	// iam role arn
	// Required: true
	IamRoleArn *string `json:"iam_role_arn"`

	// is master
	// Required: true
	IsMaster *bool `json:"is_master"`

	// organization id
	OrganizationID string `json:"organization_id,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this k8sreg a w s account resp
func (m *K8sregAWSAccountResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsPermissionsStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromCspm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIamRoleArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsMaster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sregAWSAccountResp) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *K8sregAWSAccountResp) validateAwsPermissionsStatus(formats strfmt.Registry) error {

	if err := validate.Required("aws_permissions_status", "body", m.AwsPermissionsStatus); err != nil {
		return err
	}

	for i := 0; i < len(m.AwsPermissionsStatus); i++ {
		if swag.IsZero(m.AwsPermissionsStatus[i]) { // not required
			continue
		}

		if m.AwsPermissionsStatus[i] != nil {
			if err := m.AwsPermissionsStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aws_permissions_status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aws_permissions_status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sregAWSAccountResp) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *K8sregAWSAccountResp) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *K8sregAWSAccountResp) validateFromCspm(formats strfmt.Registry) error {

	if err := validate.Required("from_cspm", "body", m.FromCspm); err != nil {
		return err
	}

	return nil
}

func (m *K8sregAWSAccountResp) validateIamRoleArn(formats strfmt.Registry) error {

	if err := validate.Required("iam_role_arn", "body", m.IamRoleArn); err != nil {
		return err
	}

	return nil
}

func (m *K8sregAWSAccountResp) validateIsMaster(formats strfmt.Registry) error {

	if err := validate.Required("is_master", "body", m.IsMaster); err != nil {
		return err
	}

	return nil
}

func (m *K8sregAWSAccountResp) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *K8sregAWSAccountResp) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this k8sreg a w s account resp based on the context it is used
func (m *K8sregAWSAccountResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsPermissionsStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sregAWSAccountResp) contextValidateAwsPermissionsStatus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AwsPermissionsStatus); i++ {

		if m.AwsPermissionsStatus[i] != nil {

			if swag.IsZero(m.AwsPermissionsStatus[i]) { // not required
				return nil
			}

			if err := m.AwsPermissionsStatus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aws_permissions_status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aws_permissions_status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sregAWSAccountResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sregAWSAccountResp) UnmarshalBinary(b []byte) error {
	var res K8sregAWSAccountResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
