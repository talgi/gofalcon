// Code generated by go-swagger; DO NOT EDIT.

package response_policies

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

// QueryCombinedRTResponsePoliciesReader is a Reader for the QueryCombinedRTResponsePolicies structure.
type QueryCombinedRTResponsePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedRTResponsePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedRTResponsePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedRTResponsePoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedRTResponsePoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedRTResponsePoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedRTResponsePoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryCombinedRTResponsePoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryCombinedRTResponsePoliciesOK creates a QueryCombinedRTResponsePoliciesOK with default headers values
func NewQueryCombinedRTResponsePoliciesOK() *QueryCombinedRTResponsePoliciesOK {
	return &QueryCombinedRTResponsePoliciesOK{}
}

/*
QueryCombinedRTResponsePoliciesOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedRTResponsePoliciesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesRTResponsePoliciesV1
}

// IsSuccess returns true when this query combined r t response policies o k response has a 2xx status code
func (o *QueryCombinedRTResponsePoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query combined r t response policies o k response has a 3xx status code
func (o *QueryCombinedRTResponsePoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policies o k response has a 4xx status code
func (o *QueryCombinedRTResponsePoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined r t response policies o k response has a 5xx status code
func (o *QueryCombinedRTResponsePoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined r t response policies o k response a status code equal to that given
func (o *QueryCombinedRTResponsePoliciesOK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryCombinedRTResponsePoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesOK) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedRTResponsePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedRTResponsePoliciesBadRequest creates a QueryCombinedRTResponsePoliciesBadRequest with default headers values
func NewQueryCombinedRTResponsePoliciesBadRequest() *QueryCombinedRTResponsePoliciesBadRequest {
	return &QueryCombinedRTResponsePoliciesBadRequest{}
}

/*
QueryCombinedRTResponsePoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedRTResponsePoliciesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesRTResponsePoliciesV1
}

// IsSuccess returns true when this query combined r t response policies bad request response has a 2xx status code
func (o *QueryCombinedRTResponsePoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined r t response policies bad request response has a 3xx status code
func (o *QueryCombinedRTResponsePoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policies bad request response has a 4xx status code
func (o *QueryCombinedRTResponsePoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined r t response policies bad request response has a 5xx status code
func (o *QueryCombinedRTResponsePoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined r t response policies bad request response a status code equal to that given
func (o *QueryCombinedRTResponsePoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *QueryCombinedRTResponsePoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesBadRequest) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedRTResponsePoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedRTResponsePoliciesForbidden creates a QueryCombinedRTResponsePoliciesForbidden with default headers values
func NewQueryCombinedRTResponsePoliciesForbidden() *QueryCombinedRTResponsePoliciesForbidden {
	return &QueryCombinedRTResponsePoliciesForbidden{}
}

/*
QueryCombinedRTResponsePoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedRTResponsePoliciesForbidden struct {

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

// IsSuccess returns true when this query combined r t response policies forbidden response has a 2xx status code
func (o *QueryCombinedRTResponsePoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined r t response policies forbidden response has a 3xx status code
func (o *QueryCombinedRTResponsePoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policies forbidden response has a 4xx status code
func (o *QueryCombinedRTResponsePoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined r t response policies forbidden response has a 5xx status code
func (o *QueryCombinedRTResponsePoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined r t response policies forbidden response a status code equal to that given
func (o *QueryCombinedRTResponsePoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *QueryCombinedRTResponsePoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedRTResponsePoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedRTResponsePoliciesTooManyRequests creates a QueryCombinedRTResponsePoliciesTooManyRequests with default headers values
func NewQueryCombinedRTResponsePoliciesTooManyRequests() *QueryCombinedRTResponsePoliciesTooManyRequests {
	return &QueryCombinedRTResponsePoliciesTooManyRequests{}
}

/*
QueryCombinedRTResponsePoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedRTResponsePoliciesTooManyRequests struct {

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

// IsSuccess returns true when this query combined r t response policies too many requests response has a 2xx status code
func (o *QueryCombinedRTResponsePoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined r t response policies too many requests response has a 3xx status code
func (o *QueryCombinedRTResponsePoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policies too many requests response has a 4xx status code
func (o *QueryCombinedRTResponsePoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined r t response policies too many requests response has a 5xx status code
func (o *QueryCombinedRTResponsePoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined r t response policies too many requests response a status code equal to that given
func (o *QueryCombinedRTResponsePoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *QueryCombinedRTResponsePoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedRTResponsePoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedRTResponsePoliciesInternalServerError creates a QueryCombinedRTResponsePoliciesInternalServerError with default headers values
func NewQueryCombinedRTResponsePoliciesInternalServerError() *QueryCombinedRTResponsePoliciesInternalServerError {
	return &QueryCombinedRTResponsePoliciesInternalServerError{}
}

/*
QueryCombinedRTResponsePoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedRTResponsePoliciesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesRTResponsePoliciesV1
}

// IsSuccess returns true when this query combined r t response policies internal server error response has a 2xx status code
func (o *QueryCombinedRTResponsePoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined r t response policies internal server error response has a 3xx status code
func (o *QueryCombinedRTResponsePoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policies internal server error response has a 4xx status code
func (o *QueryCombinedRTResponsePoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined r t response policies internal server error response has a 5xx status code
func (o *QueryCombinedRTResponsePoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query combined r t response policies internal server error response a status code equal to that given
func (o *QueryCombinedRTResponsePoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *QueryCombinedRTResponsePoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesInternalServerError) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedRTResponsePoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedRTResponsePoliciesDefault creates a QueryCombinedRTResponsePoliciesDefault with default headers values
func NewQueryCombinedRTResponsePoliciesDefault(code int) *QueryCombinedRTResponsePoliciesDefault {
	return &QueryCombinedRTResponsePoliciesDefault{
		_statusCode: code,
	}
}

/*
QueryCombinedRTResponsePoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type QueryCombinedRTResponsePoliciesDefault struct {
	_statusCode int

	Payload *models.ResponsesRTResponsePoliciesV1
}

// Code gets the status code for the query combined r t response policies default response
func (o *QueryCombinedRTResponsePoliciesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this query combined r t response policies default response has a 2xx status code
func (o *QueryCombinedRTResponsePoliciesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query combined r t response policies default response has a 3xx status code
func (o *QueryCombinedRTResponsePoliciesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query combined r t response policies default response has a 4xx status code
func (o *QueryCombinedRTResponsePoliciesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query combined r t response policies default response has a 5xx status code
func (o *QueryCombinedRTResponsePoliciesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query combined r t response policies default response a status code equal to that given
func (o *QueryCombinedRTResponsePoliciesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QueryCombinedRTResponsePoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePolicies default  %+v", o._statusCode, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesDefault) String() string {
	return fmt.Sprintf("[GET /policy/combined/response/v1][%d] queryCombinedRTResponsePolicies default  %+v", o._statusCode, o.Payload)
}

func (o *QueryCombinedRTResponsePoliciesDefault) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedRTResponsePoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
