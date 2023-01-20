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

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserOK creates a DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {
	return &DeleteUserOK{}
}

/*
DeleteUserOK describes a response with status code 200, with default header values.

OK
*/
type DeleteUserOK struct {

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

// IsSuccess returns true when this delete user o k response has a 2xx status code
func (o *DeleteUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user o k response has a 3xx status code
func (o *DeleteUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user o k response has a 4xx status code
func (o *DeleteUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user o k response has a 5xx status code
func (o *DeleteUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user o k response a status code equal to that given
func (o *DeleteUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete user o k response
func (o *DeleteUserOK) Code() int {
	return 200
}

func (o *DeleteUserOK) Error() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] deleteUserOK  %+v", 200, o.Payload)
}

func (o *DeleteUserOK) String() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] deleteUserOK  %+v", 200, o.Payload)
}

func (o *DeleteUserOK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserBadRequest creates a DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {
	return &DeleteUserBadRequest{}
}

/*
DeleteUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUserBadRequest struct {

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

// IsSuccess returns true when this delete user bad request response has a 2xx status code
func (o *DeleteUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user bad request response has a 3xx status code
func (o *DeleteUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user bad request response has a 4xx status code
func (o *DeleteUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user bad request response has a 5xx status code
func (o *DeleteUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user bad request response a status code equal to that given
func (o *DeleteUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete user bad request response
func (o *DeleteUserBadRequest) Code() int {
	return 400
}

func (o *DeleteUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] deleteUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserBadRequest) String() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] deleteUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserForbidden creates a DeleteUserForbidden with default headers values
func NewDeleteUserForbidden() *DeleteUserForbidden {
	return &DeleteUserForbidden{}
}

/*
DeleteUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteUserForbidden struct {

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

// IsSuccess returns true when this delete user forbidden response has a 2xx status code
func (o *DeleteUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user forbidden response has a 3xx status code
func (o *DeleteUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user forbidden response has a 4xx status code
func (o *DeleteUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user forbidden response has a 5xx status code
func (o *DeleteUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user forbidden response a status code equal to that given
func (o *DeleteUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete user forbidden response
func (o *DeleteUserForbidden) Code() int {
	return 403
}

func (o *DeleteUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] deleteUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] deleteUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserNotFound creates a DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {
	return &DeleteUserNotFound{}
}

/*
DeleteUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUserNotFound struct {

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

// IsSuccess returns true when this delete user not found response has a 2xx status code
func (o *DeleteUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user not found response has a 3xx status code
func (o *DeleteUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user not found response has a 4xx status code
func (o *DeleteUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user not found response has a 5xx status code
func (o *DeleteUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user not found response a status code equal to that given
func (o *DeleteUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete user not found response
func (o *DeleteUserNotFound) Code() int {
	return 404
}

func (o *DeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] deleteUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserNotFound) String() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] deleteUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserNotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserTooManyRequests creates a DeleteUserTooManyRequests with default headers values
func NewDeleteUserTooManyRequests() *DeleteUserTooManyRequests {
	return &DeleteUserTooManyRequests{}
}

/*
DeleteUserTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteUserTooManyRequests struct {

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

// IsSuccess returns true when this delete user too many requests response has a 2xx status code
func (o *DeleteUserTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user too many requests response has a 3xx status code
func (o *DeleteUserTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user too many requests response has a 4xx status code
func (o *DeleteUserTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user too many requests response has a 5xx status code
func (o *DeleteUserTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user too many requests response a status code equal to that given
func (o *DeleteUserTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete user too many requests response
func (o *DeleteUserTooManyRequests) Code() int {
	return 429
}

func (o *DeleteUserTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] deleteUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteUserTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] deleteUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteUserTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserDefault creates a DeleteUserDefault with default headers values
func NewDeleteUserDefault(code int) *DeleteUserDefault {
	return &DeleteUserDefault{
		_statusCode: code,
	}
}

/*
DeleteUserDefault describes a response with status code -1, with default header values.

OK
*/
type DeleteUserDefault struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this delete user default response has a 2xx status code
func (o *DeleteUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete user default response has a 3xx status code
func (o *DeleteUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete user default response has a 4xx status code
func (o *DeleteUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete user default response has a 5xx status code
func (o *DeleteUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete user default response a status code equal to that given
func (o *DeleteUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete user default response
func (o *DeleteUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] DeleteUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUserDefault) String() string {
	return fmt.Sprintf("[DELETE /users/entities/users/v1][%d] DeleteUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUserDefault) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
