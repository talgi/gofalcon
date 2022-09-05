// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// CreateSensorUpdatePoliciesV2Reader is a Reader for the CreateSensorUpdatePoliciesV2 structure.
type CreateSensorUpdatePoliciesV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSensorUpdatePoliciesV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSensorUpdatePoliciesV2Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSensorUpdatePoliciesV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateSensorUpdatePoliciesV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSensorUpdatePoliciesV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateSensorUpdatePoliciesV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateSensorUpdatePoliciesV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSensorUpdatePoliciesV2Created creates a CreateSensorUpdatePoliciesV2Created with default headers values
func NewCreateSensorUpdatePoliciesV2Created() *CreateSensorUpdatePoliciesV2Created {
	return &CreateSensorUpdatePoliciesV2Created{}
}

/*
CreateSensorUpdatePoliciesV2Created describes a response with status code 201, with default header values.

Created
*/
type CreateSensorUpdatePoliciesV2Created struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

// IsSuccess returns true when this create sensor update policies v2 created response has a 2xx status code
func (o *CreateSensorUpdatePoliciesV2Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create sensor update policies v2 created response has a 3xx status code
func (o *CreateSensorUpdatePoliciesV2Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies v2 created response has a 4xx status code
func (o *CreateSensorUpdatePoliciesV2Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create sensor update policies v2 created response has a 5xx status code
func (o *CreateSensorUpdatePoliciesV2Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create sensor update policies v2 created response a status code equal to that given
func (o *CreateSensorUpdatePoliciesV2Created) IsCode(code int) bool {
	return code == 201
}

func (o *CreateSensorUpdatePoliciesV2Created) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2Created  %+v", 201, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2Created) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2Created  %+v", 201, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2Created) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesV2Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSensorUpdatePoliciesV2BadRequest creates a CreateSensorUpdatePoliciesV2BadRequest with default headers values
func NewCreateSensorUpdatePoliciesV2BadRequest() *CreateSensorUpdatePoliciesV2BadRequest {
	return &CreateSensorUpdatePoliciesV2BadRequest{}
}

/*
CreateSensorUpdatePoliciesV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateSensorUpdatePoliciesV2BadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

// IsSuccess returns true when this create sensor update policies v2 bad request response has a 2xx status code
func (o *CreateSensorUpdatePoliciesV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sensor update policies v2 bad request response has a 3xx status code
func (o *CreateSensorUpdatePoliciesV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies v2 bad request response has a 4xx status code
func (o *CreateSensorUpdatePoliciesV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create sensor update policies v2 bad request response has a 5xx status code
func (o *CreateSensorUpdatePoliciesV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create sensor update policies v2 bad request response a status code equal to that given
func (o *CreateSensorUpdatePoliciesV2BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateSensorUpdatePoliciesV2BadRequest) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2BadRequest  %+v", 400, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2BadRequest) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2BadRequest  %+v", 400, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2BadRequest) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSensorUpdatePoliciesV2Forbidden creates a CreateSensorUpdatePoliciesV2Forbidden with default headers values
func NewCreateSensorUpdatePoliciesV2Forbidden() *CreateSensorUpdatePoliciesV2Forbidden {
	return &CreateSensorUpdatePoliciesV2Forbidden{}
}

/*
CreateSensorUpdatePoliciesV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateSensorUpdatePoliciesV2Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this create sensor update policies v2 forbidden response has a 2xx status code
func (o *CreateSensorUpdatePoliciesV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sensor update policies v2 forbidden response has a 3xx status code
func (o *CreateSensorUpdatePoliciesV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies v2 forbidden response has a 4xx status code
func (o *CreateSensorUpdatePoliciesV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create sensor update policies v2 forbidden response has a 5xx status code
func (o *CreateSensorUpdatePoliciesV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create sensor update policies v2 forbidden response a status code equal to that given
func (o *CreateSensorUpdatePoliciesV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateSensorUpdatePoliciesV2Forbidden) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2Forbidden  %+v", 403, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2Forbidden) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2Forbidden  %+v", 403, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSensorUpdatePoliciesV2NotFound creates a CreateSensorUpdatePoliciesV2NotFound with default headers values
func NewCreateSensorUpdatePoliciesV2NotFound() *CreateSensorUpdatePoliciesV2NotFound {
	return &CreateSensorUpdatePoliciesV2NotFound{}
}

/*
CreateSensorUpdatePoliciesV2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateSensorUpdatePoliciesV2NotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

// IsSuccess returns true when this create sensor update policies v2 not found response has a 2xx status code
func (o *CreateSensorUpdatePoliciesV2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sensor update policies v2 not found response has a 3xx status code
func (o *CreateSensorUpdatePoliciesV2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies v2 not found response has a 4xx status code
func (o *CreateSensorUpdatePoliciesV2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create sensor update policies v2 not found response has a 5xx status code
func (o *CreateSensorUpdatePoliciesV2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create sensor update policies v2 not found response a status code equal to that given
func (o *CreateSensorUpdatePoliciesV2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateSensorUpdatePoliciesV2NotFound) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2NotFound  %+v", 404, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2NotFound) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2NotFound  %+v", 404, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2NotFound) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSensorUpdatePoliciesV2TooManyRequests creates a CreateSensorUpdatePoliciesV2TooManyRequests with default headers values
func NewCreateSensorUpdatePoliciesV2TooManyRequests() *CreateSensorUpdatePoliciesV2TooManyRequests {
	return &CreateSensorUpdatePoliciesV2TooManyRequests{}
}

/*
CreateSensorUpdatePoliciesV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateSensorUpdatePoliciesV2TooManyRequests struct {

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

// IsSuccess returns true when this create sensor update policies v2 too many requests response has a 2xx status code
func (o *CreateSensorUpdatePoliciesV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sensor update policies v2 too many requests response has a 3xx status code
func (o *CreateSensorUpdatePoliciesV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies v2 too many requests response has a 4xx status code
func (o *CreateSensorUpdatePoliciesV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create sensor update policies v2 too many requests response has a 5xx status code
func (o *CreateSensorUpdatePoliciesV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create sensor update policies v2 too many requests response a status code equal to that given
func (o *CreateSensorUpdatePoliciesV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateSensorUpdatePoliciesV2TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2TooManyRequests) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSensorUpdatePoliciesV2InternalServerError creates a CreateSensorUpdatePoliciesV2InternalServerError with default headers values
func NewCreateSensorUpdatePoliciesV2InternalServerError() *CreateSensorUpdatePoliciesV2InternalServerError {
	return &CreateSensorUpdatePoliciesV2InternalServerError{}
}

/*
CreateSensorUpdatePoliciesV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateSensorUpdatePoliciesV2InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

// IsSuccess returns true when this create sensor update policies v2 internal server error response has a 2xx status code
func (o *CreateSensorUpdatePoliciesV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sensor update policies v2 internal server error response has a 3xx status code
func (o *CreateSensorUpdatePoliciesV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies v2 internal server error response has a 4xx status code
func (o *CreateSensorUpdatePoliciesV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create sensor update policies v2 internal server error response has a 5xx status code
func (o *CreateSensorUpdatePoliciesV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create sensor update policies v2 internal server error response a status code equal to that given
func (o *CreateSensorUpdatePoliciesV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateSensorUpdatePoliciesV2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2InternalServerError) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v2][%d] createSensorUpdatePoliciesV2InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateSensorUpdatePoliciesV2InternalServerError) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
