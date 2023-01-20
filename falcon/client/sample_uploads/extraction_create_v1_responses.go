// Code generated by go-swagger; DO NOT EDIT.

package sample_uploads

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

// ExtractionCreateV1Reader is a Reader for the ExtractionCreateV1 structure.
type ExtractionCreateV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtractionCreateV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtractionCreateV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewExtractionCreateV1Accepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtractionCreateV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExtractionCreateV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewExtractionCreateV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExtractionCreateV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtractionCreateV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtractionCreateV1OK creates a ExtractionCreateV1OK with default headers values
func NewExtractionCreateV1OK() *ExtractionCreateV1OK {
	return &ExtractionCreateV1OK{}
}

/*
ExtractionCreateV1OK describes a response with status code 200, with default header values.

OK
*/
type ExtractionCreateV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientExtractionCreateResponseV1
}

// IsSuccess returns true when this extraction create v1 o k response has a 2xx status code
func (o *ExtractionCreateV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extraction create v1 o k response has a 3xx status code
func (o *ExtractionCreateV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction create v1 o k response has a 4xx status code
func (o *ExtractionCreateV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extraction create v1 o k response has a 5xx status code
func (o *ExtractionCreateV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this extraction create v1 o k response a status code equal to that given
func (o *ExtractionCreateV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extraction create v1 o k response
func (o *ExtractionCreateV1OK) Code() int {
	return 200
}

func (o *ExtractionCreateV1OK) Error() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1OK  %+v", 200, o.Payload)
}

func (o *ExtractionCreateV1OK) String() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1OK  %+v", 200, o.Payload)
}

func (o *ExtractionCreateV1OK) GetPayload() *models.ClientExtractionCreateResponseV1 {
	return o.Payload
}

func (o *ExtractionCreateV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientExtractionCreateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtractionCreateV1Accepted creates a ExtractionCreateV1Accepted with default headers values
func NewExtractionCreateV1Accepted() *ExtractionCreateV1Accepted {
	return &ExtractionCreateV1Accepted{}
}

/*
ExtractionCreateV1Accepted describes a response with status code 202, with default header values.

OK
*/
type ExtractionCreateV1Accepted struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientExtractionCreateResponseV1
}

// IsSuccess returns true when this extraction create v1 accepted response has a 2xx status code
func (o *ExtractionCreateV1Accepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extraction create v1 accepted response has a 3xx status code
func (o *ExtractionCreateV1Accepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction create v1 accepted response has a 4xx status code
func (o *ExtractionCreateV1Accepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this extraction create v1 accepted response has a 5xx status code
func (o *ExtractionCreateV1Accepted) IsServerError() bool {
	return false
}

// IsCode returns true when this extraction create v1 accepted response a status code equal to that given
func (o *ExtractionCreateV1Accepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the extraction create v1 accepted response
func (o *ExtractionCreateV1Accepted) Code() int {
	return 202
}

func (o *ExtractionCreateV1Accepted) Error() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1Accepted  %+v", 202, o.Payload)
}

func (o *ExtractionCreateV1Accepted) String() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1Accepted  %+v", 202, o.Payload)
}

func (o *ExtractionCreateV1Accepted) GetPayload() *models.ClientExtractionCreateResponseV1 {
	return o.Payload
}

func (o *ExtractionCreateV1Accepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientExtractionCreateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtractionCreateV1BadRequest creates a ExtractionCreateV1BadRequest with default headers values
func NewExtractionCreateV1BadRequest() *ExtractionCreateV1BadRequest {
	return &ExtractionCreateV1BadRequest{}
}

/*
ExtractionCreateV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtractionCreateV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientExtractionCreateResponseV1
}

// IsSuccess returns true when this extraction create v1 bad request response has a 2xx status code
func (o *ExtractionCreateV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extraction create v1 bad request response has a 3xx status code
func (o *ExtractionCreateV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction create v1 bad request response has a 4xx status code
func (o *ExtractionCreateV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extraction create v1 bad request response has a 5xx status code
func (o *ExtractionCreateV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extraction create v1 bad request response a status code equal to that given
func (o *ExtractionCreateV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the extraction create v1 bad request response
func (o *ExtractionCreateV1BadRequest) Code() int {
	return 400
}

func (o *ExtractionCreateV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1BadRequest  %+v", 400, o.Payload)
}

func (o *ExtractionCreateV1BadRequest) String() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1BadRequest  %+v", 400, o.Payload)
}

func (o *ExtractionCreateV1BadRequest) GetPayload() *models.ClientExtractionCreateResponseV1 {
	return o.Payload
}

func (o *ExtractionCreateV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientExtractionCreateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtractionCreateV1Forbidden creates a ExtractionCreateV1Forbidden with default headers values
func NewExtractionCreateV1Forbidden() *ExtractionCreateV1Forbidden {
	return &ExtractionCreateV1Forbidden{}
}

/*
ExtractionCreateV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExtractionCreateV1Forbidden struct {

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

// IsSuccess returns true when this extraction create v1 forbidden response has a 2xx status code
func (o *ExtractionCreateV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extraction create v1 forbidden response has a 3xx status code
func (o *ExtractionCreateV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction create v1 forbidden response has a 4xx status code
func (o *ExtractionCreateV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this extraction create v1 forbidden response has a 5xx status code
func (o *ExtractionCreateV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this extraction create v1 forbidden response a status code equal to that given
func (o *ExtractionCreateV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the extraction create v1 forbidden response
func (o *ExtractionCreateV1Forbidden) Code() int {
	return 403
}

func (o *ExtractionCreateV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1Forbidden  %+v", 403, o.Payload)
}

func (o *ExtractionCreateV1Forbidden) String() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1Forbidden  %+v", 403, o.Payload)
}

func (o *ExtractionCreateV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ExtractionCreateV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExtractionCreateV1TooManyRequests creates a ExtractionCreateV1TooManyRequests with default headers values
func NewExtractionCreateV1TooManyRequests() *ExtractionCreateV1TooManyRequests {
	return &ExtractionCreateV1TooManyRequests{}
}

/*
ExtractionCreateV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ExtractionCreateV1TooManyRequests struct {

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

// IsSuccess returns true when this extraction create v1 too many requests response has a 2xx status code
func (o *ExtractionCreateV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extraction create v1 too many requests response has a 3xx status code
func (o *ExtractionCreateV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction create v1 too many requests response has a 4xx status code
func (o *ExtractionCreateV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this extraction create v1 too many requests response has a 5xx status code
func (o *ExtractionCreateV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this extraction create v1 too many requests response a status code equal to that given
func (o *ExtractionCreateV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the extraction create v1 too many requests response
func (o *ExtractionCreateV1TooManyRequests) Code() int {
	return 429
}

func (o *ExtractionCreateV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *ExtractionCreateV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *ExtractionCreateV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ExtractionCreateV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExtractionCreateV1InternalServerError creates a ExtractionCreateV1InternalServerError with default headers values
func NewExtractionCreateV1InternalServerError() *ExtractionCreateV1InternalServerError {
	return &ExtractionCreateV1InternalServerError{}
}

/*
ExtractionCreateV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ExtractionCreateV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientExtractionCreateResponseV1
}

// IsSuccess returns true when this extraction create v1 internal server error response has a 2xx status code
func (o *ExtractionCreateV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extraction create v1 internal server error response has a 3xx status code
func (o *ExtractionCreateV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction create v1 internal server error response has a 4xx status code
func (o *ExtractionCreateV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this extraction create v1 internal server error response has a 5xx status code
func (o *ExtractionCreateV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this extraction create v1 internal server error response a status code equal to that given
func (o *ExtractionCreateV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the extraction create v1 internal server error response
func (o *ExtractionCreateV1InternalServerError) Code() int {
	return 500
}

func (o *ExtractionCreateV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1InternalServerError  %+v", 500, o.Payload)
}

func (o *ExtractionCreateV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] extractionCreateV1InternalServerError  %+v", 500, o.Payload)
}

func (o *ExtractionCreateV1InternalServerError) GetPayload() *models.ClientExtractionCreateResponseV1 {
	return o.Payload
}

func (o *ExtractionCreateV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientExtractionCreateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtractionCreateV1Default creates a ExtractionCreateV1Default with default headers values
func NewExtractionCreateV1Default(code int) *ExtractionCreateV1Default {
	return &ExtractionCreateV1Default{
		_statusCode: code,
	}
}

/*
ExtractionCreateV1Default describes a response with status code -1, with default header values.

OK
*/
type ExtractionCreateV1Default struct {
	_statusCode int

	Payload *models.ClientExtractionCreateResponseV1
}

// IsSuccess returns true when this extraction create v1 default response has a 2xx status code
func (o *ExtractionCreateV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extraction create v1 default response has a 3xx status code
func (o *ExtractionCreateV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extraction create v1 default response has a 4xx status code
func (o *ExtractionCreateV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extraction create v1 default response has a 5xx status code
func (o *ExtractionCreateV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extraction create v1 default response a status code equal to that given
func (o *ExtractionCreateV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extraction create v1 default response
func (o *ExtractionCreateV1Default) Code() int {
	return o._statusCode
}

func (o *ExtractionCreateV1Default) Error() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] ExtractionCreateV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ExtractionCreateV1Default) String() string {
	return fmt.Sprintf("[POST /archives/entities/extractions/v1][%d] ExtractionCreateV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ExtractionCreateV1Default) GetPayload() *models.ClientExtractionCreateResponseV1 {
	return o.Payload
}

func (o *ExtractionCreateV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientExtractionCreateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
