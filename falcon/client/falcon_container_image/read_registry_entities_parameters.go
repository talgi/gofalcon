// Code generated by go-swagger; DO NOT EDIT.

package falcon_container_image

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

// NewReadRegistryEntitiesParams creates a new ReadRegistryEntitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadRegistryEntitiesParams() *ReadRegistryEntitiesParams {
	return &ReadRegistryEntitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadRegistryEntitiesParamsWithTimeout creates a new ReadRegistryEntitiesParams object
// with the ability to set a timeout on a request.
func NewReadRegistryEntitiesParamsWithTimeout(timeout time.Duration) *ReadRegistryEntitiesParams {
	return &ReadRegistryEntitiesParams{
		timeout: timeout,
	}
}

// NewReadRegistryEntitiesParamsWithContext creates a new ReadRegistryEntitiesParams object
// with the ability to set a context for a request.
func NewReadRegistryEntitiesParamsWithContext(ctx context.Context) *ReadRegistryEntitiesParams {
	return &ReadRegistryEntitiesParams{
		Context: ctx,
	}
}

// NewReadRegistryEntitiesParamsWithHTTPClient creates a new ReadRegistryEntitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadRegistryEntitiesParamsWithHTTPClient(client *http.Client) *ReadRegistryEntitiesParams {
	return &ReadRegistryEntitiesParams{
		HTTPClient: client,
	}
}

/*
ReadRegistryEntitiesParams contains all the parameters to send to the API endpoint

	for the read registry entities operation.

	Typically these are written to a http.Request.
*/
type ReadRegistryEntitiesParams struct {

	/* Limit.

	   The upper-bound on the number of records to retrieve.
	*/
	Limit *int64

	/* Offset.

	   The offset from where to begin.
	*/
	Offset *int64

	/* Sort.

	   The field to sort on, e.g. id.desc or id.asc.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read registry entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadRegistryEntitiesParams) WithDefaults() *ReadRegistryEntitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read registry entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadRegistryEntitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read registry entities params
func (o *ReadRegistryEntitiesParams) WithTimeout(timeout time.Duration) *ReadRegistryEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read registry entities params
func (o *ReadRegistryEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read registry entities params
func (o *ReadRegistryEntitiesParams) WithContext(ctx context.Context) *ReadRegistryEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read registry entities params
func (o *ReadRegistryEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read registry entities params
func (o *ReadRegistryEntitiesParams) WithHTTPClient(client *http.Client) *ReadRegistryEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read registry entities params
func (o *ReadRegistryEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the read registry entities params
func (o *ReadRegistryEntitiesParams) WithLimit(limit *int64) *ReadRegistryEntitiesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the read registry entities params
func (o *ReadRegistryEntitiesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the read registry entities params
func (o *ReadRegistryEntitiesParams) WithOffset(offset *int64) *ReadRegistryEntitiesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the read registry entities params
func (o *ReadRegistryEntitiesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the read registry entities params
func (o *ReadRegistryEntitiesParams) WithSort(sort *string) *ReadRegistryEntitiesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the read registry entities params
func (o *ReadRegistryEntitiesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ReadRegistryEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
