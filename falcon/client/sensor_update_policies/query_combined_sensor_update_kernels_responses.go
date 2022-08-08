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

// QueryCombinedSensorUpdateKernelsReader is a Reader for the QueryCombinedSensorUpdateKernels structure.
type QueryCombinedSensorUpdateKernelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedSensorUpdateKernelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedSensorUpdateKernelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedSensorUpdateKernelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedSensorUpdateKernelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedSensorUpdateKernelsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedSensorUpdateKernelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryCombinedSensorUpdateKernelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryCombinedSensorUpdateKernelsOK creates a QueryCombinedSensorUpdateKernelsOK with default headers values
func NewQueryCombinedSensorUpdateKernelsOK() *QueryCombinedSensorUpdateKernelsOK {
	return &QueryCombinedSensorUpdateKernelsOK{}
}

/*
	QueryCombinedSensorUpdateKernelsOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedSensorUpdateKernelsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdateKernelsV1
}

func (o *QueryCombinedSensorUpdateKernelsOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-kernels/v1][%d] queryCombinedSensorUpdateKernelsOK  %+v", 200, o.Payload)
}
func (o *QueryCombinedSensorUpdateKernelsOK) GetPayload() *models.ResponsesSensorUpdateKernelsV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdateKernelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdateKernelsV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedSensorUpdateKernelsBadRequest creates a QueryCombinedSensorUpdateKernelsBadRequest with default headers values
func NewQueryCombinedSensorUpdateKernelsBadRequest() *QueryCombinedSensorUpdateKernelsBadRequest {
	return &QueryCombinedSensorUpdateKernelsBadRequest{}
}

/*
	QueryCombinedSensorUpdateKernelsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedSensorUpdateKernelsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdateKernelsV1
}

func (o *QueryCombinedSensorUpdateKernelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-kernels/v1][%d] queryCombinedSensorUpdateKernelsBadRequest  %+v", 400, o.Payload)
}
func (o *QueryCombinedSensorUpdateKernelsBadRequest) GetPayload() *models.ResponsesSensorUpdateKernelsV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdateKernelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdateKernelsV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedSensorUpdateKernelsForbidden creates a QueryCombinedSensorUpdateKernelsForbidden with default headers values
func NewQueryCombinedSensorUpdateKernelsForbidden() *QueryCombinedSensorUpdateKernelsForbidden {
	return &QueryCombinedSensorUpdateKernelsForbidden{}
}

/*
	QueryCombinedSensorUpdateKernelsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedSensorUpdateKernelsForbidden struct {

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

func (o *QueryCombinedSensorUpdateKernelsForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-kernels/v1][%d] queryCombinedSensorUpdateKernelsForbidden  %+v", 403, o.Payload)
}
func (o *QueryCombinedSensorUpdateKernelsForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedSensorUpdateKernelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedSensorUpdateKernelsTooManyRequests creates a QueryCombinedSensorUpdateKernelsTooManyRequests with default headers values
func NewQueryCombinedSensorUpdateKernelsTooManyRequests() *QueryCombinedSensorUpdateKernelsTooManyRequests {
	return &QueryCombinedSensorUpdateKernelsTooManyRequests{}
}

/*
	QueryCombinedSensorUpdateKernelsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedSensorUpdateKernelsTooManyRequests struct {

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

func (o *QueryCombinedSensorUpdateKernelsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-kernels/v1][%d] queryCombinedSensorUpdateKernelsTooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryCombinedSensorUpdateKernelsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedSensorUpdateKernelsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedSensorUpdateKernelsInternalServerError creates a QueryCombinedSensorUpdateKernelsInternalServerError with default headers values
func NewQueryCombinedSensorUpdateKernelsInternalServerError() *QueryCombinedSensorUpdateKernelsInternalServerError {
	return &QueryCombinedSensorUpdateKernelsInternalServerError{}
}

/*
	QueryCombinedSensorUpdateKernelsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedSensorUpdateKernelsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdateKernelsV1
}

func (o *QueryCombinedSensorUpdateKernelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-kernels/v1][%d] queryCombinedSensorUpdateKernelsInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryCombinedSensorUpdateKernelsInternalServerError) GetPayload() *models.ResponsesSensorUpdateKernelsV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdateKernelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdateKernelsV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedSensorUpdateKernelsDefault creates a QueryCombinedSensorUpdateKernelsDefault with default headers values
func NewQueryCombinedSensorUpdateKernelsDefault(code int) *QueryCombinedSensorUpdateKernelsDefault {
	return &QueryCombinedSensorUpdateKernelsDefault{
		_statusCode: code,
	}
}

/*
	QueryCombinedSensorUpdateKernelsDefault describes a response with status code -1, with default header values.

OK
*/
type QueryCombinedSensorUpdateKernelsDefault struct {
	_statusCode int

	Payload *models.ResponsesSensorUpdateKernelsV1
}

// Code gets the status code for the query combined sensor update kernels default response
func (o *QueryCombinedSensorUpdateKernelsDefault) Code() int {
	return o._statusCode
}

func (o *QueryCombinedSensorUpdateKernelsDefault) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-kernels/v1][%d] queryCombinedSensorUpdateKernels default  %+v", o._statusCode, o.Payload)
}
func (o *QueryCombinedSensorUpdateKernelsDefault) GetPayload() *models.ResponsesSensorUpdateKernelsV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdateKernelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesSensorUpdateKernelsV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
