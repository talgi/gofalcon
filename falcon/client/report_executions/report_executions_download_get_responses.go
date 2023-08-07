// Code generated by go-swagger; DO NOT EDIT.

package report_executions

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

// ReportExecutionsDownloadGetReader is a Reader for the ReportExecutionsDownloadGet structure.
type ReportExecutionsDownloadGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportExecutionsDownloadGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportExecutionsDownloadGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReportExecutionsDownloadGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReportExecutionsDownloadGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReportExecutionsDownloadGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReportExecutionsDownloadGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /reports/entities/report-executions-download/v1] report-executions-download.get", response, response.Code())
	}
}

// NewReportExecutionsDownloadGetOK creates a ReportExecutionsDownloadGetOK with default headers values
func NewReportExecutionsDownloadGetOK() *ReportExecutionsDownloadGetOK {
	return &ReportExecutionsDownloadGetOK{}
}

/*
ReportExecutionsDownloadGetOK describes a response with status code 200, with default header values.

OK
*/
type ReportExecutionsDownloadGetOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload []int64
}

// IsSuccess returns true when this report executions download get o k response has a 2xx status code
func (o *ReportExecutionsDownloadGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this report executions download get o k response has a 3xx status code
func (o *ReportExecutionsDownloadGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions download get o k response has a 4xx status code
func (o *ReportExecutionsDownloadGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this report executions download get o k response has a 5xx status code
func (o *ReportExecutionsDownloadGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions download get o k response a status code equal to that given
func (o *ReportExecutionsDownloadGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the report executions download get o k response
func (o *ReportExecutionsDownloadGetOK) Code() int {
	return 200
}

func (o *ReportExecutionsDownloadGetOK) Error() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions-download/v1][%d] reportExecutionsDownloadGetOK  %+v", 200, o.Payload)
}

func (o *ReportExecutionsDownloadGetOK) String() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions-download/v1][%d] reportExecutionsDownloadGetOK  %+v", 200, o.Payload)
}

func (o *ReportExecutionsDownloadGetOK) GetPayload() []int64 {
	return o.Payload
}

func (o *ReportExecutionsDownloadGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReportExecutionsDownloadGetBadRequest creates a ReportExecutionsDownloadGetBadRequest with default headers values
func NewReportExecutionsDownloadGetBadRequest() *ReportExecutionsDownloadGetBadRequest {
	return &ReportExecutionsDownloadGetBadRequest{}
}

/*
ReportExecutionsDownloadGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ReportExecutionsDownloadGetBadRequest struct {

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

// IsSuccess returns true when this report executions download get bad request response has a 2xx status code
func (o *ReportExecutionsDownloadGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report executions download get bad request response has a 3xx status code
func (o *ReportExecutionsDownloadGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions download get bad request response has a 4xx status code
func (o *ReportExecutionsDownloadGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this report executions download get bad request response has a 5xx status code
func (o *ReportExecutionsDownloadGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions download get bad request response a status code equal to that given
func (o *ReportExecutionsDownloadGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the report executions download get bad request response
func (o *ReportExecutionsDownloadGetBadRequest) Code() int {
	return 400
}

func (o *ReportExecutionsDownloadGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions-download/v1][%d] reportExecutionsDownloadGetBadRequest  %+v", 400, o.Payload)
}

func (o *ReportExecutionsDownloadGetBadRequest) String() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions-download/v1][%d] reportExecutionsDownloadGetBadRequest  %+v", 400, o.Payload)
}

func (o *ReportExecutionsDownloadGetBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReportExecutionsDownloadGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReportExecutionsDownloadGetForbidden creates a ReportExecutionsDownloadGetForbidden with default headers values
func NewReportExecutionsDownloadGetForbidden() *ReportExecutionsDownloadGetForbidden {
	return &ReportExecutionsDownloadGetForbidden{}
}

/*
ReportExecutionsDownloadGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReportExecutionsDownloadGetForbidden struct {

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

// IsSuccess returns true when this report executions download get forbidden response has a 2xx status code
func (o *ReportExecutionsDownloadGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report executions download get forbidden response has a 3xx status code
func (o *ReportExecutionsDownloadGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions download get forbidden response has a 4xx status code
func (o *ReportExecutionsDownloadGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this report executions download get forbidden response has a 5xx status code
func (o *ReportExecutionsDownloadGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions download get forbidden response a status code equal to that given
func (o *ReportExecutionsDownloadGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the report executions download get forbidden response
func (o *ReportExecutionsDownloadGetForbidden) Code() int {
	return 403
}

func (o *ReportExecutionsDownloadGetForbidden) Error() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions-download/v1][%d] reportExecutionsDownloadGetForbidden  %+v", 403, o.Payload)
}

func (o *ReportExecutionsDownloadGetForbidden) String() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions-download/v1][%d] reportExecutionsDownloadGetForbidden  %+v", 403, o.Payload)
}

func (o *ReportExecutionsDownloadGetForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReportExecutionsDownloadGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReportExecutionsDownloadGetTooManyRequests creates a ReportExecutionsDownloadGetTooManyRequests with default headers values
func NewReportExecutionsDownloadGetTooManyRequests() *ReportExecutionsDownloadGetTooManyRequests {
	return &ReportExecutionsDownloadGetTooManyRequests{}
}

/*
ReportExecutionsDownloadGetTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ReportExecutionsDownloadGetTooManyRequests struct {

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

// IsSuccess returns true when this report executions download get too many requests response has a 2xx status code
func (o *ReportExecutionsDownloadGetTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report executions download get too many requests response has a 3xx status code
func (o *ReportExecutionsDownloadGetTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions download get too many requests response has a 4xx status code
func (o *ReportExecutionsDownloadGetTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this report executions download get too many requests response has a 5xx status code
func (o *ReportExecutionsDownloadGetTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions download get too many requests response a status code equal to that given
func (o *ReportExecutionsDownloadGetTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the report executions download get too many requests response
func (o *ReportExecutionsDownloadGetTooManyRequests) Code() int {
	return 429
}

func (o *ReportExecutionsDownloadGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions-download/v1][%d] reportExecutionsDownloadGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReportExecutionsDownloadGetTooManyRequests) String() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions-download/v1][%d] reportExecutionsDownloadGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReportExecutionsDownloadGetTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReportExecutionsDownloadGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReportExecutionsDownloadGetInternalServerError creates a ReportExecutionsDownloadGetInternalServerError with default headers values
func NewReportExecutionsDownloadGetInternalServerError() *ReportExecutionsDownloadGetInternalServerError {
	return &ReportExecutionsDownloadGetInternalServerError{}
}

/*
ReportExecutionsDownloadGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReportExecutionsDownloadGetInternalServerError struct {

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

// IsSuccess returns true when this report executions download get internal server error response has a 2xx status code
func (o *ReportExecutionsDownloadGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report executions download get internal server error response has a 3xx status code
func (o *ReportExecutionsDownloadGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions download get internal server error response has a 4xx status code
func (o *ReportExecutionsDownloadGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this report executions download get internal server error response has a 5xx status code
func (o *ReportExecutionsDownloadGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this report executions download get internal server error response a status code equal to that given
func (o *ReportExecutionsDownloadGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the report executions download get internal server error response
func (o *ReportExecutionsDownloadGetInternalServerError) Code() int {
	return 500
}

func (o *ReportExecutionsDownloadGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions-download/v1][%d] reportExecutionsDownloadGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ReportExecutionsDownloadGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions-download/v1][%d] reportExecutionsDownloadGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ReportExecutionsDownloadGetInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReportExecutionsDownloadGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
