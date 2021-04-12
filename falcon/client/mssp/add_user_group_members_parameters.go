// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// NewAddUserGroupMembersParams creates a new AddUserGroupMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddUserGroupMembersParams() *AddUserGroupMembersParams {
	return &AddUserGroupMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddUserGroupMembersParamsWithTimeout creates a new AddUserGroupMembersParams object
// with the ability to set a timeout on a request.
func NewAddUserGroupMembersParamsWithTimeout(timeout time.Duration) *AddUserGroupMembersParams {
	return &AddUserGroupMembersParams{
		timeout: timeout,
	}
}

// NewAddUserGroupMembersParamsWithContext creates a new AddUserGroupMembersParams object
// with the ability to set a context for a request.
func NewAddUserGroupMembersParamsWithContext(ctx context.Context) *AddUserGroupMembersParams {
	return &AddUserGroupMembersParams{
		Context: ctx,
	}
}

// NewAddUserGroupMembersParamsWithHTTPClient creates a new AddUserGroupMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddUserGroupMembersParamsWithHTTPClient(client *http.Client) *AddUserGroupMembersParams {
	return &AddUserGroupMembersParams{
		HTTPClient: client,
	}
}

/* AddUserGroupMembersParams contains all the parameters to send to the API endpoint
   for the add user group members operation.

   Typically these are written to a http.Request.
*/
type AddUserGroupMembersParams struct {

	// Body.
	Body *models.DomainUserGroupMembersRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add user group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddUserGroupMembersParams) WithDefaults() *AddUserGroupMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add user group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddUserGroupMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add user group members params
func (o *AddUserGroupMembersParams) WithTimeout(timeout time.Duration) *AddUserGroupMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add user group members params
func (o *AddUserGroupMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add user group members params
func (o *AddUserGroupMembersParams) WithContext(ctx context.Context) *AddUserGroupMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add user group members params
func (o *AddUserGroupMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add user group members params
func (o *AddUserGroupMembersParams) WithHTTPClient(client *http.Client) *AddUserGroupMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add user group members params
func (o *AddUserGroupMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add user group members params
func (o *AddUserGroupMembersParams) WithBody(body *models.DomainUserGroupMembersRequestV1) *AddUserGroupMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add user group members params
func (o *AddUserGroupMembersParams) SetBody(body *models.DomainUserGroupMembersRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddUserGroupMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
