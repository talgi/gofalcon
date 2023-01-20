// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// QueryChildrenReader is a Reader for the QueryChildren structure.
type QueryChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryChildrenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryChildrenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryChildrenTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryChildrenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryChildrenOK creates a QueryChildrenOK with default headers values
func NewQueryChildrenOK() *QueryChildrenOK {
	return &QueryChildrenOK{}
}

/*
QueryChildrenOK describes a response with status code 200, with default header values.

OK
*/
type QueryChildrenOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query children o k response has a 2xx status code
func (o *QueryChildrenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query children o k response has a 3xx status code
func (o *QueryChildrenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query children o k response has a 4xx status code
func (o *QueryChildrenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query children o k response has a 5xx status code
func (o *QueryChildrenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query children o k response a status code equal to that given
func (o *QueryChildrenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query children o k response
func (o *QueryChildrenOK) Code() int {
	return 200
}

func (o *QueryChildrenOK) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/children/v1][%d] queryChildrenOK  %+v", 200, o.Payload)
}

func (o *QueryChildrenOK) String() string {
	return fmt.Sprintf("[GET /mssp/queries/children/v1][%d] queryChildrenOK  %+v", 200, o.Payload)
}

func (o *QueryChildrenOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryChildrenBadRequest creates a QueryChildrenBadRequest with default headers values
func NewQueryChildrenBadRequest() *QueryChildrenBadRequest {
	return &QueryChildrenBadRequest{}
}

/*
QueryChildrenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryChildrenBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query children bad request response has a 2xx status code
func (o *QueryChildrenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query children bad request response has a 3xx status code
func (o *QueryChildrenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query children bad request response has a 4xx status code
func (o *QueryChildrenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query children bad request response has a 5xx status code
func (o *QueryChildrenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query children bad request response a status code equal to that given
func (o *QueryChildrenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query children bad request response
func (o *QueryChildrenBadRequest) Code() int {
	return 400
}

func (o *QueryChildrenBadRequest) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/children/v1][%d] queryChildrenBadRequest  %+v", 400, o.Payload)
}

func (o *QueryChildrenBadRequest) String() string {
	return fmt.Sprintf("[GET /mssp/queries/children/v1][%d] queryChildrenBadRequest  %+v", 400, o.Payload)
}

func (o *QueryChildrenBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryChildrenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryChildrenForbidden creates a QueryChildrenForbidden with default headers values
func NewQueryChildrenForbidden() *QueryChildrenForbidden {
	return &QueryChildrenForbidden{}
}

/*
QueryChildrenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryChildrenForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query children forbidden response has a 2xx status code
func (o *QueryChildrenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query children forbidden response has a 3xx status code
func (o *QueryChildrenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query children forbidden response has a 4xx status code
func (o *QueryChildrenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query children forbidden response has a 5xx status code
func (o *QueryChildrenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query children forbidden response a status code equal to that given
func (o *QueryChildrenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query children forbidden response
func (o *QueryChildrenForbidden) Code() int {
	return 403
}

func (o *QueryChildrenForbidden) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/children/v1][%d] queryChildrenForbidden  %+v", 403, o.Payload)
}

func (o *QueryChildrenForbidden) String() string {
	return fmt.Sprintf("[GET /mssp/queries/children/v1][%d] queryChildrenForbidden  %+v", 403, o.Payload)
}

func (o *QueryChildrenForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryChildrenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryChildrenTooManyRequests creates a QueryChildrenTooManyRequests with default headers values
func NewQueryChildrenTooManyRequests() *QueryChildrenTooManyRequests {
	return &QueryChildrenTooManyRequests{}
}

/*
QueryChildrenTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryChildrenTooManyRequests struct {

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

// IsSuccess returns true when this query children too many requests response has a 2xx status code
func (o *QueryChildrenTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query children too many requests response has a 3xx status code
func (o *QueryChildrenTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query children too many requests response has a 4xx status code
func (o *QueryChildrenTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query children too many requests response has a 5xx status code
func (o *QueryChildrenTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query children too many requests response a status code equal to that given
func (o *QueryChildrenTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query children too many requests response
func (o *QueryChildrenTooManyRequests) Code() int {
	return 429
}

func (o *QueryChildrenTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/children/v1][%d] queryChildrenTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryChildrenTooManyRequests) String() string {
	return fmt.Sprintf("[GET /mssp/queries/children/v1][%d] queryChildrenTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryChildrenTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryChildrenTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryChildrenDefault creates a QueryChildrenDefault with default headers values
func NewQueryChildrenDefault(code int) *QueryChildrenDefault {
	return &QueryChildrenDefault{
		_statusCode: code,
	}
}

/*
QueryChildrenDefault describes a response with status code -1, with default header values.

OK
*/
type QueryChildrenDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query children default response has a 2xx status code
func (o *QueryChildrenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query children default response has a 3xx status code
func (o *QueryChildrenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query children default response has a 4xx status code
func (o *QueryChildrenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query children default response has a 5xx status code
func (o *QueryChildrenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query children default response a status code equal to that given
func (o *QueryChildrenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query children default response
func (o *QueryChildrenDefault) Code() int {
	return o._statusCode
}

func (o *QueryChildrenDefault) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/children/v1][%d] queryChildren default  %+v", o._statusCode, o.Payload)
}

func (o *QueryChildrenDefault) String() string {
	return fmt.Sprintf("[GET /mssp/queries/children/v1][%d] queryChildren default  %+v", o._statusCode, o.Payload)
}

func (o *QueryChildrenDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryChildrenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
