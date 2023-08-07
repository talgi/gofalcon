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

// GetMalQueryEntitiesSamplesFetchV1Reader is a Reader for the GetMalQueryEntitiesSamplesFetchV1 structure.
type GetMalQueryEntitiesSamplesFetchV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMalQueryEntitiesSamplesFetchV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMalQueryEntitiesSamplesFetchV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMalQueryEntitiesSamplesFetchV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMalQueryEntitiesSamplesFetchV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMalQueryEntitiesSamplesFetchV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMalQueryEntitiesSamplesFetchV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /malquery/entities/samples-fetch/v1] GetMalQueryEntitiesSamplesFetchV1", response, response.Code())
	}
}

// NewGetMalQueryEntitiesSamplesFetchV1OK creates a GetMalQueryEntitiesSamplesFetchV1OK with default headers values
func NewGetMalQueryEntitiesSamplesFetchV1OK() *GetMalQueryEntitiesSamplesFetchV1OK {
	return &GetMalQueryEntitiesSamplesFetchV1OK{}
}

/*
GetMalQueryEntitiesSamplesFetchV1OK describes a response with status code 200, with default header values.

The zip archive if it's ready for download
*/
type GetMalQueryEntitiesSamplesFetchV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this get mal query entities samples fetch v1 o k response has a 2xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get mal query entities samples fetch v1 o k response has a 3xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query entities samples fetch v1 o k response has a 4xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mal query entities samples fetch v1 o k response has a 5xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query entities samples fetch v1 o k response a status code equal to that given
func (o *GetMalQueryEntitiesSamplesFetchV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get mal query entities samples fetch v1 o k response
func (o *GetMalQueryEntitiesSamplesFetchV1OK) Code() int {
	return 200
}

func (o *GetMalQueryEntitiesSamplesFetchV1OK) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/samples-fetch/v1][%d] getMalQueryEntitiesSamplesFetchV1OK ", 200)
}

func (o *GetMalQueryEntitiesSamplesFetchV1OK) String() string {
	return fmt.Sprintf("[GET /malquery/entities/samples-fetch/v1][%d] getMalQueryEntitiesSamplesFetchV1OK ", 200)
}

