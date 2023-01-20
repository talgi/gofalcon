// Code generated by go-swagger; DO NOT EDIT.

package event_streams

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

// ListAvailableStreamsOAuth2Reader is a Reader for the ListAvailableStreamsOAuth2 structure.
type ListAvailableStreamsOAuth2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAvailableStreamsOAuth2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAvailableStreamsOAuth2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAvailableStreamsOAuth2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAvailableStreamsOAuth2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListAvailableStreamsOAuth2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAvailableStreamsOAuth2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAvailableStreamsOAuth2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAvailableStreamsOAuth2OK creates a ListAvailableStreamsOAuth2OK with default headers values
func NewListAvailableStreamsOAuth2OK() *ListAvailableStreamsOAuth2OK {
	return &ListAvailableStreamsOAuth2OK{}
}

/*
ListAvailableStreamsOAuth2OK describes a response with status code 200, with default header values.

OK
*/
type ListAvailableStreamsOAuth2OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MainDiscoveryResponseV2
}

// IsSuccess returns true when this list available streams o auth2 o k response has a 2xx status code
func (o *ListAvailableStreamsOAuth2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list available streams o auth2 o k response has a 3xx status code
func (o *ListAvailableStreamsOAuth2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list available streams o auth2 o k response has a 4xx status code
func (o *ListAvailableStreamsOAuth2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list available streams o auth2 o k response has a 5xx status code
func (o *ListAvailableStreamsOAuth2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list available streams o auth2 o k response a status code equal to that given
func (o *ListAvailableStreamsOAuth2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list available streams o auth2 o k response
func (o *ListAvailableStreamsOAuth2OK) Code() int {
	return 200
}

func (o *ListAvailableStreamsOAuth2OK) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2OK  %+v", 200, o.Payload)
}

func (o *ListAvailableStreamsOAuth2OK) String() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2OK  %+v", 200, o.Payload)
}

func (o *ListAvailableStreamsOAuth2OK) GetPayload() *models.MainDiscoveryResponseV2 {
	return o.Payload
}

func (o *ListAvailableStreamsOAuth2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MainDiscoveryResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAvailableStreamsOAuth2BadRequest creates a ListAvailableStreamsOAuth2BadRequest with default headers values
func NewListAvailableStreamsOAuth2BadRequest() *ListAvailableStreamsOAuth2BadRequest {
	return &ListAvailableStreamsOAuth2BadRequest{}
}

/*
ListAvailableStreamsOAuth2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListAvailableStreamsOAuth2BadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MainDiscoveryResponseV2
}

// IsSuccess returns true when this list available streams o auth2 bad request response has a 2xx status code
func (o *ListAvailableStreamsOAuth2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list available streams o auth2 bad request response has a 3xx status code
func (o *ListAvailableStreamsOAuth2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list available streams o auth2 bad request response has a 4xx status code
func (o *ListAvailableStreamsOAuth2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list available streams o auth2 bad request response has a 5xx status code
func (o *ListAvailableStreamsOAuth2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list available streams o auth2 bad request response a status code equal to that given
func (o *ListAvailableStreamsOAuth2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list available streams o auth2 bad request response
func (o *ListAvailableStreamsOAuth2BadRequest) Code() int {
	return 400
}

func (o *ListAvailableStreamsOAuth2BadRequest) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2BadRequest  %+v", 400, o.Payload)
}

func (o *ListAvailableStreamsOAuth2BadRequest) String() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2BadRequest  %+v", 400, o.Payload)
}

func (o *ListAvailableStreamsOAuth2BadRequest) GetPayload() *models.MainDiscoveryResponseV2 {
	return o.Payload
}

func (o *ListAvailableStreamsOAuth2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MainDiscoveryResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAvailableStreamsOAuth2Forbidden creates a ListAvailableStreamsOAuth2Forbidden with default headers values
func NewListAvailableStreamsOAuth2Forbidden() *ListAvailableStreamsOAuth2Forbidden {
	return &ListAvailableStreamsOAuth2Forbidden{}
}

/*
ListAvailableStreamsOAuth2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListAvailableStreamsOAuth2Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this list available streams o auth2 forbidden response has a 2xx status code
func (o *ListAvailableStreamsOAuth2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list available streams o auth2 forbidden response has a 3xx status code
func (o *ListAvailableStreamsOAuth2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list available streams o auth2 forbidden response has a 4xx status code
func (o *ListAvailableStreamsOAuth2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list available streams o auth2 forbidden response has a 5xx status code
func (o *ListAvailableStreamsOAuth2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list available streams o auth2 forbidden response a status code equal to that given
func (o *ListAvailableStreamsOAuth2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list available streams o auth2 forbidden response
func (o *ListAvailableStreamsOAuth2Forbidden) Code() int {
	return 403
}

func (o *ListAvailableStreamsOAuth2Forbidden) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2Forbidden  %+v", 403, o.Payload)
}

func (o *ListAvailableStreamsOAuth2Forbidden) String() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2Forbidden  %+v", 403, o.Payload)
}

func (o *ListAvailableStreamsOAuth2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ListAvailableStreamsOAuth2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAvailableStreamsOAuth2TooManyRequests creates a ListAvailableStreamsOAuth2TooManyRequests with default headers values
func NewListAvailableStreamsOAuth2TooManyRequests() *ListAvailableStreamsOAuth2TooManyRequests {
	return &ListAvailableStreamsOAuth2TooManyRequests{}
}

/*
ListAvailableStreamsOAuth2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ListAvailableStreamsOAuth2TooManyRequests struct {

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

// IsSuccess returns true when this list available streams o auth2 too many requests response has a 2xx status code
func (o *ListAvailableStreamsOAuth2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list available streams o auth2 too many requests response has a 3xx status code
func (o *ListAvailableStreamsOAuth2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list available streams o auth2 too many requests response has a 4xx status code
func (o *ListAvailableStreamsOAuth2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list available streams o auth2 too many requests response has a 5xx status code
func (o *ListAvailableStreamsOAuth2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list available streams o auth2 too many requests response a status code equal to that given
func (o *ListAvailableStreamsOAuth2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the list available streams o auth2 too many requests response
func (o *ListAvailableStreamsOAuth2TooManyRequests) Code() int {
	return 429
}

func (o *ListAvailableStreamsOAuth2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2TooManyRequests  %+v", 429, o.Payload)
}

func (o *ListAvailableStreamsOAuth2TooManyRequests) String() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2TooManyRequests  %+v", 429, o.Payload)
}

func (o *ListAvailableStreamsOAuth2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ListAvailableStreamsOAuth2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAvailableStreamsOAuth2InternalServerError creates a ListAvailableStreamsOAuth2InternalServerError with default headers values
func NewListAvailableStreamsOAuth2InternalServerError() *ListAvailableStreamsOAuth2InternalServerError {
	return &ListAvailableStreamsOAuth2InternalServerError{}
}

/*
ListAvailableStreamsOAuth2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ListAvailableStreamsOAuth2InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MainDiscoveryResponseV2
}

// IsSuccess returns true when this list available streams o auth2 internal server error response has a 2xx status code
func (o *ListAvailableStreamsOAuth2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list available streams o auth2 internal server error response has a 3xx status code
func (o *ListAvailableStreamsOAuth2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list available streams o auth2 internal server error response has a 4xx status code
func (o *ListAvailableStreamsOAuth2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list available streams o auth2 internal server error response has a 5xx status code
func (o *ListAvailableStreamsOAuth2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list available streams o auth2 internal server error response a status code equal to that given
func (o *ListAvailableStreamsOAuth2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list available streams o auth2 internal server error response
func (o *ListAvailableStreamsOAuth2InternalServerError) Code() int {
	return 500
}

func (o *ListAvailableStreamsOAuth2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2InternalServerError  %+v", 500, o.Payload)
}

func (o *ListAvailableStreamsOAuth2InternalServerError) String() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2InternalServerError  %+v", 500, o.Payload)
}

func (o *ListAvailableStreamsOAuth2InternalServerError) GetPayload() *models.MainDiscoveryResponseV2 {
	return o.Payload
}

func (o *ListAvailableStreamsOAuth2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MainDiscoveryResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAvailableStreamsOAuth2Default creates a ListAvailableStreamsOAuth2Default with default headers values
func NewListAvailableStreamsOAuth2Default(code int) *ListAvailableStreamsOAuth2Default {
	return &ListAvailableStreamsOAuth2Default{
		_statusCode: code,
	}
}

/*
ListAvailableStreamsOAuth2Default describes a response with status code -1, with default header values.

OK
*/
type ListAvailableStreamsOAuth2Default struct {
	_statusCode int

	Payload *models.MainDiscoveryResponseV2
}

// IsSuccess returns true when this list available streams o auth2 default response has a 2xx status code
func (o *ListAvailableStreamsOAuth2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list available streams o auth2 default response has a 3xx status code
func (o *ListAvailableStreamsOAuth2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list available streams o auth2 default response has a 4xx status code
func (o *ListAvailableStreamsOAuth2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list available streams o auth2 default response has a 5xx status code
func (o *ListAvailableStreamsOAuth2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list available streams o auth2 default response a status code equal to that given
func (o *ListAvailableStreamsOAuth2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list available streams o auth2 default response
func (o *ListAvailableStreamsOAuth2Default) Code() int {
	return o._statusCode
}

func (o *ListAvailableStreamsOAuth2Default) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListAvailableStreamsOAuth2Default) String() string {
	return fmt.Sprintf("[GET /sensors/entities/datafeed/v2][%d] listAvailableStreamsOAuth2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListAvailableStreamsOAuth2Default) GetPayload() *models.MainDiscoveryResponseV2 {
	return o.Payload
}

func (o *ListAvailableStreamsOAuth2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MainDiscoveryResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
