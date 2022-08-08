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
	"github.com/go-openapi/swag"
)

// NewQueryIntelIndicatorEntitiesParams creates a new QueryIntelIndicatorEntitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryIntelIndicatorEntitiesParams() *QueryIntelIndicatorEntitiesParams {
	return &QueryIntelIndicatorEntitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryIntelIndicatorEntitiesParamsWithTimeout creates a new QueryIntelIndicatorEntitiesParams object
// with the ability to set a timeout on a request.
func NewQueryIntelIndicatorEntitiesParamsWithTimeout(timeout time.Duration) *QueryIntelIndicatorEntitiesParams {
	return &QueryIntelIndicatorEntitiesParams{
		timeout: timeout,
	}
}

// NewQueryIntelIndicatorEntitiesParamsWithContext creates a new QueryIntelIndicatorEntitiesParams object
// with the ability to set a context for a request.
func NewQueryIntelIndicatorEntitiesParamsWithContext(ctx context.Context) *QueryIntelIndicatorEntitiesParams {
	return &QueryIntelIndicatorEntitiesParams{
		Context: ctx,
	}
}

// NewQueryIntelIndicatorEntitiesParamsWithHTTPClient creates a new QueryIntelIndicatorEntitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryIntelIndicatorEntitiesParamsWithHTTPClient(client *http.Client) *QueryIntelIndicatorEntitiesParams {
	return &QueryIntelIndicatorEntitiesParams{
		HTTPClient: client,
	}
}

/*
QueryIntelIndicatorEntitiesParams contains all the parameters to send to the API endpoint

	for the query intel indicator entities operation.

	Typically these are written to a http.Request.
*/
type QueryIntelIndicatorEntitiesParams struct {

	/* Filter.

	     Filter your query by specifying FQL filter parameters. Filter parameters include:

	_marker, actors, deleted, domain_types, id, indicator, ip_address_types, kill_chains, labels, labels.created_on, labels.last_valid_on, labels.name, last_updated, malicious_confidence, malware_families, published_date, reports, scope, targets, threat_types, type, vulnerabilities.
	*/
	Filter *string

	/* IncludeDeleted.

	   If true, include both published and deleted indicators in the response. Defaults to false.
	*/
	IncludeDeleted *bool

	/* IncludeRelations.

	   If true, include related indicators in the response. Defaults to true.
	*/
	IncludeRelations *bool

	/* Limit.

	   Set the number of indicators to return. The number must be between 1 and 10000
	*/
	Limit *int64

	/* Offset.

	   Set the starting row number to return indicators from. Defaults to 0.
	*/
	Offset *int64

	/* Q.

	   Perform a generic substring search across all fields.
	*/
	Q *string

	/* Sort.

	     Order fields in ascending or descending order.

	Ex: published_date|asc.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query intel indicator entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryIntelIndicatorEntitiesParams) WithDefaults() *QueryIntelIndicatorEntitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query intel indicator entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryIntelIndicatorEntitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) WithTimeout(timeout time.Duration) *QueryIntelIndicatorEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) WithContext(ctx context.Context) *QueryIntelIndicatorEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) WithHTTPClient(client *http.Client) *QueryIntelIndicatorEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) WithFilter(filter *string) *QueryIntelIndicatorEntitiesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithIncludeDeleted adds the includeDeleted to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) WithIncludeDeleted(includeDeleted *bool) *QueryIntelIndicatorEntitiesParams {
	o.SetIncludeDeleted(includeDeleted)
	return o
}

// SetIncludeDeleted adds the includeDeleted to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) SetIncludeDeleted(includeDeleted *bool) {
	o.IncludeDeleted = includeDeleted
}

// WithIncludeRelations adds the includeRelations to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) WithIncludeRelations(includeRelations *bool) *QueryIntelIndicatorEntitiesParams {
	o.SetIncludeRelations(includeRelations)
	return o
}

// SetIncludeRelations adds the includeRelations to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) SetIncludeRelations(includeRelations *bool) {
	o.IncludeRelations = includeRelations
}

// WithLimit adds the limit to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) WithLimit(limit *int64) *QueryIntelIndicatorEntitiesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) WithOffset(offset *int64) *QueryIntelIndicatorEntitiesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) WithQ(q *string) *QueryIntelIndicatorEntitiesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) WithSort(sort *string) *QueryIntelIndicatorEntitiesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query intel indicator entities params
func (o *QueryIntelIndicatorEntitiesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryIntelIndicatorEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludeDeleted != nil {

		// query param include_deleted
		var qrIncludeDeleted bool

		if o.IncludeDeleted != nil {
			qrIncludeDeleted = *o.IncludeDeleted
		}
		qIncludeDeleted := swag.FormatBool(qrIncludeDeleted)
		if qIncludeDeleted != "" {

			if err := r.SetQueryParam("include_deleted", qIncludeDeleted); err != nil {
				return err
			}
		}
	}

	if o.IncludeRelations != nil {

		// query param include_relations
		var qrIncludeRelations bool

		if o.IncludeRelations != nil {
			qrIncludeRelations = *o.IncludeRelations
		}
		qIncludeRelations := swag.FormatBool(qrIncludeRelations)
		if qIncludeRelations != "" {

			if err := r.SetQueryParam("include_relations", qIncludeRelations); err != nil {
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

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
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
