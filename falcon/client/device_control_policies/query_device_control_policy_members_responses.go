// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

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

// QueryDeviceControlPolicyMembersReader is a Reader for the QueryDeviceControlPolicyMembers structure.
type QueryDeviceControlPolicyMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDeviceControlPolicyMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryDeviceControlPolicyMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryDeviceControlPolicyMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryDeviceControlPolicyMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryDeviceControlPolicyMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryDeviceControlPolicyMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryDeviceControlPolicyMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /policy/queries/device-control-members/v1] queryDeviceControlPolicyMembers", response, response.Code())
	}
}

// NewQueryDeviceControlPolicyMembersOK creates a QueryDeviceControlPolicyMembersOK with default headers values
func NewQueryDeviceControlPolicyMembersOK() *QueryDeviceControlPolicyMembersOK {
	return &QueryDeviceControlPolicyMembersOK{}
}

/*
QueryDeviceControlPolicyMembersOK describes a response with status code 200, with default header values.

OK
*/
type QueryDeviceControlPolicyMembersOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query device control policy members o k response has a 2xx status code
func (o *QueryDeviceControlPolicyMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query device control policy members o k response has a 3xx status code
func (o *QueryDeviceControlPolicyMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policy members o k response has a 4xx status code
func (o *QueryDeviceControlPolicyMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query device control policy members o k response has a 5xx status code
func (o *QueryDeviceControlPolicyMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query device control policy members o k response a status code equal to that given
func (o *QueryDeviceControlPolicyMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query device control policy members o k response
func (o *QueryDeviceControlPolicyMembersOK) Code() int {
	return 200
}

func (o *QueryDeviceControlPolicyMembersOK) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersOK  %+v", 200, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersOK) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersOK  %+v", 200, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDeviceControlPolicyMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDeviceControlPolicyMembersBadRequest creates a QueryDeviceControlPolicyMembersBadRequest with default headers values
func NewQueryDeviceControlPolicyMembersBadRequest() *QueryDeviceControlPolicyMembersBadRequest {
	return &QueryDeviceControlPolicyMembersBadRequest{}
}

/*
QueryDeviceControlPolicyMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryDeviceControlPolicyMembersBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query device control policy members bad request response has a 2xx status code
func (o *QueryDeviceControlPolicyMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device control policy members bad request response has a 3xx status code
func (o *QueryDeviceControlPolicyMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policy members bad request response has a 4xx status code
func (o *QueryDeviceControlPolicyMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query device control policy members bad request response has a 5xx status code
func (o *QueryDeviceControlPolicyMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query device control policy members bad request response a status code equal to that given
func (o *QueryDeviceControlPolicyMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query device control policy members bad request response
func (o *QueryDeviceControlPolicyMembersBadRequest) Code() int {
	return 400
}

func (o *QueryDeviceControlPolicyMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDeviceControlPolicyMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDeviceControlPolicyMembersForbidden creates a QueryDeviceControlPolicyMembersForbidden with default headers values
func NewQueryDeviceControlPolicyMembersForbidden() *QueryDeviceControlPolicyMembersForbidden {
	return &QueryDeviceControlPolicyMembersForbidden{}
}

/*
QueryDeviceControlPolicyMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryDeviceControlPolicyMembersForbidden struct {

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

// IsSuccess returns true when this query device control policy members forbidden response has a 2xx status code
func (o *QueryDeviceControlPolicyMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device control policy members forbidden response has a 3xx status code
func (o *QueryDeviceControlPolicyMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policy members forbidden response has a 4xx status code
func (o *QueryDeviceControlPolicyMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query device control policy members forbidden response has a 5xx status code
func (o *QueryDeviceControlPolicyMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query device control policy members forbidden response a status code equal to that given
func (o *QueryDeviceControlPolicyMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query device control policy members forbidden response
func (o *QueryDeviceControlPolicyMembersForbidden) Code() int {
	return 403
}

func (o *QueryDeviceControlPolicyMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersForbidden) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryDeviceControlPolicyMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDeviceControlPolicyMembersNotFound creates a QueryDeviceControlPolicyMembersNotFound with default headers values
func NewQueryDeviceControlPolicyMembersNotFound() *QueryDeviceControlPolicyMembersNotFound {
	return &QueryDeviceControlPolicyMembersNotFound{}
}

/*
QueryDeviceControlPolicyMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryDeviceControlPolicyMembersNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query device control policy members not found response has a 2xx status code
func (o *QueryDeviceControlPolicyMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device control policy members not found response has a 3xx status code
func (o *QueryDeviceControlPolicyMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policy members not found response has a 4xx status code
func (o *QueryDeviceControlPolicyMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query device control policy members not found response has a 5xx status code
func (o *QueryDeviceControlPolicyMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query device control policy members not found response a status code equal to that given
func (o *QueryDeviceControlPolicyMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the query device control policy members not found response
func (o *QueryDeviceControlPolicyMembersNotFound) Code() int {
	return 404
}

func (o *QueryDeviceControlPolicyMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersNotFound) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDeviceControlPolicyMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDeviceControlPolicyMembersTooManyRequests creates a QueryDeviceControlPolicyMembersTooManyRequests with default headers values
func NewQueryDeviceControlPolicyMembersTooManyRequests() *QueryDeviceControlPolicyMembersTooManyRequests {
	return &QueryDeviceControlPolicyMembersTooManyRequests{}
}

/*
QueryDeviceControlPolicyMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryDeviceControlPolicyMembersTooManyRequests struct {

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

// IsSuccess returns true when this query device control policy members too many requests response has a 2xx status code
func (o *QueryDeviceControlPolicyMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device control policy members too many requests response has a 3xx status code
func (o *QueryDeviceControlPolicyMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policy members too many requests response has a 4xx status code
func (o *QueryDeviceControlPolicyMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query device control policy members too many requests response has a 5xx status code
func (o *QueryDeviceControlPolicyMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query device control policy members too many requests response a status code equal to that given
func (o *QueryDeviceControlPolicyMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query device control policy members too many requests response
func (o *QueryDeviceControlPolicyMembersTooManyRequests) Code() int {
	return 429
}

func (o *QueryDeviceControlPolicyMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryDeviceControlPolicyMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDeviceControlPolicyMembersInternalServerError creates a QueryDeviceControlPolicyMembersInternalServerError with default headers values
func NewQueryDeviceControlPolicyMembersInternalServerError() *QueryDeviceControlPolicyMembersInternalServerError {
	return &QueryDeviceControlPolicyMembersInternalServerError{}
}

/*
QueryDeviceControlPolicyMembersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryDeviceControlPolicyMembersInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query device control policy members internal server error response has a 2xx status code
func (o *QueryDeviceControlPolicyMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device control policy members internal server error response has a 3xx status code
func (o *QueryDeviceControlPolicyMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policy members internal server error response has a 4xx status code
func (o *QueryDeviceControlPolicyMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query device control policy members internal server error response has a 5xx status code
func (o *QueryDeviceControlPolicyMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query device control policy members internal server error response a status code equal to that given
func (o *QueryDeviceControlPolicyMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query device control policy members internal server error response
func (o *QueryDeviceControlPolicyMembersInternalServerError) Code() int {
	return 500
}

func (o *QueryDeviceControlPolicyMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control-members/v1][%d] queryDeviceControlPolicyMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryDeviceControlPolicyMembersInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDeviceControlPolicyMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
