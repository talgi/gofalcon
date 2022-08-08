// Code generated by go-swagger; DO NOT EDIT.

package sensor_visibility_exclusions

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

// UpdateSensorVisibilityExclusionsV1Reader is a Reader for the UpdateSensorVisibilityExclusionsV1 structure.
type UpdateSensorVisibilityExclusionsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSensorVisibilityExclusionsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSensorVisibilityExclusionsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSensorVisibilityExclusionsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateSensorVisibilityExclusionsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateSensorVisibilityExclusionsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSensorVisibilityExclusionsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateSensorVisibilityExclusionsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSensorVisibilityExclusionsV1OK creates a UpdateSensorVisibilityExclusionsV1OK with default headers values
func NewUpdateSensorVisibilityExclusionsV1OK() *UpdateSensorVisibilityExclusionsV1OK {
	return &UpdateSensorVisibilityExclusionsV1OK{}
}

/*
	UpdateSensorVisibilityExclusionsV1OK describes a response with status code 200, with default header values.

OK
*/
type UpdateSensorVisibilityExclusionsV1OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSvExclusionRespV1
}

func (o *UpdateSensorVisibilityExclusionsV1OK) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sv-exclusions/v1][%d] updateSensorVisibilityExclusionsV1OK  %+v", 200, o.Payload)
}
func (o *UpdateSensorVisibilityExclusionsV1OK) GetPayload() *models.ResponsesSvExclusionRespV1 {
	return o.Payload
}

func (o *UpdateSensorVisibilityExclusionsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSvExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorVisibilityExclusionsV1BadRequest creates a UpdateSensorVisibilityExclusionsV1BadRequest with default headers values
func NewUpdateSensorVisibilityExclusionsV1BadRequest() *UpdateSensorVisibilityExclusionsV1BadRequest {
	return &UpdateSensorVisibilityExclusionsV1BadRequest{}
}

/*
	UpdateSensorVisibilityExclusionsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateSensorVisibilityExclusionsV1BadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSvExclusionRespV1
}

func (o *UpdateSensorVisibilityExclusionsV1BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sv-exclusions/v1][%d] updateSensorVisibilityExclusionsV1BadRequest  %+v", 400, o.Payload)
}
func (o *UpdateSensorVisibilityExclusionsV1BadRequest) GetPayload() *models.ResponsesSvExclusionRespV1 {
	return o.Payload
}

func (o *UpdateSensorVisibilityExclusionsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSvExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorVisibilityExclusionsV1Forbidden creates a UpdateSensorVisibilityExclusionsV1Forbidden with default headers values
func NewUpdateSensorVisibilityExclusionsV1Forbidden() *UpdateSensorVisibilityExclusionsV1Forbidden {
	return &UpdateSensorVisibilityExclusionsV1Forbidden{}
}

/*
	UpdateSensorVisibilityExclusionsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateSensorVisibilityExclusionsV1Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *UpdateSensorVisibilityExclusionsV1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sv-exclusions/v1][%d] updateSensorVisibilityExclusionsV1Forbidden  %+v", 403, o.Payload)
}
func (o *UpdateSensorVisibilityExclusionsV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateSensorVisibilityExclusionsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateSensorVisibilityExclusionsV1TooManyRequests creates a UpdateSensorVisibilityExclusionsV1TooManyRequests with default headers values
func NewUpdateSensorVisibilityExclusionsV1TooManyRequests() *UpdateSensorVisibilityExclusionsV1TooManyRequests {
	return &UpdateSensorVisibilityExclusionsV1TooManyRequests{}
}

/*
	UpdateSensorVisibilityExclusionsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateSensorVisibilityExclusionsV1TooManyRequests struct {

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

func (o *UpdateSensorVisibilityExclusionsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sv-exclusions/v1][%d] updateSensorVisibilityExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateSensorVisibilityExclusionsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateSensorVisibilityExclusionsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateSensorVisibilityExclusionsV1InternalServerError creates a UpdateSensorVisibilityExclusionsV1InternalServerError with default headers values
func NewUpdateSensorVisibilityExclusionsV1InternalServerError() *UpdateSensorVisibilityExclusionsV1InternalServerError {
	return &UpdateSensorVisibilityExclusionsV1InternalServerError{}
}

/*
	UpdateSensorVisibilityExclusionsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateSensorVisibilityExclusionsV1InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSvExclusionRespV1
}

func (o *UpdateSensorVisibilityExclusionsV1InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sv-exclusions/v1][%d] updateSensorVisibilityExclusionsV1InternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateSensorVisibilityExclusionsV1InternalServerError) GetPayload() *models.ResponsesSvExclusionRespV1 {
	return o.Payload
}

func (o *UpdateSensorVisibilityExclusionsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSvExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorVisibilityExclusionsV1Default creates a UpdateSensorVisibilityExclusionsV1Default with default headers values
func NewUpdateSensorVisibilityExclusionsV1Default(code int) *UpdateSensorVisibilityExclusionsV1Default {
	return &UpdateSensorVisibilityExclusionsV1Default{
		_statusCode: code,
	}
}

/*
	UpdateSensorVisibilityExclusionsV1Default describes a response with status code -1, with default header values.

OK
*/
type UpdateSensorVisibilityExclusionsV1Default struct {
	_statusCode int

	Payload *models.ResponsesSvExclusionRespV1
}

// Code gets the status code for the update sensor visibility exclusions v1 default response
func (o *UpdateSensorVisibilityExclusionsV1Default) Code() int {
	return o._statusCode
}

func (o *UpdateSensorVisibilityExclusionsV1Default) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sv-exclusions/v1][%d] updateSensorVisibilityExclusionsV1 default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateSensorVisibilityExclusionsV1Default) GetPayload() *models.ResponsesSvExclusionRespV1 {
	return o.Payload
}

func (o *UpdateSensorVisibilityExclusionsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesSvExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
