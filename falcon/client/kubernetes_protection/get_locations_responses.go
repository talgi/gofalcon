// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// GetLocationsReader is a Reader for the GetLocations structure.
type GetLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetLocationsMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLocationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLocationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLocationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLocationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetLocationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLocationsOK creates a GetLocationsOK with default headers values
func NewGetLocationsOK() *GetLocationsOK {
	return &GetLocationsOK{}
}

/*
GetLocationsOK describes a response with status code 200, with default header values.

OK
*/
type GetLocationsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetLocationsResp
}

// IsSuccess returns true when this get locations o k response has a 2xx status code
func (o *GetLocationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get locations o k response has a 3xx status code
func (o *GetLocationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get locations o k response has a 4xx status code
func (o *GetLocationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get locations o k response has a 5xx status code
func (o *GetLocationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get locations o k response a status code equal to that given
func (o *GetLocationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get locations o k response
func (o *GetLocationsOK) Code() int {
	return 200
}

func (o *GetLocationsOK) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsOK  %+v", 200, o.Payload)
}

func (o *GetLocationsOK) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsOK  %+v", 200, o.Payload)
}

func (o *GetLocationsOK) GetPayload() *models.K8sregGetLocationsResp {
	return o.Payload
}

func (o *GetLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetLocationsResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsMultiStatus creates a GetLocationsMultiStatus with default headers values
func NewGetLocationsMultiStatus() *GetLocationsMultiStatus {
	return &GetLocationsMultiStatus{}
}

/*
GetLocationsMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetLocationsMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetLocationsResp
}

// IsSuccess returns true when this get locations multi status response has a 2xx status code
func (o *GetLocationsMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get locations multi status response has a 3xx status code
func (o *GetLocationsMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get locations multi status response has a 4xx status code
func (o *GetLocationsMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get locations multi status response has a 5xx status code
func (o *GetLocationsMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get locations multi status response a status code equal to that given
func (o *GetLocationsMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the get locations multi status response
func (o *GetLocationsMultiStatus) Code() int {
	return 207
}

func (o *GetLocationsMultiStatus) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsMultiStatus  %+v", 207, o.Payload)
}

func (o *GetLocationsMultiStatus) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsMultiStatus  %+v", 207, o.Payload)
}

func (o *GetLocationsMultiStatus) GetPayload() *models.K8sregGetLocationsResp {
	return o.Payload
}

func (o *GetLocationsMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetLocationsResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsBadRequest creates a GetLocationsBadRequest with default headers values
func NewGetLocationsBadRequest() *GetLocationsBadRequest {
	return &GetLocationsBadRequest{}
}

/*
GetLocationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetLocationsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetLocationsResp
}

// IsSuccess returns true when this get locations bad request response has a 2xx status code
func (o *GetLocationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get locations bad request response has a 3xx status code
func (o *GetLocationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get locations bad request response has a 4xx status code
func (o *GetLocationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get locations bad request response has a 5xx status code
func (o *GetLocationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get locations bad request response a status code equal to that given
func (o *GetLocationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get locations bad request response
func (o *GetLocationsBadRequest) Code() int {
	return 400
}

func (o *GetLocationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLocationsBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLocationsBadRequest) GetPayload() *models.K8sregGetLocationsResp {
	return o.Payload
}

func (o *GetLocationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetLocationsResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsForbidden creates a GetLocationsForbidden with default headers values
func NewGetLocationsForbidden() *GetLocationsForbidden {
	return &GetLocationsForbidden{}
}

/*
GetLocationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLocationsForbidden struct {

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

// IsSuccess returns true when this get locations forbidden response has a 2xx status code
func (o *GetLocationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get locations forbidden response has a 3xx status code
func (o *GetLocationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get locations forbidden response has a 4xx status code
func (o *GetLocationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get locations forbidden response has a 5xx status code
func (o *GetLocationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get locations forbidden response a status code equal to that given
func (o *GetLocationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get locations forbidden response
func (o *GetLocationsForbidden) Code() int {
	return 403
}

func (o *GetLocationsForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsForbidden  %+v", 403, o.Payload)
}

func (o *GetLocationsForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsForbidden  %+v", 403, o.Payload)
}

func (o *GetLocationsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetLocationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLocationsTooManyRequests creates a GetLocationsTooManyRequests with default headers values
func NewGetLocationsTooManyRequests() *GetLocationsTooManyRequests {
	return &GetLocationsTooManyRequests{}
}

/*
GetLocationsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetLocationsTooManyRequests struct {

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

// IsSuccess returns true when this get locations too many requests response has a 2xx status code
func (o *GetLocationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get locations too many requests response has a 3xx status code
func (o *GetLocationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get locations too many requests response has a 4xx status code
func (o *GetLocationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get locations too many requests response has a 5xx status code
func (o *GetLocationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get locations too many requests response a status code equal to that given
func (o *GetLocationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get locations too many requests response
func (o *GetLocationsTooManyRequests) Code() int {
	return 429
}

func (o *GetLocationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLocationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLocationsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetLocationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLocationsInternalServerError creates a GetLocationsInternalServerError with default headers values
func NewGetLocationsInternalServerError() *GetLocationsInternalServerError {
	return &GetLocationsInternalServerError{}
}

/*
GetLocationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetLocationsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetLocationsResp
}

// IsSuccess returns true when this get locations internal server error response has a 2xx status code
func (o *GetLocationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get locations internal server error response has a 3xx status code
func (o *GetLocationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get locations internal server error response has a 4xx status code
func (o *GetLocationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get locations internal server error response has a 5xx status code
func (o *GetLocationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get locations internal server error response a status code equal to that given
func (o *GetLocationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get locations internal server error response
func (o *GetLocationsInternalServerError) Code() int {
	return 500
}

func (o *GetLocationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLocationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] getLocationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLocationsInternalServerError) GetPayload() *models.K8sregGetLocationsResp {
	return o.Payload
}

func (o *GetLocationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetLocationsResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsDefault creates a GetLocationsDefault with default headers values
func NewGetLocationsDefault(code int) *GetLocationsDefault {
	return &GetLocationsDefault{
		_statusCode: code,
	}
}

/*
GetLocationsDefault describes a response with status code -1, with default header values.

OK
*/
type GetLocationsDefault struct {
	_statusCode int

	Payload *models.K8sregGetLocationsResp
}

// IsSuccess returns true when this get locations default response has a 2xx status code
func (o *GetLocationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get locations default response has a 3xx status code
func (o *GetLocationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get locations default response has a 4xx status code
func (o *GetLocationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get locations default response has a 5xx status code
func (o *GetLocationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get locations default response a status code equal to that given
func (o *GetLocationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get locations default response
func (o *GetLocationsDefault) Code() int {
	return o._statusCode
}

func (o *GetLocationsDefault) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] GetLocations default  %+v", o._statusCode, o.Payload)
}

func (o *GetLocationsDefault) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/cloud-locations/v1][%d] GetLocations default  %+v", o._statusCode, o.Payload)
}

func (o *GetLocationsDefault) GetPayload() *models.K8sregGetLocationsResp {
	return o.Payload
}

func (o *GetLocationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8sregGetLocationsResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
