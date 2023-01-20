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

// DeleteD4CAwsAccountReader is a Reader for the DeleteD4CAwsAccount structure.
type DeleteD4CAwsAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteD4CAwsAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteD4CAwsAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewDeleteD4CAwsAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteD4CAwsAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteD4CAwsAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteD4CAwsAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteD4CAwsAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteD4CAwsAccountOK creates a DeleteD4CAwsAccountOK with default headers values
func NewDeleteD4CAwsAccountOK() *DeleteD4CAwsAccountOK {
	return &DeleteD4CAwsAccountOK{}
}

/*
DeleteD4CAwsAccountOK describes a response with status code 200, with default header values.

OK
*/
type DeleteD4CAwsAccountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete d4 c aws account o k response has a 2xx status code
func (o *DeleteD4CAwsAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete d4 c aws account o k response has a 3xx status code
func (o *DeleteD4CAwsAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete d4 c aws account o k response has a 4xx status code
func (o *DeleteD4CAwsAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete d4 c aws account o k response has a 5xx status code
func (o *DeleteD4CAwsAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete d4 c aws account o k response a status code equal to that given
func (o *DeleteD4CAwsAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete d4 c aws account o k response
func (o *DeleteD4CAwsAccountOK) Code() int {
	return 200
}

func (o *DeleteD4CAwsAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteD4CAwsAccountOK) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteD4CAwsAccountOK) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteD4CAwsAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteD4CAwsAccountMultiStatus creates a DeleteD4CAwsAccountMultiStatus with default headers values
func NewDeleteD4CAwsAccountMultiStatus() *DeleteD4CAwsAccountMultiStatus {
	return &DeleteD4CAwsAccountMultiStatus{}
}

/*
DeleteD4CAwsAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type DeleteD4CAwsAccountMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete d4 c aws account multi status response has a 2xx status code
func (o *DeleteD4CAwsAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete d4 c aws account multi status response has a 3xx status code
func (o *DeleteD4CAwsAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete d4 c aws account multi status response has a 4xx status code
func (o *DeleteD4CAwsAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete d4 c aws account multi status response has a 5xx status code
func (o *DeleteD4CAwsAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this delete d4 c aws account multi status response a status code equal to that given
func (o *DeleteD4CAwsAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the delete d4 c aws account multi status response
func (o *DeleteD4CAwsAccountMultiStatus) Code() int {
	return 207
}

func (o *DeleteD4CAwsAccountMultiStatus) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteD4CAwsAccountMultiStatus) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteD4CAwsAccountMultiStatus) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteD4CAwsAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteD4CAwsAccountBadRequest creates a DeleteD4CAwsAccountBadRequest with default headers values
func NewDeleteD4CAwsAccountBadRequest() *DeleteD4CAwsAccountBadRequest {
	return &DeleteD4CAwsAccountBadRequest{}
}

/*
DeleteD4CAwsAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteD4CAwsAccountBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete d4 c aws account bad request response has a 2xx status code
func (o *DeleteD4CAwsAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete d4 c aws account bad request response has a 3xx status code
func (o *DeleteD4CAwsAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete d4 c aws account bad request response has a 4xx status code
func (o *DeleteD4CAwsAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete d4 c aws account bad request response has a 5xx status code
func (o *DeleteD4CAwsAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete d4 c aws account bad request response a status code equal to that given
func (o *DeleteD4CAwsAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete d4 c aws account bad request response
func (o *DeleteD4CAwsAccountBadRequest) Code() int {
	return 400
}

func (o *DeleteD4CAwsAccountBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteD4CAwsAccountBadRequest) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteD4CAwsAccountBadRequest) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteD4CAwsAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteD4CAwsAccountForbidden creates a DeleteD4CAwsAccountForbidden with default headers values
func NewDeleteD4CAwsAccountForbidden() *DeleteD4CAwsAccountForbidden {
	return &DeleteD4CAwsAccountForbidden{}
}

/*
DeleteD4CAwsAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteD4CAwsAccountForbidden struct {

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

// IsSuccess returns true when this delete d4 c aws account forbidden response has a 2xx status code
func (o *DeleteD4CAwsAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete d4 c aws account forbidden response has a 3xx status code
func (o *DeleteD4CAwsAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete d4 c aws account forbidden response has a 4xx status code
func (o *DeleteD4CAwsAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete d4 c aws account forbidden response has a 5xx status code
func (o *DeleteD4CAwsAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete d4 c aws account forbidden response a status code equal to that given
func (o *DeleteD4CAwsAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete d4 c aws account forbidden response
func (o *DeleteD4CAwsAccountForbidden) Code() int {
	return 403
}

func (o *DeleteD4CAwsAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteD4CAwsAccountForbidden) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteD4CAwsAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteD4CAwsAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteD4CAwsAccountTooManyRequests creates a DeleteD4CAwsAccountTooManyRequests with default headers values
func NewDeleteD4CAwsAccountTooManyRequests() *DeleteD4CAwsAccountTooManyRequests {
	return &DeleteD4CAwsAccountTooManyRequests{}
}

/*
DeleteD4CAwsAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteD4CAwsAccountTooManyRequests struct {

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

// IsSuccess returns true when this delete d4 c aws account too many requests response has a 2xx status code
func (o *DeleteD4CAwsAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete d4 c aws account too many requests response has a 3xx status code
func (o *DeleteD4CAwsAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete d4 c aws account too many requests response has a 4xx status code
func (o *DeleteD4CAwsAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete d4 c aws account too many requests response has a 5xx status code
func (o *DeleteD4CAwsAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete d4 c aws account too many requests response a status code equal to that given
func (o *DeleteD4CAwsAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete d4 c aws account too many requests response
func (o *DeleteD4CAwsAccountTooManyRequests) Code() int {
	return 429
}

func (o *DeleteD4CAwsAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteD4CAwsAccountTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteD4CAwsAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteD4CAwsAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteD4CAwsAccountInternalServerError creates a DeleteD4CAwsAccountInternalServerError with default headers values
func NewDeleteD4CAwsAccountInternalServerError() *DeleteD4CAwsAccountInternalServerError {
	return &DeleteD4CAwsAccountInternalServerError{}
}

/*
DeleteD4CAwsAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteD4CAwsAccountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete d4 c aws account internal server error response has a 2xx status code
func (o *DeleteD4CAwsAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete d4 c aws account internal server error response has a 3xx status code
func (o *DeleteD4CAwsAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete d4 c aws account internal server error response has a 4xx status code
func (o *DeleteD4CAwsAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete d4 c aws account internal server error response has a 5xx status code
func (o *DeleteD4CAwsAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete d4 c aws account internal server error response a status code equal to that given
func (o *DeleteD4CAwsAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete d4 c aws account internal server error response
func (o *DeleteD4CAwsAccountInternalServerError) Code() int {
	return 500
}

func (o *DeleteD4CAwsAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteD4CAwsAccountInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/account/v2][%d] deleteD4CAwsAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteD4CAwsAccountInternalServerError) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteD4CAwsAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
