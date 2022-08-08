// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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
	"github.com/go-openapi/swag"
)

// NewGetLocationsParams creates a new GetLocationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLocationsParams() *GetLocationsParams {
	return &GetLocationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLocationsParamsWithTimeout creates a new GetLocationsParams object
// with the ability to set a timeout on a request.
func NewGetLocationsParamsWithTimeout(timeout time.Duration) *GetLocationsParams {
	return &GetLocationsParams{
		timeout: timeout,
	}
}

// NewGetLocationsParamsWithContext creates a new GetLocationsParams object
// with the ability to set a context for a request.
func NewGetLocationsParamsWithContext(ctx context.Context) *GetLocationsParams {
	return &GetLocationsParams{
		Context: ctx,
	}
}

// NewGetLocationsParamsWithHTTPClient creates a new GetLocationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLocationsParamsWithHTTPClient(client *http.Client) *GetLocationsParams {
	return &GetLocationsParams{
		HTTPClient: client,
	}
}

/*
GetLocationsParams contains all the parameters to send to the API endpoint

	for the get locations operation.

	Typically these are written to a http.Request.
*/
type GetLocationsParams struct {

	/* Clouds.

	   Cloud Provider
	*/
	Clouds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get locations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLocationsParams) WithDefaults() *GetLocationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get locations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLocationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get locations params
func (o *GetLocationsParams) WithTimeout(timeout time.Duration) *GetLocationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get locations params
func (o *GetLocationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get locations params
func (o *GetLocationsParams) WithContext(ctx context.Context) *GetLocationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get locations params
func (o *GetLocationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get locations params
func (o *GetLocationsParams) WithHTTPClient(client *http.Client) *GetLocationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get locations params
func (o *GetLocationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClouds adds the clouds to the get locations params
func (o *GetLocationsParams) WithClouds(clouds []string) *GetLocationsParams {
	o.SetClouds(clouds)
	return o
}

// SetClouds adds the clouds to the get locations params
func (o *GetLocationsParams) SetClouds(clouds []string) {
	o.Clouds = clouds
}

// WriteToRequest writes these params to a swagger request
func (o *GetLocationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Clouds != nil {

		// binding items for clouds
		joinedClouds := o.bindParamClouds(reg)

		// query array param clouds
		if err := r.SetQueryParam("clouds", joinedClouds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetLocations binds the parameter clouds
func (o *GetLocationsParams) bindParamClouds(formats strfmt.Registry) []string {
	cloudsIR := o.Clouds

	var cloudsIC []string
	for _, cloudsIIR := range cloudsIR { // explode []string

		cloudsIIV := cloudsIIR // string as string
		cloudsIC = append(cloudsIC, cloudsIIV)
	}

	// items.CollectionFormat: "csv"
	cloudsIS := swag.JoinByFormat(cloudsIC, "csv")

	return cloudsIS
}
