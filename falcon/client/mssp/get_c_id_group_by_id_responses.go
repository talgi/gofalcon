// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// GetCIDGroupByIDReader is a Reader for the GetCIDGroupByID structure.
type GetCIDGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCIDGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCIDGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetCIDGroupByIDMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCIDGroupByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCIDGroupByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCIDGroupByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCIDGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCIDGroupByIDOK creates a GetCIDGroupByIDOK with default headers values
func NewGetCIDGroupByIDOK() *GetCIDGroupByIDOK {
	return &GetCIDGroupByIDOK{}
}

/*
GetCIDGroupByIDOK describes a response with status code 200, with default header values.

OK
*/
type GetCIDGroupByIDOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupsResponseV1
}

// IsSuccess returns true when this get c Id group by Id o k response has a 2xx status code
func (o *GetCIDGroupByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c Id group by Id o k response has a 3xx status code
func (o *GetCIDGroupByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group by Id o k response has a 4xx status code
func (o *GetCIDGroupByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c Id group by Id o k response has a 5xx status code
func (o *GetCIDGroupByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group by Id o k response a status code equal to that given
func (o *GetCIDGroupByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCIDGroupByIDOK) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIdGroupByIdOK  %+v", 200, o.Payload)
}

func (o *GetCIDGroupByIDOK) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIdGroupByIdOK  %+v", 200, o.Payload)
}

func (o *GetCIDGroupByIDOK) GetPayload() *models.DomainCIDGroupsResponseV1 {
	return o.Payload
}

func (o *GetCIDGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCIDGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIDGroupByIDMultiStatus creates a GetCIDGroupByIDMultiStatus with default headers values
func NewGetCIDGroupByIDMultiStatus() *GetCIDGroupByIDMultiStatus {
	return &GetCIDGroupByIDMultiStatus{}
}

/*
GetCIDGroupByIDMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetCIDGroupByIDMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupsResponseV1
}

// IsSuccess returns true when this get c Id group by Id multi status response has a 2xx status code
func (o *GetCIDGroupByIDMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c Id group by Id multi status response has a 3xx status code
func (o *GetCIDGroupByIDMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group by Id multi status response has a 4xx status code
func (o *GetCIDGroupByIDMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c Id group by Id multi status response has a 5xx status code
func (o *GetCIDGroupByIDMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group by Id multi status response a status code equal to that given
func (o *GetCIDGroupByIDMultiStatus) IsCode(code int) bool {
	return code == 207
}

func (o *GetCIDGroupByIDMultiStatus) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIdGroupByIdMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCIDGroupByIDMultiStatus) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIdGroupByIdMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCIDGroupByIDMultiStatus) GetPayload() *models.DomainCIDGroupsResponseV1 {
	return o.Payload
}

func (o *GetCIDGroupByIDMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCIDGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIDGroupByIDBadRequest creates a GetCIDGroupByIDBadRequest with default headers values
func NewGetCIDGroupByIDBadRequest() *GetCIDGroupByIDBadRequest {
	return &GetCIDGroupByIDBadRequest{}
}

/*
GetCIDGroupByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCIDGroupByIDBadRequest struct {

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

// IsSuccess returns true when this get c Id group by Id bad request response has a 2xx status code
func (o *GetCIDGroupByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c Id group by Id bad request response has a 3xx status code
func (o *GetCIDGroupByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group by Id bad request response has a 4xx status code
func (o *GetCIDGroupByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c Id group by Id bad request response has a 5xx status code
func (o *GetCIDGroupByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group by Id bad request response a status code equal to that given
func (o *GetCIDGroupByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCIDGroupByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIdGroupByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetCIDGroupByIDBadRequest) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIdGroupByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetCIDGroupByIDBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetCIDGroupByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCIDGroupByIDForbidden creates a GetCIDGroupByIDForbidden with default headers values
func NewGetCIDGroupByIDForbidden() *GetCIDGroupByIDForbidden {
	return &GetCIDGroupByIDForbidden{}
}

/*
GetCIDGroupByIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCIDGroupByIDForbidden struct {

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

// IsSuccess returns true when this get c Id group by Id forbidden response has a 2xx status code
func (o *GetCIDGroupByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c Id group by Id forbidden response has a 3xx status code
func (o *GetCIDGroupByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group by Id forbidden response has a 4xx status code
func (o *GetCIDGroupByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c Id group by Id forbidden response has a 5xx status code
func (o *GetCIDGroupByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group by Id forbidden response a status code equal to that given
func (o *GetCIDGroupByIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCIDGroupByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIdGroupByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetCIDGroupByIDForbidden) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIdGroupByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetCIDGroupByIDForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetCIDGroupByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCIDGroupByIDTooManyRequests creates a GetCIDGroupByIDTooManyRequests with default headers values
func NewGetCIDGroupByIDTooManyRequests() *GetCIDGroupByIDTooManyRequests {
	return &GetCIDGroupByIDTooManyRequests{}
}

/*
GetCIDGroupByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCIDGroupByIDTooManyRequests struct {

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

// IsSuccess returns true when this get c Id group by Id too many requests response has a 2xx status code
func (o *GetCIDGroupByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c Id group by Id too many requests response has a 3xx status code
func (o *GetCIDGroupByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group by Id too many requests response has a 4xx status code
func (o *GetCIDGroupByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c Id group by Id too many requests response has a 5xx status code
func (o *GetCIDGroupByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group by Id too many requests response a status code equal to that given
func (o *GetCIDGroupByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCIDGroupByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIdGroupByIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCIDGroupByIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIdGroupByIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCIDGroupByIDTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCIDGroupByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCIDGroupByIDDefault creates a GetCIDGroupByIDDefault with default headers values
func NewGetCIDGroupByIDDefault(code int) *GetCIDGroupByIDDefault {
	return &GetCIDGroupByIDDefault{
		_statusCode: code,
	}
}

/*
GetCIDGroupByIDDefault describes a response with status code -1, with default header values.

OK
*/
type GetCIDGroupByIDDefault struct {
	_statusCode int

	Payload *models.DomainCIDGroupsResponseV1
}

// Code gets the status code for the get c ID group by Id default response
func (o *GetCIDGroupByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get c ID group by Id default response has a 2xx status code
func (o *GetCIDGroupByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get c ID group by Id default response has a 3xx status code
func (o *GetCIDGroupByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get c ID group by Id default response has a 4xx status code
func (o *GetCIDGroupByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get c ID group by Id default response has a 5xx status code
func (o *GetCIDGroupByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get c ID group by Id default response a status code equal to that given
func (o *GetCIDGroupByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetCIDGroupByIDDefault) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIDGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *GetCIDGroupByIDDefault) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-groups/v1][%d] getCIDGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *GetCIDGroupByIDDefault) GetPayload() *models.DomainCIDGroupsResponseV1 {
	return o.Payload
}

func (o *GetCIDGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainCIDGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
