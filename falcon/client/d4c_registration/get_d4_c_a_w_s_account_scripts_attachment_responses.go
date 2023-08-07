// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// GetD4CAWSAccountScriptsAttachmentReader is a Reader for the GetD4CAWSAccountScriptsAttachment structure.
type GetD4CAWSAccountScriptsAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetD4CAWSAccountScriptsAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetD4CAWSAccountScriptsAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetD4CAWSAccountScriptsAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetD4CAWSAccountScriptsAttachmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetD4CAWSAccountScriptsAttachmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetD4CAWSAccountScriptsAttachmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud-connect-aws/entities/user-scripts-download/v1] GetD4CAWSAccountScriptsAttachment", response, response.Code())
	}
}

// NewGetD4CAWSAccountScriptsAttachmentOK creates a GetD4CAWSAccountScriptsAttachmentOK with default headers values
func NewGetD4CAWSAccountScriptsAttachmentOK() *GetD4CAWSAccountScriptsAttachmentOK {
	return &GetD4CAWSAccountScriptsAttachmentOK{}
}

/*
GetD4CAWSAccountScriptsAttachmentOK describes a response with status code 200, with default header values.

OK
*/
type GetD4CAWSAccountScriptsAttachmentOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSProvisionGetAccountScriptResponseV2
}

