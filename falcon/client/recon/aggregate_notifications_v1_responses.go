// Code generated by go-swagger; DO NOT EDIT.

package recon

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

// AggregateNotificationsV1Reader is a Reader for the AggregateNotificationsV1 structure.
type AggregateNotificationsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateNotificationsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateNotificationsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAggregateNotificationsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAggregateNotificationsV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAggregateNotificationsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateNotificationsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAggregateNotificationsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAggregateNotificationsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregateNotificationsV1OK creates a AggregateNotificationsV1OK with default headers values
func NewAggregateNotificationsV1OK() *AggregateNotificationsV1OK {
	return &AggregateNotificationsV1OK{}
}

/*
AggregateNotificationsV1OK describes a response with status code 200, with default header values.

OK
*/
type AggregateNotificationsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregatesResponse
}

// IsSuccess returns true when this aggregate notifications v1 o k response has a 2xx status code
func (o *AggregateNotificationsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate notifications v1 o k response has a 3xx status code
func (o *AggregateNotificationsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate notifications v1 o k response has a 4xx status code
func (o *AggregateNotificationsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate notifications v1 o k response has a 5xx status code
func (o *AggregateNotificationsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate notifications v1 o k response a status code equal to that given
func (o *AggregateNotificationsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate notifications v1 o k response
func (o *AggregateNotificationsV1OK) Code() int {
	return 200
}

func (o *AggregateNotificationsV1OK) Error() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1OK  %+v", 200, o.Payload)
}

func (o *AggregateNotificationsV1OK) String() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1OK  %+v", 200, o.Payload)
}

func (o *AggregateNotificationsV1OK) GetPayload() *models.DomainAggregatesResponse {
	return o.Payload
}

func (o *AggregateNotificationsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateNotificationsV1BadRequest creates a AggregateNotificationsV1BadRequest with default headers values
func NewAggregateNotificationsV1BadRequest() *AggregateNotificationsV1BadRequest {
	return &AggregateNotificationsV1BadRequest{}
}

/*
AggregateNotificationsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AggregateNotificationsV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this aggregate notifications v1 bad request response has a 2xx status code
func (o *AggregateNotificationsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate notifications v1 bad request response has a 3xx status code
func (o *AggregateNotificationsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate notifications v1 bad request response has a 4xx status code
func (o *AggregateNotificationsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate notifications v1 bad request response has a 5xx status code
func (o *AggregateNotificationsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate notifications v1 bad request response a status code equal to that given
func (o *AggregateNotificationsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the aggregate notifications v1 bad request response
func (o *AggregateNotificationsV1BadRequest) Code() int {
	return 400
}

func (o *AggregateNotificationsV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1BadRequest  %+v", 400, o.Payload)
}

func (o *AggregateNotificationsV1BadRequest) String() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1BadRequest  %+v", 400, o.Payload)
}

func (o *AggregateNotificationsV1BadRequest) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *AggregateNotificationsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateNotificationsV1Unauthorized creates a AggregateNotificationsV1Unauthorized with default headers values
func NewAggregateNotificationsV1Unauthorized() *AggregateNotificationsV1Unauthorized {
	return &AggregateNotificationsV1Unauthorized{}
}

/*
AggregateNotificationsV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AggregateNotificationsV1Unauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this aggregate notifications v1 unauthorized response has a 2xx status code
func (o *AggregateNotificationsV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate notifications v1 unauthorized response has a 3xx status code
func (o *AggregateNotificationsV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate notifications v1 unauthorized response has a 4xx status code
func (o *AggregateNotificationsV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate notifications v1 unauthorized response has a 5xx status code
func (o *AggregateNotificationsV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate notifications v1 unauthorized response a status code equal to that given
func (o *AggregateNotificationsV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the aggregate notifications v1 unauthorized response
func (o *AggregateNotificationsV1Unauthorized) Code() int {
	return 401
}

func (o *AggregateNotificationsV1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1Unauthorized  %+v", 401, o.Payload)
}

func (o *AggregateNotificationsV1Unauthorized) String() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1Unauthorized  %+v", 401, o.Payload)
}

func (o *AggregateNotificationsV1Unauthorized) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *AggregateNotificationsV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateNotificationsV1Forbidden creates a AggregateNotificationsV1Forbidden with default headers values
func NewAggregateNotificationsV1Forbidden() *AggregateNotificationsV1Forbidden {
	return &AggregateNotificationsV1Forbidden{}
}

/*
AggregateNotificationsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateNotificationsV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this aggregate notifications v1 forbidden response has a 2xx status code
func (o *AggregateNotificationsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate notifications v1 forbidden response has a 3xx status code
func (o *AggregateNotificationsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate notifications v1 forbidden response has a 4xx status code
func (o *AggregateNotificationsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate notifications v1 forbidden response has a 5xx status code
func (o *AggregateNotificationsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate notifications v1 forbidden response a status code equal to that given
func (o *AggregateNotificationsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate notifications v1 forbidden response
func (o *AggregateNotificationsV1Forbidden) Code() int {
	return 403
}

func (o *AggregateNotificationsV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1Forbidden  %+v", 403, o.Payload)
}

func (o *AggregateNotificationsV1Forbidden) String() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1Forbidden  %+v", 403, o.Payload)
}

func (o *AggregateNotificationsV1Forbidden) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *AggregateNotificationsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateNotificationsV1TooManyRequests creates a AggregateNotificationsV1TooManyRequests with default headers values
func NewAggregateNotificationsV1TooManyRequests() *AggregateNotificationsV1TooManyRequests {
	return &AggregateNotificationsV1TooManyRequests{}
}

/*
AggregateNotificationsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateNotificationsV1TooManyRequests struct {

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

// IsSuccess returns true when this aggregate notifications v1 too many requests response has a 2xx status code
func (o *AggregateNotificationsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate notifications v1 too many requests response has a 3xx status code
func (o *AggregateNotificationsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate notifications v1 too many requests response has a 4xx status code
func (o *AggregateNotificationsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate notifications v1 too many requests response has a 5xx status code
func (o *AggregateNotificationsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate notifications v1 too many requests response a status code equal to that given
func (o *AggregateNotificationsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate notifications v1 too many requests response
func (o *AggregateNotificationsV1TooManyRequests) Code() int {
	return 429
}

func (o *AggregateNotificationsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateNotificationsV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateNotificationsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateNotificationsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateNotificationsV1InternalServerError creates a AggregateNotificationsV1InternalServerError with default headers values
func NewAggregateNotificationsV1InternalServerError() *AggregateNotificationsV1InternalServerError {
	return &AggregateNotificationsV1InternalServerError{}
}

/*
AggregateNotificationsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AggregateNotificationsV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this aggregate notifications v1 internal server error response has a 2xx status code
func (o *AggregateNotificationsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate notifications v1 internal server error response has a 3xx status code
func (o *AggregateNotificationsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate notifications v1 internal server error response has a 4xx status code
func (o *AggregateNotificationsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate notifications v1 internal server error response has a 5xx status code
func (o *AggregateNotificationsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aggregate notifications v1 internal server error response a status code equal to that given
func (o *AggregateNotificationsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the aggregate notifications v1 internal server error response
func (o *AggregateNotificationsV1InternalServerError) Code() int {
	return 500
}

func (o *AggregateNotificationsV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateNotificationsV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] aggregateNotificationsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateNotificationsV1InternalServerError) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *AggregateNotificationsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateNotificationsV1Default creates a AggregateNotificationsV1Default with default headers values
func NewAggregateNotificationsV1Default(code int) *AggregateNotificationsV1Default {
	return &AggregateNotificationsV1Default{
		_statusCode: code,
	}
}

/*
AggregateNotificationsV1Default describes a response with status code -1, with default header values.

OK
*/
type AggregateNotificationsV1Default struct {
	_statusCode int

	Payload *models.DomainAggregatesResponse
}

// IsSuccess returns true when this aggregate notifications v1 default response has a 2xx status code
func (o *AggregateNotificationsV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aggregate notifications v1 default response has a 3xx status code
func (o *AggregateNotificationsV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aggregate notifications v1 default response has a 4xx status code
func (o *AggregateNotificationsV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aggregate notifications v1 default response has a 5xx status code
func (o *AggregateNotificationsV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aggregate notifications v1 default response a status code equal to that given
func (o *AggregateNotificationsV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the aggregate notifications v1 default response
func (o *AggregateNotificationsV1Default) Code() int {
	return o._statusCode
}

func (o *AggregateNotificationsV1Default) Error() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] AggregateNotificationsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *AggregateNotificationsV1Default) String() string {
	return fmt.Sprintf("[POST /recon/aggregates/notifications/GET/v1][%d] AggregateNotificationsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *AggregateNotificationsV1Default) GetPayload() *models.DomainAggregatesResponse {
	return o.Payload
}

func (o *AggregateNotificationsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
