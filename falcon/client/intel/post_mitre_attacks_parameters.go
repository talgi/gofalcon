// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// NewPostMitreAttacksParams creates a new PostMitreAttacksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostMitreAttacksParams() *PostMitreAttacksParams {
	return &PostMitreAttacksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostMitreAttacksParamsWithTimeout creates a new PostMitreAttacksParams object
// with the ability to set a timeout on a request.
func NewPostMitreAttacksParamsWithTimeout(timeout time.Duration) *PostMitreAttacksParams {
	return &PostMitreAttacksParams{
		timeout: timeout,
	}
}

// NewPostMitreAttacksParamsWithContext creates a new PostMitreAttacksParams object
// with the ability to set a context for a request.
func NewPostMitreAttacksParamsWithContext(ctx context.Context) *PostMitreAttacksParams {
	return &PostMitreAttacksParams{
		Context: ctx,
	}
}

// NewPostMitreAttacksParamsWithHTTPClient creates a new PostMitreAttacksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostMitreAttacksParamsWithHTTPClient(client *http.Client) *PostMitreAttacksParams {
	return &PostMitreAttacksParams{
		HTTPClient: client,
	}
}

/*
PostMitreAttacksParams contains all the parameters to send to the API endpoint

	for the post mitre attacks operation.

	Typically these are written to a http.Request.
*/
type PostMitreAttacksParams struct {

	// Body.
	Body *models.MsaIdsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post mitre attacks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostMitreAttacksParams) WithDefaults() *PostMitreAttacksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post mitre attacks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostMitreAttacksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post mitre attacks params
func (o *PostMitreAttacksParams) WithTimeout(timeout time.Duration) *PostMitreAttacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post mitre attacks params
func (o *PostMitreAttacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post mitre attacks params
func (o *PostMitreAttacksParams) WithContext(ctx context.Context) *PostMitreAttacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post mitre attacks params
func (o *PostMitreAttacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post mitre attacks params
func (o *PostMitreAttacksParams) WithHTTPClient(client *http.Client) *PostMitreAttacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post mitre attacks params
func (o *PostMitreAttacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post mitre attacks params
func (o *PostMitreAttacksParams) WithBody(body *models.MsaIdsRequest) *PostMitreAttacksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post mitre attacks params
func (o *PostMitreAttacksParams) SetBody(body *models.MsaIdsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostMitreAttacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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