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

// UpdateActionV1Reader is a Reader for the UpdateActionV1 structure.
type UpdateActionV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateActionV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateActionV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateActionV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateActionV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateActionV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateActionV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateActionV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateActionV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateActionV1OK creates a UpdateActionV1OK with default headers values
func NewUpdateActionV1OK() *UpdateActionV1OK {
	return &UpdateActionV1OK{}
}

/*
UpdateActionV1OK describes a response with status code 200, with default header values.

OK
*/
type UpdateActionV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainActionEntitiesResponseV1
}

// IsSuccess returns true when this update action v1 o k response has a 2xx status code
func (o *UpdateActionV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update action v1 o k response has a 3xx status code
func (o *UpdateActionV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update action v1 o k response has a 4xx status code
func (o *UpdateActionV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update action v1 o k response has a 5xx status code
func (o *UpdateActionV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this update action v1 o k response a status code equal to that given
func (o *UpdateActionV1OK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateActionV1OK) Error() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1OK  %+v", 200, o.Payload)
}

func (o *UpdateActionV1OK) String() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1OK  %+v", 200, o.Payload)
}

func (o *UpdateActionV1OK) GetPayload() *models.DomainActionEntitiesResponseV1 {
	return o.Payload
}

func (o *UpdateActionV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainActionEntitiesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateActionV1BadRequest creates a UpdateActionV1BadRequest with default headers values
func NewUpdateActionV1BadRequest() *UpdateActionV1BadRequest {
	return &UpdateActionV1BadRequest{}
}

/*
UpdateActionV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateActionV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this update action v1 bad request response has a 2xx status code
func (o *UpdateActionV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update action v1 bad request response has a 3xx status code
func (o *UpdateActionV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update action v1 bad request response has a 4xx status code
func (o *UpdateActionV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update action v1 bad request response has a 5xx status code
func (o *UpdateActionV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update action v1 bad request response a status code equal to that given
func (o *UpdateActionV1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateActionV1BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateActionV1BadRequest) String() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateActionV1BadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateActionV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateActionV1Unauthorized creates a UpdateActionV1Unauthorized with default headers values
func NewUpdateActionV1Unauthorized() *UpdateActionV1Unauthorized {
	return &UpdateActionV1Unauthorized{}
}

/*
UpdateActionV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateActionV1Unauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this update action v1 unauthorized response has a 2xx status code
func (o *UpdateActionV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update action v1 unauthorized response has a 3xx status code
func (o *UpdateActionV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update action v1 unauthorized response has a 4xx status code
func (o *UpdateActionV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update action v1 unauthorized response has a 5xx status code
func (o *UpdateActionV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update action v1 unauthorized response a status code equal to that given
func (o *UpdateActionV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateActionV1Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1Unauthorized  %+v", 401, o.Payload)
}

func (o *UpdateActionV1Unauthorized) String() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1Unauthorized  %+v", 401, o.Payload)
}

func (o *UpdateActionV1Unauthorized) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateActionV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateActionV1Forbidden creates a UpdateActionV1Forbidden with default headers values
func NewUpdateActionV1Forbidden() *UpdateActionV1Forbidden {
	return &UpdateActionV1Forbidden{}
}

/*
UpdateActionV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateActionV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this update action v1 forbidden response has a 2xx status code
func (o *UpdateActionV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update action v1 forbidden response has a 3xx status code
func (o *UpdateActionV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update action v1 forbidden response has a 4xx status code
func (o *UpdateActionV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update action v1 forbidden response has a 5xx status code
func (o *UpdateActionV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update action v1 forbidden response a status code equal to that given
func (o *UpdateActionV1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateActionV1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateActionV1Forbidden) String() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateActionV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateActionV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateActionV1TooManyRequests creates a UpdateActionV1TooManyRequests with default headers values
func NewUpdateActionV1TooManyRequests() *UpdateActionV1TooManyRequests {
	return &UpdateActionV1TooManyRequests{}
}

/*
UpdateActionV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateActionV1TooManyRequests struct {

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

// IsSuccess returns true when this update action v1 too many requests response has a 2xx status code
func (o *UpdateActionV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update action v1 too many requests response has a 3xx status code
func (o *UpdateActionV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update action v1 too many requests response has a 4xx status code
func (o *UpdateActionV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update action v1 too many requests response has a 5xx status code
func (o *UpdateActionV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update action v1 too many requests response a status code equal to that given
func (o *UpdateActionV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateActionV1TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateActionV1TooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateActionV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateActionV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateActionV1InternalServerError creates a UpdateActionV1InternalServerError with default headers values
func NewUpdateActionV1InternalServerError() *UpdateActionV1InternalServerError {
	return &UpdateActionV1InternalServerError{}
}

/*
UpdateActionV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateActionV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this update action v1 internal server error response has a 2xx status code
func (o *UpdateActionV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update action v1 internal server error response has a 3xx status code
func (o *UpdateActionV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update action v1 internal server error response has a 4xx status code
func (o *UpdateActionV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update action v1 internal server error response has a 5xx status code
func (o *UpdateActionV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update action v1 internal server error response a status code equal to that given
func (o *UpdateActionV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateActionV1InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateActionV1InternalServerError) String() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] updateActionV1InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateActionV1InternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateActionV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateActionV1Default creates a UpdateActionV1Default with default headers values
func NewUpdateActionV1Default(code int) *UpdateActionV1Default {
	return &UpdateActionV1Default{
		_statusCode: code,
	}
}

/*
UpdateActionV1Default describes a response with status code -1, with default header values.

OK
*/
type UpdateActionV1Default struct {
	_statusCode int

	Payload *models.DomainActionEntitiesResponseV1
}

// Code gets the status code for the update action v1 default response
func (o *UpdateActionV1Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update action v1 default response has a 2xx status code
func (o *UpdateActionV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update action v1 default response has a 3xx status code
func (o *UpdateActionV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update action v1 default response has a 4xx status code
func (o *UpdateActionV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update action v1 default response has a 5xx status code
func (o *UpdateActionV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update action v1 default response a status code equal to that given
func (o *UpdateActionV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateActionV1Default) Error() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] UpdateActionV1 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateActionV1Default) String() string {
	return fmt.Sprintf("[PATCH /recon/entities/actions/v1][%d] UpdateActionV1 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateActionV1Default) GetPayload() *models.DomainActionEntitiesResponseV1 {
	return o.Payload
}

func (o *UpdateActionV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainActionEntitiesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
