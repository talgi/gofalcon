// Code generated by go-swagger; DO NOT EDIT.

package ods

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

// DeleteScheduledScansReader is a Reader for the DeleteScheduledScans structure.
type DeleteScheduledScansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteScheduledScansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteScheduledScansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteScheduledScansForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteScheduledScansNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteScheduledScansTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteScheduledScansOK creates a DeleteScheduledScansOK with default headers values
func NewDeleteScheduledScansOK() *DeleteScheduledScansOK {
	return &DeleteScheduledScansOK{}
}

/*
DeleteScheduledScansOK describes a response with status code 200, with default header values.

OK
*/
type DeleteScheduledScansOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this delete scheduled scans o k response has a 2xx status code
func (o *DeleteScheduledScansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete scheduled scans o k response has a 3xx status code
func (o *DeleteScheduledScansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete scheduled scans o k response has a 4xx status code
func (o *DeleteScheduledScansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete scheduled scans o k response has a 5xx status code
func (o *DeleteScheduledScansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete scheduled scans o k response a status code equal to that given
func (o *DeleteScheduledScansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete scheduled scans o k response
func (o *DeleteScheduledScansOK) Code() int {
	return 200
}

func (o *DeleteScheduledScansOK) Error() string {
	return fmt.Sprintf("[DELETE /ods/entities/scheduled-scans/v1][%d] deleteScheduledScansOK  %+v", 200, o.Payload)
}

func (o *DeleteScheduledScansOK) String() string {
	return fmt.Sprintf("[DELETE /ods/entities/scheduled-scans/v1][%d] deleteScheduledScansOK  %+v", 200, o.Payload)
}

func (o *DeleteScheduledScansOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *DeleteScheduledScansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduledScansForbidden creates a DeleteScheduledScansForbidden with default headers values
func NewDeleteScheduledScansForbidden() *DeleteScheduledScansForbidden {
	return &DeleteScheduledScansForbidden{}
}

/*
DeleteScheduledScansForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteScheduledScansForbidden struct {

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

// IsSuccess returns true when this delete scheduled scans forbidden response has a 2xx status code
func (o *DeleteScheduledScansForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete scheduled scans forbidden response has a 3xx status code
func (o *DeleteScheduledScansForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete scheduled scans forbidden response has a 4xx status code
func (o *DeleteScheduledScansForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete scheduled scans forbidden response has a 5xx status code
func (o *DeleteScheduledScansForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete scheduled scans forbidden response a status code equal to that given
func (o *DeleteScheduledScansForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete scheduled scans forbidden response
func (o *DeleteScheduledScansForbidden) Code() int {
	return 403
}

func (o *DeleteScheduledScansForbidden) Error() string {
	return fmt.Sprintf("[DELETE /ods/entities/scheduled-scans/v1][%d] deleteScheduledScansForbidden  %+v", 403, o.Payload)
}

func (o *DeleteScheduledScansForbidden) String() string {
	return fmt.Sprintf("[DELETE /ods/entities/scheduled-scans/v1][%d] deleteScheduledScansForbidden  %+v", 403, o.Payload)
}

func (o *DeleteScheduledScansForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteScheduledScansForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteScheduledScansNotFound creates a DeleteScheduledScansNotFound with default headers values
func NewDeleteScheduledScansNotFound() *DeleteScheduledScansNotFound {
	return &DeleteScheduledScansNotFound{}
}

/*
DeleteScheduledScansNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteScheduledScansNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this delete scheduled scans not found response has a 2xx status code
func (o *DeleteScheduledScansNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete scheduled scans not found response has a 3xx status code
func (o *DeleteScheduledScansNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete scheduled scans not found response has a 4xx status code
func (o *DeleteScheduledScansNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete scheduled scans not found response has a 5xx status code
func (o *DeleteScheduledScansNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete scheduled scans not found response a status code equal to that given
func (o *DeleteScheduledScansNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete scheduled scans not found response
func (o *DeleteScheduledScansNotFound) Code() int {
	return 404
}

func (o *DeleteScheduledScansNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ods/entities/scheduled-scans/v1][%d] deleteScheduledScansNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScheduledScansNotFound) String() string {
	return fmt.Sprintf("[DELETE /ods/entities/scheduled-scans/v1][%d] deleteScheduledScansNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScheduledScansNotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *DeleteScheduledScansNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduledScansTooManyRequests creates a DeleteScheduledScansTooManyRequests with default headers values
func NewDeleteScheduledScansTooManyRequests() *DeleteScheduledScansTooManyRequests {
	return &DeleteScheduledScansTooManyRequests{}
}

/*
DeleteScheduledScansTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteScheduledScansTooManyRequests struct {

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

// IsSuccess returns true when this delete scheduled scans too many requests response has a 2xx status code
func (o *DeleteScheduledScansTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete scheduled scans too many requests response has a 3xx status code
func (o *DeleteScheduledScansTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete scheduled scans too many requests response has a 4xx status code
func (o *DeleteScheduledScansTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete scheduled scans too many requests response has a 5xx status code
func (o *DeleteScheduledScansTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete scheduled scans too many requests response a status code equal to that given
func (o *DeleteScheduledScansTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete scheduled scans too many requests response
func (o *DeleteScheduledScansTooManyRequests) Code() int {
	return 429
}

func (o *DeleteScheduledScansTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /ods/entities/scheduled-scans/v1][%d] deleteScheduledScansTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteScheduledScansTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /ods/entities/scheduled-scans/v1][%d] deleteScheduledScansTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteScheduledScansTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteScheduledScansTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
