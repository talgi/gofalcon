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

// DomainDiscoverAPIApplicationHost Represents information about an application's host'.
//
// swagger:model domain.DiscoverAPIApplicationHost
type DomainDiscoverAPIApplicationHost struct {

	// The version of the Falcon sensor that's installed on the asset.
	AgentVersion string `json:"agent_version,omitempty"`

	// The agent ID of the Falcon sensor installed on the asset.
	Aid string `json:"aid,omitempty"`

	// The name of the country where the asset is located.
	Country string `json:"country,omitempty"`

	// The last seen MAC address of the asset.
	CurrentMacAddress string `json:"current_mac_address,omitempty"`

	// The last seen network prefix of the asset.
	CurrentNetworkPrefix string `json:"current_network_prefix,omitempty"`

	// The external IPv4 address of the asset.
	ExternalIP string `json:"external_ip,omitempty"`

	// The host management groups the asset is part of.
	Groups []string `json:"groups"`

	// The asset's hostname.
	Hostname string `json:"hostname,omitempty"`

	// The unique ID of the asset.
	// Required: true
	ID *string `json:"id"`

	// Whether the asset is exposed to the internet (Yes or Unknown).
	InternetExposure string `json:"internet_exposure,omitempty"`

	// For Linux and Mac hosts: the major version, minor version, and patch version of the kernel for the asset. For Windows hosts: the build number of the asset.
	KernelVersion string `json:"kernel_version,omitempty"`

	// The domain name the asset is currently joined to.
	MachineDomain string `json:"machine_domain,omitempty"`

	// The OS version of the asset.
	OsVersion string `json:"os_version,omitempty"`

	// The organizational unit of the asset.
	Ou string `json:"ou,omitempty"`

	// The platform name of the asset (Windows, Mac, Linux).
	PlatformName string `json:"platform_name,omitempty"`

	// The product type of the asset (Workstation, Domain Controller, Server).
	ProductTypeDesc string `json:"product_type_desc,omitempty"`

	// The site name of the domain the asset is joined to (applies only to Windows hosts).
	SiteName string `json:"site_name,omitempty"`

	// The asset's system manufacturer.
	SystemManufacturer string `json:"system_manufacturer,omitempty"`

	// The sensor and cloud tags of the asset.
	Tags []string `json:"tags"`
}

// Validate validates this domain discover API application host
func (m *DomainDiscoverAPIApplicationHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDiscoverAPIApplicationHost) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain discover API application host based on context it is used
func (m *DomainDiscoverAPIApplicationHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainDiscoverAPIApplicationHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDiscoverAPIApplicationHost) UnmarshalBinary(b []byte) error {
	var res DomainDiscoverAPIApplicationHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
