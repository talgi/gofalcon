// Code generated by go-swagger; DO NOT EDIT.

package zero_trust_assessment

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

// GetAssessmentsByScoreV1Reader is a Reader for the GetAssessmentsByScoreV1 structure.
type GetAssessmentsByScoreV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssessmentsByScoreV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssessmentsByScoreV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAssessmentsByScoreV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAssessmentsByScoreV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAssessmentsByScoreV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAssessmentsByScoreV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /zero-trust-assessment/queries/assessments/v1] getAssessmentsByScoreV1", response, response.Code())
	}
}

// NewGetAssessmentsByScoreV1OK creates a GetAssessmentsByScoreV1OK with default headers values
func NewGetAssessmentsByScoreV1OK() *GetAssessmentsByScoreV1OK {
	return &GetAssessmentsByScoreV1OK{}
}

/*
GetAssessmentsByScoreV1OK describes a response with status code 200, with default header values.

OK
*/
type GetAssessmentsByScoreV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAssessmentsByScoreResponse
}

// IsSuccess returns true when this get assessments by score v1 o k response has a 2xx status code
func (o *GetAssessmentsByScoreV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get assessments by score v1 o k response has a 3xx status code
func (o *GetAssessmentsByScoreV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assessments by score v1 o k response has a 4xx status code
func (o *GetAssessmentsByScoreV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get assessments by score v1 o k response has a 5xx status code
func (o *GetAssessmentsByScoreV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get assessments by score v1 o k response a status code equal to that given
func (o *GetAssessmentsByScoreV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get assessments by score v1 o k response
func (o *GetAssessmentsByScoreV1OK) Code() int {
	return 200
}

func (o *GetAssessmentsByScoreV1OK) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/queries/assessments/v1][%d] getAssessmentsByScoreV1OK  %+v", 200, o.Payload)
}

func (o *GetAssessmentsByScoreV1OK) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/queries/assessments/v1][%d] getAssessmentsByScoreV1OK  %+v", 200, o.Payload)
}

func (o *GetAssessmentsByScoreV1OK) GetPayload() *models.DomainAssessmentsByScoreResponse {
	return o.Payload
}

func (o *GetAssessmentsByScoreV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAssessmentsByScoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssessmentsByScoreV1BadRequest creates a GetAssessmentsByScoreV1BadRequest with default headers values
func NewGetAssessmentsByScoreV1BadRequest() *GetAssessmentsByScoreV1BadRequest {
	return &GetAssessmentsByScoreV1BadRequest{}
}

/*
GetAssessmentsByScoreV1BadRequest describes a response with status code 400, with default header values.

Number of agent IDs exceeds the limit of 1000.
*/
type GetAssessmentsByScoreV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAssessmentsByScoreResponse
}

// IsSuccess returns true when this get assessments by score v1 bad request response has a 2xx status code
func (o *GetAssessmentsByScoreV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assessments by score v1 bad request response has a 3xx status code
func (o *GetAssessmentsByScoreV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assessments by score v1 bad request response has a 4xx status code
func (o *GetAssessmentsByScoreV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get assessments by score v1 bad request response has a 5xx status code
func (o *GetAssessmentsByScoreV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get assessments by score v1 bad request response a status code equal to that given
func (o *GetAssessmentsByScoreV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get assessments by score v1 bad request response
func (o *GetAssessmentsByScoreV1BadRequest) Code() int {
	return 400
}

func (o *GetAssessmentsByScoreV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/queries/assessments/v1][%d] getAssessmentsByScoreV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetAssessmentsByScoreV1BadRequest) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/queries/assessments/v1][%d] getAssessmentsByScoreV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetAssessmentsByScoreV1BadRequest) GetPayload() *models.DomainAssessmentsByScoreResponse {
	return o.Payload
}

func (o *GetAssessmentsByScoreV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAssessmentsByScoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssessmentsByScoreV1Forbidden creates a GetAssessmentsByScoreV1Forbidden with default headers values
func NewGetAssessmentsByScoreV1Forbidden() *GetAssessmentsByScoreV1Forbidden {
	return &GetAssessmentsByScoreV1Forbidden{}
}

/*
GetAssessmentsByScoreV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAssessmentsByScoreV1Forbidden struct {

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

// IsSuccess returns true when this get assessments by score v1 forbidden response has a 2xx status code
func (o *GetAssessmentsByScoreV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assessments by score v1 forbidden response has a 3xx status code
func (o *GetAssessmentsByScoreV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assessments by score v1 forbidden response has a 4xx status code
func (o *GetAssessmentsByScoreV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get assessments by score v1 forbidden response has a 5xx status code
func (o *GetAssessmentsByScoreV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get assessments by score v1 forbidden response a status code equal to that given
func (o *GetAssessmentsByScoreV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get assessments by score v1 forbidden response
func (o *GetAssessmentsByScoreV1Forbidden) Code() int {
	return 403
}

func (o *GetAssessmentsByScoreV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/queries/assessments/v1][%d] getAssessmentsByScoreV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetAssessmentsByScoreV1Forbidden) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/queries/assessments/v1][%d] getAssessmentsByScoreV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetAssessmentsByScoreV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetAssessmentsByScoreV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAssessmentsByScoreV1NotFound creates a GetAssessmentsByScoreV1NotFound with default headers values
func NewGetAssessmentsByScoreV1NotFound() *GetAssessmentsByScoreV1NotFound {
	return &GetAssessmentsByScoreV1NotFound{}
}

/*
GetAssessmentsByScoreV1NotFound describes a response with status code 404, with default header values.

One or more of the specified agent IDs is not found.
*/
type GetAssessmentsByScoreV1NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAssessmentsByScoreResponse
}

// IsSuccess returns true when this get assessments by score v1 not found response has a 2xx status code
func (o *GetAssessmentsByScoreV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assessments by score v1 not found response has a 3xx status code
func (o *GetAssessmentsByScoreV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assessments by score v1 not found response has a 4xx status code
func (o *GetAssessmentsByScoreV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get assessments by score v1 not found response has a 5xx status code
func (o *GetAssessmentsByScoreV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get assessments by score v1 not found response a status code equal to that given
func (o *GetAssessmentsByScoreV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get assessments by score v1 not found response
func (o *GetAssessmentsByScoreV1NotFound) Code() int {
	return 404
}

func (o *GetAssessmentsByScoreV1NotFound) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/queries/assessments/v1][%d] getAssessmentsByScoreV1NotFound  %+v", 404, o.Payload)
}

func (o *GetAssessmentsByScoreV1NotFound) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/queries/assessments/v1][%d] getAssessmentsByScoreV1NotFound  %+v", 404, o.Payload)
}

func (o *GetAssessmentsByScoreV1NotFound) GetPayload() *models.DomainAssessmentsByScoreResponse {
	return o.Payload
}

func (o *GetAssessmentsByScoreV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAssessmentsByScoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssessmentsByScoreV1TooManyRequests creates a GetAssessmentsByScoreV1TooManyRequests with default headers values
func NewGetAssessmentsByScoreV1TooManyRequests() *GetAssessmentsByScoreV1TooManyRequests {
	return &GetAssessmentsByScoreV1TooManyRequests{}
}

/*
GetAssessmentsByScoreV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAssessmentsByScoreV1TooManyRequests struct {

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

// IsSuccess returns true when this get assessments by score v1 too many requests response has a 2xx status code
func (o *GetAssessmentsByScoreV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assessments by score v1 too many requests response has a 3xx status code
func (o *GetAssessmentsByScoreV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assessments by score v1 too many requests response has a 4xx status code
func (o *GetAssessmentsByScoreV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get assessments by score v1 too many requests response has a 5xx status code
func (o *GetAssessmentsByScoreV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get assessments by score v1 too many requests response a status code equal to that given
func (o *GetAssessmentsByScoreV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get assessments by score v1 too many requests response
func (o *GetAssessmentsByScoreV1TooManyRequests) Code() int {
	return 429
}

func (o *GetAssessmentsByScoreV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/queries/assessments/v1][%d] getAssessmentsByScoreV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAssessmentsByScoreV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/queries/assessments/v1][%d] getAssessmentsByScoreV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAssessmentsByScoreV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetAssessmentsByScoreV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
