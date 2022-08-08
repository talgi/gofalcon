// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

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

// NewSetDeviceControlPoliciesPrecedenceParams creates a new SetDeviceControlPoliciesPrecedenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetDeviceControlPoliciesPrecedenceParams() *SetDeviceControlPoliciesPrecedenceParams {
	return &SetDeviceControlPoliciesPrecedenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetDeviceControlPoliciesPrecedenceParamsWithTimeout creates a new SetDeviceControlPoliciesPrecedenceParams object
// with the ability to set a timeout on a request.
func NewSetDeviceControlPoliciesPrecedenceParamsWithTimeout(timeout time.Duration) *SetDeviceControlPoliciesPrecedenceParams {
	return &SetDeviceControlPoliciesPrecedenceParams{
		timeout: timeout,
	}
}

// NewSetDeviceControlPoliciesPrecedenceParamsWithContext creates a new SetDeviceControlPoliciesPrecedenceParams object
// with the ability to set a context for a request.
func NewSetDeviceControlPoliciesPrecedenceParamsWithContext(ctx context.Context) *SetDeviceControlPoliciesPrecedenceParams {
	return &SetDeviceControlPoliciesPrecedenceParams{
		Context: ctx,
	}
}

// NewSetDeviceControlPoliciesPrecedenceParamsWithHTTPClient creates a new SetDeviceControlPoliciesPrecedenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetDeviceControlPoliciesPrecedenceParamsWithHTTPClient(client *http.Client) *SetDeviceControlPoliciesPrecedenceParams {
	return &SetDeviceControlPoliciesPrecedenceParams{
		HTTPClient: client,
	}
}

/*
SetDeviceControlPoliciesPrecedenceParams contains all the parameters to send to the API endpoint

	for the set device control policies precedence operation.

	Typically these are written to a http.Request.
*/
type SetDeviceControlPoliciesPrecedenceParams struct {

	// Body.
	Body *models.RequestsSetPolicyPrecedenceReqV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set device control policies precedence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDeviceControlPoliciesPrecedenceParams) WithDefaults() *SetDeviceControlPoliciesPrecedenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set device control policies precedence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDeviceControlPoliciesPrecedenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set device control policies precedence params
func (o *SetDeviceControlPoliciesPrecedenceParams) WithTimeout(timeout time.Duration) *SetDeviceControlPoliciesPrecedenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set device control policies precedence params
func (o *SetDeviceControlPoliciesPrecedenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set device control policies precedence params
func (o *SetDeviceControlPoliciesPrecedenceParams) WithContext(ctx context.Context) *SetDeviceControlPoliciesPrecedenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set device control policies precedence params
func (o *SetDeviceControlPoliciesPrecedenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set device control policies precedence params
func (o *SetDeviceControlPoliciesPrecedenceParams) WithHTTPClient(client *http.Client) *SetDeviceControlPoliciesPrecedenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set device control policies precedence params
func (o *SetDeviceControlPoliciesPrecedenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set device control policies precedence params
func (o *SetDeviceControlPoliciesPrecedenceParams) WithBody(body *models.RequestsSetPolicyPrecedenceReqV1) *SetDeviceControlPoliciesPrecedenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set device control policies precedence params
func (o *SetDeviceControlPoliciesPrecedenceParams) SetBody(body *models.RequestsSetPolicyPrecedenceReqV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetDeviceControlPoliciesPrecedenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
