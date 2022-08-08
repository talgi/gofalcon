// Code generated by go-swagger; DO NOT EDIT.

package quick_scan

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

// NewGetScansAggregatesParams creates a new GetScansAggregatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScansAggregatesParams() *GetScansAggregatesParams {
	return &GetScansAggregatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScansAggregatesParamsWithTimeout creates a new GetScansAggregatesParams object
// with the ability to set a timeout on a request.
func NewGetScansAggregatesParamsWithTimeout(timeout time.Duration) *GetScansAggregatesParams {
	return &GetScansAggregatesParams{
		timeout: timeout,
	}
}

// NewGetScansAggregatesParamsWithContext creates a new GetScansAggregatesParams object
// with the ability to set a context for a request.
func NewGetScansAggregatesParamsWithContext(ctx context.Context) *GetScansAggregatesParams {
	return &GetScansAggregatesParams{
		Context: ctx,
	}
}

// NewGetScansAggregatesParamsWithHTTPClient creates a new GetScansAggregatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScansAggregatesParamsWithHTTPClient(client *http.Client) *GetScansAggregatesParams {
	return &GetScansAggregatesParams{
		HTTPClient: client,
	}
}

/*
GetScansAggregatesParams contains all the parameters to send to the API endpoint

	for the get scans aggregates operation.

	Typically these are written to a http.Request.
*/
type GetScansAggregatesParams struct {

	// Body.
	Body *models.MsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scans aggregates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScansAggregatesParams) WithDefaults() *GetScansAggregatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scans aggregates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScansAggregatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scans aggregates params
func (o *GetScansAggregatesParams) WithTimeout(timeout time.Duration) *GetScansAggregatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scans aggregates params
func (o *GetScansAggregatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scans aggregates params
func (o *GetScansAggregatesParams) WithContext(ctx context.Context) *GetScansAggregatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scans aggregates params
func (o *GetScansAggregatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scans aggregates params
func (o *GetScansAggregatesParams) WithHTTPClient(client *http.Client) *GetScansAggregatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scans aggregates params
func (o *GetScansAggregatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get scans aggregates params
func (o *GetScansAggregatesParams) WithBody(body *models.MsaAggregateQueryRequest) *GetScansAggregatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get scans aggregates params
func (o *GetScansAggregatesParams) SetBody(body *models.MsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetScansAggregatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
