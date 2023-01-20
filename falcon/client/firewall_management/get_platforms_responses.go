// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// GetPlatformsReader is a Reader for the GetPlatforms structure.
type GetPlatformsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlatformsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPlatformsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPlatformsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPlatformsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPlatformsOK creates a GetPlatformsOK with default headers values
func NewGetPlatformsOK() *GetPlatformsOK {
	return &GetPlatformsOK{}
}

/*
GetPlatformsOK describes a response with status code 200, with default header values.

OK
*/
type GetPlatformsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrAPIPlatformsResponse
}

// IsSuccess returns true when this get platforms o k response has a 2xx status code
func (o *GetPlatformsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get platforms o k response has a 3xx status code
func (o *GetPlatformsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get platforms o k response has a 4xx status code
func (o *GetPlatformsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get platforms o k response has a 5xx status code
func (o *GetPlatformsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get platforms o k response a status code equal to that given
func (o *GetPlatformsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get platforms o k response
func (o *GetPlatformsOK) Code() int {
	return 200
}

func (o *GetPlatformsOK) Error() string {
	return fmt.Sprintf("[GET /fwmgr/entities/platforms/v1][%d] getPlatformsOK  %+v", 200, o.Payload)
}

func (o *GetPlatformsOK) String() string {
	return fmt.Sprintf("[GET /fwmgr/entities/platforms/v1][%d] getPlatformsOK  %+v", 200, o.Payload)
}

func (o *GetPlatformsOK) GetPayload() *models.FwmgrAPIPlatformsResponse {
	return o.Payload
}

func (o *GetPlatformsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrAPIPlatformsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformsForbidden creates a GetPlatformsForbidden with default headers values
func NewGetPlatformsForbidden() *GetPlatformsForbidden {
	return &GetPlatformsForbidden{}
}

/*
GetPlatformsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPlatformsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get platforms forbidden response has a 2xx status code
func (o *GetPlatformsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get platforms forbidden response has a 3xx status code
func (o *GetPlatformsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get platforms forbidden response has a 4xx status code
func (o *GetPlatformsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get platforms forbidden response has a 5xx status code
func (o *GetPlatformsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get platforms forbidden response a status code equal to that given
func (o *GetPlatformsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get platforms forbidden response
func (o *GetPlatformsForbidden) Code() int {
	return 403
}

func (o *GetPlatformsForbidden) Error() string {
	return fmt.Sprintf("[GET /fwmgr/entities/platforms/v1][%d] getPlatformsForbidden  %+v", 403, o.Payload)
}

func (o *GetPlatformsForbidden) String() string {
	return fmt.Sprintf("[GET /fwmgr/entities/platforms/v1][%d] getPlatformsForbidden  %+v", 403, o.Payload)
}

func (o *GetPlatformsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetPlatformsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPlatformsTooManyRequests creates a GetPlatformsTooManyRequests with default headers values
func NewGetPlatformsTooManyRequests() *GetPlatformsTooManyRequests {
	return &GetPlatformsTooManyRequests{}
}

/*
GetPlatformsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetPlatformsTooManyRequests struct {

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

// IsSuccess returns true when this get platforms too many requests response has a 2xx status code
func (o *GetPlatformsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get platforms too many requests response has a 3xx status code
func (o *GetPlatformsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get platforms too many requests response has a 4xx status code
func (o *GetPlatformsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get platforms too many requests response has a 5xx status code
func (o *GetPlatformsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get platforms too many requests response a status code equal to that given
func (o *GetPlatformsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get platforms too many requests response
func (o *GetPlatformsTooManyRequests) Code() int {
	return 429
}

func (o *GetPlatformsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fwmgr/entities/platforms/v1][%d] getPlatformsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPlatformsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fwmgr/entities/platforms/v1][%d] getPlatformsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPlatformsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetPlatformsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPlatformsDefault creates a GetPlatformsDefault with default headers values
func NewGetPlatformsDefault(code int) *GetPlatformsDefault {
	return &GetPlatformsDefault{
		_statusCode: code,
	}
}

/*
GetPlatformsDefault describes a response with status code -1, with default header values.

OK
*/
type GetPlatformsDefault struct {
	_statusCode int

	Payload *models.FwmgrAPIPlatformsResponse
}

// IsSuccess returns true when this get platforms default response has a 2xx status code
func (o *GetPlatformsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get platforms default response has a 3xx status code
func (o *GetPlatformsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get platforms default response has a 4xx status code
func (o *GetPlatformsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get platforms default response has a 5xx status code
func (o *GetPlatformsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get platforms default response a status code equal to that given
func (o *GetPlatformsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get platforms default response
func (o *GetPlatformsDefault) Code() int {
	return o._statusCode
}

func (o *GetPlatformsDefault) Error() string {
	return fmt.Sprintf("[GET /fwmgr/entities/platforms/v1][%d] get-platforms default  %+v", o._statusCode, o.Payload)
}

func (o *GetPlatformsDefault) String() string {
	return fmt.Sprintf("[GET /fwmgr/entities/platforms/v1][%d] get-platforms default  %+v", o._statusCode, o.Payload)
}

func (o *GetPlatformsDefault) GetPayload() *models.FwmgrAPIPlatformsResponse {
	return o.Payload
}

func (o *GetPlatformsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FwmgrAPIPlatformsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
