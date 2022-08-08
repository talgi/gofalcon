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

// NewGetIntelIndicatorEntitiesParams creates a new GetIntelIndicatorEntitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntelIndicatorEntitiesParams() *GetIntelIndicatorEntitiesParams {
	return &GetIntelIndicatorEntitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntelIndicatorEntitiesParamsWithTimeout creates a new GetIntelIndicatorEntitiesParams object
// with the ability to set a timeout on a request.
func NewGetIntelIndicatorEntitiesParamsWithTimeout(timeout time.Duration) *GetIntelIndicatorEntitiesParams {
	return &GetIntelIndicatorEntitiesParams{
		timeout: timeout,
	}
}

// NewGetIntelIndicatorEntitiesParamsWithContext creates a new GetIntelIndicatorEntitiesParams object
// with the ability to set a context for a request.
func NewGetIntelIndicatorEntitiesParamsWithContext(ctx context.Context) *GetIntelIndicatorEntitiesParams {
	return &GetIntelIndicatorEntitiesParams{
		Context: ctx,
	}
}

// NewGetIntelIndicatorEntitiesParamsWithHTTPClient creates a new GetIntelIndicatorEntitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntelIndicatorEntitiesParamsWithHTTPClient(client *http.Client) *GetIntelIndicatorEntitiesParams {
	return &GetIntelIndicatorEntitiesParams{
		HTTPClient: client,
	}
}

/*
GetIntelIndicatorEntitiesParams contains all the parameters to send to the API endpoint

	for the get intel indicator entities operation.

	Typically these are written to a http.Request.
*/
type GetIntelIndicatorEntitiesParams struct {

	// Body.
	Body *models.MsaIdsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get intel indicator entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntelIndicatorEntitiesParams) WithDefaults() *GetIntelIndicatorEntitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get intel indicator entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntelIndicatorEntitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get intel indicator entities params
func (o *GetIntelIndicatorEntitiesParams) WithTimeout(timeout time.Duration) *GetIntelIndicatorEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get intel indicator entities params
func (o *GetIntelIndicatorEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get intel indicator entities params
func (o *GetIntelIndicatorEntitiesParams) WithContext(ctx context.Context) *GetIntelIndicatorEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get intel indicator entities params
func (o *GetIntelIndicatorEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get intel indicator entities params
func (o *GetIntelIndicatorEntitiesParams) WithHTTPClient(client *http.Client) *GetIntelIndicatorEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get intel indicator entities params
func (o *GetIntelIndicatorEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get intel indicator entities params
func (o *GetIntelIndicatorEntitiesParams) WithBody(body *models.MsaIdsRequest) *GetIntelIndicatorEntitiesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get intel indicator entities params
func (o *GetIntelIndicatorEntitiesParams) SetBody(body *models.MsaIdsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntelIndicatorEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
