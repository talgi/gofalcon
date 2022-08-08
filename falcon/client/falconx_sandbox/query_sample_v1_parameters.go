// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

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

// NewQuerySampleV1Params creates a new QuerySampleV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuerySampleV1Params() *QuerySampleV1Params {
	return &QuerySampleV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuerySampleV1ParamsWithTimeout creates a new QuerySampleV1Params object
// with the ability to set a timeout on a request.
func NewQuerySampleV1ParamsWithTimeout(timeout time.Duration) *QuerySampleV1Params {
	return &QuerySampleV1Params{
		timeout: timeout,
	}
}

// NewQuerySampleV1ParamsWithContext creates a new QuerySampleV1Params object
// with the ability to set a context for a request.
func NewQuerySampleV1ParamsWithContext(ctx context.Context) *QuerySampleV1Params {
	return &QuerySampleV1Params{
		Context: ctx,
	}
}

// NewQuerySampleV1ParamsWithHTTPClient creates a new QuerySampleV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewQuerySampleV1ParamsWithHTTPClient(client *http.Client) *QuerySampleV1Params {
	return &QuerySampleV1Params{
		HTTPClient: client,
	}
}

/*
QuerySampleV1Params contains all the parameters to send to the API endpoint

	for the query sample v1 operation.

	Typically these are written to a http.Request.
*/
type QuerySampleV1Params struct {

	/* Body.

	   Pass a list of sha256s to check if the exist. It will be returned the list of existing hashes.
	*/
	Body *models.ClientQuerySamplesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query sample v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuerySampleV1Params) WithDefaults() *QuerySampleV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query sample v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuerySampleV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query sample v1 params
func (o *QuerySampleV1Params) WithTimeout(timeout time.Duration) *QuerySampleV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query sample v1 params
func (o *QuerySampleV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query sample v1 params
func (o *QuerySampleV1Params) WithContext(ctx context.Context) *QuerySampleV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query sample v1 params
func (o *QuerySampleV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query sample v1 params
func (o *QuerySampleV1Params) WithHTTPClient(client *http.Client) *QuerySampleV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query sample v1 params
func (o *QuerySampleV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the query sample v1 params
func (o *QuerySampleV1Params) WithBody(body *models.ClientQuerySamplesRequest) *QuerySampleV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the query sample v1 params
func (o *QuerySampleV1Params) SetBody(body *models.ClientQuerySamplesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *QuerySampleV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
