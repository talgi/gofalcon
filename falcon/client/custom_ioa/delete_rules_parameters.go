// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// NewDeleteRulesParams creates a new DeleteRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRulesParams() *DeleteRulesParams {
	return &DeleteRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRulesParamsWithTimeout creates a new DeleteRulesParams object
// with the ability to set a timeout on a request.
func NewDeleteRulesParamsWithTimeout(timeout time.Duration) *DeleteRulesParams {
	return &DeleteRulesParams{
		timeout: timeout,
	}
}

// NewDeleteRulesParamsWithContext creates a new DeleteRulesParams object
// with the ability to set a context for a request.
func NewDeleteRulesParamsWithContext(ctx context.Context) *DeleteRulesParams {
	return &DeleteRulesParams{
		Context: ctx,
	}
}

// NewDeleteRulesParamsWithHTTPClient creates a new DeleteRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRulesParamsWithHTTPClient(client *http.Client) *DeleteRulesParams {
	return &DeleteRulesParams{
		HTTPClient: client,
	}
}

/* DeleteRulesParams contains all the parameters to send to the API endpoint
   for the delete rules operation.

   Typically these are written to a http.Request.
*/
type DeleteRulesParams struct {

	/* Comment.

	   Explains why the entity is being deleted
	*/
	Comment *string

	/* Ids.

	   The IDs of the entities
	*/
	Ids []string

	/* RuleGroupID.

	   The parent rule group
	*/
	RuleGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRulesParams) WithDefaults() *DeleteRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete rules params
func (o *DeleteRulesParams) WithTimeout(timeout time.Duration) *DeleteRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete rules params
func (o *DeleteRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete rules params
func (o *DeleteRulesParams) WithContext(ctx context.Context) *DeleteRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete rules params
func (o *DeleteRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete rules params
func (o *DeleteRulesParams) WithHTTPClient(client *http.Client) *DeleteRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete rules params
func (o *DeleteRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the delete rules params
func (o *DeleteRulesParams) WithComment(comment *string) *DeleteRulesParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the delete rules params
func (o *DeleteRulesParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithIds adds the ids to the delete rules params
func (o *DeleteRulesParams) WithIds(ids []string) *DeleteRulesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete rules params
func (o *DeleteRulesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithRuleGroupID adds the ruleGroupID to the delete rules params
func (o *DeleteRulesParams) WithRuleGroupID(ruleGroupID string) *DeleteRulesParams {
	o.SetRuleGroupID(ruleGroupID)
	return o
}

// SetRuleGroupID adds the ruleGroupId to the delete rules params
func (o *DeleteRulesParams) SetRuleGroupID(ruleGroupID string) {
	o.RuleGroupID = ruleGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	// query param rule_group_id
	qrRuleGroupID := o.RuleGroupID
	qRuleGroupID := qrRuleGroupID
	if qRuleGroupID != "" {

		if err := r.SetQueryParam("rule_group_id", qRuleGroupID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeleteRules binds the parameter ids
func (o *DeleteRulesParams) bindParamIds(formats strfmt.Registry) []string {
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
