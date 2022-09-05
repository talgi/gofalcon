// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

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

// PerformFirewallPoliciesActionReader is a Reader for the PerformFirewallPoliciesAction structure.
type PerformFirewallPoliciesActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformFirewallPoliciesActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformFirewallPoliciesActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPerformFirewallPoliciesActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPerformFirewallPoliciesActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPerformFirewallPoliciesActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPerformFirewallPoliciesActionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPerformFirewallPoliciesActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPerformFirewallPoliciesActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformFirewallPoliciesActionOK creates a PerformFirewallPoliciesActionOK with default headers values
func NewPerformFirewallPoliciesActionOK() *PerformFirewallPoliciesActionOK {
	return &PerformFirewallPoliciesActionOK{}
}

/*
PerformFirewallPoliciesActionOK describes a response with status code 200, with default header values.

OK
*/
type PerformFirewallPoliciesActionOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

// IsSuccess returns true when this perform firewall policies action o k response has a 2xx status code
func (o *PerformFirewallPoliciesActionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this perform firewall policies action o k response has a 3xx status code
func (o *PerformFirewallPoliciesActionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform firewall policies action o k response has a 4xx status code
func (o *PerformFirewallPoliciesActionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform firewall policies action o k response has a 5xx status code
func (o *PerformFirewallPoliciesActionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this perform firewall policies action o k response a status code equal to that given
func (o *PerformFirewallPoliciesActionOK) IsCode(code int) bool {
	return code == 200
}

func (o *PerformFirewallPoliciesActionOK) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionOK  %+v", 200, o.Payload)
}

func (o *PerformFirewallPoliciesActionOK) String() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionOK  %+v", 200, o.Payload)
}

func (o *PerformFirewallPoliciesActionOK) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformFirewallPoliciesActionBadRequest creates a PerformFirewallPoliciesActionBadRequest with default headers values
func NewPerformFirewallPoliciesActionBadRequest() *PerformFirewallPoliciesActionBadRequest {
	return &PerformFirewallPoliciesActionBadRequest{}
}

/*
PerformFirewallPoliciesActionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PerformFirewallPoliciesActionBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

// IsSuccess returns true when this perform firewall policies action bad request response has a 2xx status code
func (o *PerformFirewallPoliciesActionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform firewall policies action bad request response has a 3xx status code
func (o *PerformFirewallPoliciesActionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform firewall policies action bad request response has a 4xx status code
func (o *PerformFirewallPoliciesActionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform firewall policies action bad request response has a 5xx status code
func (o *PerformFirewallPoliciesActionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this perform firewall policies action bad request response a status code equal to that given
func (o *PerformFirewallPoliciesActionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PerformFirewallPoliciesActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionBadRequest  %+v", 400, o.Payload)
}

func (o *PerformFirewallPoliciesActionBadRequest) String() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionBadRequest  %+v", 400, o.Payload)
}

func (o *PerformFirewallPoliciesActionBadRequest) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformFirewallPoliciesActionForbidden creates a PerformFirewallPoliciesActionForbidden with default headers values
func NewPerformFirewallPoliciesActionForbidden() *PerformFirewallPoliciesActionForbidden {
	return &PerformFirewallPoliciesActionForbidden{}
}

/*
PerformFirewallPoliciesActionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PerformFirewallPoliciesActionForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this perform firewall policies action forbidden response has a 2xx status code
func (o *PerformFirewallPoliciesActionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform firewall policies action forbidden response has a 3xx status code
func (o *PerformFirewallPoliciesActionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform firewall policies action forbidden response has a 4xx status code
func (o *PerformFirewallPoliciesActionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform firewall policies action forbidden response has a 5xx status code
func (o *PerformFirewallPoliciesActionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this perform firewall policies action forbidden response a status code equal to that given
func (o *PerformFirewallPoliciesActionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PerformFirewallPoliciesActionForbidden) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionForbidden  %+v", 403, o.Payload)
}

func (o *PerformFirewallPoliciesActionForbidden) String() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionForbidden  %+v", 403, o.Payload)
}

func (o *PerformFirewallPoliciesActionForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPerformFirewallPoliciesActionNotFound creates a PerformFirewallPoliciesActionNotFound with default headers values
func NewPerformFirewallPoliciesActionNotFound() *PerformFirewallPoliciesActionNotFound {
	return &PerformFirewallPoliciesActionNotFound{}
}

/*
PerformFirewallPoliciesActionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PerformFirewallPoliciesActionNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

// IsSuccess returns true when this perform firewall policies action not found response has a 2xx status code
func (o *PerformFirewallPoliciesActionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform firewall policies action not found response has a 3xx status code
func (o *PerformFirewallPoliciesActionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform firewall policies action not found response has a 4xx status code
func (o *PerformFirewallPoliciesActionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform firewall policies action not found response has a 5xx status code
func (o *PerformFirewallPoliciesActionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this perform firewall policies action not found response a status code equal to that given
func (o *PerformFirewallPoliciesActionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PerformFirewallPoliciesActionNotFound) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionNotFound  %+v", 404, o.Payload)
}

func (o *PerformFirewallPoliciesActionNotFound) String() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionNotFound  %+v", 404, o.Payload)
}

func (o *PerformFirewallPoliciesActionNotFound) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformFirewallPoliciesActionTooManyRequests creates a PerformFirewallPoliciesActionTooManyRequests with default headers values
func NewPerformFirewallPoliciesActionTooManyRequests() *PerformFirewallPoliciesActionTooManyRequests {
	return &PerformFirewallPoliciesActionTooManyRequests{}
}

/*
PerformFirewallPoliciesActionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PerformFirewallPoliciesActionTooManyRequests struct {

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

// IsSuccess returns true when this perform firewall policies action too many requests response has a 2xx status code
func (o *PerformFirewallPoliciesActionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform firewall policies action too many requests response has a 3xx status code
func (o *PerformFirewallPoliciesActionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform firewall policies action too many requests response has a 4xx status code
func (o *PerformFirewallPoliciesActionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform firewall policies action too many requests response has a 5xx status code
func (o *PerformFirewallPoliciesActionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this perform firewall policies action too many requests response a status code equal to that given
func (o *PerformFirewallPoliciesActionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PerformFirewallPoliciesActionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PerformFirewallPoliciesActionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PerformFirewallPoliciesActionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPerformFirewallPoliciesActionInternalServerError creates a PerformFirewallPoliciesActionInternalServerError with default headers values
func NewPerformFirewallPoliciesActionInternalServerError() *PerformFirewallPoliciesActionInternalServerError {
	return &PerformFirewallPoliciesActionInternalServerError{}
}

/*
PerformFirewallPoliciesActionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PerformFirewallPoliciesActionInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

// IsSuccess returns true when this perform firewall policies action internal server error response has a 2xx status code
func (o *PerformFirewallPoliciesActionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform firewall policies action internal server error response has a 3xx status code
func (o *PerformFirewallPoliciesActionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform firewall policies action internal server error response has a 4xx status code
func (o *PerformFirewallPoliciesActionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform firewall policies action internal server error response has a 5xx status code
func (o *PerformFirewallPoliciesActionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this perform firewall policies action internal server error response a status code equal to that given
func (o *PerformFirewallPoliciesActionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PerformFirewallPoliciesActionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformFirewallPoliciesActionInternalServerError) String() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformFirewallPoliciesActionInternalServerError) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformFirewallPoliciesActionDefault creates a PerformFirewallPoliciesActionDefault with default headers values
func NewPerformFirewallPoliciesActionDefault(code int) *PerformFirewallPoliciesActionDefault {
	return &PerformFirewallPoliciesActionDefault{
		_statusCode: code,
	}
}

/*
PerformFirewallPoliciesActionDefault describes a response with status code -1, with default header values.

OK
*/
type PerformFirewallPoliciesActionDefault struct {
	_statusCode int

	Payload *models.ResponsesFirewallPoliciesV1
}

// Code gets the status code for the perform firewall policies action default response
func (o *PerformFirewallPoliciesActionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this perform firewall policies action default response has a 2xx status code
func (o *PerformFirewallPoliciesActionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this perform firewall policies action default response has a 3xx status code
func (o *PerformFirewallPoliciesActionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this perform firewall policies action default response has a 4xx status code
func (o *PerformFirewallPoliciesActionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this perform firewall policies action default response has a 5xx status code
func (o *PerformFirewallPoliciesActionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this perform firewall policies action default response a status code equal to that given
func (o *PerformFirewallPoliciesActionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PerformFirewallPoliciesActionDefault) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesAction default  %+v", o._statusCode, o.Payload)
}

func (o *PerformFirewallPoliciesActionDefault) String() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesAction default  %+v", o._statusCode, o.Payload)
}

func (o *PerformFirewallPoliciesActionDefault) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
