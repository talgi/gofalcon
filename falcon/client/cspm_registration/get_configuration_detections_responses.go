// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// GetConfigurationDetectionsReader is a Reader for the GetConfigurationDetections structure.
type GetConfigurationDetectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationDetectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationDetectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConfigurationDetectionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConfigurationDetectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConfigurationDetectionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConfigurationDetectionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetConfigurationDetectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetConfigurationDetectionsOK creates a GetConfigurationDetectionsOK with default headers values
func NewGetConfigurationDetectionsOK() *GetConfigurationDetectionsOK {
	return &GetConfigurationDetectionsOK{}
}

/*
GetConfigurationDetectionsOK describes a response with status code 200, with default header values.

OK
*/
type GetConfigurationDetectionsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationExternalIOMEventResponse
}

// IsSuccess returns true when this get configuration detections o k response has a 2xx status code
func (o *GetConfigurationDetectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get configuration detections o k response has a 3xx status code
func (o *GetConfigurationDetectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detections o k response has a 4xx status code
func (o *GetConfigurationDetectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configuration detections o k response has a 5xx status code
func (o *GetConfigurationDetectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detections o k response a status code equal to that given
func (o *GetConfigurationDetectionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConfigurationDetectionsOK) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] getConfigurationDetectionsOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationDetectionsOK) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] getConfigurationDetectionsOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationDetectionsOK) GetPayload() *models.RegistrationExternalIOMEventResponse {
	return o.Payload
}

func (o *GetConfigurationDetectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationExternalIOMEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationDetectionsBadRequest creates a GetConfigurationDetectionsBadRequest with default headers values
func NewGetConfigurationDetectionsBadRequest() *GetConfigurationDetectionsBadRequest {
	return &GetConfigurationDetectionsBadRequest{}
}

/*
GetConfigurationDetectionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetConfigurationDetectionsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationExternalIOMEventResponse
}

// IsSuccess returns true when this get configuration detections bad request response has a 2xx status code
func (o *GetConfigurationDetectionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detections bad request response has a 3xx status code
func (o *GetConfigurationDetectionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detections bad request response has a 4xx status code
func (o *GetConfigurationDetectionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration detections bad request response has a 5xx status code
func (o *GetConfigurationDetectionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detections bad request response a status code equal to that given
func (o *GetConfigurationDetectionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConfigurationDetectionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] getConfigurationDetectionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationDetectionsBadRequest) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] getConfigurationDetectionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationDetectionsBadRequest) GetPayload() *models.RegistrationExternalIOMEventResponse {
	return o.Payload
}

func (o *GetConfigurationDetectionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationExternalIOMEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationDetectionsForbidden creates a GetConfigurationDetectionsForbidden with default headers values
func NewGetConfigurationDetectionsForbidden() *GetConfigurationDetectionsForbidden {
	return &GetConfigurationDetectionsForbidden{}
}

/*
GetConfigurationDetectionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetConfigurationDetectionsForbidden struct {

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

// IsSuccess returns true when this get configuration detections forbidden response has a 2xx status code
func (o *GetConfigurationDetectionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detections forbidden response has a 3xx status code
func (o *GetConfigurationDetectionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detections forbidden response has a 4xx status code
func (o *GetConfigurationDetectionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration detections forbidden response has a 5xx status code
func (o *GetConfigurationDetectionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detections forbidden response a status code equal to that given
func (o *GetConfigurationDetectionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConfigurationDetectionsForbidden) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] getConfigurationDetectionsForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationDetectionsForbidden) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] getConfigurationDetectionsForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationDetectionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetConfigurationDetectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConfigurationDetectionsTooManyRequests creates a GetConfigurationDetectionsTooManyRequests with default headers values
func NewGetConfigurationDetectionsTooManyRequests() *GetConfigurationDetectionsTooManyRequests {
	return &GetConfigurationDetectionsTooManyRequests{}
}

/*
GetConfigurationDetectionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetConfigurationDetectionsTooManyRequests struct {

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

// IsSuccess returns true when this get configuration detections too many requests response has a 2xx status code
func (o *GetConfigurationDetectionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detections too many requests response has a 3xx status code
func (o *GetConfigurationDetectionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detections too many requests response has a 4xx status code
func (o *GetConfigurationDetectionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration detections too many requests response has a 5xx status code
func (o *GetConfigurationDetectionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detections too many requests response a status code equal to that given
func (o *GetConfigurationDetectionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConfigurationDetectionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] getConfigurationDetectionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfigurationDetectionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] getConfigurationDetectionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfigurationDetectionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetConfigurationDetectionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConfigurationDetectionsInternalServerError creates a GetConfigurationDetectionsInternalServerError with default headers values
func NewGetConfigurationDetectionsInternalServerError() *GetConfigurationDetectionsInternalServerError {
	return &GetConfigurationDetectionsInternalServerError{}
}

/*
GetConfigurationDetectionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetConfigurationDetectionsInternalServerError struct {

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

// IsSuccess returns true when this get configuration detections internal server error response has a 2xx status code
func (o *GetConfigurationDetectionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detections internal server error response has a 3xx status code
func (o *GetConfigurationDetectionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detections internal server error response has a 4xx status code
func (o *GetConfigurationDetectionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configuration detections internal server error response has a 5xx status code
func (o *GetConfigurationDetectionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get configuration detections internal server error response a status code equal to that given
func (o *GetConfigurationDetectionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConfigurationDetectionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] getConfigurationDetectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigurationDetectionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] getConfigurationDetectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigurationDetectionsInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetConfigurationDetectionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConfigurationDetectionsDefault creates a GetConfigurationDetectionsDefault with default headers values
func NewGetConfigurationDetectionsDefault(code int) *GetConfigurationDetectionsDefault {
	return &GetConfigurationDetectionsDefault{
		_statusCode: code,
	}
}

/*
GetConfigurationDetectionsDefault describes a response with status code -1, with default header values.

OK
*/
type GetConfigurationDetectionsDefault struct {
	_statusCode int

	Payload *models.RegistrationExternalIOMEventResponse
}

// Code gets the status code for the get configuration detections default response
func (o *GetConfigurationDetectionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get configuration detections default response has a 2xx status code
func (o *GetConfigurationDetectionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get configuration detections default response has a 3xx status code
func (o *GetConfigurationDetectionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get configuration detections default response has a 4xx status code
func (o *GetConfigurationDetectionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get configuration detections default response has a 5xx status code
func (o *GetConfigurationDetectionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get configuration detections default response a status code equal to that given
func (o *GetConfigurationDetectionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetConfigurationDetectionsDefault) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] GetConfigurationDetections default  %+v", o._statusCode, o.Payload)
}

func (o *GetConfigurationDetectionsDefault) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v1][%d] GetConfigurationDetections default  %+v", o._statusCode, o.Payload)
}

func (o *GetConfigurationDetectionsDefault) GetPayload() *models.RegistrationExternalIOMEventResponse {
	return o.Payload
}

func (o *GetConfigurationDetectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationExternalIOMEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
