// Code generated by go-swagger; DO NOT EDIT.

package prevention_policies

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

// NewQueryPreventionPoliciesParams creates a new QueryPreventionPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryPreventionPoliciesParams() *QueryPreventionPoliciesParams {
	return &QueryPreventionPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryPreventionPoliciesParamsWithTimeout creates a new QueryPreventionPoliciesParams object
// with the ability to set a timeout on a request.
func NewQueryPreventionPoliciesParamsWithTimeout(timeout time.Duration) *QueryPreventionPoliciesParams {
	return &QueryPreventionPoliciesParams{
		timeout: timeout,
	}
}

// NewQueryPreventionPoliciesParamsWithContext creates a new QueryPreventionPoliciesParams object
// with the ability to set a context for a request.
func NewQueryPreventionPoliciesParamsWithContext(ctx context.Context) *QueryPreventionPoliciesParams {
	return &QueryPreventionPoliciesParams{
		Context: ctx,
	}
}

// NewQueryPreventionPoliciesParamsWithHTTPClient creates a new QueryPreventionPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryPreventionPoliciesParamsWithHTTPClient(client *http.Client) *QueryPreventionPoliciesParams {
	return &QueryPreventionPoliciesParams{
		HTTPClient: client,
	}
}

/*
QueryPreventionPoliciesParams contains all the parameters to send to the API endpoint

	for the query prevention policies operation.

	Typically these are written to a http.Request.
*/
type QueryPreventionPoliciesParams struct {

	/* Filter.

	   The filter expression that should be used to limit the results
	*/
	Filter *string

	/* Limit.

	   The maximum records to return. [1-5000]
	*/
	Limit *int64

	/* Offset.

	   The offset to start retrieving records from
	*/
	Offset *int64

	/* Sort.

	   The property to sort by
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query prevention policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryPreventionPoliciesParams) WithDefaults() *QueryPreventionPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query prevention policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryPreventionPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query prevention policies params
func (o *QueryPreventionPoliciesParams) WithTimeout(timeout time.Duration) *QueryPreventionPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query prevention policies params
func (o *QueryPreventionPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query prevention policies params
func (o *QueryPreventionPoliciesParams) WithContext(ctx context.Context) *QueryPreventionPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query prevention policies params
func (o *QueryPreventionPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query prevention policies params
func (o *QueryPreventionPoliciesParams) WithHTTPClient(client *http.Client) *QueryPreventionPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query prevention policies params
func (o *QueryPreventionPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query prevention policies params
func (o *QueryPreventionPoliciesParams) WithFilter(filter *string) *QueryPreventionPoliciesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query prevention policies params
func (o *QueryPreventionPoliciesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query prevention policies params
func (o *QueryPreventionPoliciesParams) WithLimit(limit *int64) *QueryPreventionPoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query prevention policies params
func (o *QueryPreventionPoliciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query prevention policies params
func (o *QueryPreventionPoliciesParams) WithOffset(offset *int64) *QueryPreventionPoliciesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query prevention policies params
func (o *QueryPreventionPoliciesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the query prevention policies params
func (o *QueryPreventionPoliciesParams) WithSort(sort *string) *QueryPreventionPoliciesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query prevention policies params
func (o *QueryPreventionPoliciesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryPreventionPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

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
