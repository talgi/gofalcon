// Code generated by go-swagger; DO NOT EDIT.

package host_group

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

// QueryGroupMembersReader is a Reader for the QueryGroupMembers structure.
type QueryGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryGroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryGroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryGroupMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryGroupMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryGroupMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryGroupMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryGroupMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryGroupMembersOK creates a QueryGroupMembersOK with default headers values
func NewQueryGroupMembersOK() *QueryGroupMembersOK {
	return &QueryGroupMembersOK{}
}

/*
QueryGroupMembersOK describes a response with status code 200, with default header values.

OK
*/
type QueryGroupMembersOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query group members o k response has a 2xx status code
func (o *QueryGroupMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query group members o k response has a 3xx status code
func (o *QueryGroupMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query group members o k response has a 4xx status code
func (o *QueryGroupMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query group members o k response has a 5xx status code
func (o *QueryGroupMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query group members o k response a status code equal to that given
func (o *QueryGroupMembersOK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryGroupMembersOK) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersOK  %+v", 200, o.Payload)
}

func (o *QueryGroupMembersOK) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersOK  %+v", 200, o.Payload)
}

func (o *QueryGroupMembersOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryGroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryGroupMembersBadRequest creates a QueryGroupMembersBadRequest with default headers values
func NewQueryGroupMembersBadRequest() *QueryGroupMembersBadRequest {
	return &QueryGroupMembersBadRequest{}
}

/*
QueryGroupMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryGroupMembersBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query group members bad request response has a 2xx status code
func (o *QueryGroupMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query group members bad request response has a 3xx status code
func (o *QueryGroupMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query group members bad request response has a 4xx status code
func (o *QueryGroupMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query group members bad request response has a 5xx status code
func (o *QueryGroupMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query group members bad request response a status code equal to that given
func (o *QueryGroupMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *QueryGroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryGroupMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryGroupMembersBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryGroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryGroupMembersForbidden creates a QueryGroupMembersForbidden with default headers values
func NewQueryGroupMembersForbidden() *QueryGroupMembersForbidden {
	return &QueryGroupMembersForbidden{}
}

/*
QueryGroupMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryGroupMembersForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query group members forbidden response has a 2xx status code
func (o *QueryGroupMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query group members forbidden response has a 3xx status code
func (o *QueryGroupMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query group members forbidden response has a 4xx status code
func (o *QueryGroupMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query group members forbidden response has a 5xx status code
func (o *QueryGroupMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query group members forbidden response a status code equal to that given
func (o *QueryGroupMembersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *QueryGroupMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryGroupMembersForbidden) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryGroupMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryGroupMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryGroupMembersNotFound creates a QueryGroupMembersNotFound with default headers values
func NewQueryGroupMembersNotFound() *QueryGroupMembersNotFound {
	return &QueryGroupMembersNotFound{}
}

/*
QueryGroupMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryGroupMembersNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query group members not found response has a 2xx status code
func (o *QueryGroupMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query group members not found response has a 3xx status code
func (o *QueryGroupMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query group members not found response has a 4xx status code
func (o *QueryGroupMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query group members not found response has a 5xx status code
func (o *QueryGroupMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query group members not found response a status code equal to that given
func (o *QueryGroupMembersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *QueryGroupMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryGroupMembersNotFound) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryGroupMembersNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryGroupMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryGroupMembersTooManyRequests creates a QueryGroupMembersTooManyRequests with default headers values
func NewQueryGroupMembersTooManyRequests() *QueryGroupMembersTooManyRequests {
	return &QueryGroupMembersTooManyRequests{}
}

/*
QueryGroupMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryGroupMembersTooManyRequests struct {

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

// IsSuccess returns true when this query group members too many requests response has a 2xx status code
func (o *QueryGroupMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query group members too many requests response has a 3xx status code
func (o *QueryGroupMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query group members too many requests response has a 4xx status code
func (o *QueryGroupMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query group members too many requests response has a 5xx status code
func (o *QueryGroupMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query group members too many requests response a status code equal to that given
func (o *QueryGroupMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *QueryGroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryGroupMembersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryGroupMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryGroupMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryGroupMembersInternalServerError creates a QueryGroupMembersInternalServerError with default headers values
func NewQueryGroupMembersInternalServerError() *QueryGroupMembersInternalServerError {
	return &QueryGroupMembersInternalServerError{}
}

/*
QueryGroupMembersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryGroupMembersInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query group members internal server error response has a 2xx status code
func (o *QueryGroupMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query group members internal server error response has a 3xx status code
func (o *QueryGroupMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query group members internal server error response has a 4xx status code
func (o *QueryGroupMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query group members internal server error response has a 5xx status code
func (o *QueryGroupMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query group members internal server error response a status code equal to that given
func (o *QueryGroupMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *QueryGroupMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryGroupMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryGroupMembersInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryGroupMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryGroupMembersDefault creates a QueryGroupMembersDefault with default headers values
func NewQueryGroupMembersDefault(code int) *QueryGroupMembersDefault {
	return &QueryGroupMembersDefault{
		_statusCode: code,
	}
}

/*
QueryGroupMembersDefault describes a response with status code -1, with default header values.

OK
*/
type QueryGroupMembersDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query group members default response
func (o *QueryGroupMembersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this query group members default response has a 2xx status code
func (o *QueryGroupMembersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query group members default response has a 3xx status code
func (o *QueryGroupMembersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query group members default response has a 4xx status code
func (o *QueryGroupMembersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query group members default response has a 5xx status code
func (o *QueryGroupMembersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query group members default response a status code equal to that given
func (o *QueryGroupMembersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QueryGroupMembersDefault) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryGroupMembersDefault) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryGroupMembersDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryGroupMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
