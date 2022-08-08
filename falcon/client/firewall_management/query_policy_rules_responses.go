// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// QueryPolicyRulesReader is a Reader for the QueryPolicyRules structure.
type QueryPolicyRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryPolicyRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryPolicyRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryPolicyRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryPolicyRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryPolicyRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryPolicyRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryPolicyRulesOK creates a QueryPolicyRulesOK with default headers values
func NewQueryPolicyRulesOK() *QueryPolicyRulesOK {
	return &QueryPolicyRulesOK{}
}

/*
	QueryPolicyRulesOK describes a response with status code 200, with default header values.

OK
*/
type QueryPolicyRulesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrAPIQueryResponse
}

func (o *QueryPolicyRulesOK) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/policy-rules/v1][%d] queryPolicyRulesOK  %+v", 200, o.Payload)
}
func (o *QueryPolicyRulesOK) GetPayload() *models.FwmgrAPIQueryResponse {
	return o.Payload
}

func (o *QueryPolicyRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.FwmgrAPIQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPolicyRulesBadRequest creates a QueryPolicyRulesBadRequest with default headers values
func NewQueryPolicyRulesBadRequest() *QueryPolicyRulesBadRequest {
	return &QueryPolicyRulesBadRequest{}
}

/*
	QueryPolicyRulesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryPolicyRulesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaReplyMetaOnly
}

func (o *QueryPolicyRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/policy-rules/v1][%d] queryPolicyRulesBadRequest  %+v", 400, o.Payload)
}
func (o *QueryPolicyRulesBadRequest) GetPayload() *models.FwmgrMsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryPolicyRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.FwmgrMsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPolicyRulesForbidden creates a QueryPolicyRulesForbidden with default headers values
func NewQueryPolicyRulesForbidden() *QueryPolicyRulesForbidden {
	return &QueryPolicyRulesForbidden{}
}

/*
	QueryPolicyRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryPolicyRulesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryPolicyRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/policy-rules/v1][%d] queryPolicyRulesForbidden  %+v", 403, o.Payload)
}
func (o *QueryPolicyRulesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryPolicyRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPolicyRulesTooManyRequests creates a QueryPolicyRulesTooManyRequests with default headers values
func NewQueryPolicyRulesTooManyRequests() *QueryPolicyRulesTooManyRequests {
	return &QueryPolicyRulesTooManyRequests{}
}

/*
	QueryPolicyRulesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryPolicyRulesTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryPolicyRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/policy-rules/v1][%d] queryPolicyRulesTooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryPolicyRulesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryPolicyRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPolicyRulesDefault creates a QueryPolicyRulesDefault with default headers values
func NewQueryPolicyRulesDefault(code int) *QueryPolicyRulesDefault {
	return &QueryPolicyRulesDefault{
		_statusCode: code,
	}
}

/*
	QueryPolicyRulesDefault describes a response with status code -1, with default header values.

OK
*/
type QueryPolicyRulesDefault struct {
	_statusCode int

	Payload *models.FwmgrAPIQueryResponse
}

// Code gets the status code for the query policy rules default response
func (o *QueryPolicyRulesDefault) Code() int {
	return o._statusCode
}

func (o *QueryPolicyRulesDefault) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/policy-rules/v1][%d] query-policy-rules default  %+v", o._statusCode, o.Payload)
}
func (o *QueryPolicyRulesDefault) GetPayload() *models.FwmgrAPIQueryResponse {
	return o.Payload
}

func (o *QueryPolicyRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FwmgrAPIQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
