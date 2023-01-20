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

// CreateUserGroupsReader is a Reader for the CreateUserGroups structure.
type CreateUserGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewCreateUserGroupsMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateUserGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateUserGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateUserGroupsOK creates a CreateUserGroupsOK with default headers values
func NewCreateUserGroupsOK() *CreateUserGroupsOK {
	return &CreateUserGroupsOK{}
}

/*
CreateUserGroupsOK describes a response with status code 200, with default header values.

OK
*/
type CreateUserGroupsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainUserGroupsResponseV1
}

// IsSuccess returns true when this create user groups o k response has a 2xx status code
func (o *CreateUserGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user groups o k response has a 3xx status code
func (o *CreateUserGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user groups o k response has a 4xx status code
func (o *CreateUserGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user groups o k response has a 5xx status code
func (o *CreateUserGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create user groups o k response a status code equal to that given
func (o *CreateUserGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create user groups o k response
func (o *CreateUserGroupsOK) Code() int {
	return 200
}

func (o *CreateUserGroupsOK) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroupsOK  %+v", 200, o.Payload)
}

func (o *CreateUserGroupsOK) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroupsOK  %+v", 200, o.Payload)
}

func (o *CreateUserGroupsOK) GetPayload() *models.DomainUserGroupsResponseV1 {
	return o.Payload
}

func (o *CreateUserGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainUserGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserGroupsMultiStatus creates a CreateUserGroupsMultiStatus with default headers values
func NewCreateUserGroupsMultiStatus() *CreateUserGroupsMultiStatus {
	return &CreateUserGroupsMultiStatus{}
}

/*
CreateUserGroupsMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type CreateUserGroupsMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainUserGroupsResponseV1
}

// IsSuccess returns true when this create user groups multi status response has a 2xx status code
func (o *CreateUserGroupsMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user groups multi status response has a 3xx status code
func (o *CreateUserGroupsMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user groups multi status response has a 4xx status code
func (o *CreateUserGroupsMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user groups multi status response has a 5xx status code
func (o *CreateUserGroupsMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this create user groups multi status response a status code equal to that given
func (o *CreateUserGroupsMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the create user groups multi status response
func (o *CreateUserGroupsMultiStatus) Code() int {
	return 207
}

func (o *CreateUserGroupsMultiStatus) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroupsMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateUserGroupsMultiStatus) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroupsMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateUserGroupsMultiStatus) GetPayload() *models.DomainUserGroupsResponseV1 {
	return o.Payload
}

func (o *CreateUserGroupsMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainUserGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserGroupsBadRequest creates a CreateUserGroupsBadRequest with default headers values
func NewCreateUserGroupsBadRequest() *CreateUserGroupsBadRequest {
	return &CreateUserGroupsBadRequest{}
}

/*
CreateUserGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateUserGroupsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this create user groups bad request response has a 2xx status code
func (o *CreateUserGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user groups bad request response has a 3xx status code
func (o *CreateUserGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user groups bad request response has a 4xx status code
func (o *CreateUserGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user groups bad request response has a 5xx status code
func (o *CreateUserGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create user groups bad request response a status code equal to that given
func (o *CreateUserGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create user groups bad request response
func (o *CreateUserGroupsBadRequest) Code() int {
	return 400
}

func (o *CreateUserGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserGroupsBadRequest) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserGroupsBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *CreateUserGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupsForbidden creates a CreateUserGroupsForbidden with default headers values
func NewCreateUserGroupsForbidden() *CreateUserGroupsForbidden {
	return &CreateUserGroupsForbidden{}
}

/*
CreateUserGroupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateUserGroupsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this create user groups forbidden response has a 2xx status code
func (o *CreateUserGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user groups forbidden response has a 3xx status code
func (o *CreateUserGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user groups forbidden response has a 4xx status code
func (o *CreateUserGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user groups forbidden response has a 5xx status code
func (o *CreateUserGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create user groups forbidden response a status code equal to that given
func (o *CreateUserGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create user groups forbidden response
func (o *CreateUserGroupsForbidden) Code() int {
	return 403
}

func (o *CreateUserGroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroupsForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserGroupsForbidden) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroupsForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserGroupsForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *CreateUserGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupsTooManyRequests creates a CreateUserGroupsTooManyRequests with default headers values
func NewCreateUserGroupsTooManyRequests() *CreateUserGroupsTooManyRequests {
	return &CreateUserGroupsTooManyRequests{}
}

/*
CreateUserGroupsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateUserGroupsTooManyRequests struct {

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

// IsSuccess returns true when this create user groups too many requests response has a 2xx status code
func (o *CreateUserGroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user groups too many requests response has a 3xx status code
func (o *CreateUserGroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user groups too many requests response has a 4xx status code
func (o *CreateUserGroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user groups too many requests response has a 5xx status code
func (o *CreateUserGroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create user groups too many requests response a status code equal to that given
func (o *CreateUserGroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create user groups too many requests response
func (o *CreateUserGroupsTooManyRequests) Code() int {
	return 429
}

func (o *CreateUserGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateUserGroupsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateUserGroupsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateUserGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupsDefault creates a CreateUserGroupsDefault with default headers values
func NewCreateUserGroupsDefault(code int) *CreateUserGroupsDefault {
	return &CreateUserGroupsDefault{
		_statusCode: code,
	}
}

/*
CreateUserGroupsDefault describes a response with status code -1, with default header values.

OK
*/
type CreateUserGroupsDefault struct {
	_statusCode int

	Payload *models.DomainUserGroupsResponseV1
}

// IsSuccess returns true when this create user groups default response has a 2xx status code
func (o *CreateUserGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create user groups default response has a 3xx status code
func (o *CreateUserGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create user groups default response has a 4xx status code
func (o *CreateUserGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create user groups default response has a 5xx status code
func (o *CreateUserGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create user groups default response a status code equal to that given
func (o *CreateUserGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create user groups default response
func (o *CreateUserGroupsDefault) Code() int {
	return o._statusCode
}

func (o *CreateUserGroupsDefault) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroups default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserGroupsDefault) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-groups/v1][%d] createUserGroups default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserGroupsDefault) GetPayload() *models.DomainUserGroupsResponseV1 {
	return o.Payload
}

func (o *CreateUserGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainUserGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
