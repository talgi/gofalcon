// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// GetIntelReportPDFReader is a Reader for the GetIntelReportPDF structure.
type GetIntelReportPDFReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetIntelReportPDFReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntelReportPDFOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntelReportPDFBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntelReportPDFForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntelReportPDFTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntelReportPDFInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /intel/entities/report-files/v1] GetIntelReportPDF", response, response.Code())
	}
}

// NewGetIntelReportPDFOK creates a GetIntelReportPDFOK with default headers values
func NewGetIntelReportPDFOK(writer io.Writer) *GetIntelReportPDFOK {
	return &GetIntelReportPDFOK{

		Payload: writer,
	}
}

/*
GetIntelReportPDFOK describes a response with status code 200, with default header values.

OK
*/
type GetIntelReportPDFOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload io.Writer
}

// IsSuccess returns true when this get intel report p d f o k response has a 2xx status code
func (o *GetIntelReportPDFOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get intel report p d f o k response has a 3xx status code
func (o *GetIntelReportPDFOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel report p d f o k response has a 4xx status code
func (o *GetIntelReportPDFOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get intel report p d f o k response has a 5xx status code
func (o *GetIntelReportPDFOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel report p d f o k response a status code equal to that given
func (o *GetIntelReportPDFOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get intel report p d f o k response
func (o *GetIntelReportPDFOK) Code() int {
	return 200
}

func (o *GetIntelReportPDFOK) Error() string {
	return fmt.Sprintf("[GET /intel/entities/report-files/v1][%d] getIntelReportPDFOK  %+v", 200, o.Payload)
}

func (o *GetIntelReportPDFOK) String() string {
	return fmt.Sprintf("[GET /intel/entities/report-files/v1][%d] getIntelReportPDFOK  %+v", 200, o.Payload)
}

func (o *GetIntelReportPDFOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetIntelReportPDFOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntelReportPDFBadRequest creates a GetIntelReportPDFBadRequest with default headers values
func NewGetIntelReportPDFBadRequest() *GetIntelReportPDFBadRequest {
	return &GetIntelReportPDFBadRequest{}
}

/*
GetIntelReportPDFBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetIntelReportPDFBadRequest struct {

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

// IsSuccess returns true when this get intel report p d f bad request response has a 2xx status code
func (o *GetIntelReportPDFBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel report p d f bad request response has a 3xx status code
func (o *GetIntelReportPDFBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel report p d f bad request response has a 4xx status code
func (o *GetIntelReportPDFBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get intel report p d f bad request response has a 5xx status code
func (o *GetIntelReportPDFBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel report p d f bad request response a status code equal to that given
func (o *GetIntelReportPDFBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get intel report p d f bad request response
func (o *GetIntelReportPDFBadRequest) Code() int {
	return 400
}

func (o *GetIntelReportPDFBadRequest) Error() string {
	return fmt.Sprintf("[GET /intel/entities/report-files/v1][%d] getIntelReportPDFBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntelReportPDFBadRequest) String() string {
	return fmt.Sprintf("[GET /intel/entities/report-files/v1][%d] getIntelReportPDFBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntelReportPDFBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetIntelReportPDFBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelReportPDFForbidden creates a GetIntelReportPDFForbidden with default headers values
func NewGetIntelReportPDFForbidden() *GetIntelReportPDFForbidden {
	return &GetIntelReportPDFForbidden{}
}

/*
GetIntelReportPDFForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetIntelReportPDFForbidden struct {

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

// IsSuccess returns true when this get intel report p d f forbidden response has a 2xx status code
func (o *GetIntelReportPDFForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel report p d f forbidden response has a 3xx status code
func (o *GetIntelReportPDFForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel report p d f forbidden response has a 4xx status code
func (o *GetIntelReportPDFForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get intel report p d f forbidden response has a 5xx status code
func (o *GetIntelReportPDFForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel report p d f forbidden response a status code equal to that given
func (o *GetIntelReportPDFForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get intel report p d f forbidden response
func (o *GetIntelReportPDFForbidden) Code() int {
	return 403
}

func (o *GetIntelReportPDFForbidden) Error() string {
	return fmt.Sprintf("[GET /intel/entities/report-files/v1][%d] getIntelReportPDFForbidden  %+v", 403, o.Payload)
}

func (o *GetIntelReportPDFForbidden) String() string {
	return fmt.Sprintf("[GET /intel/entities/report-files/v1][%d] getIntelReportPDFForbidden  %+v", 403, o.Payload)
}

func (o *GetIntelReportPDFForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIntelReportPDFForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelReportPDFTooManyRequests creates a GetIntelReportPDFTooManyRequests with default headers values
func NewGetIntelReportPDFTooManyRequests() *GetIntelReportPDFTooManyRequests {
	return &GetIntelReportPDFTooManyRequests{}
}

/*
GetIntelReportPDFTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetIntelReportPDFTooManyRequests struct {

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

// IsSuccess returns true when this get intel report p d f too many requests response has a 2xx status code
func (o *GetIntelReportPDFTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel report p d f too many requests response has a 3xx status code
func (o *GetIntelReportPDFTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel report p d f too many requests response has a 4xx status code
func (o *GetIntelReportPDFTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get intel report p d f too many requests response has a 5xx status code
func (o *GetIntelReportPDFTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel report p d f too many requests response a status code equal to that given
func (o *GetIntelReportPDFTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get intel report p d f too many requests response
func (o *GetIntelReportPDFTooManyRequests) Code() int {
	return 429
}

func (o *GetIntelReportPDFTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /intel/entities/report-files/v1][%d] getIntelReportPDFTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntelReportPDFTooManyRequests) String() string {
	return fmt.Sprintf("[GET /intel/entities/report-files/v1][%d] getIntelReportPDFTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntelReportPDFTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIntelReportPDFTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelReportPDFInternalServerError creates a GetIntelReportPDFInternalServerError with default headers values
func NewGetIntelReportPDFInternalServerError() *GetIntelReportPDFInternalServerError {
	return &GetIntelReportPDFInternalServerError{}
}

/*
GetIntelReportPDFInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetIntelReportPDFInternalServerError struct {

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

// IsSuccess returns true when this get intel report p d f internal server error response has a 2xx status code
func (o *GetIntelReportPDFInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel report p d f internal server error response has a 3xx status code
func (o *GetIntelReportPDFInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel report p d f internal server error response has a 4xx status code
func (o *GetIntelReportPDFInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get intel report p d f internal server error response has a 5xx status code
func (o *GetIntelReportPDFInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get intel report p d f internal server error response a status code equal to that given
func (o *GetIntelReportPDFInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get intel report p d f internal server error response
func (o *GetIntelReportPDFInternalServerError) Code() int {
	return 500
}

func (o *GetIntelReportPDFInternalServerError) Error() string {
	return fmt.Sprintf("[GET /intel/entities/report-files/v1][%d] getIntelReportPDFInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntelReportPDFInternalServerError) String() string {
	return fmt.Sprintf("[GET /intel/entities/report-files/v1][%d] getIntelReportPDFInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntelReportPDFInternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetIntelReportPDFInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
