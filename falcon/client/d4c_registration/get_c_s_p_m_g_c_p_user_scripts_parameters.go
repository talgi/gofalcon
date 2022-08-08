// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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
)

// NewGetCSPMGCPUserScriptsParams creates a new GetCSPMGCPUserScriptsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCSPMGCPUserScriptsParams() *GetCSPMGCPUserScriptsParams {
	return &GetCSPMGCPUserScriptsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCSPMGCPUserScriptsParamsWithTimeout creates a new GetCSPMGCPUserScriptsParams object
// with the ability to set a timeout on a request.
func NewGetCSPMGCPUserScriptsParamsWithTimeout(timeout time.Duration) *GetCSPMGCPUserScriptsParams {
	return &GetCSPMGCPUserScriptsParams{
		timeout: timeout,
	}
}

// NewGetCSPMGCPUserScriptsParamsWithContext creates a new GetCSPMGCPUserScriptsParams object
// with the ability to set a context for a request.
func NewGetCSPMGCPUserScriptsParamsWithContext(ctx context.Context) *GetCSPMGCPUserScriptsParams {
	return &GetCSPMGCPUserScriptsParams{
		Context: ctx,
	}
}

// NewGetCSPMGCPUserScriptsParamsWithHTTPClient creates a new GetCSPMGCPUserScriptsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCSPMGCPUserScriptsParamsWithHTTPClient(client *http.Client) *GetCSPMGCPUserScriptsParams {
	return &GetCSPMGCPUserScriptsParams{
		HTTPClient: client,
	}
}

/*
GetCSPMGCPUserScriptsParams contains all the parameters to send to the API endpoint

	for the get c s p m g c p user scripts operation.

	Typically these are written to a http.Request.
*/
type GetCSPMGCPUserScriptsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get c s p m g c p user scripts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCSPMGCPUserScriptsParams) WithDefaults() *GetCSPMGCPUserScriptsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get c s p m g c p user scripts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCSPMGCPUserScriptsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get c s p m g c p user scripts params
func (o *GetCSPMGCPUserScriptsParams) WithTimeout(timeout time.Duration) *GetCSPMGCPUserScriptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c s p m g c p user scripts params
func (o *GetCSPMGCPUserScriptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c s p m g c p user scripts params
func (o *GetCSPMGCPUserScriptsParams) WithContext(ctx context.Context) *GetCSPMGCPUserScriptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c s p m g c p user scripts params
func (o *GetCSPMGCPUserScriptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c s p m g c p user scripts params
func (o *GetCSPMGCPUserScriptsParams) WithHTTPClient(client *http.Client) *GetCSPMGCPUserScriptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c s p m g c p user scripts params
func (o *GetCSPMGCPUserScriptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCSPMGCPUserScriptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