func (o *GetMalQueryEntitiesSamplesFetchV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewGetMalQueryEntitiesSamplesFetchV1Unauthorized creates a GetMalQueryEntitiesSamplesFetchV1Unauthorized with default headers values
func NewGetMalQueryEntitiesSamplesFetchV1Unauthorized() *GetMalQueryEntitiesSamplesFetchV1Unauthorized {
	return &GetMalQueryEntitiesSamplesFetchV1Unauthorized{}
}

/*
GetMalQueryEntitiesSamplesFetchV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetMalQueryEntitiesSamplesFetchV1Unauthorized struct {

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

// IsSuccess returns true when this get mal query entities samples fetch v1 unauthorized response has a 2xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query entities samples fetch v1 unauthorized response has a 3xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query entities samples fetch v1 unauthorized response has a 4xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query entities samples fetch v1 unauthorized response has a 5xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query entities samples fetch v1 unauthorized response a status code equal to that given
func (o *GetMalQueryEntitiesSamplesFetchV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get mal query entities samples fetch v1 unauthorized response
func (o *GetMalQueryEntitiesSamplesFetchV1Unauthorized) Code() int {
	return 401
}

func (o *GetMalQueryEntitiesSamplesFetchV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/samples-fetch/v1][%d] getMalQueryEntitiesSamplesFetchV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetMalQueryEntitiesSamplesFetchV1Unauthorized) String() string {
	return fmt.Sprintf("[GET /malquery/entities/samples-fetch/v1][%d] getMalQueryEntitiesSamplesFetchV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetMalQueryEntitiesSamplesFetchV1Unauthorized) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryEntitiesSamplesFetchV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryEntitiesSamplesFetchV1Forbidden creates a GetMalQueryEntitiesSamplesFetchV1Forbidden with default headers values
func NewGetMalQueryEntitiesSamplesFetchV1Forbidden() *GetMalQueryEntitiesSamplesFetchV1Forbidden {
	return &GetMalQueryEntitiesSamplesFetchV1Forbidden{}
}

/*
GetMalQueryEntitiesSamplesFetchV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetMalQueryEntitiesSamplesFetchV1Forbidden struct {

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

// IsSuccess returns true when this get mal query entities samples fetch v1 forbidden response has a 2xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query entities samples fetch v1 forbidden response has a 3xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query entities samples fetch v1 forbidden response has a 4xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query entities samples fetch v1 forbidden response has a 5xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query entities samples fetch v1 forbidden response a status code equal to that given
func (o *GetMalQueryEntitiesSamplesFetchV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get mal query entities samples fetch v1 forbidden response
func (o *GetMalQueryEntitiesSamplesFetchV1Forbidden) Code() int {
	return 403
}

func (o *GetMalQueryEntitiesSamplesFetchV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/samples-fetch/v1][%d] getMalQueryEntitiesSamplesFetchV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMalQueryEntitiesSamplesFetchV1Forbidden) String() string {
	return fmt.Sprintf("[GET /malquery/entities/samples-fetch/v1][%d] getMalQueryEntitiesSamplesFetchV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMalQueryEntitiesSamplesFetchV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryEntitiesSamplesFetchV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryEntitiesSamplesFetchV1TooManyRequests creates a GetMalQueryEntitiesSamplesFetchV1TooManyRequests with default headers values
func NewGetMalQueryEntitiesSamplesFetchV1TooManyRequests() *GetMalQueryEntitiesSamplesFetchV1TooManyRequests {
	return &GetMalQueryEntitiesSamplesFetchV1TooManyRequests{}
}

/*
GetMalQueryEntitiesSamplesFetchV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetMalQueryEntitiesSamplesFetchV1TooManyRequests struct {

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

// IsSuccess returns true when this get mal query entities samples fetch v1 too many requests response has a 2xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query entities samples fetch v1 too many requests response has a 3xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query entities samples fetch v1 too many requests response has a 4xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query entities samples fetch v1 too many requests response has a 5xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query entities samples fetch v1 too many requests response a status code equal to that given
func (o *GetMalQueryEntitiesSamplesFetchV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get mal query entities samples fetch v1 too many requests response
func (o *GetMalQueryEntitiesSamplesFetchV1TooManyRequests) Code() int {
	return 429
}

func (o *GetMalQueryEntitiesSamplesFetchV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/samples-fetch/v1][%d] getMalQueryEntitiesSamplesFetchV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMalQueryEntitiesSamplesFetchV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /malquery/entities/samples-fetch/v1][%d] getMalQueryEntitiesSamplesFetchV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMalQueryEntitiesSamplesFetchV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMalQueryEntitiesSamplesFetchV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryEntitiesSamplesFetchV1InternalServerError creates a GetMalQueryEntitiesSamplesFetchV1InternalServerError with default headers values
func NewGetMalQueryEntitiesSamplesFetchV1InternalServerError() *GetMalQueryEntitiesSamplesFetchV1InternalServerError {
	return &GetMalQueryEntitiesSamplesFetchV1InternalServerError{}
}

/*
GetMalQueryEntitiesSamplesFetchV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetMalQueryEntitiesSamplesFetchV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this get mal query entities samples fetch v1 internal server error response has a 2xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query entities samples fetch v1 internal server error response has a 3xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query entities samples fetch v1 internal server error response has a 4xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mal query entities samples fetch v1 internal server error response has a 5xx status code
func (o *GetMalQueryEntitiesSamplesFetchV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get mal query entities samples fetch v1 internal server error response a status code equal to that given
func (o *GetMalQueryEntitiesSamplesFetchV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get mal query entities samples fetch v1 internal server error response
func (o *GetMalQueryEntitiesSamplesFetchV1InternalServerError) Code() int {
	return 500
}

func (o *GetMalQueryEntitiesSamplesFetchV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/samples-fetch/v1][%d] getMalQueryEntitiesSamplesFetchV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMalQueryEntitiesSamplesFetchV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /malquery/entities/samples-fetch/v1][%d] getMalQueryEntitiesSamplesFetchV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMalQueryEntitiesSamplesFetchV1InternalServerError) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *GetMalQueryEntitiesSamplesFetchV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
