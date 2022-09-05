// Code generated by go-swagger; DO NOT EDIT.

package prevention_policies

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

// PerformPreventionPoliciesActionReader is a Reader for the PerformPreventionPoliciesAction structure.
type PerformPreventionPoliciesActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformPreventionPoliciesActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformPreventionPoliciesActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPerformPreventionPoliciesActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPerformPreventionPoliciesActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPerformPreventionPoliciesActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPerformPreventionPoliciesActionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPerformPreventionPoliciesActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPerformPreventionPoliciesActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformPreventionPoliciesActionOK creates a PerformPreventionPoliciesActionOK with default headers values
func NewPerformPreventionPoliciesActionOK() *PerformPreventionPoliciesActionOK {
	return &PerformPreventionPoliciesActionOK{}
}

/*
PerformPreventionPoliciesActionOK describes a response with status code 200, with default header values.

OK
*/
type PerformPreventionPoliciesActionOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this perform prevention policies action o k response has a 2xx status code
func (o *PerformPreventionPoliciesActionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this perform prevention policies action o k response has a 3xx status code
func (o *PerformPreventionPoliciesActionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prevention policies action o k response has a 4xx status code
func (o *PerformPreventionPoliciesActionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform prevention policies action o k response has a 5xx status code
func (o *PerformPreventionPoliciesActionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this perform prevention policies action o k response a status code equal to that given
func (o *PerformPreventionPoliciesActionOK) IsCode(code int) bool {
	return code == 200
}

func (o *PerformPreventionPoliciesActionOK) Error() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionOK  %+v", 200, o.Payload)
}

func (o *PerformPreventionPoliciesActionOK) String() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionOK  %+v", 200, o.Payload)
}

func (o *PerformPreventionPoliciesActionOK) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *PerformPreventionPoliciesActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPreventionPoliciesActionBadRequest creates a PerformPreventionPoliciesActionBadRequest with default headers values
func NewPerformPreventionPoliciesActionBadRequest() *PerformPreventionPoliciesActionBadRequest {
	return &PerformPreventionPoliciesActionBadRequest{}
}

/*
PerformPreventionPoliciesActionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PerformPreventionPoliciesActionBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this perform prevention policies action bad request response has a 2xx status code
func (o *PerformPreventionPoliciesActionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform prevention policies action bad request response has a 3xx status code
func (o *PerformPreventionPoliciesActionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prevention policies action bad request response has a 4xx status code
func (o *PerformPreventionPoliciesActionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform prevention policies action bad request response has a 5xx status code
func (o *PerformPreventionPoliciesActionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this perform prevention policies action bad request response a status code equal to that given
func (o *PerformPreventionPoliciesActionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PerformPreventionPoliciesActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionBadRequest  %+v", 400, o.Payload)
}

func (o *PerformPreventionPoliciesActionBadRequest) String() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionBadRequest  %+v", 400, o.Payload)
}

func (o *PerformPreventionPoliciesActionBadRequest) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *PerformPreventionPoliciesActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPreventionPoliciesActionForbidden creates a PerformPreventionPoliciesActionForbidden with default headers values
func NewPerformPreventionPoliciesActionForbidden() *PerformPreventionPoliciesActionForbidden {
	return &PerformPreventionPoliciesActionForbidden{}
}

/*
PerformPreventionPoliciesActionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PerformPreventionPoliciesActionForbidden struct {

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

// IsSuccess returns true when this perform prevention policies action forbidden response has a 2xx status code
func (o *PerformPreventionPoliciesActionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform prevention policies action forbidden response has a 3xx status code
func (o *PerformPreventionPoliciesActionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prevention policies action forbidden response has a 4xx status code
func (o *PerformPreventionPoliciesActionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform prevention policies action forbidden response has a 5xx status code
func (o *PerformPreventionPoliciesActionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this perform prevention policies action forbidden response a status code equal to that given
func (o *PerformPreventionPoliciesActionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PerformPreventionPoliciesActionForbidden) Error() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionForbidden  %+v", 403, o.Payload)
}

func (o *PerformPreventionPoliciesActionForbidden) String() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionForbidden  %+v", 403, o.Payload)
}

func (o *PerformPreventionPoliciesActionForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *PerformPreventionPoliciesActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPerformPreventionPoliciesActionNotFound creates a PerformPreventionPoliciesActionNotFound with default headers values
func NewPerformPreventionPoliciesActionNotFound() *PerformPreventionPoliciesActionNotFound {
	return &PerformPreventionPoliciesActionNotFound{}
}

/*
PerformPreventionPoliciesActionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PerformPreventionPoliciesActionNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this perform prevention policies action not found response has a 2xx status code
func (o *PerformPreventionPoliciesActionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform prevention policies action not found response has a 3xx status code
func (o *PerformPreventionPoliciesActionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prevention policies action not found response has a 4xx status code
func (o *PerformPreventionPoliciesActionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform prevention policies action not found response has a 5xx status code
func (o *PerformPreventionPoliciesActionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this perform prevention policies action not found response a status code equal to that given
func (o *PerformPreventionPoliciesActionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PerformPreventionPoliciesActionNotFound) Error() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionNotFound  %+v", 404, o.Payload)
}

func (o *PerformPreventionPoliciesActionNotFound) String() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionNotFound  %+v", 404, o.Payload)
}

func (o *PerformPreventionPoliciesActionNotFound) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *PerformPreventionPoliciesActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPreventionPoliciesActionTooManyRequests creates a PerformPreventionPoliciesActionTooManyRequests with default headers values
func NewPerformPreventionPoliciesActionTooManyRequests() *PerformPreventionPoliciesActionTooManyRequests {
	return &PerformPreventionPoliciesActionTooManyRequests{}
}

/*
PerformPreventionPoliciesActionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PerformPreventionPoliciesActionTooManyRequests struct {

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

// IsSuccess returns true when this perform prevention policies action too many requests response has a 2xx status code
func (o *PerformPreventionPoliciesActionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform prevention policies action too many requests response has a 3xx status code
func (o *PerformPreventionPoliciesActionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prevention policies action too many requests response has a 4xx status code
func (o *PerformPreventionPoliciesActionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform prevention policies action too many requests response has a 5xx status code
func (o *PerformPreventionPoliciesActionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this perform prevention policies action too many requests response a status code equal to that given
func (o *PerformPreventionPoliciesActionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PerformPreventionPoliciesActionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PerformPreventionPoliciesActionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PerformPreventionPoliciesActionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PerformPreventionPoliciesActionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPerformPreventionPoliciesActionInternalServerError creates a PerformPreventionPoliciesActionInternalServerError with default headers values
func NewPerformPreventionPoliciesActionInternalServerError() *PerformPreventionPoliciesActionInternalServerError {
	return &PerformPreventionPoliciesActionInternalServerError{}
}

/*
PerformPreventionPoliciesActionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PerformPreventionPoliciesActionInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this perform prevention policies action internal server error response has a 2xx status code
func (o *PerformPreventionPoliciesActionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform prevention policies action internal server error response has a 3xx status code
func (o *PerformPreventionPoliciesActionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prevention policies action internal server error response has a 4xx status code
func (o *PerformPreventionPoliciesActionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform prevention policies action internal server error response has a 5xx status code
func (o *PerformPreventionPoliciesActionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this perform prevention policies action internal server error response a status code equal to that given
func (o *PerformPreventionPoliciesActionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PerformPreventionPoliciesActionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformPreventionPoliciesActionInternalServerError) String() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesActionInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformPreventionPoliciesActionInternalServerError) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *PerformPreventionPoliciesActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPreventionPoliciesActionDefault creates a PerformPreventionPoliciesActionDefault with default headers values
func NewPerformPreventionPoliciesActionDefault(code int) *PerformPreventionPoliciesActionDefault {
	return &PerformPreventionPoliciesActionDefault{
		_statusCode: code,
	}
}

/*
PerformPreventionPoliciesActionDefault describes a response with status code -1, with default header values.

OK
*/
type PerformPreventionPoliciesActionDefault struct {
	_statusCode int

	Payload *models.ResponsesPreventionPoliciesV1
}

// Code gets the status code for the perform prevention policies action default response
func (o *PerformPreventionPoliciesActionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this perform prevention policies action default response has a 2xx status code
func (o *PerformPreventionPoliciesActionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this perform prevention policies action default response has a 3xx status code
func (o *PerformPreventionPoliciesActionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this perform prevention policies action default response has a 4xx status code
func (o *PerformPreventionPoliciesActionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this perform prevention policies action default response has a 5xx status code
func (o *PerformPreventionPoliciesActionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this perform prevention policies action default response a status code equal to that given
func (o *PerformPreventionPoliciesActionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PerformPreventionPoliciesActionDefault) Error() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesAction default  %+v", o._statusCode, o.Payload)
}

func (o *PerformPreventionPoliciesActionDefault) String() string {
	return fmt.Sprintf("[POST /policy/entities/prevention-actions/v1][%d] performPreventionPoliciesAction default  %+v", o._statusCode, o.Payload)
}

func (o *PerformPreventionPoliciesActionDefault) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *PerformPreventionPoliciesActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
