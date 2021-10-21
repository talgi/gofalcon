// Code generated by go-swagger; DO NOT EDIT.

package report_executions

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

// NewReportExecutionsRetryParams creates a new ReportExecutionsRetryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReportExecutionsRetryParams() *ReportExecutionsRetryParams {
	return &ReportExecutionsRetryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReportExecutionsRetryParamsWithTimeout creates a new ReportExecutionsRetryParams object
// with the ability to set a timeout on a request.
func NewReportExecutionsRetryParamsWithTimeout(timeout time.Duration) *ReportExecutionsRetryParams {
	return &ReportExecutionsRetryParams{
		timeout: timeout,
	}
}

// NewReportExecutionsRetryParamsWithContext creates a new ReportExecutionsRetryParams object
// with the ability to set a context for a request.
func NewReportExecutionsRetryParamsWithContext(ctx context.Context) *ReportExecutionsRetryParams {
	return &ReportExecutionsRetryParams{
		Context: ctx,
	}
}

// NewReportExecutionsRetryParamsWithHTTPClient creates a new ReportExecutionsRetryParams object
// with the ability to set a custom HTTPClient for a request.
func NewReportExecutionsRetryParamsWithHTTPClient(client *http.Client) *ReportExecutionsRetryParams {
	return &ReportExecutionsRetryParams{
		HTTPClient: client,
	}
}

/* ReportExecutionsRetryParams contains all the parameters to send to the API endpoint
   for the report executions retry operation.

   Typically these are written to a http.Request.
*/
type ReportExecutionsRetryParams struct {

	/* XCSUSERID.

	   The user id
	*/
	XCSUSERID *string

	/* XCSUSERUUID.

	   The user uuid
	*/
	XCSUSERUUID string

	// Body.
	Body []*models.APIReportExecutionRetryRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the report executions retry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportExecutionsRetryParams) WithDefaults() *ReportExecutionsRetryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the report executions retry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportExecutionsRetryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the report executions retry params
func (o *ReportExecutionsRetryParams) WithTimeout(timeout time.Duration) *ReportExecutionsRetryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report executions retry params
func (o *ReportExecutionsRetryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report executions retry params
func (o *ReportExecutionsRetryParams) WithContext(ctx context.Context) *ReportExecutionsRetryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report executions retry params
func (o *ReportExecutionsRetryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report executions retry params
func (o *ReportExecutionsRetryParams) WithHTTPClient(client *http.Client) *ReportExecutionsRetryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report executions retry params
func (o *ReportExecutionsRetryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERID adds the xCSUSERID to the report executions retry params
func (o *ReportExecutionsRetryParams) WithXCSUSERID(xCSUSERID *string) *ReportExecutionsRetryParams {
	o.SetXCSUSERID(xCSUSERID)
	return o
}

// SetXCSUSERID adds the xCSUSERId to the report executions retry params
func (o *ReportExecutionsRetryParams) SetXCSUSERID(xCSUSERID *string) {
	o.XCSUSERID = xCSUSERID
}

// WithXCSUSERUUID adds the xCSUSERUUID to the report executions retry params
func (o *ReportExecutionsRetryParams) WithXCSUSERUUID(xCSUSERUUID string) *ReportExecutionsRetryParams {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the report executions retry params
func (o *ReportExecutionsRetryParams) SetXCSUSERUUID(xCSUSERUUID string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithBody adds the body to the report executions retry params
func (o *ReportExecutionsRetryParams) WithBody(body []*models.APIReportExecutionRetryRequestV1) *ReportExecutionsRetryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the report executions retry params
func (o *ReportExecutionsRetryParams) SetBody(body []*models.APIReportExecutionRetryRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ReportExecutionsRetryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XCSUSERID != nil {

		// header param X-CS-USERID
		if err := r.SetHeaderParam("X-CS-USERID", *o.XCSUSERID); err != nil {
			return err
		}
	}

	// header param X-CS-USERUUID
	if err := r.SetHeaderParam("X-CS-USERUUID", o.XCSUSERUUID); err != nil {
		return err
	}
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
