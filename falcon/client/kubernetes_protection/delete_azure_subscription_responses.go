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

// DeleteAzureSubscriptionReader is a Reader for the DeleteAzureSubscription structure.
type DeleteAzureSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAzureSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAzureSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewDeleteAzureSubscriptionMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAzureSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAzureSubscriptionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAzureSubscriptionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAzureSubscriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /kubernetes-protection/entities/accounts/azure/v1] DeleteAzureSubscription", response, response.Code())
	}
}

// NewDeleteAzureSubscriptionOK creates a DeleteAzureSubscriptionOK with default headers values
func NewDeleteAzureSubscriptionOK() *DeleteAzureSubscriptionOK {
	return &DeleteAzureSubscriptionOK{}
}

/*
DeleteAzureSubscriptionOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAzureSubscriptionOK struct {

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

// IsSuccess returns true when this delete azure subscription o k response has a 2xx status code
func (o *DeleteAzureSubscriptionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete azure subscription o k response has a 3xx status code
func (o *DeleteAzureSubscriptionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete azure subscription o k response has a 4xx status code
func (o *DeleteAzureSubscriptionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete azure subscription o k response has a 5xx status code
func (o *DeleteAzureSubscriptionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete azure subscription o k response a status code equal to that given
func (o *DeleteAzureSubscriptionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete azure subscription o k response
func (o *DeleteAzureSubscriptionOK) Code() int {
	return 200
}

func (o *DeleteAzureSubscriptionOK) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionOK  %+v", 200, o.Payload)
}

func (o *DeleteAzureSubscriptionOK) String() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionOK  %+v", 200, o.Payload)
}

func (o *DeleteAzureSubscriptionOK) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteAzureSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAzureSubscriptionMultiStatus creates a DeleteAzureSubscriptionMultiStatus with default headers values
func NewDeleteAzureSubscriptionMultiStatus() *DeleteAzureSubscriptionMultiStatus {
	return &DeleteAzureSubscriptionMultiStatus{}
}

/*
DeleteAzureSubscriptionMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type DeleteAzureSubscriptionMultiStatus struct {

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

// IsSuccess returns true when this delete azure subscription multi status response has a 2xx status code
func (o *DeleteAzureSubscriptionMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete azure subscription multi status response has a 3xx status code
func (o *DeleteAzureSubscriptionMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete azure subscription multi status response has a 4xx status code
func (o *DeleteAzureSubscriptionMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete azure subscription multi status response has a 5xx status code
func (o *DeleteAzureSubscriptionMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this delete azure subscription multi status response a status code equal to that given
func (o *DeleteAzureSubscriptionMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the delete azure subscription multi status response
func (o *DeleteAzureSubscriptionMultiStatus) Code() int {
	return 207
}

func (o *DeleteAzureSubscriptionMultiStatus) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteAzureSubscriptionMultiStatus) String() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteAzureSubscriptionMultiStatus) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteAzureSubscriptionMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAzureSubscriptionBadRequest creates a DeleteAzureSubscriptionBadRequest with default headers values
func NewDeleteAzureSubscriptionBadRequest() *DeleteAzureSubscriptionBadRequest {
	return &DeleteAzureSubscriptionBadRequest{}
}

/*
DeleteAzureSubscriptionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAzureSubscriptionBadRequest struct {

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

// IsSuccess returns true when this delete azure subscription bad request response has a 2xx status code
func (o *DeleteAzureSubscriptionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete azure subscription bad request response has a 3xx status code
func (o *DeleteAzureSubscriptionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete azure subscription bad request response has a 4xx status code
func (o *DeleteAzureSubscriptionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete azure subscription bad request response has a 5xx status code
func (o *DeleteAzureSubscriptionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete azure subscription bad request response a status code equal to that given
func (o *DeleteAzureSubscriptionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete azure subscription bad request response
func (o *DeleteAzureSubscriptionBadRequest) Code() int {
	return 400
}

func (o *DeleteAzureSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAzureSubscriptionBadRequest) String() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAzureSubscriptionBadRequest) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteAzureSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAzureSubscriptionForbidden creates a DeleteAzureSubscriptionForbidden with default headers values
func NewDeleteAzureSubscriptionForbidden() *DeleteAzureSubscriptionForbidden {
	return &DeleteAzureSubscriptionForbidden{}
}

/*
DeleteAzureSubscriptionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAzureSubscriptionForbidden struct {

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

// IsSuccess returns true when this delete azure subscription forbidden response has a 2xx status code
func (o *DeleteAzureSubscriptionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete azure subscription forbidden response has a 3xx status code
func (o *DeleteAzureSubscriptionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete azure subscription forbidden response has a 4xx status code
func (o *DeleteAzureSubscriptionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete azure subscription forbidden response has a 5xx status code
func (o *DeleteAzureSubscriptionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete azure subscription forbidden response a status code equal to that given
func (o *DeleteAzureSubscriptionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete azure subscription forbidden response
func (o *DeleteAzureSubscriptionForbidden) Code() int {
	return 403
}

func (o *DeleteAzureSubscriptionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAzureSubscriptionForbidden) String() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAzureSubscriptionForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteAzureSubscriptionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAzureSubscriptionTooManyRequests creates a DeleteAzureSubscriptionTooManyRequests with default headers values
func NewDeleteAzureSubscriptionTooManyRequests() *DeleteAzureSubscriptionTooManyRequests {
	return &DeleteAzureSubscriptionTooManyRequests{}
}

/*
DeleteAzureSubscriptionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteAzureSubscriptionTooManyRequests struct {

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

// IsSuccess returns true when this delete azure subscription too many requests response has a 2xx status code
func (o *DeleteAzureSubscriptionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete azure subscription too many requests response has a 3xx status code
func (o *DeleteAzureSubscriptionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete azure subscription too many requests response has a 4xx status code
func (o *DeleteAzureSubscriptionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete azure subscription too many requests response has a 5xx status code
func (o *DeleteAzureSubscriptionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete azure subscription too many requests response a status code equal to that given
func (o *DeleteAzureSubscriptionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete azure subscription too many requests response
func (o *DeleteAzureSubscriptionTooManyRequests) Code() int {
	return 429
}

func (o *DeleteAzureSubscriptionTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteAzureSubscriptionTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteAzureSubscriptionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteAzureSubscriptionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAzureSubscriptionInternalServerError creates a DeleteAzureSubscriptionInternalServerError with default headers values
func NewDeleteAzureSubscriptionInternalServerError() *DeleteAzureSubscriptionInternalServerError {
	return &DeleteAzureSubscriptionInternalServerError{}
}

/*
DeleteAzureSubscriptionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAzureSubscriptionInternalServerError struct {

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

// IsSuccess returns true when this delete azure subscription internal server error response has a 2xx status code
func (o *DeleteAzureSubscriptionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete azure subscription internal server error response has a 3xx status code
func (o *DeleteAzureSubscriptionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete azure subscription internal server error response has a 4xx status code
func (o *DeleteAzureSubscriptionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete azure subscription internal server error response has a 5xx status code
func (o *DeleteAzureSubscriptionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete azure subscription internal server error response a status code equal to that given
func (o *DeleteAzureSubscriptionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete azure subscription internal server error response
func (o *DeleteAzureSubscriptionInternalServerError) Code() int {
	return 500
}

func (o *DeleteAzureSubscriptionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAzureSubscriptionInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/azure/v1][%d] deleteAzureSubscriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAzureSubscriptionInternalServerError) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteAzureSubscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
