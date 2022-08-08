// Code generated by go-swagger; DO NOT EDIT.

package filevantage

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

// NewQueryChangesParams creates a new QueryChangesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryChangesParams() *QueryChangesParams {
	return &QueryChangesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryChangesParamsWithTimeout creates a new QueryChangesParams object
// with the ability to set a timeout on a request.
func NewQueryChangesParamsWithTimeout(timeout time.Duration) *QueryChangesParams {
	return &QueryChangesParams{
		timeout: timeout,
	}
}

// NewQueryChangesParamsWithContext creates a new QueryChangesParams object
// with the ability to set a context for a request.
func NewQueryChangesParamsWithContext(ctx context.Context) *QueryChangesParams {
	return &QueryChangesParams{
		Context: ctx,
	}
}

// NewQueryChangesParamsWithHTTPClient creates a new QueryChangesParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryChangesParamsWithHTTPClient(client *http.Client) *QueryChangesParams {
	return &QueryChangesParams{
		HTTPClient: client,
	}
}

/*
QueryChangesParams contains all the parameters to send to the API endpoint

	for the query changes operation.

	Typically these are written to a http.Request.
*/
type QueryChangesParams struct {

	/* Filter.

	     Filter changes using a query in Falcon Query Language (FQL).

	Common filter options include:

	 - `host.host_name`
	 - `action_timestamp`

	 The full list of allowed filter parameters can be reviewed in our API documentation.
	*/
	Filter *string

	/* Limit.

	   The maximum number of changes to return in the response (default: 100; max: 500). Use with the `offset` parameter to manage pagination of results
	*/
	Limit *int64

	/* Offset.

	   The first change index to return in the response. If not provided it will default to '0'. Use with the `limit` parameter to manage pagination of results.
	*/
	Offset *int64

	/* Sort.

	     Sort changes using options like:

	- `action_timestamp` (timestamp of the change occurrence)

	 Sort either `asc` (ascending) or `desc` (descending). For example: `action_timestamp|asc`.
	The full list of allowed sorting options can be reviewed in our API documentation.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query changes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryChangesParams) WithDefaults() *QueryChangesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query changes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryChangesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query changes params
func (o *QueryChangesParams) WithTimeout(timeout time.Duration) *QueryChangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query changes params
func (o *QueryChangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query changes params
func (o *QueryChangesParams) WithContext(ctx context.Context) *QueryChangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query changes params
func (o *QueryChangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query changes params
func (o *QueryChangesParams) WithHTTPClient(client *http.Client) *QueryChangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query changes params
func (o *QueryChangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query changes params
func (o *QueryChangesParams) WithFilter(filter *string) *QueryChangesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query changes params
func (o *QueryChangesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query changes params
func (o *QueryChangesParams) WithLimit(limit *int64) *QueryChangesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query changes params
func (o *QueryChangesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query changes params
func (o *QueryChangesParams) WithOffset(offset *int64) *QueryChangesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query changes params
func (o *QueryChangesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the query changes params
func (o *QueryChangesParams) WithSort(sort *string) *QueryChangesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query changes params
func (o *QueryChangesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryChangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
