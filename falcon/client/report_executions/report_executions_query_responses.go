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

// ReportExecutionsQueryReader is a Reader for the ReportExecutionsQuery structure.
type ReportExecutionsQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportExecutionsQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportExecutionsQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReportExecutionsQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReportExecutionsQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReportExecutionsQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReportExecutionsQueryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReportExecutionsQueryOK creates a ReportExecutionsQueryOK with default headers values
func NewReportExecutionsQueryOK() *ReportExecutionsQueryOK {
	return &ReportExecutionsQueryOK{}
}

/*
ReportExecutionsQueryOK describes a response with status code 200, with default header values.

OK
*/
type ReportExecutionsQueryOK struct {

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

// IsSuccess returns true when this report executions query o k response has a 2xx status code
func (o *ReportExecutionsQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this report executions query o k response has a 3xx status code
func (o *ReportExecutionsQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions query o k response has a 4xx status code
func (o *ReportExecutionsQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this report executions query o k response has a 5xx status code
func (o *ReportExecutionsQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions query o k response a status code equal to that given
func (o *ReportExecutionsQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *ReportExecutionsQueryOK) Error() string {
	return fmt.Sprintf("[GET /reports/queries/report-executions/v1][%d] reportExecutionsQueryOK  %+v", 200, o.Payload)
}

func (o *ReportExecutionsQueryOK) String() string {
	return fmt.Sprintf("[GET /reports/queries/report-executions/v1][%d] reportExecutionsQueryOK  %+v", 200, o.Payload)
}

func (o *ReportExecutionsQueryOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *ReportExecutionsQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReportExecutionsQueryBadRequest creates a ReportExecutionsQueryBadRequest with default headers values
func NewReportExecutionsQueryBadRequest() *ReportExecutionsQueryBadRequest {
	return &ReportExecutionsQueryBadRequest{}
}

/*
ReportExecutionsQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ReportExecutionsQueryBadRequest struct {

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

// IsSuccess returns true when this report executions query bad request response has a 2xx status code
func (o *ReportExecutionsQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report executions query bad request response has a 3xx status code
func (o *ReportExecutionsQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions query bad request response has a 4xx status code
func (o *ReportExecutionsQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this report executions query bad request response has a 5xx status code
func (o *ReportExecutionsQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions query bad request response a status code equal to that given
func (o *ReportExecutionsQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ReportExecutionsQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/queries/report-executions/v1][%d] reportExecutionsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *ReportExecutionsQueryBadRequest) String() string {
	return fmt.Sprintf("[GET /reports/queries/report-executions/v1][%d] reportExecutionsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *ReportExecutionsQueryBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReportExecutionsQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReportExecutionsQueryForbidden creates a ReportExecutionsQueryForbidden with default headers values
func NewReportExecutionsQueryForbidden() *ReportExecutionsQueryForbidden {
	return &ReportExecutionsQueryForbidden{}
}

/*
ReportExecutionsQueryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReportExecutionsQueryForbidden struct {

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

// IsSuccess returns true when this report executions query forbidden response has a 2xx status code
func (o *ReportExecutionsQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report executions query forbidden response has a 3xx status code
func (o *ReportExecutionsQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions query forbidden response has a 4xx status code
func (o *ReportExecutionsQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this report executions query forbidden response has a 5xx status code
func (o *ReportExecutionsQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions query forbidden response a status code equal to that given
func (o *ReportExecutionsQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ReportExecutionsQueryForbidden) Error() string {
	return fmt.Sprintf("[GET /reports/queries/report-executions/v1][%d] reportExecutionsQueryForbidden  %+v", 403, o.Payload)
}

func (o *ReportExecutionsQueryForbidden) String() string {
	return fmt.Sprintf("[GET /reports/queries/report-executions/v1][%d] reportExecutionsQueryForbidden  %+v", 403, o.Payload)
}

func (o *ReportExecutionsQueryForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReportExecutionsQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReportExecutionsQueryTooManyRequests creates a ReportExecutionsQueryTooManyRequests with default headers values
func NewReportExecutionsQueryTooManyRequests() *ReportExecutionsQueryTooManyRequests {
	return &ReportExecutionsQueryTooManyRequests{}
}

/*
ReportExecutionsQueryTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ReportExecutionsQueryTooManyRequests struct {

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

// IsSuccess returns true when this report executions query too many requests response has a 2xx status code
func (o *ReportExecutionsQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report executions query too many requests response has a 3xx status code
func (o *ReportExecutionsQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions query too many requests response has a 4xx status code
func (o *ReportExecutionsQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this report executions query too many requests response has a 5xx status code
func (o *ReportExecutionsQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions query too many requests response a status code equal to that given
func (o *ReportExecutionsQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ReportExecutionsQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reports/queries/report-executions/v1][%d] reportExecutionsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReportExecutionsQueryTooManyRequests) String() string {
	return fmt.Sprintf("[GET /reports/queries/report-executions/v1][%d] reportExecutionsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReportExecutionsQueryTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReportExecutionsQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReportExecutionsQueryDefault creates a ReportExecutionsQueryDefault with default headers values
func NewReportExecutionsQueryDefault(code int) *ReportExecutionsQueryDefault {
	return &ReportExecutionsQueryDefault{
		_statusCode: code,
	}
}

/*
ReportExecutionsQueryDefault describes a response with status code -1, with default header values.

OK
*/
type ReportExecutionsQueryDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the report executions query default response
func (o *ReportExecutionsQueryDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this report executions query default response has a 2xx status code
func (o *ReportExecutionsQueryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this report executions query default response has a 3xx status code
func (o *ReportExecutionsQueryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this report executions query default response has a 4xx status code
func (o *ReportExecutionsQueryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this report executions query default response has a 5xx status code
func (o *ReportExecutionsQueryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this report executions query default response a status code equal to that given
func (o *ReportExecutionsQueryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ReportExecutionsQueryDefault) Error() string {
	return fmt.Sprintf("[GET /reports/queries/report-executions/v1][%d] report-executions.query default  %+v", o._statusCode, o.Payload)
}

func (o *ReportExecutionsQueryDefault) String() string {
	return fmt.Sprintf("[GET /reports/queries/report-executions/v1][%d] report-executions.query default  %+v", o._statusCode, o.Payload)
}

func (o *ReportExecutionsQueryDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *ReportExecutionsQueryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
