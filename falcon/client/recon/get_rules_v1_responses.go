// Code generated by go-swagger; DO NOT EDIT.

package recon

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

// GetRulesV1Reader is a Reader for the GetRulesV1 structure.
type GetRulesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRulesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRulesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRulesV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRulesV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRulesV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRulesV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRulesV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRulesV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRulesV1OK creates a GetRulesV1OK with default headers values
func NewGetRulesV1OK() *GetRulesV1OK {
	return &GetRulesV1OK{}
}

/*
GetRulesV1OK describes a response with status code 200, with default header values.

OK
*/
type GetRulesV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainRulesEntitiesResponseV1
}

// IsSuccess returns true when this get rules v1 o k response has a 2xx status code
func (o *GetRulesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get rules v1 o k response has a 3xx status code
func (o *GetRulesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules v1 o k response has a 4xx status code
func (o *GetRulesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rules v1 o k response has a 5xx status code
func (o *GetRulesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules v1 o k response a status code equal to that given
func (o *GetRulesV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get rules v1 o k response
func (o *GetRulesV1OK) Code() int {
	return 200
}

func (o *GetRulesV1OK) Error() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1OK  %+v", 200, o.Payload)
}

func (o *GetRulesV1OK) String() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1OK  %+v", 200, o.Payload)
}

func (o *GetRulesV1OK) GetPayload() *models.DomainRulesEntitiesResponseV1 {
	return o.Payload
}

func (o *GetRulesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainRulesEntitiesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRulesV1BadRequest creates a GetRulesV1BadRequest with default headers values
func NewGetRulesV1BadRequest() *GetRulesV1BadRequest {
	return &GetRulesV1BadRequest{}
}

/*
GetRulesV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRulesV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this get rules v1 bad request response has a 2xx status code
func (o *GetRulesV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules v1 bad request response has a 3xx status code
func (o *GetRulesV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules v1 bad request response has a 4xx status code
func (o *GetRulesV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rules v1 bad request response has a 5xx status code
func (o *GetRulesV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules v1 bad request response a status code equal to that given
func (o *GetRulesV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get rules v1 bad request response
func (o *GetRulesV1BadRequest) Code() int {
	return 400
}

func (o *GetRulesV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetRulesV1BadRequest) String() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetRulesV1BadRequest) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetRulesV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRulesV1Unauthorized creates a GetRulesV1Unauthorized with default headers values
func NewGetRulesV1Unauthorized() *GetRulesV1Unauthorized {
	return &GetRulesV1Unauthorized{}
}

/*
GetRulesV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRulesV1Unauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this get rules v1 unauthorized response has a 2xx status code
func (o *GetRulesV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules v1 unauthorized response has a 3xx status code
func (o *GetRulesV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules v1 unauthorized response has a 4xx status code
func (o *GetRulesV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rules v1 unauthorized response has a 5xx status code
func (o *GetRulesV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules v1 unauthorized response a status code equal to that given
func (o *GetRulesV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get rules v1 unauthorized response
func (o *GetRulesV1Unauthorized) Code() int {
	return 401
}

func (o *GetRulesV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetRulesV1Unauthorized) String() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetRulesV1Unauthorized) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetRulesV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRulesV1Forbidden creates a GetRulesV1Forbidden with default headers values
func NewGetRulesV1Forbidden() *GetRulesV1Forbidden {
	return &GetRulesV1Forbidden{}
}

/*
GetRulesV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRulesV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this get rules v1 forbidden response has a 2xx status code
func (o *GetRulesV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules v1 forbidden response has a 3xx status code
func (o *GetRulesV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules v1 forbidden response has a 4xx status code
func (o *GetRulesV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rules v1 forbidden response has a 5xx status code
func (o *GetRulesV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules v1 forbidden response a status code equal to that given
func (o *GetRulesV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get rules v1 forbidden response
func (o *GetRulesV1Forbidden) Code() int {
	return 403
}

func (o *GetRulesV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetRulesV1Forbidden) String() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetRulesV1Forbidden) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetRulesV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRulesV1TooManyRequests creates a GetRulesV1TooManyRequests with default headers values
func NewGetRulesV1TooManyRequests() *GetRulesV1TooManyRequests {
	return &GetRulesV1TooManyRequests{}
}

/*
GetRulesV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRulesV1TooManyRequests struct {

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

// IsSuccess returns true when this get rules v1 too many requests response has a 2xx status code
func (o *GetRulesV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules v1 too many requests response has a 3xx status code
func (o *GetRulesV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules v1 too many requests response has a 4xx status code
func (o *GetRulesV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rules v1 too many requests response has a 5xx status code
func (o *GetRulesV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules v1 too many requests response a status code equal to that given
func (o *GetRulesV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get rules v1 too many requests response
func (o *GetRulesV1TooManyRequests) Code() int {
	return 429
}

func (o *GetRulesV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRulesV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRulesV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRulesV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRulesV1InternalServerError creates a GetRulesV1InternalServerError with default headers values
func NewGetRulesV1InternalServerError() *GetRulesV1InternalServerError {
	return &GetRulesV1InternalServerError{}
}

/*
GetRulesV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRulesV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this get rules v1 internal server error response has a 2xx status code
func (o *GetRulesV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules v1 internal server error response has a 3xx status code
func (o *GetRulesV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules v1 internal server error response has a 4xx status code
func (o *GetRulesV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rules v1 internal server error response has a 5xx status code
func (o *GetRulesV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get rules v1 internal server error response a status code equal to that given
func (o *GetRulesV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get rules v1 internal server error response
func (o *GetRulesV1InternalServerError) Code() int {
	return 500
}

func (o *GetRulesV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetRulesV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] getRulesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetRulesV1InternalServerError) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetRulesV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRulesV1Default creates a GetRulesV1Default with default headers values
func NewGetRulesV1Default(code int) *GetRulesV1Default {
	return &GetRulesV1Default{
		_statusCode: code,
	}
}

/*
GetRulesV1Default describes a response with status code -1, with default header values.

OK
*/
type GetRulesV1Default struct {
	_statusCode int

	Payload *models.DomainRulesEntitiesResponseV1
}

// IsSuccess returns true when this get rules v1 default response has a 2xx status code
func (o *GetRulesV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get rules v1 default response has a 3xx status code
func (o *GetRulesV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get rules v1 default response has a 4xx status code
func (o *GetRulesV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get rules v1 default response has a 5xx status code
func (o *GetRulesV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get rules v1 default response a status code equal to that given
func (o *GetRulesV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get rules v1 default response
func (o *GetRulesV1Default) Code() int {
	return o._statusCode
}

func (o *GetRulesV1Default) Error() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] GetRulesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetRulesV1Default) String() string {
	return fmt.Sprintf("[GET /recon/entities/rules/v1][%d] GetRulesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetRulesV1Default) GetPayload() *models.DomainRulesEntitiesResponseV1 {
	return o.Payload
}

func (o *GetRulesV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainRulesEntitiesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
