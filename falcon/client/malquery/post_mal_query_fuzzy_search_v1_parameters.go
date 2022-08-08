// Code generated by go-swagger; DO NOT EDIT.

package malquery

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

// NewPostMalQueryFuzzySearchV1Params creates a new PostMalQueryFuzzySearchV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostMalQueryFuzzySearchV1Params() *PostMalQueryFuzzySearchV1Params {
	return &PostMalQueryFuzzySearchV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostMalQueryFuzzySearchV1ParamsWithTimeout creates a new PostMalQueryFuzzySearchV1Params object
// with the ability to set a timeout on a request.
func NewPostMalQueryFuzzySearchV1ParamsWithTimeout(timeout time.Duration) *PostMalQueryFuzzySearchV1Params {
	return &PostMalQueryFuzzySearchV1Params{
		timeout: timeout,
	}
}

// NewPostMalQueryFuzzySearchV1ParamsWithContext creates a new PostMalQueryFuzzySearchV1Params object
// with the ability to set a context for a request.
func NewPostMalQueryFuzzySearchV1ParamsWithContext(ctx context.Context) *PostMalQueryFuzzySearchV1Params {
	return &PostMalQueryFuzzySearchV1Params{
		Context: ctx,
	}
}

// NewPostMalQueryFuzzySearchV1ParamsWithHTTPClient creates a new PostMalQueryFuzzySearchV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewPostMalQueryFuzzySearchV1ParamsWithHTTPClient(client *http.Client) *PostMalQueryFuzzySearchV1Params {
	return &PostMalQueryFuzzySearchV1Params{
		HTTPClient: client,
	}
}

/*
PostMalQueryFuzzySearchV1Params contains all the parameters to send to the API endpoint

	for the post mal query fuzzy search v1 operation.

	Typically these are written to a http.Request.
*/
type PostMalQueryFuzzySearchV1Params struct {

	/* Body.

	   Fuzzy search parameters. See model for more details.
	*/
	Body *models.MalqueryFuzzySearchParametersV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post mal query fuzzy search v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostMalQueryFuzzySearchV1Params) WithDefaults() *PostMalQueryFuzzySearchV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post mal query fuzzy search v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostMalQueryFuzzySearchV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post mal query fuzzy search v1 params
func (o *PostMalQueryFuzzySearchV1Params) WithTimeout(timeout time.Duration) *PostMalQueryFuzzySearchV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post mal query fuzzy search v1 params
func (o *PostMalQueryFuzzySearchV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post mal query fuzzy search v1 params
func (o *PostMalQueryFuzzySearchV1Params) WithContext(ctx context.Context) *PostMalQueryFuzzySearchV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post mal query fuzzy search v1 params
func (o *PostMalQueryFuzzySearchV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post mal query fuzzy search v1 params
func (o *PostMalQueryFuzzySearchV1Params) WithHTTPClient(client *http.Client) *PostMalQueryFuzzySearchV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post mal query fuzzy search v1 params
func (o *PostMalQueryFuzzySearchV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post mal query fuzzy search v1 params
func (o *PostMalQueryFuzzySearchV1Params) WithBody(body *models.MalqueryFuzzySearchParametersV1) *PostMalQueryFuzzySearchV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post mal query fuzzy search v1 params
func (o *PostMalQueryFuzzySearchV1Params) SetBody(body *models.MalqueryFuzzySearchParametersV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostMalQueryFuzzySearchV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
