// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoaCloudAccountID ioa cloud account ID
//
// swagger:model ioa.CloudAccountID
type IoaCloudAccountID struct {

	// aws account id
	AwsAccountID string `json:"aws_account_id,omitempty"`

	// azure subscription id
	AzureSubscriptionID string `json:"azure_subscription_id,omitempty"`

	// azure tenant id
	AzureTenantID string `json:"azure_tenant_id,omitempty"`
}

// Validate validates this ioa cloud account ID
func (m *IoaCloudAccountID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ioa cloud account ID based on context it is used
func (m *IoaCloudAccountID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoaCloudAccountID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoaCloudAccountID) UnmarshalBinary(b []byte) error {
	var res IoaCloudAccountID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}