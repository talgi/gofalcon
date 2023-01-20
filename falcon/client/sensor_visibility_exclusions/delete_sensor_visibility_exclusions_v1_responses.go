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

// DeleteSensorVisibilityExclusionsV1Reader is a Reader for the DeleteSensorVisibilityExclusionsV1 structure.
type DeleteSensorVisibilityExclusionsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSensorVisibilityExclusionsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSensorVisibilityExclusionsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSensorVisibilityExclusionsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSensorVisibilityExclusionsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteSensorVisibilityExclusionsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSensorVisibilityExclusionsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSensorVisibilityExclusionsV1OK creates a DeleteSensorVisibilityExclusionsV1OK with default headers values
func NewDeleteSensorVisibilityExclusionsV1OK() *DeleteSensorVisibilityExclusionsV1OK {
	return &DeleteSensorVisibilityExclusionsV1OK{}
}

/*
DeleteSensorVisibilityExclusionsV1OK describes a response with status code 200, with default header values.

OK
*/
type DeleteSensorVisibilityExclusionsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this delete sensor visibility exclusions v1 o k response has a 2xx status code
func (o *DeleteSensorVisibilityExclusionsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete sensor visibility exclusions v1 o k response has a 3xx status code
func (o *DeleteSensorVisibilityExclusionsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sensor visibility exclusions v1 o k response has a 4xx status code
func (o *DeleteSensorVisibilityExclusionsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete sensor visibility exclusions v1 o k response has a 5xx status code
func (o *DeleteSensorVisibilityExclusionsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sensor visibility exclusions v1 o k response a status code equal to that given
func (o *DeleteSensorVisibilityExclusionsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete sensor visibility exclusions v1 o k response
func (o *DeleteSensorVisibilityExclusionsV1OK) Code() int {
	return 200
}

func (o *DeleteSensorVisibilityExclusionsV1OK) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sv-exclusions/v1][%d] deleteSensorVisibilityExclusionsV1OK  %+v", 200, o.Payload)
}

func (o *DeleteSensorVisibilityExclusionsV1OK) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/sv-exclusions/v1][%d] deleteSensorVisibilityExclusionsV1OK  %+v", 200, o.Payload)
}

func (o *DeleteSensorVisibilityExclusionsV1OK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSensorVisibilityExclusionsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSensorVisibilityExclusionsV1BadRequest creates a DeleteSensorVisibilityExclusionsV1BadRequest with default headers values
func NewDeleteSensorVisibilityExclusionsV1BadRequest() *DeleteSensorVisibilityExclusionsV1BadRequest {
	return &DeleteSensorVisibilityExclusionsV1BadRequest{}
}

/*
DeleteSensorVisibilityExclusionsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteSensorVisibilityExclusionsV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this delete sensor visibility exclusions v1 bad request response has a 2xx status code
func (o *DeleteSensorVisibilityExclusionsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sensor visibility exclusions v1 bad request response has a 3xx status code
func (o *DeleteSensorVisibilityExclusionsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sensor visibility exclusions v1 bad request response has a 4xx status code
func (o *DeleteSensorVisibilityExclusionsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete sensor visibility exclusions v1 bad request response has a 5xx status code
func (o *DeleteSensorVisibilityExclusionsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sensor visibility exclusions v1 bad request response a status code equal to that given
func (o *DeleteSensorVisibilityExclusionsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete sensor visibility exclusions v1 bad request response
func (o *DeleteSensorVisibilityExclusionsV1BadRequest) Code() int {
	return 400
}

func (o *DeleteSensorVisibilityExclusionsV1BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sv-exclusions/v1][%d] deleteSensorVisibilityExclusionsV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSensorVisibilityExclusionsV1BadRequest) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/sv-exclusions/v1][%d] deleteSensorVisibilityExclusionsV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSensorVisibilityExclusionsV1BadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSensorVisibilityExclusionsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSensorVisibilityExclusionsV1Forbidden creates a DeleteSensorVisibilityExclusionsV1Forbidden with default headers values
func NewDeleteSensorVisibilityExclusionsV1Forbidden() *DeleteSensorVisibilityExclusionsV1Forbidden {
	return &DeleteSensorVisibilityExclusionsV1Forbidden{}
}

/*
DeleteSensorVisibilityExclusionsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteSensorVisibilityExclusionsV1Forbidden struct {

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

// IsSuccess returns true when this delete sensor visibility exclusions v1 forbidden response has a 2xx status code
func (o *DeleteSensorVisibilityExclusionsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sensor visibility exclusions v1 forbidden response has a 3xx status code
func (o *DeleteSensorVisibilityExclusionsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sensor visibility exclusions v1 forbidden response has a 4xx status code
func (o *DeleteSensorVisibilityExclusionsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete sensor visibility exclusions v1 forbidden response has a 5xx status code
func (o *DeleteSensorVisibilityExclusionsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sensor visibility exclusions v1 forbidden response a status code equal to that given
func (o *DeleteSensorVisibilityExclusionsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete sensor visibility exclusions v1 forbidden response
func (o *DeleteSensorVisibilityExclusionsV1Forbidden) Code() int {
	return 403
}

func (o *DeleteSensorVisibilityExclusionsV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sv-exclusions/v1][%d] deleteSensorVisibilityExclusionsV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteSensorVisibilityExclusionsV1Forbidden) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/sv-exclusions/v1][%d] deleteSensorVisibilityExclusionsV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteSensorVisibilityExclusionsV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *DeleteSensorVisibilityExclusionsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSensorVisibilityExclusionsV1TooManyRequests creates a DeleteSensorVisibilityExclusionsV1TooManyRequests with default headers values
func NewDeleteSensorVisibilityExclusionsV1TooManyRequests() *DeleteSensorVisibilityExclusionsV1TooManyRequests {
	return &DeleteSensorVisibilityExclusionsV1TooManyRequests{}
}

/*
DeleteSensorVisibilityExclusionsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteSensorVisibilityExclusionsV1TooManyRequests struct {

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

// IsSuccess returns true when this delete sensor visibility exclusions v1 too many requests response has a 2xx status code
func (o *DeleteSensorVisibilityExclusionsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sensor visibility exclusions v1 too many requests response has a 3xx status code
func (o *DeleteSensorVisibilityExclusionsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sensor visibility exclusions v1 too many requests response has a 4xx status code
func (o *DeleteSensorVisibilityExclusionsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete sensor visibility exclusions v1 too many requests response has a 5xx status code
func (o *DeleteSensorVisibilityExclusionsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sensor visibility exclusions v1 too many requests response a status code equal to that given
func (o *DeleteSensorVisibilityExclusionsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete sensor visibility exclusions v1 too many requests response
func (o *DeleteSensorVisibilityExclusionsV1TooManyRequests) Code() int {
	return 429
}

func (o *DeleteSensorVisibilityExclusionsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sv-exclusions/v1][%d] deleteSensorVisibilityExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteSensorVisibilityExclusionsV1TooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/sv-exclusions/v1][%d] deleteSensorVisibilityExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteSensorVisibilityExclusionsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteSensorVisibilityExclusionsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSensorVisibilityExclusionsV1InternalServerError creates a DeleteSensorVisibilityExclusionsV1InternalServerError with default headers values
func NewDeleteSensorVisibilityExclusionsV1InternalServerError() *DeleteSensorVisibilityExclusionsV1InternalServerError {
	return &DeleteSensorVisibilityExclusionsV1InternalServerError{}
}

/*
DeleteSensorVisibilityExclusionsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteSensorVisibilityExclusionsV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this delete sensor visibility exclusions v1 internal server error response has a 2xx status code
func (o *DeleteSensorVisibilityExclusionsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sensor visibility exclusions v1 internal server error response has a 3xx status code
func (o *DeleteSensorVisibilityExclusionsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sensor visibility exclusions v1 internal server error response has a 4xx status code
func (o *DeleteSensorVisibilityExclusionsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete sensor visibility exclusions v1 internal server error response has a 5xx status code
func (o *DeleteSensorVisibilityExclusionsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete sensor visibility exclusions v1 internal server error response a status code equal to that given
func (o *DeleteSensorVisibilityExclusionsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete sensor visibility exclusions v1 internal server error response
func (o *DeleteSensorVisibilityExclusionsV1InternalServerError) Code() int {
	return 500
}

func (o *DeleteSensorVisibilityExclusionsV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sv-exclusions/v1][%d] deleteSensorVisibilityExclusionsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSensorVisibilityExclusionsV1InternalServerError) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/sv-exclusions/v1][%d] deleteSensorVisibilityExclusionsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSensorVisibilityExclusionsV1InternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSensorVisibilityExclusionsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
