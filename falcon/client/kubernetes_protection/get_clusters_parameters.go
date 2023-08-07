// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// NewGetClustersParams creates a new GetClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClustersParams() *GetClustersParams {
	return &GetClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClustersParamsWithTimeout creates a new GetClustersParams object
// with the ability to set a timeout on a request.
func NewGetClustersParamsWithTimeout(timeout time.Duration) *GetClustersParams {
	return &GetClustersParams{
		timeout: timeout,
	}
}

// NewGetClustersParamsWithContext creates a new GetClustersParams object
// with the ability to set a context for a request.
func NewGetClustersParamsWithContext(ctx context.Context) *GetClustersParams {
	return &GetClustersParams{
		Context: ctx,
	}
}

// NewGetClustersParamsWithHTTPClient creates a new GetClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClustersParamsWithHTTPClient(client *http.Client) *GetClustersParams {
	return &GetClustersParams{
		HTTPClient: client,
	}
}

/*
GetClustersParams contains all the parameters to send to the API endpoint

	for the get clusters operation.

	Typically these are written to a http.Request.
*/
type GetClustersParams struct {

	/* AccountIds.

	   Cluster Account id. For EKS it will be AWS account ID.
	*/
	AccountIds []string

	/* ClusterNames.

	   Cluster name. For EKS it will be cluster ARN.
	*/
	ClusterNames []string

	/* ClusterService.

	   Cluster Service
	*/
	ClusterService *string

	/* Limit.

	   Limit returned accounts
	*/
	Limit *int64

	/* Locations.

	   Cloud location
	*/
	Locations []string

	/* Offset.

	   Offset returned accounts
	*/
	Offset *int64

	/* Status.

	   Cluster Status
	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClustersParams) WithDefaults() *GetClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get clusters params
func (o *GetClustersParams) WithTimeout(timeout time.Duration) *GetClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get clusters params
func (o *GetClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get clusters params
func (o *GetClustersParams) WithContext(ctx context.Context) *GetClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get clusters params
func (o *GetClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get clusters params
func (o *GetClustersParams) WithHTTPClient(client *http.Client) *GetClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get clusters params
func (o *GetClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountIds adds the accountIds to the get clusters params
func (o *GetClustersParams) WithAccountIds(accountIds []string) *GetClustersParams {
	o.SetAccountIds(accountIds)
	return o
}

// SetAccountIds adds the accountIds to the get clusters params
func (o *GetClustersParams) SetAccountIds(accountIds []string) {
	o.AccountIds = accountIds
}

// WithClusterNames adds the clusterNames to the get clusters params
func (o *GetClustersParams) WithClusterNames(clusterNames []string) *GetClustersParams {
	o.SetClusterNames(clusterNames)
	return o
}

// SetClusterNames adds the clusterNames to the get clusters params
func (o *GetClustersParams) SetClusterNames(clusterNames []string) {
	o.ClusterNames = clusterNames
}

// WithClusterService adds the clusterService to the get clusters params
func (o *GetClustersParams) WithClusterService(clusterService *string) *GetClustersParams {
	o.SetClusterService(clusterService)
	return o
}

// SetClusterService adds the clusterService to the get clusters params
func (o *GetClustersParams) SetClusterService(clusterService *string) {
	o.ClusterService = clusterService
}

// WithLimit adds the limit to the get clusters params
func (o *GetClustersParams) WithLimit(limit *int64) *GetClustersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get clusters params
func (o *GetClustersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLocations adds the locations to the get clusters params
func (o *GetClustersParams) WithLocations(locations []string) *GetClustersParams {
	o.SetLocations(locations)
	return o
}

// SetLocations adds the locations to the get clusters params
func (o *GetClustersParams) SetLocations(locations []string) {
	o.Locations = locations
}

// WithOffset adds the offset to the get clusters params
func (o *GetClustersParams) WithOffset(offset *int64) *GetClustersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get clusters params
func (o *GetClustersParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithStatus adds the status to the get clusters params
func (o *GetClustersParams) WithStatus(status []string) *GetClustersParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get clusters params
func (o *GetClustersParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountIds != nil {

		// binding items for account_ids
		joinedAccountIds := o.bindParamAccountIds(reg)

		// query array param account_ids
		if err := r.SetQueryParam("account_ids", joinedAccountIds...); err != nil {
			return err
		}
	}

	if o.ClusterNames != nil {

		// binding items for cluster_names
		joinedClusterNames := o.bindParamClusterNames(reg)

		// query array param cluster_names
		if err := r.SetQueryParam("cluster_names", joinedClusterNames...); err != nil {
			return err
		}
	}

	if o.ClusterService != nil {

		// query param cluster_service
		var qrClusterService string

		if o.ClusterService != nil {
			qrClusterService = *o.ClusterService
		}
		qClusterService := qrClusterService
		if qClusterService != "" {

			if err := r.SetQueryParam("cluster_service", qClusterService); err != nil {
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

	if o.Locations != nil {

		// binding items for locations
		joinedLocations := o.bindParamLocations(reg)

		// query array param locations
		if err := r.SetQueryParam("locations", joinedLocations...); err != nil {
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

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetClusters binds the parameter account_ids
func (o *GetClustersParams) bindParamAccountIds(formats strfmt.Registry) []string {
	accountIdsIR := o.AccountIds

	var accountIdsIC []string
	for _, accountIdsIIR := range accountIdsIR { // explode []string

		accountIdsIIV := accountIdsIIR // string as string
		accountIdsIC = append(accountIdsIC, accountIdsIIV)
	}

	// items.CollectionFormat: "csv"
	accountIdsIS := swag.JoinByFormat(accountIdsIC, "csv")

	return accountIdsIS
}

// bindParamGetClusters binds the parameter cluster_names
func (o *GetClustersParams) bindParamClusterNames(formats strfmt.Registry) []string {
	clusterNamesIR := o.ClusterNames

	var clusterNamesIC []string
	for _, clusterNamesIIR := range clusterNamesIR { // explode []string

		clusterNamesIIV := clusterNamesIIR // string as string
		clusterNamesIC = append(clusterNamesIC, clusterNamesIIV)
	}

	// items.CollectionFormat: "csv"
	clusterNamesIS := swag.JoinByFormat(clusterNamesIC, "csv")

	return clusterNamesIS
}

// bindParamGetClusters binds the parameter locations
func (o *GetClustersParams) bindParamLocations(formats strfmt.Registry) []string {
	locationsIR := o.Locations

	var locationsIC []string
	for _, locationsIIR := range locationsIR { // explode []string

		locationsIIV := locationsIIR // string as string
		locationsIC = append(locationsIC, locationsIIV)
	}

	// items.CollectionFormat: "csv"
	locationsIS := swag.JoinByFormat(locationsIC, "csv")

	return locationsIS
}

// bindParamGetClusters binds the parameter status
func (o *GetClustersParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []string

		statusIIV := statusIIR // string as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: "csv"
	statusIS := swag.JoinByFormat(statusIC, "csv")

	return statusIS
}
