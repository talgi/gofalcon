// Code generated by go-swagger; DO NOT EDIT.

package spotlight_vulnerabilities

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

// GetRemediationsV2Reader is a Reader for the GetRemediationsV2 structure.
type GetRemediationsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRemediationsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRemediationsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRemediationsV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRemediationsV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRemediationsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRemediationsV2OK creates a GetRemediationsV2OK with default headers values
func NewGetRemediationsV2OK() *GetRemediationsV2OK {
	return &GetRemediationsV2OK{}
}

/*
GetRemediationsV2OK describes a response with status code 200, with default header values.

OK
*/
type GetRemediationsV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPIRemediationEntitiesResponseV2
}

// IsSuccess returns true when this get remediations v2 o k response has a 2xx status code
func (o *GetRemediationsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get remediations v2 o k response has a 3xx status code
func (o *GetRemediationsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remediations v2 o k response has a 4xx status code
func (o *GetRemediationsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get remediations v2 o k response has a 5xx status code
func (o *GetRemediationsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get remediations v2 o k response a status code equal to that given
func (o *GetRemediationsV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRemediationsV2OK) Error() string {
	return fmt.Sprintf("[GET /spotlight/entities/remediations/v2][%d] getRemediationsV2OK  %+v", 200, o.Payload)
}

func (o *GetRemediationsV2OK) String() string {
	return fmt.Sprintf("[GET /spotlight/entities/remediations/v2][%d] getRemediationsV2OK  %+v", 200, o.Payload)
}

func (o *GetRemediationsV2OK) GetPayload() *models.DomainSPAPIRemediationEntitiesResponseV2 {
	return o.Payload
}

func (o *GetRemediationsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPIRemediationEntitiesResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemediationsV2Forbidden creates a GetRemediationsV2Forbidden with default headers values
func NewGetRemediationsV2Forbidden() *GetRemediationsV2Forbidden {
	return &GetRemediationsV2Forbidden{}
}

/*
GetRemediationsV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRemediationsV2Forbidden struct {

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

// IsSuccess returns true when this get remediations v2 forbidden response has a 2xx status code
func (o *GetRemediationsV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remediations v2 forbidden response has a 3xx status code
func (o *GetRemediationsV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remediations v2 forbidden response has a 4xx status code
func (o *GetRemediationsV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remediations v2 forbidden response has a 5xx status code
func (o *GetRemediationsV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get remediations v2 forbidden response a status code equal to that given
func (o *GetRemediationsV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRemediationsV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /spotlight/entities/remediations/v2][%d] getRemediationsV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetRemediationsV2Forbidden) String() string {
	return fmt.Sprintf("[GET /spotlight/entities/remediations/v2][%d] getRemediationsV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetRemediationsV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRemediationsV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRemediationsV2TooManyRequests creates a GetRemediationsV2TooManyRequests with default headers values
func NewGetRemediationsV2TooManyRequests() *GetRemediationsV2TooManyRequests {
	return &GetRemediationsV2TooManyRequests{}
}

/*
GetRemediationsV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRemediationsV2TooManyRequests struct {

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

// IsSuccess returns true when this get remediations v2 too many requests response has a 2xx status code
func (o *GetRemediationsV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remediations v2 too many requests response has a 3xx status code
func (o *GetRemediationsV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remediations v2 too many requests response has a 4xx status code
func (o *GetRemediationsV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remediations v2 too many requests response has a 5xx status code
func (o *GetRemediationsV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get remediations v2 too many requests response a status code equal to that given
func (o *GetRemediationsV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRemediationsV2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /spotlight/entities/remediations/v2][%d] getRemediationsV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRemediationsV2TooManyRequests) String() string {
	return fmt.Sprintf("[GET /spotlight/entities/remediations/v2][%d] getRemediationsV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRemediationsV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRemediationsV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRemediationsV2Default creates a GetRemediationsV2Default with default headers values
func NewGetRemediationsV2Default(code int) *GetRemediationsV2Default {
	return &GetRemediationsV2Default{
		_statusCode: code,
	}
}

/*
GetRemediationsV2Default describes a response with status code -1, with default header values.

OK
*/
type GetRemediationsV2Default struct {
	_statusCode int

	Payload *models.DomainSPAPIRemediationEntitiesResponseV2
}

// Code gets the status code for the get remediations v2 default response
func (o *GetRemediationsV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get remediations v2 default response has a 2xx status code
func (o *GetRemediationsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get remediations v2 default response has a 3xx status code
func (o *GetRemediationsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get remediations v2 default response has a 4xx status code
func (o *GetRemediationsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get remediations v2 default response has a 5xx status code
func (o *GetRemediationsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get remediations v2 default response a status code equal to that given
func (o *GetRemediationsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetRemediationsV2Default) Error() string {
	return fmt.Sprintf("[GET /spotlight/entities/remediations/v2][%d] getRemediationsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *GetRemediationsV2Default) String() string {
	return fmt.Sprintf("[GET /spotlight/entities/remediations/v2][%d] getRemediationsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *GetRemediationsV2Default) GetPayload() *models.DomainSPAPIRemediationEntitiesResponseV2 {
	return o.Payload
}

func (o *GetRemediationsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainSPAPIRemediationEntitiesResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
