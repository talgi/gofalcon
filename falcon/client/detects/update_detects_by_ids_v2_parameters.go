// Code generated by go-swagger; DO NOT EDIT.

package detects

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

// NewUpdateDetectsByIdsV2Params creates a new UpdateDetectsByIdsV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDetectsByIdsV2Params() *UpdateDetectsByIdsV2Params {
	return &UpdateDetectsByIdsV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDetectsByIdsV2ParamsWithTimeout creates a new UpdateDetectsByIdsV2Params object
// with the ability to set a timeout on a request.
func NewUpdateDetectsByIdsV2ParamsWithTimeout(timeout time.Duration) *UpdateDetectsByIdsV2Params {
	return &UpdateDetectsByIdsV2Params{
		timeout: timeout,
	}
}

// NewUpdateDetectsByIdsV2ParamsWithContext creates a new UpdateDetectsByIdsV2Params object
// with the ability to set a context for a request.
func NewUpdateDetectsByIdsV2ParamsWithContext(ctx context.Context) *UpdateDetectsByIdsV2Params {
	return &UpdateDetectsByIdsV2Params{
		Context: ctx,
	}
}

// NewUpdateDetectsByIdsV2ParamsWithHTTPClient creates a new UpdateDetectsByIdsV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDetectsByIdsV2ParamsWithHTTPClient(client *http.Client) *UpdateDetectsByIdsV2Params {
	return &UpdateDetectsByIdsV2Params{
		HTTPClient: client,
	}
}

/*
UpdateDetectsByIdsV2Params contains all the parameters to send to the API endpoint

	for the update detects by ids v2 operation.

	Typically these are written to a http.Request.
*/
type UpdateDetectsByIdsV2Params struct {

	/* Body.

	     This endpoint modifies attributes (state and assignee) of detections.

	This endpoint accepts a query formatted as a JSON array of key-value pairs. You can update one or more attributes one or more detections with a single request.

	**`assigned_to_uuid` values**

	A user ID, such as `1234567891234567891`

	**`ids` values**

	One or more detection IDs, which you can find with the `/detects/queries/detects/v1` endpoint, the Falcon console, or the Streaming API.

	**`show_in_ui` values**

	- `true`: This detection is displayed in Falcon
	- `false`: This detection is not displayed in Falcon. Most commonly used together with the `status` key's `false_positive` value.

	**`status` values**

	- `new`
	- `in_progress`
	- `true_positive`
	- `false_positive`
	- `ignored`

	**`comment` values**
	Optional comment to add to the detection. Comments are displayed with the detection in Falcon and usually used to provide context or notes for other Falcon users. A detection can have multiple comments over time.
	*/
	Body *models.DomainDetectsEntitiesPatchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update detects by ids v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDetectsByIdsV2Params) WithDefaults() *UpdateDetectsByIdsV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update detects by ids v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDetectsByIdsV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update detects by ids v2 params
func (o *UpdateDetectsByIdsV2Params) WithTimeout(timeout time.Duration) *UpdateDetectsByIdsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update detects by ids v2 params
func (o *UpdateDetectsByIdsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update detects by ids v2 params
func (o *UpdateDetectsByIdsV2Params) WithContext(ctx context.Context) *UpdateDetectsByIdsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update detects by ids v2 params
func (o *UpdateDetectsByIdsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update detects by ids v2 params
func (o *UpdateDetectsByIdsV2Params) WithHTTPClient(client *http.Client) *UpdateDetectsByIdsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update detects by ids v2 params
func (o *UpdateDetectsByIdsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update detects by ids v2 params
func (o *UpdateDetectsByIdsV2Params) WithBody(body *models.DomainDetectsEntitiesPatchRequest) *UpdateDetectsByIdsV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update detects by ids v2 params
func (o *UpdateDetectsByIdsV2Params) SetBody(body *models.DomainDetectsEntitiesPatchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDetectsByIdsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
