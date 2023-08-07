// Code generated by go-swagger; DO NOT EDIT.

package malquery

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

// GetMalQueryRequestV1Reader is a Reader for the GetMalQueryRequestV1 structure.
type GetMalQueryRequestV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMalQueryRequestV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMalQueryRequestV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMalQueryRequestV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMalQueryRequestV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMalQueryRequestV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMalQueryRequestV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMalQueryRequestV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /malquery/entities/requests/v1] GetMalQueryRequestV1", response, response.Code())
	}
}

// NewGetMalQueryRequestV1OK creates a GetMalQueryRequestV1OK with default headers values
func NewGetMalQueryRequestV1OK() *GetMalQueryRequestV1OK {
	return &GetMalQueryRequestV1OK{}
}

/*
GetMalQueryRequestV1OK describes a response with status code 200, with default header values.

OK
*/
type GetMalQueryRequestV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryRequestResponse
}

// IsSuccess returns true when this get mal query request v1 o k response has a 2xx status code
func (o *GetMalQueryRequestV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get mal query request v1 o k response has a 3xx status code
func (o *GetMalQueryRequestV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query request v1 o k response has a 4xx status code
func (o *GetMalQueryRequestV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mal query request v1 o k response has a 5xx status code
func (o *GetMalQueryRequestV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query request v1 o k response a status code equal to that given
func (o *GetMalQueryRequestV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get mal query request v1 o k response
func (o *GetMalQueryRequestV1OK) Code() int {
	return 200
}

func (o *GetMalQueryRequestV1OK) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1OK  %+v", 200, o.Payload)
}

func (o *GetMalQueryRequestV1OK) String() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1OK  %+v", 200, o.Payload)
}

func (o *GetMalQueryRequestV1OK) GetPayload() *models.MalqueryRequestResponse {
	return o.Payload
}

func (o *GetMalQueryRequestV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryRequestV1BadRequest creates a GetMalQueryRequestV1BadRequest with default headers values
func NewGetMalQueryRequestV1BadRequest() *GetMalQueryRequestV1BadRequest {
	return &GetMalQueryRequestV1BadRequest{}
}

/*
GetMalQueryRequestV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetMalQueryRequestV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryRequestResponse
}

// IsSuccess returns true when this get mal query request v1 bad request response has a 2xx status code
func (o *GetMalQueryRequestV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query request v1 bad request response has a 3xx status code
func (o *GetMalQueryRequestV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query request v1 bad request response has a 4xx status code
func (o *GetMalQueryRequestV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query request v1 bad request response has a 5xx status code
func (o *GetMalQueryRequestV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query request v1 bad request response a status code equal to that given
func (o *GetMalQueryRequestV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get mal query request v1 bad request response
func (o *GetMalQueryRequestV1BadRequest) Code() int {
	return 400
}

func (o *GetMalQueryRequestV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMalQueryRequestV1BadRequest) String() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMalQueryRequestV1BadRequest) GetPayload() *models.MalqueryRequestResponse {
	return o.Payload
}

func (o *GetMalQueryRequestV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryRequestV1Unauthorized creates a GetMalQueryRequestV1Unauthorized with default headers values
func NewGetMalQueryRequestV1Unauthorized() *GetMalQueryRequestV1Unauthorized {
	return &GetMalQueryRequestV1Unauthorized{}
}

/*
GetMalQueryRequestV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetMalQueryRequestV1Unauthorized struct {

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

// IsSuccess returns true when this get mal query request v1 unauthorized response has a 2xx status code
func (o *GetMalQueryRequestV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query request v1 unauthorized response has a 3xx status code
func (o *GetMalQueryRequestV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query request v1 unauthorized response has a 4xx status code
func (o *GetMalQueryRequestV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query request v1 unauthorized response has a 5xx status code
func (o *GetMalQueryRequestV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query request v1 unauthorized response a status code equal to that given
func (o *GetMalQueryRequestV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get mal query request v1 unauthorized response
func (o *GetMalQueryRequestV1Unauthorized) Code() int {
	return 401
}

func (o *GetMalQueryRequestV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetMalQueryRequestV1Unauthorized) String() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetMalQueryRequestV1Unauthorized) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryRequestV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryRequestV1Forbidden creates a GetMalQueryRequestV1Forbidden with default headers values
func NewGetMalQueryRequestV1Forbidden() *GetMalQueryRequestV1Forbidden {
	return &GetMalQueryRequestV1Forbidden{}
}

/*
GetMalQueryRequestV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetMalQueryRequestV1Forbidden struct {

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

// IsSuccess returns true when this get mal query request v1 forbidden response has a 2xx status code
func (o *GetMalQueryRequestV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query request v1 forbidden response has a 3xx status code
func (o *GetMalQueryRequestV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query request v1 forbidden response has a 4xx status code
func (o *GetMalQueryRequestV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query request v1 forbidden response has a 5xx status code
func (o *GetMalQueryRequestV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query request v1 forbidden response a status code equal to that given
func (o *GetMalQueryRequestV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get mal query request v1 forbidden response
func (o *GetMalQueryRequestV1Forbidden) Code() int {
	return 403
}

func (o *GetMalQueryRequestV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMalQueryRequestV1Forbidden) String() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMalQueryRequestV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryRequestV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryRequestV1TooManyRequests creates a GetMalQueryRequestV1TooManyRequests with default headers values
func NewGetMalQueryRequestV1TooManyRequests() *GetMalQueryRequestV1TooManyRequests {
	return &GetMalQueryRequestV1TooManyRequests{}
}

/*
GetMalQueryRequestV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetMalQueryRequestV1TooManyRequests struct {

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

// IsSuccess returns true when this get mal query request v1 too many requests response has a 2xx status code
func (o *GetMalQueryRequestV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query request v1 too many requests response has a 3xx status code
func (o *GetMalQueryRequestV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query request v1 too many requests response has a 4xx status code
func (o *GetMalQueryRequestV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query request v1 too many requests response has a 5xx status code
func (o *GetMalQueryRequestV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query request v1 too many requests response a status code equal to that given
func (o *GetMalQueryRequestV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get mal query request v1 too many requests response
func (o *GetMalQueryRequestV1TooManyRequests) Code() int {
	return 429
}

func (o *GetMalQueryRequestV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMalQueryRequestV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMalQueryRequestV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMalQueryRequestV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryRequestV1InternalServerError creates a GetMalQueryRequestV1InternalServerError with default headers values
func NewGetMalQueryRequestV1InternalServerError() *GetMalQueryRequestV1InternalServerError {
	return &GetMalQueryRequestV1InternalServerError{}
}

/*
GetMalQueryRequestV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetMalQueryRequestV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryRequestResponse
}

// IsSuccess returns true when this get mal query request v1 internal server error response has a 2xx status code
func (o *GetMalQueryRequestV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query request v1 internal server error response has a 3xx status code
func (o *GetMalQueryRequestV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query request v1 internal server error response has a 4xx status code
func (o *GetMalQueryRequestV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mal query request v1 internal server error response has a 5xx status code
func (o *GetMalQueryRequestV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get mal query request v1 internal server error response a status code equal to that given
func (o *GetMalQueryRequestV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get mal query request v1 internal server error response
func (o *GetMalQueryRequestV1InternalServerError) Code() int {
	return 500
}

func (o *GetMalQueryRequestV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMalQueryRequestV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMalQueryRequestV1InternalServerError) GetPayload() *models.MalqueryRequestResponse {
	return o.Payload
}

func (o *GetMalQueryRequestV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
