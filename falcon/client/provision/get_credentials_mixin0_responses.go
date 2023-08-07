// Code generated by go-swagger; DO NOT EDIT.

package provision

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

// GetCredentialsMixin0Reader is a Reader for the GetCredentialsMixin0 structure.
type GetCredentialsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCredentialsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCredentialsMixin0BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCredentialsMixin0Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCredentialsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCredentialsMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCredentialsMixin0InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /snapshots/entities/image-registry-credentials/v1] GetCredentialsMixin0", response, response.Code())
	}
}

// NewGetCredentialsMixin0OK creates a GetCredentialsMixin0OK with default headers values
func NewGetCredentialsMixin0OK() *GetCredentialsMixin0OK {
	return &GetCredentialsMixin0OK{}
}

/*
GetCredentialsMixin0OK describes a response with status code 200, with default header values.

OK
*/
type GetCredentialsMixin0OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsRegistryCredentialsResponse
}

// IsSuccess returns true when this get credentials mixin0 o k response has a 2xx status code
func (o *GetCredentialsMixin0OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get credentials mixin0 o k response has a 3xx status code
func (o *GetCredentialsMixin0OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials mixin0 o k response has a 4xx status code
func (o *GetCredentialsMixin0OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credentials mixin0 o k response has a 5xx status code
func (o *GetCredentialsMixin0OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials mixin0 o k response a status code equal to that given
func (o *GetCredentialsMixin0OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get credentials mixin0 o k response
func (o *GetCredentialsMixin0OK) Code() int {
	return 200
}

func (o *GetCredentialsMixin0OK) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0OK  %+v", 200, o.Payload)
}

func (o *GetCredentialsMixin0OK) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0OK  %+v", 200, o.Payload)
}

func (o *GetCredentialsMixin0OK) GetPayload() *models.ModelsRegistryCredentialsResponse {
	return o.Payload
}

func (o *GetCredentialsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsRegistryCredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsMixin0BadRequest creates a GetCredentialsMixin0BadRequest with default headers values
func NewGetCredentialsMixin0BadRequest() *GetCredentialsMixin0BadRequest {
	return &GetCredentialsMixin0BadRequest{}
}

/*
GetCredentialsMixin0BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCredentialsMixin0BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAPIError
}

// IsSuccess returns true when this get credentials mixin0 bad request response has a 2xx status code
func (o *GetCredentialsMixin0BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials mixin0 bad request response has a 3xx status code
func (o *GetCredentialsMixin0BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials mixin0 bad request response has a 4xx status code
func (o *GetCredentialsMixin0BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credentials mixin0 bad request response has a 5xx status code
func (o *GetCredentialsMixin0BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials mixin0 bad request response a status code equal to that given
func (o *GetCredentialsMixin0BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get credentials mixin0 bad request response
func (o *GetCredentialsMixin0BadRequest) Code() int {
	return 400
}

func (o *GetCredentialsMixin0BadRequest) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0BadRequest  %+v", 400, o.Payload)
}

func (o *GetCredentialsMixin0BadRequest) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0BadRequest  %+v", 400, o.Payload)
}

func (o *GetCredentialsMixin0BadRequest) GetPayload() *models.MsaAPIError {
	return o.Payload
}

func (o *GetCredentialsMixin0BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsMixin0Unauthorized creates a GetCredentialsMixin0Unauthorized with default headers values
func NewGetCredentialsMixin0Unauthorized() *GetCredentialsMixin0Unauthorized {
	return &GetCredentialsMixin0Unauthorized{}
}

/*
GetCredentialsMixin0Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCredentialsMixin0Unauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAPIError
}

// IsSuccess returns true when this get credentials mixin0 unauthorized response has a 2xx status code
func (o *GetCredentialsMixin0Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials mixin0 unauthorized response has a 3xx status code
func (o *GetCredentialsMixin0Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials mixin0 unauthorized response has a 4xx status code
func (o *GetCredentialsMixin0Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credentials mixin0 unauthorized response has a 5xx status code
func (o *GetCredentialsMixin0Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials mixin0 unauthorized response a status code equal to that given
func (o *GetCredentialsMixin0Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get credentials mixin0 unauthorized response
func (o *GetCredentialsMixin0Unauthorized) Code() int {
	return 401
}

func (o *GetCredentialsMixin0Unauthorized) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0Unauthorized  %+v", 401, o.Payload)
}

func (o *GetCredentialsMixin0Unauthorized) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0Unauthorized  %+v", 401, o.Payload)
}

func (o *GetCredentialsMixin0Unauthorized) GetPayload() *models.MsaAPIError {
	return o.Payload
}

func (o *GetCredentialsMixin0Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsMixin0Forbidden creates a GetCredentialsMixin0Forbidden with default headers values
func NewGetCredentialsMixin0Forbidden() *GetCredentialsMixin0Forbidden {
	return &GetCredentialsMixin0Forbidden{}
}

/*
GetCredentialsMixin0Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCredentialsMixin0Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAPIError
}

// IsSuccess returns true when this get credentials mixin0 forbidden response has a 2xx status code
func (o *GetCredentialsMixin0Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials mixin0 forbidden response has a 3xx status code
func (o *GetCredentialsMixin0Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials mixin0 forbidden response has a 4xx status code
func (o *GetCredentialsMixin0Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credentials mixin0 forbidden response has a 5xx status code
func (o *GetCredentialsMixin0Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials mixin0 forbidden response a status code equal to that given
func (o *GetCredentialsMixin0Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get credentials mixin0 forbidden response
func (o *GetCredentialsMixin0Forbidden) Code() int {
	return 403
}

func (o *GetCredentialsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *GetCredentialsMixin0Forbidden) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *GetCredentialsMixin0Forbidden) GetPayload() *models.MsaAPIError {
	return o.Payload
}

func (o *GetCredentialsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsMixin0TooManyRequests creates a GetCredentialsMixin0TooManyRequests with default headers values
func NewGetCredentialsMixin0TooManyRequests() *GetCredentialsMixin0TooManyRequests {
	return &GetCredentialsMixin0TooManyRequests{}
}

/*
GetCredentialsMixin0TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCredentialsMixin0TooManyRequests struct {

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

// IsSuccess returns true when this get credentials mixin0 too many requests response has a 2xx status code
func (o *GetCredentialsMixin0TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials mixin0 too many requests response has a 3xx status code
func (o *GetCredentialsMixin0TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials mixin0 too many requests response has a 4xx status code
func (o *GetCredentialsMixin0TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credentials mixin0 too many requests response has a 5xx status code
func (o *GetCredentialsMixin0TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials mixin0 too many requests response a status code equal to that given
func (o *GetCredentialsMixin0TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get credentials mixin0 too many requests response
func (o *GetCredentialsMixin0TooManyRequests) Code() int {
	return 429
}

func (o *GetCredentialsMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCredentialsMixin0TooManyRequests) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCredentialsMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCredentialsMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCredentialsMixin0InternalServerError creates a GetCredentialsMixin0InternalServerError with default headers values
func NewGetCredentialsMixin0InternalServerError() *GetCredentialsMixin0InternalServerError {
	return &GetCredentialsMixin0InternalServerError{}
}

/*
GetCredentialsMixin0InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCredentialsMixin0InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsRegistryCredentialsResponse
}

// IsSuccess returns true when this get credentials mixin0 internal server error response has a 2xx status code
func (o *GetCredentialsMixin0InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials mixin0 internal server error response has a 3xx status code
func (o *GetCredentialsMixin0InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials mixin0 internal server error response has a 4xx status code
func (o *GetCredentialsMixin0InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credentials mixin0 internal server error response has a 5xx status code
func (o *GetCredentialsMixin0InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get credentials mixin0 internal server error response a status code equal to that given
func (o *GetCredentialsMixin0InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get credentials mixin0 internal server error response
func (o *GetCredentialsMixin0InternalServerError) Code() int {
	return 500
}

func (o *GetCredentialsMixin0InternalServerError) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0InternalServerError  %+v", 500, o.Payload)
}

func (o *GetCredentialsMixin0InternalServerError) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/image-registry-credentials/v1][%d] getCredentialsMixin0InternalServerError  %+v", 500, o.Payload)
}

func (o *GetCredentialsMixin0InternalServerError) GetPayload() *models.ModelsRegistryCredentialsResponse {
	return o.Payload
}

func (o *GetCredentialsMixin0InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsRegistryCredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
