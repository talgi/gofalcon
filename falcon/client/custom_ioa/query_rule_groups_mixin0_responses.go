// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// QueryRuleGroupsMixin0Reader is a Reader for the QueryRuleGroupsMixin0 structure.
type QueryRuleGroupsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryRuleGroupsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryRuleGroupsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryRuleGroupsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryRuleGroupsMixin0NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryRuleGroupsMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryRuleGroupsMixin0Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryRuleGroupsMixin0OK creates a QueryRuleGroupsMixin0OK with default headers values
func NewQueryRuleGroupsMixin0OK() *QueryRuleGroupsMixin0OK {
	return &QueryRuleGroupsMixin0OK{}
}

/*
QueryRuleGroupsMixin0OK describes a response with status code 200, with default header values.

OK
*/
type QueryRuleGroupsMixin0OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query rule groups mixin0 o k response has a 2xx status code
func (o *QueryRuleGroupsMixin0OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query rule groups mixin0 o k response has a 3xx status code
func (o *QueryRuleGroupsMixin0OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule groups mixin0 o k response has a 4xx status code
func (o *QueryRuleGroupsMixin0OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query rule groups mixin0 o k response has a 5xx status code
func (o *QueryRuleGroupsMixin0OK) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule groups mixin0 o k response a status code equal to that given
func (o *QueryRuleGroupsMixin0OK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryRuleGroupsMixin0OK) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups/v1][%d] queryRuleGroupsMixin0OK  %+v", 200, o.Payload)
}

func (o *QueryRuleGroupsMixin0OK) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups/v1][%d] queryRuleGroupsMixin0OK  %+v", 200, o.Payload)
}

func (o *QueryRuleGroupsMixin0OK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryRuleGroupsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRuleGroupsMixin0Forbidden creates a QueryRuleGroupsMixin0Forbidden with default headers values
func NewQueryRuleGroupsMixin0Forbidden() *QueryRuleGroupsMixin0Forbidden {
	return &QueryRuleGroupsMixin0Forbidden{}
}

/*
QueryRuleGroupsMixin0Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryRuleGroupsMixin0Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query rule groups mixin0 forbidden response has a 2xx status code
func (o *QueryRuleGroupsMixin0Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query rule groups mixin0 forbidden response has a 3xx status code
func (o *QueryRuleGroupsMixin0Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule groups mixin0 forbidden response has a 4xx status code
func (o *QueryRuleGroupsMixin0Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query rule groups mixin0 forbidden response has a 5xx status code
func (o *QueryRuleGroupsMixin0Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule groups mixin0 forbidden response a status code equal to that given
func (o *QueryRuleGroupsMixin0Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *QueryRuleGroupsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups/v1][%d] queryRuleGroupsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *QueryRuleGroupsMixin0Forbidden) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups/v1][%d] queryRuleGroupsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *QueryRuleGroupsMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryRuleGroupsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryRuleGroupsMixin0NotFound creates a QueryRuleGroupsMixin0NotFound with default headers values
func NewQueryRuleGroupsMixin0NotFound() *QueryRuleGroupsMixin0NotFound {
	return &QueryRuleGroupsMixin0NotFound{}
}

/*
QueryRuleGroupsMixin0NotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryRuleGroupsMixin0NotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query rule groups mixin0 not found response has a 2xx status code
func (o *QueryRuleGroupsMixin0NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query rule groups mixin0 not found response has a 3xx status code
func (o *QueryRuleGroupsMixin0NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule groups mixin0 not found response has a 4xx status code
func (o *QueryRuleGroupsMixin0NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query rule groups mixin0 not found response has a 5xx status code
func (o *QueryRuleGroupsMixin0NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule groups mixin0 not found response a status code equal to that given
func (o *QueryRuleGroupsMixin0NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *QueryRuleGroupsMixin0NotFound) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups/v1][%d] queryRuleGroupsMixin0NotFound  %+v", 404, o.Payload)
}

func (o *QueryRuleGroupsMixin0NotFound) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups/v1][%d] queryRuleGroupsMixin0NotFound  %+v", 404, o.Payload)
}

func (o *QueryRuleGroupsMixin0NotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryRuleGroupsMixin0NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryRuleGroupsMixin0TooManyRequests creates a QueryRuleGroupsMixin0TooManyRequests with default headers values
func NewQueryRuleGroupsMixin0TooManyRequests() *QueryRuleGroupsMixin0TooManyRequests {
	return &QueryRuleGroupsMixin0TooManyRequests{}
}

/*
QueryRuleGroupsMixin0TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryRuleGroupsMixin0TooManyRequests struct {

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

// IsSuccess returns true when this query rule groups mixin0 too many requests response has a 2xx status code
func (o *QueryRuleGroupsMixin0TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query rule groups mixin0 too many requests response has a 3xx status code
func (o *QueryRuleGroupsMixin0TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule groups mixin0 too many requests response has a 4xx status code
func (o *QueryRuleGroupsMixin0TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query rule groups mixin0 too many requests response has a 5xx status code
func (o *QueryRuleGroupsMixin0TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule groups mixin0 too many requests response a status code equal to that given
func (o *QueryRuleGroupsMixin0TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *QueryRuleGroupsMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups/v1][%d] queryRuleGroupsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryRuleGroupsMixin0TooManyRequests) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups/v1][%d] queryRuleGroupsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryRuleGroupsMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryRuleGroupsMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryRuleGroupsMixin0Default creates a QueryRuleGroupsMixin0Default with default headers values
func NewQueryRuleGroupsMixin0Default(code int) *QueryRuleGroupsMixin0Default {
	return &QueryRuleGroupsMixin0Default{
		_statusCode: code,
	}
}

/*
QueryRuleGroupsMixin0Default describes a response with status code -1, with default header values.

OK
*/
type QueryRuleGroupsMixin0Default struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query rule groups mixin0 default response
func (o *QueryRuleGroupsMixin0Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this query rule groups mixin0 default response has a 2xx status code
func (o *QueryRuleGroupsMixin0Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query rule groups mixin0 default response has a 3xx status code
func (o *QueryRuleGroupsMixin0Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query rule groups mixin0 default response has a 4xx status code
func (o *QueryRuleGroupsMixin0Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query rule groups mixin0 default response has a 5xx status code
func (o *QueryRuleGroupsMixin0Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query rule groups mixin0 default response a status code equal to that given
func (o *QueryRuleGroupsMixin0Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QueryRuleGroupsMixin0Default) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups/v1][%d] query-rule-groupsMixin0 default  %+v", o._statusCode, o.Payload)
}

func (o *QueryRuleGroupsMixin0Default) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-groups/v1][%d] query-rule-groupsMixin0 default  %+v", o._statusCode, o.Payload)
}

func (o *QueryRuleGroupsMixin0Default) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryRuleGroupsMixin0Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
