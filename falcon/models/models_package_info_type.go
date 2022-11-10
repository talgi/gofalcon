// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelsPackageInfoType models package info type
//
// swagger:model models.PackageInfoType
type ModelsPackageInfoType struct {

	// layer hash
	LayerHash string `json:"LayerHash,omitempty"`

	// layer index
	LayerIndex int32 `json:"LayerIndex,omitempty"`

	// major version
	MajorVersion string `json:"MajorVersion,omitempty"`

	// package hash
	PackageHash string `json:"PackageHash,omitempty"`

	// package provider
	PackageProvider string `json:"PackageProvider,omitempty"`

	// package source
	PackageSource string `json:"PackageSource,omitempty"`

	// product
	Product string `json:"Product,omitempty"`

	// software architecture
	SoftwareArchitecture string `json:"SoftwareArchitecture,omitempty"`

	// status
	Status string `json:"Status,omitempty"`

	// vendor
	Vendor string `json:"Vendor,omitempty"`
}

// Validate validates this models package info type
func (m *ModelsPackageInfoType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this models package info type based on context it is used
func (m *ModelsPackageInfoType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsPackageInfoType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsPackageInfoType) UnmarshalBinary(b []byte) error {
	var res ModelsPackageInfoType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}