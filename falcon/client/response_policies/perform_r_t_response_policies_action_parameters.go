// Code generated by go-swagger; DO NOT EDIT.

package response_policies

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

// NewPerformRTResponsePoliciesActionParams creates a new PerformRTResponsePoliciesActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformRTResponsePoliciesActionParams() *PerformRTResponsePoliciesActionParams {
	return &PerformRTResponsePoliciesActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformRTResponsePoliciesActionParamsWithTimeout creates a new PerformRTResponsePoliciesActionParams object
// with the ability to set a timeout on a request.
func NewPerformRTResponsePoliciesActionParamsWithTimeout(timeout time.Duration) *PerformRTResponsePoliciesActionParams {
	return &PerformRTResponsePoliciesActionParams{
		timeout: timeout,
	}
}

// NewPerformRTResponsePoliciesActionParamsWithContext creates a new PerformRTResponsePoliciesActionParams object
// with the ability to set a context for a request.
func NewPerformRTResponsePoliciesActionParamsWithContext(ctx context.Context) *PerformRTResponsePoliciesActionParams {
	return &PerformRTResponsePoliciesActionParams{
		Context: ctx,
	}
}

// NewPerformRTResponsePoliciesActionParamsWithHTTPClient creates a new PerformRTResponsePoliciesActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformRTResponsePoliciesActionParamsWithHTTPClient(client *http.Client) *PerformRTResponsePoliciesActionParams {
	return &PerformRTResponsePoliciesActionParams{
		HTTPClient: client,
	}
}

/*
PerformRTResponsePoliciesActionParams contains all the parameters to send to the API endpoint

	for the perform r t response policies action operation.

	Typically these are written to a http.Request.
*/
type PerformRTResponsePoliciesActionParams struct {

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

// WithDefaults hydrates default values in the perform r t response policies action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformRTResponsePoliciesActionParams) WithDefaults() *PerformRTResponsePoliciesActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform r t response policies action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformRTResponsePoliciesActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the perform r t response policies action params
func (o *PerformRTResponsePoliciesActionParams) WithTimeout(timeout time.Duration) *PerformRTResponsePoliciesActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform r t response policies action params
func (o *PerformRTResponsePoliciesActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform r t response policies action params
func (o *PerformRTResponsePoliciesActionParams) WithContext(ctx context.Context) *PerformRTResponsePoliciesActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform r t response policies action params
func (o *PerformRTResponsePoliciesActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform r t response policies action params
func (o *PerformRTResponsePoliciesActionParams) WithHTTPClient(client *http.Client) *PerformRTResponsePoliciesActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform r t response policies action params
func (o *PerformRTResponsePoliciesActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionName adds the actionName to the perform r t response policies action params
func (o *PerformRTResponsePoliciesActionParams) WithActionName(actionName string) *PerformRTResponsePoliciesActionParams {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the perform r t response policies action params
func (o *PerformRTResponsePoliciesActionParams) SetActionName(actionName string) {
	o.ActionName = actionName
}

// WithBody adds the body to the perform r t response policies action params
func (o *PerformRTResponsePoliciesActionParams) WithBody(body *models.MsaEntityActionRequestV2) *PerformRTResponsePoliciesActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the perform r t response policies action params
func (o *PerformRTResponsePoliciesActionParams) SetBody(body *models.MsaEntityActionRequestV2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PerformRTResponsePoliciesActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
