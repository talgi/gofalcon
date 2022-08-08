// Code generated by go-swagger; DO NOT EDIT.

package response_policies

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

// NewDeleteRTResponsePoliciesParams creates a new DeleteRTResponsePoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRTResponsePoliciesParams() *DeleteRTResponsePoliciesParams {
	return &DeleteRTResponsePoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRTResponsePoliciesParamsWithTimeout creates a new DeleteRTResponsePoliciesParams object
// with the ability to set a timeout on a request.
func NewDeleteRTResponsePoliciesParamsWithTimeout(timeout time.Duration) *DeleteRTResponsePoliciesParams {
	return &DeleteRTResponsePoliciesParams{
		timeout: timeout,
	}
}

// NewDeleteRTResponsePoliciesParamsWithContext creates a new DeleteRTResponsePoliciesParams object
// with the ability to set a context for a request.
func NewDeleteRTResponsePoliciesParamsWithContext(ctx context.Context) *DeleteRTResponsePoliciesParams {
	return &DeleteRTResponsePoliciesParams{
		Context: ctx,
	}
}

// NewDeleteRTResponsePoliciesParamsWithHTTPClient creates a new DeleteRTResponsePoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRTResponsePoliciesParamsWithHTTPClient(client *http.Client) *DeleteRTResponsePoliciesParams {
	return &DeleteRTResponsePoliciesParams{
		HTTPClient: client,
	}
}

/*
DeleteRTResponsePoliciesParams contains all the parameters to send to the API endpoint

	for the delete r t response policies operation.

	Typically these are written to a http.Request.
*/
type DeleteRTResponsePoliciesParams struct {

	/* Ids.

	   The IDs of the Response Policies to delete
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete r t response policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRTResponsePoliciesParams) WithDefaults() *DeleteRTResponsePoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete r t response policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRTResponsePoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete r t response policies params
func (o *DeleteRTResponsePoliciesParams) WithTimeout(timeout time.Duration) *DeleteRTResponsePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete r t response policies params
func (o *DeleteRTResponsePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete r t response policies params
func (o *DeleteRTResponsePoliciesParams) WithContext(ctx context.Context) *DeleteRTResponsePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete r t response policies params
func (o *DeleteRTResponsePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete r t response policies params
func (o *DeleteRTResponsePoliciesParams) WithHTTPClient(client *http.Client) *DeleteRTResponsePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete r t response policies params
func (o *DeleteRTResponsePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the delete r t response policies params
func (o *DeleteRTResponsePoliciesParams) WithIds(ids []string) *DeleteRTResponsePoliciesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete r t response policies params
func (o *DeleteRTResponsePoliciesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRTResponsePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeleteRTResponsePolicies binds the parameter ids
func (o *DeleteRTResponsePoliciesParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}
