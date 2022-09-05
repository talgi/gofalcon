// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// UserActionV1Reader is a Reader for the UserActionV1 structure.
type UserActionV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserActionV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserActionV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserActionV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserActionV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUserActionV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserActionV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUserActionV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserActionV1OK creates a UserActionV1OK with default headers values
func NewUserActionV1OK() *UserActionV1OK {
	return &UserActionV1OK{}
}

/*
UserActionV1OK describes a response with status code 200, with default header values.

OK
*/
type UserActionV1OK struct {

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

// IsSuccess returns true when this user action v1 o k response has a 2xx status code
func (o *UserActionV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user action v1 o k response has a 3xx status code
func (o *UserActionV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user action v1 o k response has a 4xx status code
func (o *UserActionV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user action v1 o k response has a 5xx status code
func (o *UserActionV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this user action v1 o k response a status code equal to that given
func (o *UserActionV1OK) IsCode(code int) bool {
	return code == 200
}

func (o *UserActionV1OK) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1OK  %+v", 200, o.Payload)
}

func (o *UserActionV1OK) String() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1OK  %+v", 200, o.Payload)
}

func (o *UserActionV1OK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UserActionV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUserActionV1BadRequest creates a UserActionV1BadRequest with default headers values
func NewUserActionV1BadRequest() *UserActionV1BadRequest {
	return &UserActionV1BadRequest{}
}

/*
UserActionV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserActionV1BadRequest struct {

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

// IsSuccess returns true when this user action v1 bad request response has a 2xx status code
func (o *UserActionV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user action v1 bad request response has a 3xx status code
func (o *UserActionV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user action v1 bad request response has a 4xx status code
func (o *UserActionV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user action v1 bad request response has a 5xx status code
func (o *UserActionV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user action v1 bad request response a status code equal to that given
func (o *UserActionV1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserActionV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1BadRequest  %+v", 400, o.Payload)
}

func (o *UserActionV1BadRequest) String() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1BadRequest  %+v", 400, o.Payload)
}

func (o *UserActionV1BadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UserActionV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUserActionV1Forbidden creates a UserActionV1Forbidden with default headers values
func NewUserActionV1Forbidden() *UserActionV1Forbidden {
	return &UserActionV1Forbidden{}
}

/*
UserActionV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserActionV1Forbidden struct {

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

// IsSuccess returns true when this user action v1 forbidden response has a 2xx status code
func (o *UserActionV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user action v1 forbidden response has a 3xx status code
func (o *UserActionV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user action v1 forbidden response has a 4xx status code
func (o *UserActionV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user action v1 forbidden response has a 5xx status code
func (o *UserActionV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user action v1 forbidden response a status code equal to that given
func (o *UserActionV1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserActionV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1Forbidden  %+v", 403, o.Payload)
}

func (o *UserActionV1Forbidden) String() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1Forbidden  %+v", 403, o.Payload)
}

func (o *UserActionV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UserActionV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUserActionV1TooManyRequests creates a UserActionV1TooManyRequests with default headers values
func NewUserActionV1TooManyRequests() *UserActionV1TooManyRequests {
	return &UserActionV1TooManyRequests{}
}

/*
UserActionV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UserActionV1TooManyRequests struct {

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

// IsSuccess returns true when this user action v1 too many requests response has a 2xx status code
func (o *UserActionV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user action v1 too many requests response has a 3xx status code
func (o *UserActionV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user action v1 too many requests response has a 4xx status code
func (o *UserActionV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this user action v1 too many requests response has a 5xx status code
func (o *UserActionV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this user action v1 too many requests response a status code equal to that given
func (o *UserActionV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UserActionV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *UserActionV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *UserActionV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UserActionV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUserActionV1InternalServerError creates a UserActionV1InternalServerError with default headers values
func NewUserActionV1InternalServerError() *UserActionV1InternalServerError {
	return &UserActionV1InternalServerError{}
}

/*
UserActionV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UserActionV1InternalServerError struct {

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

// IsSuccess returns true when this user action v1 internal server error response has a 2xx status code
func (o *UserActionV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user action v1 internal server error response has a 3xx status code
func (o *UserActionV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user action v1 internal server error response has a 4xx status code
func (o *UserActionV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user action v1 internal server error response has a 5xx status code
func (o *UserActionV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user action v1 internal server error response a status code equal to that given
func (o *UserActionV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserActionV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1InternalServerError  %+v", 500, o.Payload)
}

func (o *UserActionV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1InternalServerError  %+v", 500, o.Payload)
}

func (o *UserActionV1InternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UserActionV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUserActionV1Default creates a UserActionV1Default with default headers values
func NewUserActionV1Default(code int) *UserActionV1Default {
	return &UserActionV1Default{
		_statusCode: code,
	}
}

/*
UserActionV1Default describes a response with status code -1, with default header values.

OK
*/
type UserActionV1Default struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// Code gets the status code for the user action v1 default response
func (o *UserActionV1Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this user action v1 default response has a 2xx status code
func (o *UserActionV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this user action v1 default response has a 3xx status code
func (o *UserActionV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this user action v1 default response has a 4xx status code
func (o *UserActionV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this user action v1 default response has a 5xx status code
func (o *UserActionV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this user action v1 default response a status code equal to that given
func (o *UserActionV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UserActionV1Default) Error() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1 default  %+v", o._statusCode, o.Payload)
}

func (o *UserActionV1Default) String() string {
	return fmt.Sprintf("[POST /user-management/entities/user-actions/v1][%d] userActionV1 default  %+v", o._statusCode, o.Payload)
}

func (o *UserActionV1Default) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UserActionV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
