// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// NewGetRolesParams creates a new GetRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRolesParams() *GetRolesParams {
	return &GetRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRolesParamsWithTimeout creates a new GetRolesParams object
// with the ability to set a timeout on a request.
func NewGetRolesParamsWithTimeout(timeout time.Duration) *GetRolesParams {
	return &GetRolesParams{
		timeout: timeout,
	}
}

// NewGetRolesParamsWithContext creates a new GetRolesParams object
// with the ability to set a context for a request.
func NewGetRolesParamsWithContext(ctx context.Context) *GetRolesParams {
	return &GetRolesParams{
		Context: ctx,
	}
}

// NewGetRolesParamsWithHTTPClient creates a new GetRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRolesParamsWithHTTPClient(client *http.Client) *GetRolesParams {
	return &GetRolesParams{
		HTTPClient: client,
	}
}

/*
GetRolesParams contains all the parameters to send to the API endpoint

	for the get roles operation.

	Typically these are written to a http.Request.
*/
type GetRolesParams struct {

	/* Ids.

	   ID of a role. Find a role ID from `/customer/queries/roles/v1` or `/users/queries/roles/v1`.
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRolesParams) WithDefaults() *GetRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get roles params
func (o *GetRolesParams) WithTimeout(timeout time.Duration) *GetRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get roles params
func (o *GetRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get roles params
func (o *GetRolesParams) WithContext(ctx context.Context) *GetRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get roles params
func (o *GetRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get roles params
func (o *GetRolesParams) WithHTTPClient(client *http.Client) *GetRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get roles params
func (o *GetRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get roles params
func (o *GetRolesParams) WithIds(ids []string) *GetRolesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get roles params
func (o *GetRolesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetRoles binds the parameter ids
func (o *GetRolesParams) bindParamIds(formats strfmt.Registry) []string {
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
