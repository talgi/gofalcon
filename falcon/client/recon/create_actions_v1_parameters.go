// Code generated by go-swagger; DO NOT EDIT.

package recon

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

// NewCreateActionsV1Params creates a new CreateActionsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateActionsV1Params() *CreateActionsV1Params {
	return &CreateActionsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateActionsV1ParamsWithTimeout creates a new CreateActionsV1Params object
// with the ability to set a timeout on a request.
func NewCreateActionsV1ParamsWithTimeout(timeout time.Duration) *CreateActionsV1Params {
	return &CreateActionsV1Params{
		timeout: timeout,
	}
}

// NewCreateActionsV1ParamsWithContext creates a new CreateActionsV1Params object
// with the ability to set a context for a request.
func NewCreateActionsV1ParamsWithContext(ctx context.Context) *CreateActionsV1Params {
	return &CreateActionsV1Params{
		Context: ctx,
	}
}

// NewCreateActionsV1ParamsWithHTTPClient creates a new CreateActionsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateActionsV1ParamsWithHTTPClient(client *http.Client) *CreateActionsV1Params {
	return &CreateActionsV1Params{
		HTTPClient: client,
	}
}

/* CreateActionsV1Params contains all the parameters to send to the API endpoint
   for the create actions v1 operation.

   Typically these are written to a http.Request.
*/
type CreateActionsV1Params struct {

	// Body.
	Body *models.DomainRegisterActionsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create actions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateActionsV1Params) WithDefaults() *CreateActionsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create actions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateActionsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create actions v1 params
func (o *CreateActionsV1Params) WithTimeout(timeout time.Duration) *CreateActionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create actions v1 params
func (o *CreateActionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create actions v1 params
func (o *CreateActionsV1Params) WithContext(ctx context.Context) *CreateActionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create actions v1 params
func (o *CreateActionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create actions v1 params
func (o *CreateActionsV1Params) WithHTTPClient(client *http.Client) *CreateActionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create actions v1 params
func (o *CreateActionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create actions v1 params
func (o *CreateActionsV1Params) WithBody(body *models.DomainRegisterActionsRequest) *CreateActionsV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create actions v1 params
func (o *CreateActionsV1Params) SetBody(body *models.DomainRegisterActionsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateActionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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