// Code generated by go-swagger; DO NOT EDIT.

package quick_scan

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

// ScanSamplesReader is a Reader for the ScanSamples structure.
type ScanSamplesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScanSamplesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScanSamplesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewScanSamplesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewScanSamplesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewScanSamplesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewScanSamplesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewScanSamplesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScanSamplesOK creates a ScanSamplesOK with default headers values
func NewScanSamplesOK() *ScanSamplesOK {
	return &ScanSamplesOK{}
}

/*
	ScanSamplesOK describes a response with status code 200, with default header values.

OK
*/
type ScanSamplesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MlscannerQueryResponse
}

func (o *ScanSamplesOK) Error() string {
	return fmt.Sprintf("[POST /scanner/entities/scans/v1][%d] scanSamplesOK  %+v", 200, o.Payload)
}
func (o *ScanSamplesOK) GetPayload() *models.MlscannerQueryResponse {
	return o.Payload
}

func (o *ScanSamplesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MlscannerQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScanSamplesBadRequest creates a ScanSamplesBadRequest with default headers values
func NewScanSamplesBadRequest() *ScanSamplesBadRequest {
	return &ScanSamplesBadRequest{}
}

/*
	ScanSamplesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ScanSamplesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MlscannerQueryResponse
}

func (o *ScanSamplesBadRequest) Error() string {
	return fmt.Sprintf("[POST /scanner/entities/scans/v1][%d] scanSamplesBadRequest  %+v", 400, o.Payload)
}
func (o *ScanSamplesBadRequest) GetPayload() *models.MlscannerQueryResponse {
	return o.Payload
}

func (o *ScanSamplesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MlscannerQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScanSamplesForbidden creates a ScanSamplesForbidden with default headers values
func NewScanSamplesForbidden() *ScanSamplesForbidden {
	return &ScanSamplesForbidden{}
}

/*
	ScanSamplesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ScanSamplesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *ScanSamplesForbidden) Error() string {
	return fmt.Sprintf("[POST /scanner/entities/scans/v1][%d] scanSamplesForbidden  %+v", 403, o.Payload)
}
func (o *ScanSamplesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ScanSamplesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScanSamplesTooManyRequests creates a ScanSamplesTooManyRequests with default headers values
func NewScanSamplesTooManyRequests() *ScanSamplesTooManyRequests {
	return &ScanSamplesTooManyRequests{}
}

/*
	ScanSamplesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ScanSamplesTooManyRequests struct {

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

func (o *ScanSamplesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /scanner/entities/scans/v1][%d] scanSamplesTooManyRequests  %+v", 429, o.Payload)
}
func (o *ScanSamplesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ScanSamplesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScanSamplesInternalServerError creates a ScanSamplesInternalServerError with default headers values
func NewScanSamplesInternalServerError() *ScanSamplesInternalServerError {
	return &ScanSamplesInternalServerError{}
}

/*
	ScanSamplesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ScanSamplesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MlscannerQueryResponse
}

func (o *ScanSamplesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /scanner/entities/scans/v1][%d] scanSamplesInternalServerError  %+v", 500, o.Payload)
}
func (o *ScanSamplesInternalServerError) GetPayload() *models.MlscannerQueryResponse {
	return o.Payload
}

func (o *ScanSamplesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MlscannerQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScanSamplesDefault creates a ScanSamplesDefault with default headers values
func NewScanSamplesDefault(code int) *ScanSamplesDefault {
	return &ScanSamplesDefault{
		_statusCode: code,
	}
}

/*
	ScanSamplesDefault describes a response with status code -1, with default header values.

OK
*/
type ScanSamplesDefault struct {
	_statusCode int

	Payload *models.MlscannerQueryResponse
}

// Code gets the status code for the scan samples default response
func (o *ScanSamplesDefault) Code() int {
	return o._statusCode
}

func (o *ScanSamplesDefault) Error() string {
	return fmt.Sprintf("[POST /scanner/entities/scans/v1][%d] ScanSamples default  %+v", o._statusCode, o.Payload)
}
func (o *ScanSamplesDefault) GetPayload() *models.MlscannerQueryResponse {
	return o.Payload
}

func (o *ScanSamplesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MlscannerQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