// IsSuccess returns true when this get d4 c a w s account scripts attachment o k response has a 2xx status code
func (o *GetD4CAWSAccountScriptsAttachmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get d4 c a w s account scripts attachment o k response has a 3xx status code
func (o *GetD4CAWSAccountScriptsAttachmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c a w s account scripts attachment o k response has a 4xx status code
func (o *GetD4CAWSAccountScriptsAttachmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c a w s account scripts attachment o k response has a 5xx status code
func (o *GetD4CAWSAccountScriptsAttachmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c a w s account scripts attachment o k response a status code equal to that given
func (o *GetD4CAWSAccountScriptsAttachmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get d4 c a w s account scripts attachment o k response
func (o *GetD4CAWSAccountScriptsAttachmentOK) Code() int {
	return 200
}

func (o *GetD4CAWSAccountScriptsAttachmentOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/user-scripts-download/v1][%d] getD4CAWSAccountScriptsAttachmentOK  %+v", 200, o.Payload)
}

func (o *GetD4CAWSAccountScriptsAttachmentOK) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/user-scripts-download/v1][%d] getD4CAWSAccountScriptsAttachmentOK  %+v", 200, o.Payload)
}

func (o *GetD4CAWSAccountScriptsAttachmentOK) GetPayload() *models.RegistrationAWSProvisionGetAccountScriptResponseV2 {
	return o.Payload
}

func (o *GetD4CAWSAccountScriptsAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSProvisionGetAccountScriptResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CAWSAccountScriptsAttachmentBadRequest creates a GetD4CAWSAccountScriptsAttachmentBadRequest with default headers values
func NewGetD4CAWSAccountScriptsAttachmentBadRequest() *GetD4CAWSAccountScriptsAttachmentBadRequest {
	return &GetD4CAWSAccountScriptsAttachmentBadRequest{}
}

/*
GetD4CAWSAccountScriptsAttachmentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetD4CAWSAccountScriptsAttachmentBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSProvisionGetAccountScriptResponseV2
}

// IsSuccess returns true when this get d4 c a w s account scripts attachment bad request response has a 2xx status code
func (o *GetD4CAWSAccountScriptsAttachmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c a w s account scripts attachment bad request response has a 3xx status code
func (o *GetD4CAWSAccountScriptsAttachmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c a w s account scripts attachment bad request response has a 4xx status code
func (o *GetD4CAWSAccountScriptsAttachmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c a w s account scripts attachment bad request response has a 5xx status code
func (o *GetD4CAWSAccountScriptsAttachmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c a w s account scripts attachment bad request response a status code equal to that given
func (o *GetD4CAWSAccountScriptsAttachmentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get d4 c a w s account scripts attachment bad request response
func (o *GetD4CAWSAccountScriptsAttachmentBadRequest) Code() int {
	return 400
}

func (o *GetD4CAWSAccountScriptsAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/user-scripts-download/v1][%d] getD4CAWSAccountScriptsAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetD4CAWSAccountScriptsAttachmentBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/user-scripts-download/v1][%d] getD4CAWSAccountScriptsAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetD4CAWSAccountScriptsAttachmentBadRequest) GetPayload() *models.RegistrationAWSProvisionGetAccountScriptResponseV2 {
	return o.Payload
}

func (o *GetD4CAWSAccountScriptsAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSProvisionGetAccountScriptResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CAWSAccountScriptsAttachmentForbidden creates a GetD4CAWSAccountScriptsAttachmentForbidden with default headers values
func NewGetD4CAWSAccountScriptsAttachmentForbidden() *GetD4CAWSAccountScriptsAttachmentForbidden {
	return &GetD4CAWSAccountScriptsAttachmentForbidden{}
}

/*
GetD4CAWSAccountScriptsAttachmentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetD4CAWSAccountScriptsAttachmentForbidden struct {

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

// IsSuccess returns true when this get d4 c a w s account scripts attachment forbidden response has a 2xx status code
func (o *GetD4CAWSAccountScriptsAttachmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c a w s account scripts attachment forbidden response has a 3xx status code
func (o *GetD4CAWSAccountScriptsAttachmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c a w s account scripts attachment forbidden response has a 4xx status code
func (o *GetD4CAWSAccountScriptsAttachmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c a w s account scripts attachment forbidden response has a 5xx status code
func (o *GetD4CAWSAccountScriptsAttachmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c a w s account scripts attachment forbidden response a status code equal to that given
func (o *GetD4CAWSAccountScriptsAttachmentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get d4 c a w s account scripts attachment forbidden response
func (o *GetD4CAWSAccountScriptsAttachmentForbidden) Code() int {
	return 403
}

func (o *GetD4CAWSAccountScriptsAttachmentForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/user-scripts-download/v1][%d] getD4CAWSAccountScriptsAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *GetD4CAWSAccountScriptsAttachmentForbidden) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/user-scripts-download/v1][%d] getD4CAWSAccountScriptsAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *GetD4CAWSAccountScriptsAttachmentForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetD4CAWSAccountScriptsAttachmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetD4CAWSAccountScriptsAttachmentTooManyRequests creates a GetD4CAWSAccountScriptsAttachmentTooManyRequests with default headers values
func NewGetD4CAWSAccountScriptsAttachmentTooManyRequests() *GetD4CAWSAccountScriptsAttachmentTooManyRequests {
	return &GetD4CAWSAccountScriptsAttachmentTooManyRequests{}
}

/*
GetD4CAWSAccountScriptsAttachmentTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetD4CAWSAccountScriptsAttachmentTooManyRequests struct {

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

// IsSuccess returns true when this get d4 c a w s account scripts attachment too many requests response has a 2xx status code
func (o *GetD4CAWSAccountScriptsAttachmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c a w s account scripts attachment too many requests response has a 3xx status code
func (o *GetD4CAWSAccountScriptsAttachmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c a w s account scripts attachment too many requests response has a 4xx status code
func (o *GetD4CAWSAccountScriptsAttachmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c a w s account scripts attachment too many requests response has a 5xx status code
func (o *GetD4CAWSAccountScriptsAttachmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c a w s account scripts attachment too many requests response a status code equal to that given
func (o *GetD4CAWSAccountScriptsAttachmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get d4 c a w s account scripts attachment too many requests response
func (o *GetD4CAWSAccountScriptsAttachmentTooManyRequests) Code() int {
	return 429
}

func (o *GetD4CAWSAccountScriptsAttachmentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/user-scripts-download/v1][%d] getD4CAWSAccountScriptsAttachmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetD4CAWSAccountScriptsAttachmentTooManyRequests) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/user-scripts-download/v1][%d] getD4CAWSAccountScriptsAttachmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetD4CAWSAccountScriptsAttachmentTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetD4CAWSAccountScriptsAttachmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetD4CAWSAccountScriptsAttachmentInternalServerError creates a GetD4CAWSAccountScriptsAttachmentInternalServerError with default headers values
func NewGetD4CAWSAccountScriptsAttachmentInternalServerError() *GetD4CAWSAccountScriptsAttachmentInternalServerError {
	return &GetD4CAWSAccountScriptsAttachmentInternalServerError{}
}

/*
GetD4CAWSAccountScriptsAttachmentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetD4CAWSAccountScriptsAttachmentInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSProvisionGetAccountScriptResponseV2
}

// IsSuccess returns true when this get d4 c a w s account scripts attachment internal server error response has a 2xx status code
func (o *GetD4CAWSAccountScriptsAttachmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c a w s account scripts attachment internal server error response has a 3xx status code
func (o *GetD4CAWSAccountScriptsAttachmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c a w s account scripts attachment internal server error response has a 4xx status code
func (o *GetD4CAWSAccountScriptsAttachmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c a w s account scripts attachment internal server error response has a 5xx status code
func (o *GetD4CAWSAccountScriptsAttachmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get d4 c a w s account scripts attachment internal server error response a status code equal to that given
func (o *GetD4CAWSAccountScriptsAttachmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get d4 c a w s account scripts attachment internal server error response
func (o *GetD4CAWSAccountScriptsAttachmentInternalServerError) Code() int {
	return 500
}

func (o *GetD4CAWSAccountScriptsAttachmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/user-scripts-download/v1][%d] getD4CAWSAccountScriptsAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetD4CAWSAccountScriptsAttachmentInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/user-scripts-download/v1][%d] getD4CAWSAccountScriptsAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetD4CAWSAccountScriptsAttachmentInternalServerError) GetPayload() *models.RegistrationAWSProvisionGetAccountScriptResponseV2 {
	return o.Payload
}

func (o *GetD4CAWSAccountScriptsAttachmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSProvisionGetAccountScriptResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
