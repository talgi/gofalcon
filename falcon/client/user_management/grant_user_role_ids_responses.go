// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// GrantUserRoleIdsReader is a Reader for the GrantUserRoleIds structure.
type GrantUserRoleIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GrantUserRoleIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGrantUserRoleIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGrantUserRoleIdsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGrantUserRoleIdsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGrantUserRoleIdsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGrantUserRoleIdsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGrantUserRoleIdsOK creates a GrantUserRoleIdsOK with default headers values
func NewGrantUserRoleIdsOK() *GrantUserRoleIdsOK {
	return &GrantUserRoleIdsOK{}
}

/*
GrantUserRoleIdsOK describes a response with status code 200, with default header values.

OK
*/
type GrantUserRoleIdsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainUserRoleIDsResponse
}

// IsSuccess returns true when this grant user role ids o k response has a 2xx status code
func (o *GrantUserRoleIdsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this grant user role ids o k response has a 3xx status code
func (o *GrantUserRoleIdsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this grant user role ids o k response has a 4xx status code
func (o *GrantUserRoleIdsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this grant user role ids o k response has a 5xx status code
func (o *GrantUserRoleIdsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this grant user role ids o k response a status code equal to that given
func (o *GrantUserRoleIdsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GrantUserRoleIdsOK) Error() string {
	return fmt.Sprintf("[POST /user-roles/entities/user-roles/v1][%d] grantUserRoleIdsOK  %+v", 200, o.Payload)
}

func (o *GrantUserRoleIdsOK) String() string {
	return fmt.Sprintf("[POST /user-roles/entities/user-roles/v1][%d] grantUserRoleIdsOK  %+v", 200, o.Payload)
}

func (o *GrantUserRoleIdsOK) GetPayload() *models.DomainUserRoleIDsResponse {
	return o.Payload
}

func (o *GrantUserRoleIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainUserRoleIDsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGrantUserRoleIdsBadRequest creates a GrantUserRoleIdsBadRequest with default headers values
func NewGrantUserRoleIdsBadRequest() *GrantUserRoleIdsBadRequest {
	return &GrantUserRoleIdsBadRequest{}
}

/*
GrantUserRoleIdsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GrantUserRoleIdsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this grant user role ids bad request response has a 2xx status code
func (o *GrantUserRoleIdsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this grant user role ids bad request response has a 3xx status code
func (o *GrantUserRoleIdsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this grant user role ids bad request response has a 4xx status code
func (o *GrantUserRoleIdsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this grant user role ids bad request response has a 5xx status code
func (o *GrantUserRoleIdsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this grant user role ids bad request response a status code equal to that given
func (o *GrantUserRoleIdsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GrantUserRoleIdsBadRequest) Error() string {
	return fmt.Sprintf("[POST /user-roles/entities/user-roles/v1][%d] grantUserRoleIdsBadRequest  %+v", 400, o.Payload)
}

func (o *GrantUserRoleIdsBadRequest) String() string {
	return fmt.Sprintf("[POST /user-roles/entities/user-roles/v1][%d] grantUserRoleIdsBadRequest  %+v", 400, o.Payload)
}

func (o *GrantUserRoleIdsBadRequest) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *GrantUserRoleIdsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGrantUserRoleIdsForbidden creates a GrantUserRoleIdsForbidden with default headers values
func NewGrantUserRoleIdsForbidden() *GrantUserRoleIdsForbidden {
	return &GrantUserRoleIdsForbidden{}
}

/*
GrantUserRoleIdsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GrantUserRoleIdsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this grant user role ids forbidden response has a 2xx status code
func (o *GrantUserRoleIdsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this grant user role ids forbidden response has a 3xx status code
func (o *GrantUserRoleIdsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this grant user role ids forbidden response has a 4xx status code
func (o *GrantUserRoleIdsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this grant user role ids forbidden response has a 5xx status code
func (o *GrantUserRoleIdsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this grant user role ids forbidden response a status code equal to that given
func (o *GrantUserRoleIdsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GrantUserRoleIdsForbidden) Error() string {
	return fmt.Sprintf("[POST /user-roles/entities/user-roles/v1][%d] grantUserRoleIdsForbidden  %+v", 403, o.Payload)
}

func (o *GrantUserRoleIdsForbidden) String() string {
	return fmt.Sprintf("[POST /user-roles/entities/user-roles/v1][%d] grantUserRoleIdsForbidden  %+v", 403, o.Payload)
}

func (o *GrantUserRoleIdsForbidden) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *GrantUserRoleIdsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGrantUserRoleIdsTooManyRequests creates a GrantUserRoleIdsTooManyRequests with default headers values
func NewGrantUserRoleIdsTooManyRequests() *GrantUserRoleIdsTooManyRequests {
	return &GrantUserRoleIdsTooManyRequests{}
}

/*
GrantUserRoleIdsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GrantUserRoleIdsTooManyRequests struct {

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

// IsSuccess returns true when this grant user role ids too many requests response has a 2xx status code
func (o *GrantUserRoleIdsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this grant user role ids too many requests response has a 3xx status code
func (o *GrantUserRoleIdsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this grant user role ids too many requests response has a 4xx status code
func (o *GrantUserRoleIdsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this grant user role ids too many requests response has a 5xx status code
func (o *GrantUserRoleIdsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this grant user role ids too many requests response a status code equal to that given
func (o *GrantUserRoleIdsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GrantUserRoleIdsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /user-roles/entities/user-roles/v1][%d] grantUserRoleIdsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GrantUserRoleIdsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /user-roles/entities/user-roles/v1][%d] grantUserRoleIdsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GrantUserRoleIdsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GrantUserRoleIdsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGrantUserRoleIdsDefault creates a GrantUserRoleIdsDefault with default headers values
func NewGrantUserRoleIdsDefault(code int) *GrantUserRoleIdsDefault {
	return &GrantUserRoleIdsDefault{
		_statusCode: code,
	}
}

/*
GrantUserRoleIdsDefault describes a response with status code -1, with default header values.

OK
*/
type GrantUserRoleIdsDefault struct {
	_statusCode int

	Payload *models.DomainUserRoleIDsResponse
}

// Code gets the status code for the grant user role ids default response
func (o *GrantUserRoleIdsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this grant user role ids default response has a 2xx status code
func (o *GrantUserRoleIdsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this grant user role ids default response has a 3xx status code
func (o *GrantUserRoleIdsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this grant user role ids default response has a 4xx status code
func (o *GrantUserRoleIdsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this grant user role ids default response has a 5xx status code
func (o *GrantUserRoleIdsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this grant user role ids default response a status code equal to that given
func (o *GrantUserRoleIdsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GrantUserRoleIdsDefault) Error() string {
	return fmt.Sprintf("[POST /user-roles/entities/user-roles/v1][%d] GrantUserRoleIds default  %+v", o._statusCode, o.Payload)
}

func (o *GrantUserRoleIdsDefault) String() string {
	return fmt.Sprintf("[POST /user-roles/entities/user-roles/v1][%d] GrantUserRoleIds default  %+v", o._statusCode, o.Payload)
}

func (o *GrantUserRoleIdsDefault) GetPayload() *models.DomainUserRoleIDsResponse {
	return o.Payload
}

func (o *GrantUserRoleIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainUserRoleIDsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
