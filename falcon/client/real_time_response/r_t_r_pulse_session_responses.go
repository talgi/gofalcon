// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// RTRPulseSessionReader is a Reader for the RTRPulseSession structure.
type RTRPulseSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRPulseSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRTRPulseSessionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRPulseSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRPulseSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRPulseSessionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRTRPulseSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /real-time-response/entities/refresh-session/v1] RTR-PulseSession", response, response.Code())
	}
}

// NewRTRPulseSessionCreated creates a RTRPulseSessionCreated with default headers values
func NewRTRPulseSessionCreated() *RTRPulseSessionCreated {
	return &RTRPulseSessionCreated{}
}

/*
RTRPulseSessionCreated describes a response with status code 201, with default header values.

Created
*/
type RTRPulseSessionCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainInitResponseWrapper
}

// IsSuccess returns true when this r t r pulse session created response has a 2xx status code
func (o *RTRPulseSessionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r pulse session created response has a 3xx status code
func (o *RTRPulseSessionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r pulse session created response has a 4xx status code
func (o *RTRPulseSessionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r pulse session created response has a 5xx status code
func (o *RTRPulseSessionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r pulse session created response a status code equal to that given
func (o *RTRPulseSessionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the r t r pulse session created response
func (o *RTRPulseSessionCreated) Code() int {
	return 201
}

func (o *RTRPulseSessionCreated) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/refresh-session/v1][%d] rTRPulseSessionCreated  %+v", 201, o.Payload)
}

func (o *RTRPulseSessionCreated) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/refresh-session/v1][%d] rTRPulseSessionCreated  %+v", 201, o.Payload)
}

func (o *RTRPulseSessionCreated) GetPayload() *models.DomainInitResponseWrapper {
	return o.Payload
}

func (o *RTRPulseSessionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainInitResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRPulseSessionBadRequest creates a RTRPulseSessionBadRequest with default headers values
func NewRTRPulseSessionBadRequest() *RTRPulseSessionBadRequest {
	return &RTRPulseSessionBadRequest{}
}

/*
RTRPulseSessionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRPulseSessionBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r pulse session bad request response has a 2xx status code
func (o *RTRPulseSessionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r pulse session bad request response has a 3xx status code
func (o *RTRPulseSessionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r pulse session bad request response has a 4xx status code
func (o *RTRPulseSessionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r pulse session bad request response has a 5xx status code
func (o *RTRPulseSessionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r pulse session bad request response a status code equal to that given
func (o *RTRPulseSessionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r pulse session bad request response
func (o *RTRPulseSessionBadRequest) Code() int {
	return 400
}

func (o *RTRPulseSessionBadRequest) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/refresh-session/v1][%d] rTRPulseSessionBadRequest  %+v", 400, o.Payload)
}

func (o *RTRPulseSessionBadRequest) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/refresh-session/v1][%d] rTRPulseSessionBadRequest  %+v", 400, o.Payload)
}

func (o *RTRPulseSessionBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRPulseSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRPulseSessionForbidden creates a RTRPulseSessionForbidden with default headers values
func NewRTRPulseSessionForbidden() *RTRPulseSessionForbidden {
	return &RTRPulseSessionForbidden{}
}

/*
RTRPulseSessionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRPulseSessionForbidden struct {

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

// IsSuccess returns true when this r t r pulse session forbidden response has a 2xx status code
func (o *RTRPulseSessionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r pulse session forbidden response has a 3xx status code
func (o *RTRPulseSessionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r pulse session forbidden response has a 4xx status code
func (o *RTRPulseSessionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r pulse session forbidden response has a 5xx status code
func (o *RTRPulseSessionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r pulse session forbidden response a status code equal to that given
func (o *RTRPulseSessionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r pulse session forbidden response
func (o *RTRPulseSessionForbidden) Code() int {
	return 403
}

func (o *RTRPulseSessionForbidden) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/refresh-session/v1][%d] rTRPulseSessionForbidden  %+v", 403, o.Payload)
}

func (o *RTRPulseSessionForbidden) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/refresh-session/v1][%d] rTRPulseSessionForbidden  %+v", 403, o.Payload)
}

func (o *RTRPulseSessionForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRPulseSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRPulseSessionTooManyRequests creates a RTRPulseSessionTooManyRequests with default headers values
func NewRTRPulseSessionTooManyRequests() *RTRPulseSessionTooManyRequests {
	return &RTRPulseSessionTooManyRequests{}
}

/*
RTRPulseSessionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRPulseSessionTooManyRequests struct {

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

// IsSuccess returns true when this r t r pulse session too many requests response has a 2xx status code
func (o *RTRPulseSessionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r pulse session too many requests response has a 3xx status code
func (o *RTRPulseSessionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r pulse session too many requests response has a 4xx status code
func (o *RTRPulseSessionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r pulse session too many requests response has a 5xx status code
func (o *RTRPulseSessionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r pulse session too many requests response a status code equal to that given
func (o *RTRPulseSessionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r pulse session too many requests response
func (o *RTRPulseSessionTooManyRequests) Code() int {
	return 429
}

func (o *RTRPulseSessionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/refresh-session/v1][%d] rTRPulseSessionTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRPulseSessionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/refresh-session/v1][%d] rTRPulseSessionTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRPulseSessionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRPulseSessionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRPulseSessionInternalServerError creates a RTRPulseSessionInternalServerError with default headers values
func NewRTRPulseSessionInternalServerError() *RTRPulseSessionInternalServerError {
	return &RTRPulseSessionInternalServerError{}
}

/*
RTRPulseSessionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RTRPulseSessionInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r pulse session internal server error response has a 2xx status code
func (o *RTRPulseSessionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r pulse session internal server error response has a 3xx status code
func (o *RTRPulseSessionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r pulse session internal server error response has a 4xx status code
func (o *RTRPulseSessionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r pulse session internal server error response has a 5xx status code
func (o *RTRPulseSessionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this r t r pulse session internal server error response a status code equal to that given
func (o *RTRPulseSessionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the r t r pulse session internal server error response
func (o *RTRPulseSessionInternalServerError) Code() int {
	return 500
}

func (o *RTRPulseSessionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/refresh-session/v1][%d] rTRPulseSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRPulseSessionInternalServerError) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/refresh-session/v1][%d] rTRPulseSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRPulseSessionInternalServerError) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRPulseSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
