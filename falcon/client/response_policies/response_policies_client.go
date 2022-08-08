// Code generated by go-swagger; DO NOT EDIT.

package response_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new response policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for response policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRTResponsePolicies(params *CreateRTResponsePoliciesParams, opts ...ClientOption) (*CreateRTResponsePoliciesCreated, error)

	DeleteRTResponsePolicies(params *DeleteRTResponsePoliciesParams, opts ...ClientOption) (*DeleteRTResponsePoliciesOK, error)

	GetRTResponsePolicies(params *GetRTResponsePoliciesParams, opts ...ClientOption) (*GetRTResponsePoliciesOK, error)

	PerformRTResponsePoliciesAction(params *PerformRTResponsePoliciesActionParams, opts ...ClientOption) (*PerformRTResponsePoliciesActionOK, error)

	QueryCombinedRTResponsePolicies(params *QueryCombinedRTResponsePoliciesParams, opts ...ClientOption) (*QueryCombinedRTResponsePoliciesOK, error)

	QueryCombinedRTResponsePolicyMembers(params *QueryCombinedRTResponsePolicyMembersParams, opts ...ClientOption) (*QueryCombinedRTResponsePolicyMembersOK, error)

	QueryRTResponsePolicies(params *QueryRTResponsePoliciesParams, opts ...ClientOption) (*QueryRTResponsePoliciesOK, error)

	QueryRTResponsePolicyMembers(params *QueryRTResponsePolicyMembersParams, opts ...ClientOption) (*QueryRTResponsePolicyMembersOK, error)

	SetRTResponsePoliciesPrecedence(params *SetRTResponsePoliciesPrecedenceParams, opts ...ClientOption) (*SetRTResponsePoliciesPrecedenceOK, error)

	UpdateRTResponsePolicies(params *UpdateRTResponsePoliciesParams, opts ...ClientOption) (*UpdateRTResponsePoliciesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateRTResponsePolicies creates response policies by specifying details about the policy to create
*/
func (a *Client) CreateRTResponsePolicies(params *CreateRTResponsePoliciesParams, opts ...ClientOption) (*CreateRTResponsePoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRTResponsePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRTResponsePolicies",
		Method:             "POST",
		PathPattern:        "/policy/entities/response/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRTResponsePoliciesReader{formats: a.formats},
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
	success, ok := result.(*CreateRTResponsePoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createRTResponsePolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteRTResponsePolicies deletes a set of response policies by specifying their i ds
*/
func (a *Client) DeleteRTResponsePolicies(params *DeleteRTResponsePoliciesParams, opts ...ClientOption) (*DeleteRTResponsePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRTResponsePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRTResponsePolicies",
		Method:             "DELETE",
		PathPattern:        "/policy/entities/response/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRTResponsePoliciesReader{formats: a.formats},
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
	success, ok := result.(*DeleteRTResponsePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRTResponsePoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetRTResponsePolicies retrieves a set of response policies by specifying their i ds
*/
func (a *Client) GetRTResponsePolicies(params *GetRTResponsePoliciesParams, opts ...ClientOption) (*GetRTResponsePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRTResponsePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRTResponsePolicies",
		Method:             "GET",
		PathPattern:        "/policy/entities/response/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRTResponsePoliciesReader{formats: a.formats},
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
	success, ok := result.(*GetRTResponsePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRTResponsePoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PerformRTResponsePoliciesAction performs the specified action on the response policies specified in the request
*/
func (a *Client) PerformRTResponsePoliciesAction(params *PerformRTResponsePoliciesActionParams, opts ...ClientOption) (*PerformRTResponsePoliciesActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformRTResponsePoliciesActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "performRTResponsePoliciesAction",
		Method:             "POST",
		PathPattern:        "/policy/entities/response-actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PerformRTResponsePoliciesActionReader{formats: a.formats},
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
	success, ok := result.(*PerformRTResponsePoliciesActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PerformRTResponsePoliciesActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryCombinedRTResponsePolicies searches for response policies in your environment by providing an f q l filter and paging details returns a set of response policies which match the filter criteria
*/
func (a *Client) QueryCombinedRTResponsePolicies(params *QueryCombinedRTResponsePoliciesParams, opts ...ClientOption) (*QueryCombinedRTResponsePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedRTResponsePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryCombinedRTResponsePolicies",
		Method:             "GET",
		PathPattern:        "/policy/combined/response/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedRTResponsePoliciesReader{formats: a.formats},
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
	success, ok := result.(*QueryCombinedRTResponsePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedRTResponsePoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryCombinedRTResponsePolicyMembers searches for members of a response policy in your environment by providing an f q l filter and paging details returns a set of host details which match the filter criteria
*/
func (a *Client) QueryCombinedRTResponsePolicyMembers(params *QueryCombinedRTResponsePolicyMembersParams, opts ...ClientOption) (*QueryCombinedRTResponsePolicyMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedRTResponsePolicyMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryCombinedRTResponsePolicyMembers",
		Method:             "GET",
		PathPattern:        "/policy/combined/response-members/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedRTResponsePolicyMembersReader{formats: a.formats},
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
	success, ok := result.(*QueryCombinedRTResponsePolicyMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedRTResponsePolicyMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryRTResponsePolicies searches for response policies in your environment by providing an f q l filter with sort and or paging details this returns a set of response policy i ds that match the given criteria
*/
func (a *Client) QueryRTResponsePolicies(params *QueryRTResponsePoliciesParams, opts ...ClientOption) (*QueryRTResponsePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryRTResponsePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryRTResponsePolicies",
		Method:             "GET",
		PathPattern:        "/policy/queries/response/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryRTResponsePoliciesReader{formats: a.formats},
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
	success, ok := result.(*QueryRTResponsePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryRTResponsePoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryRTResponsePolicyMembers searches for members of a response policy in your environment by providing an f q l filter and paging details returns a set of agent i ds which match the filter criteria
*/
func (a *Client) QueryRTResponsePolicyMembers(params *QueryRTResponsePolicyMembersParams, opts ...ClientOption) (*QueryRTResponsePolicyMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryRTResponsePolicyMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryRTResponsePolicyMembers",
		Method:             "GET",
		PathPattern:        "/policy/queries/response-members/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryRTResponsePolicyMembersReader{formats: a.formats},
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
	success, ok := result.(*QueryRTResponsePolicyMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryRTResponsePolicyMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SetRTResponsePoliciesPrecedence sets the precedence of response policies based on the order of i ds specified in the request the first ID specified will have the highest precedence and the last ID specified will have the lowest you must specify all non default policies for a platform when updating precedence
*/
func (a *Client) SetRTResponsePoliciesPrecedence(params *SetRTResponsePoliciesPrecedenceParams, opts ...ClientOption) (*SetRTResponsePoliciesPrecedenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetRTResponsePoliciesPrecedenceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setRTResponsePoliciesPrecedence",
		Method:             "POST",
		PathPattern:        "/policy/entities/response-precedence/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetRTResponsePoliciesPrecedenceReader{formats: a.formats},
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
	success, ok := result.(*SetRTResponsePoliciesPrecedenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SetRTResponsePoliciesPrecedenceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateRTResponsePolicies updates response policies by specifying the ID of the policy and details to update
*/
func (a *Client) UpdateRTResponsePolicies(params *UpdateRTResponsePoliciesParams, opts ...ClientOption) (*UpdateRTResponsePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRTResponsePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateRTResponsePolicies",
		Method:             "PATCH",
		PathPattern:        "/policy/entities/response/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRTResponsePoliciesReader{formats: a.formats},
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
	success, ok := result.(*UpdateRTResponsePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRTResponsePoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
