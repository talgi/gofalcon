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

// GetCSPMAzureAccountReader is a Reader for the GetCSPMAzureAccount structure.
type GetCSPMAzureAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMAzureAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMAzureAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetCSPMAzureAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMAzureAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMAzureAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMAzureAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMAzureAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCSPMAzureAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCSPMAzureAccountOK creates a GetCSPMAzureAccountOK with default headers values
func NewGetCSPMAzureAccountOK() *GetCSPMAzureAccountOK {
	return &GetCSPMAzureAccountOK{}
}

/*
	GetCSPMAzureAccountOK describes a response with status code 200, with default header values.

OK
*/
type GetCSPMAzureAccountOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureAccountResponseV1
}

func (o *GetCSPMAzureAccountOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/account/v1][%d] getCSPMAzureAccountOK  %+v", 200, o.Payload)
}
func (o *GetCSPMAzureAccountOK) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMAzureAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAzureAccountMultiStatus creates a GetCSPMAzureAccountMultiStatus with default headers values
func NewGetCSPMAzureAccountMultiStatus() *GetCSPMAzureAccountMultiStatus {
	return &GetCSPMAzureAccountMultiStatus{}
}

/*
	GetCSPMAzureAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetCSPMAzureAccountMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureAccountResponseV1
}

func (o *GetCSPMAzureAccountMultiStatus) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/account/v1][%d] getCSPMAzureAccountMultiStatus  %+v", 207, o.Payload)
}
func (o *GetCSPMAzureAccountMultiStatus) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMAzureAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAzureAccountBadRequest creates a GetCSPMAzureAccountBadRequest with default headers values
func NewGetCSPMAzureAccountBadRequest() *GetCSPMAzureAccountBadRequest {
	return &GetCSPMAzureAccountBadRequest{}
}

/*
	GetCSPMAzureAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCSPMAzureAccountBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureAccountResponseV1
}

func (o *GetCSPMAzureAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/account/v1][%d] getCSPMAzureAccountBadRequest  %+v", 400, o.Payload)
}
func (o *GetCSPMAzureAccountBadRequest) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMAzureAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAzureAccountForbidden creates a GetCSPMAzureAccountForbidden with default headers values
func NewGetCSPMAzureAccountForbidden() *GetCSPMAzureAccountForbidden {
	return &GetCSPMAzureAccountForbidden{}
}

/*
	GetCSPMAzureAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCSPMAzureAccountForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetCSPMAzureAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/account/v1][%d] getCSPMAzureAccountForbidden  %+v", 403, o.Payload)
}
func (o *GetCSPMAzureAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMAzureAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAzureAccountTooManyRequests creates a GetCSPMAzureAccountTooManyRequests with default headers values
func NewGetCSPMAzureAccountTooManyRequests() *GetCSPMAzureAccountTooManyRequests {
	return &GetCSPMAzureAccountTooManyRequests{}
}

/*
	GetCSPMAzureAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCSPMAzureAccountTooManyRequests struct {

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

func (o *GetCSPMAzureAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/account/v1][%d] getCSPMAzureAccountTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetCSPMAzureAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMAzureAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAzureAccountInternalServerError creates a GetCSPMAzureAccountInternalServerError with default headers values
func NewGetCSPMAzureAccountInternalServerError() *GetCSPMAzureAccountInternalServerError {
	return &GetCSPMAzureAccountInternalServerError{}
}

/*
	GetCSPMAzureAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCSPMAzureAccountInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureAccountResponseV1
}

func (o *GetCSPMAzureAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/account/v1][%d] getCSPMAzureAccountInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCSPMAzureAccountInternalServerError) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMAzureAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAzureAccountDefault creates a GetCSPMAzureAccountDefault with default headers values
func NewGetCSPMAzureAccountDefault(code int) *GetCSPMAzureAccountDefault {
	return &GetCSPMAzureAccountDefault{
		_statusCode: code,
	}
}

/*
	GetCSPMAzureAccountDefault describes a response with status code -1, with default header values.

OK
*/
type GetCSPMAzureAccountDefault struct {
	_statusCode int

	Payload *models.RegistrationAzureAccountResponseV1
}

// Code gets the status code for the get c s p m azure account default response
func (o *GetCSPMAzureAccountDefault) Code() int {
	return o._statusCode
}

func (o *GetCSPMAzureAccountDefault) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/account/v1][%d] GetCSPMAzureAccount default  %+v", o._statusCode, o.Payload)
}
func (o *GetCSPMAzureAccountDefault) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMAzureAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationAzureAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
