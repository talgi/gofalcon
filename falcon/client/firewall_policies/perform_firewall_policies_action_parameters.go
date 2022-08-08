// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewPerformFirewallPoliciesActionParams creates a new PerformFirewallPoliciesActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformFirewallPoliciesActionParams() *PerformFirewallPoliciesActionParams {
	return &PerformFirewallPoliciesActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformFirewallPoliciesActionParamsWithTimeout creates a new PerformFirewallPoliciesActionParams object
// with the ability to set a timeout on a request.
func NewPerformFirewallPoliciesActionParamsWithTimeout(timeout time.Duration) *PerformFirewallPoliciesActionParams {
	return &PerformFirewallPoliciesActionParams{
		timeout: timeout,
	}
}

// NewPerformFirewallPoliciesActionParamsWithContext creates a new PerformFirewallPoliciesActionParams object
// with the ability to set a context for a request.
func NewPerformFirewallPoliciesActionParamsWithContext(ctx context.Context) *PerformFirewallPoliciesActionParams {
	return &PerformFirewallPoliciesActionParams{
		Context: ctx,
	}
}

// NewPerformFirewallPoliciesActionParamsWithHTTPClient creates a new PerformFirewallPoliciesActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformFirewallPoliciesActionParamsWithHTTPClient(client *http.Client) *PerformFirewallPoliciesActionParams {
	return &PerformFirewallPoliciesActionParams{
		HTTPClient: client,
	}
}

/*
PerformFirewallPoliciesActionParams contains all the parameters to send to the API endpoint

	for the perform firewall policies action operation.

	Typically these are written to a http.Request.
*/
type PerformFirewallPoliciesActionParams struct {

	/* ActionName.

	   The action to perform
	*/
	ActionName string

	// Body.
	Body *models.MsaEntityActionRequestV2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the perform firewall policies action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformFirewallPoliciesActionParams) WithDefaults() *PerformFirewallPoliciesActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform firewall policies action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformFirewallPoliciesActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the perform firewall policies action params
func (o *PerformFirewallPoliciesActionParams) WithTimeout(timeout time.Duration) *PerformFirewallPoliciesActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform firewall policies action params
func (o *PerformFirewallPoliciesActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform firewall policies action params
func (o *PerformFirewallPoliciesActionParams) WithContext(ctx context.Context) *PerformFirewallPoliciesActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform firewall policies action params
func (o *PerformFirewallPoliciesActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform firewall policies action params
func (o *PerformFirewallPoliciesActionParams) WithHTTPClient(client *http.Client) *PerformFirewallPoliciesActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform firewall policies action params
func (o *PerformFirewallPoliciesActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionName adds the actionName to the perform firewall policies action params
func (o *PerformFirewallPoliciesActionParams) WithActionName(actionName string) *PerformFirewallPoliciesActionParams {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the perform firewall policies action params
func (o *PerformFirewallPoliciesActionParams) SetActionName(actionName string) {
	o.ActionName = actionName
}

// WithBody adds the body to the perform firewall policies action params
func (o *PerformFirewallPoliciesActionParams) WithBody(body *models.MsaEntityActionRequestV2) *PerformFirewallPoliciesActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the perform firewall policies action params
func (o *PerformFirewallPoliciesActionParams) SetBody(body *models.MsaEntityActionRequestV2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PerformFirewallPoliciesActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param action_name
	qrActionName := o.ActionName
	qActionName := qrActionName
	if qActionName != "" {

		if err := r.SetQueryParam("action_name", qActionName); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
