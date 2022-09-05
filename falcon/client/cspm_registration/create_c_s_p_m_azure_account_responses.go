// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// CreateCSPMAzureAccountReader is a Reader for the CreateCSPMAzureAccount structure.
type CreateCSPMAzureAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCSPMAzureAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCSPMAzureAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewCreateCSPMAzureAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCSPMAzureAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCSPMAzureAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateCSPMAzureAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateCSPMAzureAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCSPMAzureAccountCreated creates a CreateCSPMAzureAccountCreated with default headers values
func NewCreateCSPMAzureAccountCreated() *CreateCSPMAzureAccountCreated {
	return &CreateCSPMAzureAccountCreated{}
}

/*
CreateCSPMAzureAccountCreated describes a response with status code 201, with default header values.

Created
*/
type CreateCSPMAzureAccountCreated struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureAccountResponseV1
}

// IsSuccess returns true when this create c s p m azure account created response has a 2xx status code
func (o *CreateCSPMAzureAccountCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create c s p m azure account created response has a 3xx status code
func (o *CreateCSPMAzureAccountCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c s p m azure account created response has a 4xx status code
func (o *CreateCSPMAzureAccountCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create c s p m azure account created response has a 5xx status code
func (o *CreateCSPMAzureAccountCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create c s p m azure account created response a status code equal to that given
func (o *CreateCSPMAzureAccountCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateCSPMAzureAccountCreated) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateCSPMAzureAccountCreated) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateCSPMAzureAccountCreated) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *CreateCSPMAzureAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCSPMAzureAccountMultiStatus creates a CreateCSPMAzureAccountMultiStatus with default headers values
func NewCreateCSPMAzureAccountMultiStatus() *CreateCSPMAzureAccountMultiStatus {
	return &CreateCSPMAzureAccountMultiStatus{}
}

/*
CreateCSPMAzureAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type CreateCSPMAzureAccountMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureAccountResponseV1
}

// IsSuccess returns true when this create c s p m azure account multi status response has a 2xx status code
func (o *CreateCSPMAzureAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create c s p m azure account multi status response has a 3xx status code
func (o *CreateCSPMAzureAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c s p m azure account multi status response has a 4xx status code
func (o *CreateCSPMAzureAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this create c s p m azure account multi status response has a 5xx status code
func (o *CreateCSPMAzureAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this create c s p m azure account multi status response a status code equal to that given
func (o *CreateCSPMAzureAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

func (o *CreateCSPMAzureAccountMultiStatus) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateCSPMAzureAccountMultiStatus) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateCSPMAzureAccountMultiStatus) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *CreateCSPMAzureAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCSPMAzureAccountBadRequest creates a CreateCSPMAzureAccountBadRequest with default headers values
func NewCreateCSPMAzureAccountBadRequest() *CreateCSPMAzureAccountBadRequest {
	return &CreateCSPMAzureAccountBadRequest{}
}

/*
CreateCSPMAzureAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateCSPMAzureAccountBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureAccountResponseV1
}

// IsSuccess returns true when this create c s p m azure account bad request response has a 2xx status code
func (o *CreateCSPMAzureAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create c s p m azure account bad request response has a 3xx status code
func (o *CreateCSPMAzureAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c s p m azure account bad request response has a 4xx status code
func (o *CreateCSPMAzureAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create c s p m azure account bad request response has a 5xx status code
func (o *CreateCSPMAzureAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create c s p m azure account bad request response a status code equal to that given
func (o *CreateCSPMAzureAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateCSPMAzureAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCSPMAzureAccountBadRequest) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCSPMAzureAccountBadRequest) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *CreateCSPMAzureAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCSPMAzureAccountForbidden creates a CreateCSPMAzureAccountForbidden with default headers values
func NewCreateCSPMAzureAccountForbidden() *CreateCSPMAzureAccountForbidden {
	return &CreateCSPMAzureAccountForbidden{}
}

/*
CreateCSPMAzureAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateCSPMAzureAccountForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this create c s p m azure account forbidden response has a 2xx status code
func (o *CreateCSPMAzureAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create c s p m azure account forbidden response has a 3xx status code
func (o *CreateCSPMAzureAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c s p m azure account forbidden response has a 4xx status code
func (o *CreateCSPMAzureAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create c s p m azure account forbidden response has a 5xx status code
func (o *CreateCSPMAzureAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create c s p m azure account forbidden response a status code equal to that given
func (o *CreateCSPMAzureAccountForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateCSPMAzureAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountForbidden  %+v", 403, o.Payload)
}

func (o *CreateCSPMAzureAccountForbidden) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountForbidden  %+v", 403, o.Payload)
}

func (o *CreateCSPMAzureAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateCSPMAzureAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCSPMAzureAccountTooManyRequests creates a CreateCSPMAzureAccountTooManyRequests with default headers values
func NewCreateCSPMAzureAccountTooManyRequests() *CreateCSPMAzureAccountTooManyRequests {
	return &CreateCSPMAzureAccountTooManyRequests{}
}

/*
CreateCSPMAzureAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateCSPMAzureAccountTooManyRequests struct {

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

// IsSuccess returns true when this create c s p m azure account too many requests response has a 2xx status code
func (o *CreateCSPMAzureAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create c s p m azure account too many requests response has a 3xx status code
func (o *CreateCSPMAzureAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c s p m azure account too many requests response has a 4xx status code
func (o *CreateCSPMAzureAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create c s p m azure account too many requests response has a 5xx status code
func (o *CreateCSPMAzureAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create c s p m azure account too many requests response a status code equal to that given
func (o *CreateCSPMAzureAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateCSPMAzureAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateCSPMAzureAccountTooManyRequests) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateCSPMAzureAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateCSPMAzureAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCSPMAzureAccountInternalServerError creates a CreateCSPMAzureAccountInternalServerError with default headers values
func NewCreateCSPMAzureAccountInternalServerError() *CreateCSPMAzureAccountInternalServerError {
	return &CreateCSPMAzureAccountInternalServerError{}
}

/*
CreateCSPMAzureAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateCSPMAzureAccountInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureAccountResponseV1
}

// IsSuccess returns true when this create c s p m azure account internal server error response has a 2xx status code
func (o *CreateCSPMAzureAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create c s p m azure account internal server error response has a 3xx status code
func (o *CreateCSPMAzureAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c s p m azure account internal server error response has a 4xx status code
func (o *CreateCSPMAzureAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create c s p m azure account internal server error response has a 5xx status code
func (o *CreateCSPMAzureAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create c s p m azure account internal server error response a status code equal to that given
func (o *CreateCSPMAzureAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateCSPMAzureAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCSPMAzureAccountInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-azure/entities/account/v1][%d] createCSPMAzureAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCSPMAzureAccountInternalServerError) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *CreateCSPMAzureAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
