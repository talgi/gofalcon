// Code generated by go-swagger; DO NOT EDIT.

package quarantine

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

// UpdateQfByQueryReader is a Reader for the UpdateQfByQuery structure.
type UpdateQfByQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateQfByQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateQfByQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateQfByQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateQfByQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateQfByQueryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateQfByQueryOK creates a UpdateQfByQueryOK with default headers values
func NewUpdateQfByQueryOK() *UpdateQfByQueryOK {
	return &UpdateQfByQueryOK{}
}

/*
UpdateQfByQueryOK describes a response with status code 200, with default header values.

OK
*/
type UpdateQfByQueryOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update qf by query o k response has a 2xx status code
func (o *UpdateQfByQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update qf by query o k response has a 3xx status code
func (o *UpdateQfByQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update qf by query o k response has a 4xx status code
func (o *UpdateQfByQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update qf by query o k response has a 5xx status code
func (o *UpdateQfByQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update qf by query o k response a status code equal to that given
func (o *UpdateQfByQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateQfByQueryOK) Error() string {
	return fmt.Sprintf("[PATCH /quarantine/queries/quarantined-files/v1][%d] updateQfByQueryOK  %+v", 200, o.Payload)
}

func (o *UpdateQfByQueryOK) String() string {
	return fmt.Sprintf("[PATCH /quarantine/queries/quarantined-files/v1][%d] updateQfByQueryOK  %+v", 200, o.Payload)
}

func (o *UpdateQfByQueryOK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateQfByQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewUpdateQfByQueryForbidden creates a UpdateQfByQueryForbidden with default headers values
func NewUpdateQfByQueryForbidden() *UpdateQfByQueryForbidden {
	return &UpdateQfByQueryForbidden{}
}

/*
UpdateQfByQueryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateQfByQueryForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update qf by query forbidden response has a 2xx status code
func (o *UpdateQfByQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update qf by query forbidden response has a 3xx status code
func (o *UpdateQfByQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update qf by query forbidden response has a 4xx status code
func (o *UpdateQfByQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update qf by query forbidden response has a 5xx status code
func (o *UpdateQfByQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update qf by query forbidden response a status code equal to that given
func (o *UpdateQfByQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateQfByQueryForbidden) Error() string {
	return fmt.Sprintf("[PATCH /quarantine/queries/quarantined-files/v1][%d] updateQfByQueryForbidden  %+v", 403, o.Payload)
}

func (o *UpdateQfByQueryForbidden) String() string {
	return fmt.Sprintf("[PATCH /quarantine/queries/quarantined-files/v1][%d] updateQfByQueryForbidden  %+v", 403, o.Payload)
}

func (o *UpdateQfByQueryForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateQfByQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewUpdateQfByQueryTooManyRequests creates a UpdateQfByQueryTooManyRequests with default headers values
func NewUpdateQfByQueryTooManyRequests() *UpdateQfByQueryTooManyRequests {
	return &UpdateQfByQueryTooManyRequests{}
}

/*
UpdateQfByQueryTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateQfByQueryTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

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

// IsSuccess returns true when this update qf by query too many requests response has a 2xx status code
func (o *UpdateQfByQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update qf by query too many requests response has a 3xx status code
func (o *UpdateQfByQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update qf by query too many requests response has a 4xx status code
func (o *UpdateQfByQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update qf by query too many requests response has a 5xx status code
func (o *UpdateQfByQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update qf by query too many requests response a status code equal to that given
func (o *UpdateQfByQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateQfByQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /quarantine/queries/quarantined-files/v1][%d] updateQfByQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateQfByQueryTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /quarantine/queries/quarantined-files/v1][%d] updateQfByQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateQfByQueryTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateQfByQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewUpdateQfByQueryDefault creates a UpdateQfByQueryDefault with default headers values
func NewUpdateQfByQueryDefault(code int) *UpdateQfByQueryDefault {
	return &UpdateQfByQueryDefault{
		_statusCode: code,
	}
}

/*
UpdateQfByQueryDefault describes a response with status code -1, with default header values.

OK
*/
type UpdateQfByQueryDefault struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// Code gets the status code for the update qf by query default response
func (o *UpdateQfByQueryDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update qf by query default response has a 2xx status code
func (o *UpdateQfByQueryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update qf by query default response has a 3xx status code
func (o *UpdateQfByQueryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update qf by query default response has a 4xx status code
func (o *UpdateQfByQueryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update qf by query default response has a 5xx status code
func (o *UpdateQfByQueryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update qf by query default response a status code equal to that given
func (o *UpdateQfByQueryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateQfByQueryDefault) Error() string {
	return fmt.Sprintf("[PATCH /quarantine/queries/quarantined-files/v1][%d] UpdateQfByQuery default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateQfByQueryDefault) String() string {
	return fmt.Sprintf("[PATCH /quarantine/queries/quarantined-files/v1][%d] UpdateQfByQuery default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateQfByQueryDefault) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateQfByQueryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
