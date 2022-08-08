// Code generated by go-swagger; DO NOT EDIT.

package oauth2

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
)

// NewOauth2RevokeTokenParams creates a new Oauth2RevokeTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOauth2RevokeTokenParams() *Oauth2RevokeTokenParams {
	return &Oauth2RevokeTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOauth2RevokeTokenParamsWithTimeout creates a new Oauth2RevokeTokenParams object
// with the ability to set a timeout on a request.
func NewOauth2RevokeTokenParamsWithTimeout(timeout time.Duration) *Oauth2RevokeTokenParams {
	return &Oauth2RevokeTokenParams{
		timeout: timeout,
	}
}

// NewOauth2RevokeTokenParamsWithContext creates a new Oauth2RevokeTokenParams object
// with the ability to set a context for a request.
func NewOauth2RevokeTokenParamsWithContext(ctx context.Context) *Oauth2RevokeTokenParams {
	return &Oauth2RevokeTokenParams{
		Context: ctx,
	}
}

// NewOauth2RevokeTokenParamsWithHTTPClient creates a new Oauth2RevokeTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewOauth2RevokeTokenParamsWithHTTPClient(client *http.Client) *Oauth2RevokeTokenParams {
	return &Oauth2RevokeTokenParams{
		HTTPClient: client,
	}
}

/*
Oauth2RevokeTokenParams contains all the parameters to send to the API endpoint

	for the oauth2 revoke token operation.

	Typically these are written to a http.Request.
*/
type Oauth2RevokeTokenParams struct {

	/* Token.

	     The OAuth2 access token you want to revoke.

	Include your API client ID and secret in basic auth format (`Authorization: basic <encoded API client ID and secret>`) in your request header.
	*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the oauth2 revoke token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Oauth2RevokeTokenParams) WithDefaults() *Oauth2RevokeTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the oauth2 revoke token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Oauth2RevokeTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the oauth2 revoke token params
func (o *Oauth2RevokeTokenParams) WithTimeout(timeout time.Duration) *Oauth2RevokeTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oauth2 revoke token params
func (o *Oauth2RevokeTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oauth2 revoke token params
func (o *Oauth2RevokeTokenParams) WithContext(ctx context.Context) *Oauth2RevokeTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oauth2 revoke token params
func (o *Oauth2RevokeTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oauth2 revoke token params
func (o *Oauth2RevokeTokenParams) WithHTTPClient(client *http.Client) *Oauth2RevokeTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oauth2 revoke token params
func (o *Oauth2RevokeTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the oauth2 revoke token params
func (o *Oauth2RevokeTokenParams) WithToken(token string) *Oauth2RevokeTokenParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the oauth2 revoke token params
func (o *Oauth2RevokeTokenParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *Oauth2RevokeTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param token
	frToken := o.Token
	fToken := frToken
	if fToken != "" {
		if err := r.SetFormParam("token", fToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
