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

// NewQueryIntelRuleIdsParams creates a new QueryIntelRuleIdsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryIntelRuleIdsParams() *QueryIntelRuleIdsParams {
	return &QueryIntelRuleIdsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryIntelRuleIdsParamsWithTimeout creates a new QueryIntelRuleIdsParams object
// with the ability to set a timeout on a request.
func NewQueryIntelRuleIdsParamsWithTimeout(timeout time.Duration) *QueryIntelRuleIdsParams {
	return &QueryIntelRuleIdsParams{
		timeout: timeout,
	}
}

// NewQueryIntelRuleIdsParamsWithContext creates a new QueryIntelRuleIdsParams object
// with the ability to set a context for a request.
func NewQueryIntelRuleIdsParamsWithContext(ctx context.Context) *QueryIntelRuleIdsParams {
	return &QueryIntelRuleIdsParams{
		Context: ctx,
	}
}

// NewQueryIntelRuleIdsParamsWithHTTPClient creates a new QueryIntelRuleIdsParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryIntelRuleIdsParamsWithHTTPClient(client *http.Client) *QueryIntelRuleIdsParams {
	return &QueryIntelRuleIdsParams{
		HTTPClient: client,
	}
}

/*
QueryIntelRuleIdsParams contains all the parameters to send to the API endpoint

	for the query intel rule ids operation.

	Typically these are written to a http.Request.
*/
type QueryIntelRuleIdsParams struct {

	/* Description.

	   Substring match on description field.
	*/
	Description []string

	/* Limit.

	   The number of rule IDs to return. Defaults to 10.
	*/
	Limit *int64

	/* MaxCreatedDate.

	   Filter results to those created on or before a certain date.
	*/
	MaxCreatedDate *string

	/* MinCreatedDate.

	   Filter results to those created on or after a certain date.
	*/
	MinCreatedDate *int64

	/* Name.

	   Search by rule title.
	*/
	Name []string

	/* Offset.

	   Set the starting row number to return reports from. Defaults to 0.
	*/
	Offset *int64

	/* Q.

	   Perform a generic substring search across all fields.
	*/
	Q *string

	/* Sort.

	     Order fields in ascending or descending order.

	Ex: created_date|asc.
	*/
	Sort *string

	/* Tags.

	   Search for rule tags.
	*/
	Tags []string

	/* Type.

	     The rule news report type. Accepted values:

	snort-suricata-master

	snort-suricata-update

	snort-suricata-changelog

	yara-master

	yara-update

	yara-changelog

	common-event-format

	netwitness
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query intel rule ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryIntelRuleIdsParams) WithDefaults() *QueryIntelRuleIdsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query intel rule ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryIntelRuleIdsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithTimeout(timeout time.Duration) *QueryIntelRuleIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithContext(ctx context.Context) *QueryIntelRuleIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithHTTPClient(client *http.Client) *QueryIntelRuleIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithDescription(description []string) *QueryIntelRuleIdsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetDescription(description []string) {
	o.Description = description
}

// WithLimit adds the limit to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithLimit(limit *int64) *QueryIntelRuleIdsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxCreatedDate adds the maxCreatedDate to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithMaxCreatedDate(maxCreatedDate *string) *QueryIntelRuleIdsParams {
	o.SetMaxCreatedDate(maxCreatedDate)
	return o
}

// SetMaxCreatedDate adds the maxCreatedDate to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetMaxCreatedDate(maxCreatedDate *string) {
	o.MaxCreatedDate = maxCreatedDate
}

// WithMinCreatedDate adds the minCreatedDate to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithMinCreatedDate(minCreatedDate *int64) *QueryIntelRuleIdsParams {
	o.SetMinCreatedDate(minCreatedDate)
	return o
}

// SetMinCreatedDate adds the minCreatedDate to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetMinCreatedDate(minCreatedDate *int64) {
	o.MinCreatedDate = minCreatedDate
}

// WithName adds the name to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithName(name []string) *QueryIntelRuleIdsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetName(name []string) {
	o.Name = name
}

// WithOffset adds the offset to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithOffset(offset *int64) *QueryIntelRuleIdsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithQ(q *string) *QueryIntelRuleIdsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithSort(sort *string) *QueryIntelRuleIdsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTags adds the tags to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithTags(tags []string) *QueryIntelRuleIdsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetTags(tags []string) {
	o.Tags = tags
}

// WithType adds the typeVar to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) WithType(typeVar string) *QueryIntelRuleIdsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the query intel rule ids params
func (o *QueryIntelRuleIdsParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *QueryIntelRuleIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// binding items for description
		joinedDescription := o.bindParamDescription(reg)

		// query array param description
		if err := r.SetQueryParam("description", joinedDescription...); err != nil {
			return err
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

	if o.MaxCreatedDate != nil {

		// query param max_created_date
		var qrMaxCreatedDate string

		if o.MaxCreatedDate != nil {
			qrMaxCreatedDate = *o.MaxCreatedDate
		}
		qMaxCreatedDate := qrMaxCreatedDate
		if qMaxCreatedDate != "" {

			if err := r.SetQueryParam("max_created_date", qMaxCreatedDate); err != nil {
				return err
			}
		}
	}

	if o.MinCreatedDate != nil {

		// query param min_created_date
		var qrMinCreatedDate int64

		if o.MinCreatedDate != nil {
			qrMinCreatedDate = *o.MinCreatedDate
		}
		qMinCreatedDate := swag.FormatInt64(qrMinCreatedDate)
		if qMinCreatedDate != "" {

			if err := r.SetQueryParam("min_created_date", qMinCreatedDate); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// binding items for name
		joinedName := o.bindParamName(reg)

		// query array param name
		if err := r.SetQueryParam("name", joinedName...); err != nil {
			return err
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

	if o.Tags != nil {

		// binding items for tags
		joinedTags := o.bindParamTags(reg)

		// query array param tags
		if err := r.SetQueryParam("tags", joinedTags...); err != nil {
			return err
		}
	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {

		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamQueryIntelRuleIds binds the parameter description
func (o *QueryIntelRuleIdsParams) bindParamDescription(formats strfmt.Registry) []string {
	descriptionIR := o.Description

	var descriptionIC []string
	for _, descriptionIIR := range descriptionIR { // explode []string

		descriptionIIV := descriptionIIR // string as string
		descriptionIC = append(descriptionIC, descriptionIIV)
	}

	// items.CollectionFormat: "csv"
	descriptionIS := swag.JoinByFormat(descriptionIC, "csv")

	return descriptionIS
}

// bindParamQueryIntelRuleIds binds the parameter name
func (o *QueryIntelRuleIdsParams) bindParamName(formats strfmt.Registry) []string {
	nameIR := o.Name

	var nameIC []string
	for _, nameIIR := range nameIR { // explode []string

		nameIIV := nameIIR // string as string
		nameIC = append(nameIC, nameIIV)
	}

	// items.CollectionFormat: "csv"
	nameIS := swag.JoinByFormat(nameIC, "csv")

	return nameIS
}

// bindParamQueryIntelRuleIds binds the parameter tags
func (o *QueryIntelRuleIdsParams) bindParamTags(formats strfmt.Registry) []string {
	tagsIR := o.Tags

	var tagsIC []string
	for _, tagsIIR := range tagsIR { // explode []string

		tagsIIV := tagsIIR // string as string
		tagsIC = append(tagsIC, tagsIIV)
	}

	// items.CollectionFormat: "csv"
	tagsIS := swag.JoinByFormat(tagsIC, "csv")

	return tagsIS
}
