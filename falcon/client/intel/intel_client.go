// Code generated by go-swagger; DO NOT EDIT.

package intel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new intel API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for intel API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetIntelActorEntities(params *GetIntelActorEntitiesParams, opts ...ClientOption) (*GetIntelActorEntitiesOK, error)

	GetIntelIndicatorEntities(params *GetIntelIndicatorEntitiesParams, opts ...ClientOption) (*GetIntelIndicatorEntitiesOK, error)

	GetIntelReportEntities(params *GetIntelReportEntitiesParams, opts ...ClientOption) (*GetIntelReportEntitiesOK, error)

	GetIntelReportPDF(params *GetIntelReportPDFParams, writer io.Writer, opts ...ClientOption) (*GetIntelReportPDFOK, error)

	GetIntelRuleEntities(params *GetIntelRuleEntitiesParams, opts ...ClientOption) (*GetIntelRuleEntitiesOK, error)

	GetIntelRuleFile(params *GetIntelRuleFileParams, writer io.Writer, opts ...ClientOption) (*GetIntelRuleFileOK, error)

	GetLatestIntelRuleFile(params *GetLatestIntelRuleFileParams, writer io.Writer, opts ...ClientOption) (*GetLatestIntelRuleFileOK, error)

	GetMitreReport(params *GetMitreReportParams, opts ...ClientOption) (*GetMitreReportOK, error)

	GetVulnerabilities(params *GetVulnerabilitiesParams, opts ...ClientOption) (*GetVulnerabilitiesOK, error)

	PostMitreAttacks(params *PostMitreAttacksParams, opts ...ClientOption) (*PostMitreAttacksOK, error)

	QueryIntelActorEntities(params *QueryIntelActorEntitiesParams, opts ...ClientOption) (*QueryIntelActorEntitiesOK, error)

	QueryIntelActorIds(params *QueryIntelActorIdsParams, opts ...ClientOption) (*QueryIntelActorIdsOK, error)

	QueryIntelIndicatorEntities(params *QueryIntelIndicatorEntitiesParams, opts ...ClientOption) (*QueryIntelIndicatorEntitiesOK, error)

	QueryIntelIndicatorIds(params *QueryIntelIndicatorIdsParams, opts ...ClientOption) (*QueryIntelIndicatorIdsOK, error)

	QueryIntelReportEntities(params *QueryIntelReportEntitiesParams, opts ...ClientOption) (*QueryIntelReportEntitiesOK, error)

	QueryIntelReportIds(params *QueryIntelReportIdsParams, opts ...ClientOption) (*QueryIntelReportIdsOK, error)

	QueryIntelRuleIds(params *QueryIntelRuleIdsParams, opts ...ClientOption) (*QueryIntelRuleIdsOK, error)

	QueryMitreAttacks(params *QueryMitreAttacksParams, opts ...ClientOption) (*QueryMitreAttacksOK, error)

	QueryVulnerabilities(params *QueryVulnerabilitiesParams, opts ...ClientOption) (*QueryVulnerabilitiesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetIntelActorEntities retrieves specific actors using their actor i ds
*/
func (a *Client) GetIntelActorEntities(params *GetIntelActorEntitiesParams, opts ...ClientOption) (*GetIntelActorEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntelActorEntitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIntelActorEntities",
		Method:             "GET",
		PathPattern:        "/intel/entities/actors/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIntelActorEntitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntelActorEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetIntelActorEntities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetIntelIndicatorEntities retrieves specific indicators using their indicator i ds
*/
func (a *Client) GetIntelIndicatorEntities(params *GetIntelIndicatorEntitiesParams, opts ...ClientOption) (*GetIntelIndicatorEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntelIndicatorEntitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIntelIndicatorEntities",
		Method:             "POST",
		PathPattern:        "/intel/entities/indicators/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIntelIndicatorEntitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntelIndicatorEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetIntelIndicatorEntities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetIntelReportEntities retrieves specific reports using their report i ds
*/
func (a *Client) GetIntelReportEntities(params *GetIntelReportEntitiesParams, opts ...ClientOption) (*GetIntelReportEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntelReportEntitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIntelReportEntities",
		Method:             "GET",
		PathPattern:        "/intel/entities/reports/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIntelReportEntitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntelReportEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetIntelReportEntities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetIntelReportPDF returns a report p d f attachment
*/
func (a *Client) GetIntelReportPDF(params *GetIntelReportPDFParams, writer io.Writer, opts ...ClientOption) (*GetIntelReportPDFOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntelReportPDFParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIntelReportPDF",
		Method:             "GET",
		PathPattern:        "/intel/entities/report-files/v1",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream", "application/pdf"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIntelReportPDFReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntelReportPDFOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetIntelReportPDF: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetIntelRuleEntities retrieves details for rule sets for the specified ids
*/
func (a *Client) GetIntelRuleEntities(params *GetIntelRuleEntitiesParams, opts ...ClientOption) (*GetIntelRuleEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntelRuleEntitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIntelRuleEntities",
		Method:             "GET",
		PathPattern:        "/intel/entities/rules/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIntelRuleEntitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntelRuleEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetIntelRuleEntities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetIntelRuleFile downloads earlier rule sets
*/
func (a *Client) GetIntelRuleFile(params *GetIntelRuleFileParams, writer io.Writer, opts ...ClientOption) (*GetIntelRuleFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntelRuleFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIntelRuleFile",
		Method:             "GET",
		PathPattern:        "/intel/entities/rules-files/v1",
		ProducesMediaTypes: []string{"*/*", "application/gzip", "application/json", "application/octet-stream", "application/zip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIntelRuleFileReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntelRuleFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetIntelRuleFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLatestIntelRuleFile downloads the latest rule set
*/
func (a *Client) GetLatestIntelRuleFile(params *GetLatestIntelRuleFileParams, writer io.Writer, opts ...ClientOption) (*GetLatestIntelRuleFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLatestIntelRuleFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLatestIntelRuleFile",
		Method:             "GET",
		PathPattern:        "/intel/entities/rules-latest-files/v1",
		ProducesMediaTypes: []string{"*/*", "application/gzip", "application/json", "application/octet-stream", "application/zip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLatestIntelRuleFileReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLatestIntelRuleFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLatestIntelRuleFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMitreReport exports mitre a t t and c k information for a given actor
*/
func (a *Client) GetMitreReport(params *GetMitreReportParams, opts ...ClientOption) (*GetMitreReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMitreReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetMitreReport",
		Method:             "GET",
		PathPattern:        "/intel/entities/mitre-reports/v1",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMitreReportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMitreReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMitreReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVulnerabilities gets vulnerabilities
*/
func (a *Client) GetVulnerabilities(params *GetVulnerabilitiesParams, opts ...ClientOption) (*GetVulnerabilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVulnerabilitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVulnerabilities",
		Method:             "POST",
		PathPattern:        "/intel/entities/vulnerabilities/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVulnerabilitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVulnerabilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVulnerabilities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostMitreAttacks retrieves report and observable i ds associated with the given actor and attacks
*/
func (a *Client) PostMitreAttacks(params *PostMitreAttacksParams, opts ...ClientOption) (*PostMitreAttacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostMitreAttacksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostMitreAttacks",
		Method:             "POST",
		PathPattern:        "/intel/entities/mitre/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostMitreAttacksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostMitreAttacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostMitreAttacks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryIntelActorEntities gets info about actors that match provided f q l filters
*/
func (a *Client) QueryIntelActorEntities(params *QueryIntelActorEntitiesParams, opts ...ClientOption) (*QueryIntelActorEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryIntelActorEntitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryIntelActorEntities",
		Method:             "GET",
		PathPattern:        "/intel/combined/actors/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryIntelActorEntitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryIntelActorEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryIntelActorEntities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryIntelActorIds gets actor i ds that match provided f q l filters
*/
func (a *Client) QueryIntelActorIds(params *QueryIntelActorIdsParams, opts ...ClientOption) (*QueryIntelActorIdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryIntelActorIdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryIntelActorIds",
		Method:             "GET",
		PathPattern:        "/intel/queries/actors/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryIntelActorIdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryIntelActorIdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryIntelActorIds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryIntelIndicatorEntities gets info about indicators that match provided f q l filters
*/
func (a *Client) QueryIntelIndicatorEntities(params *QueryIntelIndicatorEntitiesParams, opts ...ClientOption) (*QueryIntelIndicatorEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryIntelIndicatorEntitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryIntelIndicatorEntities",
		Method:             "GET",
		PathPattern:        "/intel/combined/indicators/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryIntelIndicatorEntitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryIntelIndicatorEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryIntelIndicatorEntities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryIntelIndicatorIds gets indicators i ds that match provided f q l filters
*/
func (a *Client) QueryIntelIndicatorIds(params *QueryIntelIndicatorIdsParams, opts ...ClientOption) (*QueryIntelIndicatorIdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryIntelIndicatorIdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryIntelIndicatorIds",
		Method:             "GET",
		PathPattern:        "/intel/queries/indicators/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryIntelIndicatorIdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryIntelIndicatorIdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryIntelIndicatorIds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryIntelReportEntities gets info about reports that match provided f q l filters
*/
func (a *Client) QueryIntelReportEntities(params *QueryIntelReportEntitiesParams, opts ...ClientOption) (*QueryIntelReportEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryIntelReportEntitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryIntelReportEntities",
		Method:             "GET",
		PathPattern:        "/intel/combined/reports/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryIntelReportEntitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryIntelReportEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryIntelReportEntities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryIntelReportIds gets report i ds that match provided f q l filters
*/
func (a *Client) QueryIntelReportIds(params *QueryIntelReportIdsParams, opts ...ClientOption) (*QueryIntelReportIdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryIntelReportIdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryIntelReportIds",
		Method:             "GET",
		PathPattern:        "/intel/queries/reports/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryIntelReportIdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryIntelReportIdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryIntelReportIds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryIntelRuleIds searches for rule i ds that match provided filter criteria
*/
func (a *Client) QueryIntelRuleIds(params *QueryIntelRuleIdsParams, opts ...ClientOption) (*QueryIntelRuleIdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryIntelRuleIdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryIntelRuleIds",
		Method:             "GET",
		PathPattern:        "/intel/queries/rules/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryIntelRuleIdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryIntelRuleIdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryIntelRuleIds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryMitreAttacks gets m i t r e tactics and techniques for the given actor returning concatenation of id and tactic and technique ids example fancy bear t a0011 t1071
*/
func (a *Client) QueryMitreAttacks(params *QueryMitreAttacksParams, opts ...ClientOption) (*QueryMitreAttacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryMitreAttacksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryMitreAttacks",
		Method:             "GET",
		PathPattern:        "/intel/queries/mitre/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryMitreAttacksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryMitreAttacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryMitreAttacks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryVulnerabilities gets vulnerabilities i ds
*/
func (a *Client) QueryVulnerabilities(params *QueryVulnerabilitiesParams, opts ...ClientOption) (*QueryVulnerabilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryVulnerabilitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryVulnerabilities",
		Method:             "GET",
		PathPattern:        "/intel/queries/vulnerabilities/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryVulnerabilitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryVulnerabilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryVulnerabilities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
