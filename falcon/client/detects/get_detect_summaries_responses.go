// Code generated by go-swagger; DO NOT EDIT.

package detects

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

// GetDetectSummariesReader is a Reader for the GetDetectSummaries structure.
type GetDetectSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDetectSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDetectSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDetectSummariesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDetectSummariesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDetectSummariesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDetectSummariesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDetectSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDetectSummariesOK creates a GetDetectSummariesOK with default headers values
func NewGetDetectSummariesOK() *GetDetectSummariesOK {
	return &GetDetectSummariesOK{}
}

/*
GetDetectSummariesOK describes a response with status code 200, with default header values.

OK
*/
type GetDetectSummariesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaDetectSummariesResponse
}

// IsSuccess returns true when this get detect summaries o k response has a 2xx status code
func (o *GetDetectSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get detect summaries o k response has a 3xx status code
func (o *GetDetectSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get detect summaries o k response has a 4xx status code
func (o *GetDetectSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get detect summaries o k response has a 5xx status code
func (o *GetDetectSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get detect summaries o k response a status code equal to that given
func (o *GetDetectSummariesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDetectSummariesOK) Error() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] getDetectSummariesOK  %+v", 200, o.Payload)
}

func (o *GetDetectSummariesOK) String() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] getDetectSummariesOK  %+v", 200, o.Payload)
}

func (o *GetDetectSummariesOK) GetPayload() *models.DomainMsaDetectSummariesResponse {
	return o.Payload
}

func (o *GetDetectSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaDetectSummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDetectSummariesBadRequest creates a GetDetectSummariesBadRequest with default headers values
func NewGetDetectSummariesBadRequest() *GetDetectSummariesBadRequest {
	return &GetDetectSummariesBadRequest{}
}

/*
GetDetectSummariesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDetectSummariesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaDetectSummariesResponse
}

// IsSuccess returns true when this get detect summaries bad request response has a 2xx status code
func (o *GetDetectSummariesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get detect summaries bad request response has a 3xx status code
func (o *GetDetectSummariesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get detect summaries bad request response has a 4xx status code
func (o *GetDetectSummariesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get detect summaries bad request response has a 5xx status code
func (o *GetDetectSummariesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get detect summaries bad request response a status code equal to that given
func (o *GetDetectSummariesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetDetectSummariesBadRequest) Error() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] getDetectSummariesBadRequest  %+v", 400, o.Payload)
}

func (o *GetDetectSummariesBadRequest) String() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] getDetectSummariesBadRequest  %+v", 400, o.Payload)
}

func (o *GetDetectSummariesBadRequest) GetPayload() *models.DomainMsaDetectSummariesResponse {
	return o.Payload
}

func (o *GetDetectSummariesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaDetectSummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDetectSummariesForbidden creates a GetDetectSummariesForbidden with default headers values
func NewGetDetectSummariesForbidden() *GetDetectSummariesForbidden {
	return &GetDetectSummariesForbidden{}
}

/*
GetDetectSummariesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDetectSummariesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get detect summaries forbidden response has a 2xx status code
func (o *GetDetectSummariesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get detect summaries forbidden response has a 3xx status code
func (o *GetDetectSummariesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get detect summaries forbidden response has a 4xx status code
func (o *GetDetectSummariesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get detect summaries forbidden response has a 5xx status code
func (o *GetDetectSummariesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get detect summaries forbidden response a status code equal to that given
func (o *GetDetectSummariesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDetectSummariesForbidden) Error() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] getDetectSummariesForbidden  %+v", 403, o.Payload)
}

func (o *GetDetectSummariesForbidden) String() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] getDetectSummariesForbidden  %+v", 403, o.Payload)
}

func (o *GetDetectSummariesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetDetectSummariesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDetectSummariesTooManyRequests creates a GetDetectSummariesTooManyRequests with default headers values
func NewGetDetectSummariesTooManyRequests() *GetDetectSummariesTooManyRequests {
	return &GetDetectSummariesTooManyRequests{}
}

/*
GetDetectSummariesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDetectSummariesTooManyRequests struct {

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

// IsSuccess returns true when this get detect summaries too many requests response has a 2xx status code
func (o *GetDetectSummariesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get detect summaries too many requests response has a 3xx status code
func (o *GetDetectSummariesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get detect summaries too many requests response has a 4xx status code
func (o *GetDetectSummariesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get detect summaries too many requests response has a 5xx status code
func (o *GetDetectSummariesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get detect summaries too many requests response a status code equal to that given
func (o *GetDetectSummariesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDetectSummariesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] getDetectSummariesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDetectSummariesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] getDetectSummariesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDetectSummariesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetDetectSummariesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDetectSummariesInternalServerError creates a GetDetectSummariesInternalServerError with default headers values
func NewGetDetectSummariesInternalServerError() *GetDetectSummariesInternalServerError {
	return &GetDetectSummariesInternalServerError{}
}

/*
GetDetectSummariesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDetectSummariesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaDetectSummariesResponse
}

// IsSuccess returns true when this get detect summaries internal server error response has a 2xx status code
func (o *GetDetectSummariesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get detect summaries internal server error response has a 3xx status code
func (o *GetDetectSummariesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get detect summaries internal server error response has a 4xx status code
func (o *GetDetectSummariesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get detect summaries internal server error response has a 5xx status code
func (o *GetDetectSummariesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get detect summaries internal server error response a status code equal to that given
func (o *GetDetectSummariesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDetectSummariesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] getDetectSummariesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDetectSummariesInternalServerError) String() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] getDetectSummariesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDetectSummariesInternalServerError) GetPayload() *models.DomainMsaDetectSummariesResponse {
	return o.Payload
}

func (o *GetDetectSummariesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaDetectSummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDetectSummariesDefault creates a GetDetectSummariesDefault with default headers values
func NewGetDetectSummariesDefault(code int) *GetDetectSummariesDefault {
	return &GetDetectSummariesDefault{
		_statusCode: code,
	}
}

/*
GetDetectSummariesDefault describes a response with status code -1, with default header values.

OK
*/
type GetDetectSummariesDefault struct {
	_statusCode int

	Payload *models.DomainMsaDetectSummariesResponse
}

// Code gets the status code for the get detect summaries default response
func (o *GetDetectSummariesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get detect summaries default response has a 2xx status code
func (o *GetDetectSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get detect summaries default response has a 3xx status code
func (o *GetDetectSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get detect summaries default response has a 4xx status code
func (o *GetDetectSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get detect summaries default response has a 5xx status code
func (o *GetDetectSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get detect summaries default response a status code equal to that given
func (o *GetDetectSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetDetectSummariesDefault) Error() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] GetDetectSummaries default  %+v", o._statusCode, o.Payload)
}

func (o *GetDetectSummariesDefault) String() string {
	return fmt.Sprintf("[POST /detects/entities/summaries/GET/v1][%d] GetDetectSummaries default  %+v", o._statusCode, o.Payload)
}

func (o *GetDetectSummariesDefault) GetPayload() *models.DomainMsaDetectSummariesResponse {
	return o.Payload
}

func (o *GetDetectSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainMsaDetectSummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
