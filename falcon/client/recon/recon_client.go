// Code generated by go-swagger; DO NOT EDIT.

package recon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new recon API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for recon API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AggregateNotificationsV1(params *AggregateNotificationsV1Params, opts ...ClientOption) (*AggregateNotificationsV1OK, error)

	CreateActionsV1(params *CreateActionsV1Params, opts ...ClientOption) (*CreateActionsV1OK, error)

	CreateRulesV1(params *CreateRulesV1Params, opts ...ClientOption) (*CreateRulesV1OK, error)

	DeleteActionV1(params *DeleteActionV1Params, opts ...ClientOption) (*DeleteActionV1OK, error)

	DeleteNotificationsV1(params *DeleteNotificationsV1Params, opts ...ClientOption) (*DeleteNotificationsV1OK, error)

	DeleteRulesV1(params *DeleteRulesV1Params, opts ...ClientOption) (*DeleteRulesV1OK, error)

	GetActionsV1(params *GetActionsV1Params, opts ...ClientOption) (*GetActionsV1OK, error)

	GetNotificationsDetailedTranslatedV1(params *GetNotificationsDetailedTranslatedV1Params, opts ...ClientOption) (*GetNotificationsDetailedTranslatedV1OK, error)

	GetNotificationsDetailedV1(params *GetNotificationsDetailedV1Params, opts ...ClientOption) (*GetNotificationsDetailedV1OK, error)

	GetNotificationsTranslatedV1(params *GetNotificationsTranslatedV1Params, opts ...ClientOption) (*GetNotificationsTranslatedV1OK, error)

	GetNotificationsV1(params *GetNotificationsV1Params, opts ...ClientOption) (*GetNotificationsV1OK, error)

	GetRulesV1(params *GetRulesV1Params, opts ...ClientOption) (*GetRulesV1OK, error)

	PreviewRuleV1(params *PreviewRuleV1Params, opts ...ClientOption) (*PreviewRuleV1OK, error)

	QueryActionsV1(params *QueryActionsV1Params, opts ...ClientOption) (*QueryActionsV1OK, error)

	QueryNotificationsV1(params *QueryNotificationsV1Params, opts ...ClientOption) (*QueryNotificationsV1OK, error)

	QueryRulesV1(params *QueryRulesV1Params, opts ...ClientOption) (*QueryRulesV1OK, error)

	UpdateActionV1(params *UpdateActionV1Params, opts ...ClientOption) (*UpdateActionV1OK, error)

	UpdateNotificationsV1(params *UpdateNotificationsV1Params, opts ...ClientOption) (*UpdateNotificationsV1OK, error)

	UpdateRulesV1(params *UpdateRulesV1Params, opts ...ClientOption) (*UpdateRulesV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AggregateNotificationsV1 gets notification aggregates as specified via JSON in request body
*/
func (a *Client) AggregateNotificationsV1(params *AggregateNotificationsV1Params, opts ...ClientOption) (*AggregateNotificationsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateNotificationsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "AggregateNotificationsV1",
		Method:             "POST",
		PathPattern:        "/recon/aggregates/notifications/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateNotificationsV1Reader{formats: a.formats},
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
	success, ok := result.(*AggregateNotificationsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AggregateNotificationsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateActionsV1 creates actions for a monitoring rule accepts a list of actions that will be attached to the monitoring rule
*/
func (a *Client) CreateActionsV1(params *CreateActionsV1Params, opts ...ClientOption) (*CreateActionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateActionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateActionsV1",
		Method:             "POST",
		PathPattern:        "/recon/entities/actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateActionsV1Reader{formats: a.formats},
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
	success, ok := result.(*CreateActionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateActionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateRulesV1 creates monitoring rules
*/
func (a *Client) CreateRulesV1(params *CreateRulesV1Params, opts ...ClientOption) (*CreateRulesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRulesV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRulesV1",
		Method:             "POST",
		PathPattern:        "/recon/entities/rules/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRulesV1Reader{formats: a.formats},
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
	success, ok := result.(*CreateRulesV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRulesV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteActionV1 deletes an action from a monitoring rule based on the action ID
*/
func (a *Client) DeleteActionV1(params *DeleteActionV1Params, opts ...ClientOption) (*DeleteActionV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteActionV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteActionV1",
		Method:             "DELETE",
		PathPattern:        "/recon/entities/actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteActionV1Reader{formats: a.formats},
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
	success, ok := result.(*DeleteActionV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteActionV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteNotificationsV1 deletes notifications based on i ds notifications cannot be recovered after they are deleted
*/
func (a *Client) DeleteNotificationsV1(params *DeleteNotificationsV1Params, opts ...ClientOption) (*DeleteNotificationsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNotificationsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteNotificationsV1",
		Method:             "DELETE",
		PathPattern:        "/recon/entities/notifications/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNotificationsV1Reader{formats: a.formats},
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
	success, ok := result.(*DeleteNotificationsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNotificationsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteRulesV1 deletes monitoring rules
*/
func (a *Client) DeleteRulesV1(params *DeleteRulesV1Params, opts ...ClientOption) (*DeleteRulesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRulesV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteRulesV1",
		Method:             "DELETE",
		PathPattern:        "/recon/entities/rules/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRulesV1Reader{formats: a.formats},
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
	success, ok := result.(*DeleteRulesV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRulesV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetActionsV1 gets actions based on their i ds i ds can be retrieved using the g e t queries actions v1 endpoint
*/
func (a *Client) GetActionsV1(params *GetActionsV1Params, opts ...ClientOption) (*GetActionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetActionsV1",
		Method:             "GET",
		PathPattern:        "/recon/entities/actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActionsV1Reader{formats: a.formats},
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
	success, ok := result.(*GetActionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetActionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNotificationsDetailedTranslatedV1 gets detailed notifications based on their i ds these include the raw intelligence content that generated the match this endpoint will return translated notification content the only target language available is english a single notification can be translated per request
*/
func (a *Client) GetNotificationsDetailedTranslatedV1(params *GetNotificationsDetailedTranslatedV1Params, opts ...ClientOption) (*GetNotificationsDetailedTranslatedV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNotificationsDetailedTranslatedV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNotificationsDetailedTranslatedV1",
		Method:             "GET",
		PathPattern:        "/recon/entities/notifications-detailed-translated/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNotificationsDetailedTranslatedV1Reader{formats: a.formats},
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
	success, ok := result.(*GetNotificationsDetailedTranslatedV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNotificationsDetailedTranslatedV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNotificationsDetailedV1 gets detailed notifications based on their i ds these include the raw intelligence content that generated the match
*/
func (a *Client) GetNotificationsDetailedV1(params *GetNotificationsDetailedV1Params, opts ...ClientOption) (*GetNotificationsDetailedV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNotificationsDetailedV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNotificationsDetailedV1",
		Method:             "GET",
		PathPattern:        "/recon/entities/notifications-detailed/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNotificationsDetailedV1Reader{formats: a.formats},
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
	success, ok := result.(*GetNotificationsDetailedV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNotificationsDetailedV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNotificationsTranslatedV1 gets notifications based on their i ds i ds can be retrieved using the g e t queries notifications v1 endpoint this endpoint will return translated notification content the only target language available is english
*/
func (a *Client) GetNotificationsTranslatedV1(params *GetNotificationsTranslatedV1Params, opts ...ClientOption) (*GetNotificationsTranslatedV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNotificationsTranslatedV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNotificationsTranslatedV1",
		Method:             "GET",
		PathPattern:        "/recon/entities/notifications-translated/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNotificationsTranslatedV1Reader{formats: a.formats},
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
	success, ok := result.(*GetNotificationsTranslatedV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNotificationsTranslatedV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNotificationsV1 gets notifications based on their i ds i ds can be retrieved using the g e t queries notifications v1 endpoint
*/
func (a *Client) GetNotificationsV1(params *GetNotificationsV1Params, opts ...ClientOption) (*GetNotificationsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNotificationsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNotificationsV1",
		Method:             "GET",
		PathPattern:        "/recon/entities/notifications/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNotificationsV1Reader{formats: a.formats},
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
	success, ok := result.(*GetNotificationsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNotificationsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetRulesV1 gets monitoring rules based on their i ds i ds can be retrieved using the g e t queries rules v1 endpoint
*/
func (a *Client) GetRulesV1(params *GetRulesV1Params, opts ...ClientOption) (*GetRulesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRulesV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRulesV1",
		Method:             "GET",
		PathPattern:        "/recon/entities/rules/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRulesV1Reader{formats: a.formats},
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
	success, ok := result.(*GetRulesV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRulesV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PreviewRuleV1 previews rules notification count and distribution this will return aggregations on channel count site
*/
func (a *Client) PreviewRuleV1(params *PreviewRuleV1Params, opts ...ClientOption) (*PreviewRuleV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPreviewRuleV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PreviewRuleV1",
		Method:             "POST",
		PathPattern:        "/recon/aggregates/rules-preview/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PreviewRuleV1Reader{formats: a.formats},
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
	success, ok := result.(*PreviewRuleV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PreviewRuleV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryActionsV1 queries actions based on provided criteria use the i ds from this response to get the action entities on g e t entities actions v1
*/
func (a *Client) QueryActionsV1(params *QueryActionsV1Params, opts ...ClientOption) (*QueryActionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryActionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryActionsV1",
		Method:             "GET",
		PathPattern:        "/recon/queries/actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryActionsV1Reader{formats: a.formats},
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
	success, ok := result.(*QueryActionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryActionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryNotificationsV1 queries notifications based on provided criteria use the i ds from this response to get the notification entities on g e t entities notifications v1 or g e t entities notifications detailed v1
*/
func (a *Client) QueryNotificationsV1(params *QueryNotificationsV1Params, opts ...ClientOption) (*QueryNotificationsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryNotificationsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryNotificationsV1",
		Method:             "GET",
		PathPattern:        "/recon/queries/notifications/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryNotificationsV1Reader{formats: a.formats},
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
	success, ok := result.(*QueryNotificationsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryNotificationsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryRulesV1 queries monitoring rules based on provided criteria use the i ds from this response to fetch the rules on entities rules v1
*/
func (a *Client) QueryRulesV1(params *QueryRulesV1Params, opts ...ClientOption) (*QueryRulesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryRulesV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryRulesV1",
		Method:             "GET",
		PathPattern:        "/recon/queries/rules/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryRulesV1Reader{formats: a.formats},
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
	success, ok := result.(*QueryRulesV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryRulesV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateActionV1 updates an action for a monitoring rule
*/
func (a *Client) UpdateActionV1(params *UpdateActionV1Params, opts ...ClientOption) (*UpdateActionV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateActionV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateActionV1",
		Method:             "PATCH",
		PathPattern:        "/recon/entities/actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateActionV1Reader{formats: a.formats},
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
	success, ok := result.(*UpdateActionV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateActionV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateNotificationsV1 updates notification status or assignee accepts bulk requests
*/
func (a *Client) UpdateNotificationsV1(params *UpdateNotificationsV1Params, opts ...ClientOption) (*UpdateNotificationsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNotificationsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateNotificationsV1",
		Method:             "PATCH",
		PathPattern:        "/recon/entities/notifications/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNotificationsV1Reader{formats: a.formats},
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
	success, ok := result.(*UpdateNotificationsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateNotificationsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateRulesV1 updates monitoring rules
*/
func (a *Client) UpdateRulesV1(params *UpdateRulesV1Params, opts ...ClientOption) (*UpdateRulesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRulesV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateRulesV1",
		Method:             "PATCH",
		PathPattern:        "/recon/entities/rules/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRulesV1Reader{formats: a.formats},
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
	success, ok := result.(*UpdateRulesV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRulesV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
