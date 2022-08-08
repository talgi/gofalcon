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

// NewSetFirewallPoliciesPrecedenceParams creates a new SetFirewallPoliciesPrecedenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetFirewallPoliciesPrecedenceParams() *SetFirewallPoliciesPrecedenceParams {
	return &SetFirewallPoliciesPrecedenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetFirewallPoliciesPrecedenceParamsWithTimeout creates a new SetFirewallPoliciesPrecedenceParams object
// with the ability to set a timeout on a request.
func NewSetFirewallPoliciesPrecedenceParamsWithTimeout(timeout time.Duration) *SetFirewallPoliciesPrecedenceParams {
	return &SetFirewallPoliciesPrecedenceParams{
		timeout: timeout,
	}
}

// NewSetFirewallPoliciesPrecedenceParamsWithContext creates a new SetFirewallPoliciesPrecedenceParams object
// with the ability to set a context for a request.
func NewSetFirewallPoliciesPrecedenceParamsWithContext(ctx context.Context) *SetFirewallPoliciesPrecedenceParams {
	return &SetFirewallPoliciesPrecedenceParams{
		Context: ctx,
	}
}

// NewSetFirewallPoliciesPrecedenceParamsWithHTTPClient creates a new SetFirewallPoliciesPrecedenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetFirewallPoliciesPrecedenceParamsWithHTTPClient(client *http.Client) *SetFirewallPoliciesPrecedenceParams {
	return &SetFirewallPoliciesPrecedenceParams{
		HTTPClient: client,
	}
}

/*
SetFirewallPoliciesPrecedenceParams contains all the parameters to send to the API endpoint

	for the set firewall policies precedence operation.

	Typically these are written to a http.Request.
*/
type SetFirewallPoliciesPrecedenceParams struct {

	// Body.
	Body *models.RequestsSetPolicyPrecedenceReqV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set firewall policies precedence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetFirewallPoliciesPrecedenceParams) WithDefaults() *SetFirewallPoliciesPrecedenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set firewall policies precedence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetFirewallPoliciesPrecedenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set firewall policies precedence params
func (o *SetFirewallPoliciesPrecedenceParams) WithTimeout(timeout time.Duration) *SetFirewallPoliciesPrecedenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set firewall policies precedence params
func (o *SetFirewallPoliciesPrecedenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set firewall policies precedence params
func (o *SetFirewallPoliciesPrecedenceParams) WithContext(ctx context.Context) *SetFirewallPoliciesPrecedenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set firewall policies precedence params
func (o *SetFirewallPoliciesPrecedenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set firewall policies precedence params
func (o *SetFirewallPoliciesPrecedenceParams) WithHTTPClient(client *http.Client) *SetFirewallPoliciesPrecedenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set firewall policies precedence params
func (o *SetFirewallPoliciesPrecedenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set firewall policies precedence params
func (o *SetFirewallPoliciesPrecedenceParams) WithBody(body *models.RequestsSetPolicyPrecedenceReqV1) *SetFirewallPoliciesPrecedenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set firewall policies precedence params
func (o *SetFirewallPoliciesPrecedenceParams) SetBody(body *models.RequestsSetPolicyPrecedenceReqV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetFirewallPoliciesPrecedenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
