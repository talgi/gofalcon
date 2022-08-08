// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// NewAggregatePolicyRulesParams creates a new AggregatePolicyRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregatePolicyRulesParams() *AggregatePolicyRulesParams {
	return &AggregatePolicyRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregatePolicyRulesParamsWithTimeout creates a new AggregatePolicyRulesParams object
// with the ability to set a timeout on a request.
func NewAggregatePolicyRulesParamsWithTimeout(timeout time.Duration) *AggregatePolicyRulesParams {
	return &AggregatePolicyRulesParams{
		timeout: timeout,
	}
}

// NewAggregatePolicyRulesParamsWithContext creates a new AggregatePolicyRulesParams object
// with the ability to set a context for a request.
func NewAggregatePolicyRulesParamsWithContext(ctx context.Context) *AggregatePolicyRulesParams {
	return &AggregatePolicyRulesParams{
		Context: ctx,
	}
}

// NewAggregatePolicyRulesParamsWithHTTPClient creates a new AggregatePolicyRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregatePolicyRulesParamsWithHTTPClient(client *http.Client) *AggregatePolicyRulesParams {
	return &AggregatePolicyRulesParams{
		HTTPClient: client,
	}
}

/*
AggregatePolicyRulesParams contains all the parameters to send to the API endpoint

	for the aggregate policy rules operation.

	Typically these are written to a http.Request.
*/
type AggregatePolicyRulesParams struct {

	/* Body.

	   Query criteria and settings
	*/
	Body []*models.FwmgrMsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate policy rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregatePolicyRulesParams) WithDefaults() *AggregatePolicyRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate policy rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregatePolicyRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate policy rules params
func (o *AggregatePolicyRulesParams) WithTimeout(timeout time.Duration) *AggregatePolicyRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate policy rules params
func (o *AggregatePolicyRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate policy rules params
func (o *AggregatePolicyRulesParams) WithContext(ctx context.Context) *AggregatePolicyRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate policy rules params
func (o *AggregatePolicyRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate policy rules params
func (o *AggregatePolicyRulesParams) WithHTTPClient(client *http.Client) *AggregatePolicyRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate policy rules params
func (o *AggregatePolicyRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aggregate policy rules params
func (o *AggregatePolicyRulesParams) WithBody(body []*models.FwmgrMsaAggregateQueryRequest) *AggregatePolicyRulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aggregate policy rules params
func (o *AggregatePolicyRulesParams) SetBody(body []*models.FwmgrMsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AggregatePolicyRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
