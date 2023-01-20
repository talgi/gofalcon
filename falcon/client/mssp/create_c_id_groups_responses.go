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

// CreateCIDGroupsReader is a Reader for the CreateCIDGroups structure.
type CreateCIDGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCIDGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCIDGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewCreateCIDGroupsMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCIDGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCIDGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateCIDGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateCIDGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCIDGroupsOK creates a CreateCIDGroupsOK with default headers values
func NewCreateCIDGroupsOK() *CreateCIDGroupsOK {
	return &CreateCIDGroupsOK{}
}

/*
CreateCIDGroupsOK describes a response with status code 200, with default header values.

OK
*/
type CreateCIDGroupsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupsResponseV1
}

// IsSuccess returns true when this create c Id groups o k response has a 2xx status code
func (o *CreateCIDGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create c Id groups o k response has a 3xx status code
func (o *CreateCIDGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c Id groups o k response has a 4xx status code
func (o *CreateCIDGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create c Id groups o k response has a 5xx status code
func (o *CreateCIDGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create c Id groups o k response a status code equal to that given
func (o *CreateCIDGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create c Id groups o k response
func (o *CreateCIDGroupsOK) Code() int {
	return 200
}

func (o *CreateCIDGroupsOK) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIdGroupsOK  %+v", 200, o.Payload)
}

func (o *CreateCIDGroupsOK) String() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIdGroupsOK  %+v", 200, o.Payload)
}

func (o *CreateCIDGroupsOK) GetPayload() *models.DomainCIDGroupsResponseV1 {
	return o.Payload
}

func (o *CreateCIDGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCIDGroupsMultiStatus creates a CreateCIDGroupsMultiStatus with default headers values
func NewCreateCIDGroupsMultiStatus() *CreateCIDGroupsMultiStatus {
	return &CreateCIDGroupsMultiStatus{}
}

/*
CreateCIDGroupsMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type CreateCIDGroupsMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupsResponseV1
}

// IsSuccess returns true when this create c Id groups multi status response has a 2xx status code
func (o *CreateCIDGroupsMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create c Id groups multi status response has a 3xx status code
func (o *CreateCIDGroupsMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c Id groups multi status response has a 4xx status code
func (o *CreateCIDGroupsMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this create c Id groups multi status response has a 5xx status code
func (o *CreateCIDGroupsMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this create c Id groups multi status response a status code equal to that given
func (o *CreateCIDGroupsMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the create c Id groups multi status response
func (o *CreateCIDGroupsMultiStatus) Code() int {
	return 207
}

func (o *CreateCIDGroupsMultiStatus) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIdGroupsMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateCIDGroupsMultiStatus) String() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIdGroupsMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateCIDGroupsMultiStatus) GetPayload() *models.DomainCIDGroupsResponseV1 {
	return o.Payload
}

func (o *CreateCIDGroupsMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCIDGroupsBadRequest creates a CreateCIDGroupsBadRequest with default headers values
func NewCreateCIDGroupsBadRequest() *CreateCIDGroupsBadRequest {
	return &CreateCIDGroupsBadRequest{}
}

/*
CreateCIDGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateCIDGroupsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this create c Id groups bad request response has a 2xx status code
func (o *CreateCIDGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create c Id groups bad request response has a 3xx status code
func (o *CreateCIDGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c Id groups bad request response has a 4xx status code
func (o *CreateCIDGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create c Id groups bad request response has a 5xx status code
func (o *CreateCIDGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create c Id groups bad request response a status code equal to that given
func (o *CreateCIDGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create c Id groups bad request response
func (o *CreateCIDGroupsBadRequest) Code() int {
	return 400
}

func (o *CreateCIDGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIdGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCIDGroupsBadRequest) String() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIdGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCIDGroupsBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *CreateCIDGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCIDGroupsForbidden creates a CreateCIDGroupsForbidden with default headers values
func NewCreateCIDGroupsForbidden() *CreateCIDGroupsForbidden {
	return &CreateCIDGroupsForbidden{}
}

/*
CreateCIDGroupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateCIDGroupsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this create c Id groups forbidden response has a 2xx status code
func (o *CreateCIDGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create c Id groups forbidden response has a 3xx status code
func (o *CreateCIDGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c Id groups forbidden response has a 4xx status code
func (o *CreateCIDGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create c Id groups forbidden response has a 5xx status code
func (o *CreateCIDGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create c Id groups forbidden response a status code equal to that given
func (o *CreateCIDGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create c Id groups forbidden response
func (o *CreateCIDGroupsForbidden) Code() int {
	return 403
}

func (o *CreateCIDGroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIdGroupsForbidden  %+v", 403, o.Payload)
}

func (o *CreateCIDGroupsForbidden) String() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIdGroupsForbidden  %+v", 403, o.Payload)
}

func (o *CreateCIDGroupsForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *CreateCIDGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCIDGroupsTooManyRequests creates a CreateCIDGroupsTooManyRequests with default headers values
func NewCreateCIDGroupsTooManyRequests() *CreateCIDGroupsTooManyRequests {
	return &CreateCIDGroupsTooManyRequests{}
}

/*
CreateCIDGroupsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateCIDGroupsTooManyRequests struct {

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

// IsSuccess returns true when this create c Id groups too many requests response has a 2xx status code
func (o *CreateCIDGroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create c Id groups too many requests response has a 3xx status code
func (o *CreateCIDGroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create c Id groups too many requests response has a 4xx status code
func (o *CreateCIDGroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create c Id groups too many requests response has a 5xx status code
func (o *CreateCIDGroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create c Id groups too many requests response a status code equal to that given
func (o *CreateCIDGroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create c Id groups too many requests response
func (o *CreateCIDGroupsTooManyRequests) Code() int {
	return 429
}

func (o *CreateCIDGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIdGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateCIDGroupsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIdGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateCIDGroupsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateCIDGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCIDGroupsDefault creates a CreateCIDGroupsDefault with default headers values
func NewCreateCIDGroupsDefault(code int) *CreateCIDGroupsDefault {
	return &CreateCIDGroupsDefault{
		_statusCode: code,
	}
}

/*
CreateCIDGroupsDefault describes a response with status code -1, with default header values.

OK
*/
type CreateCIDGroupsDefault struct {
	_statusCode int

	Payload *models.DomainCIDGroupsResponseV1
}

// IsSuccess returns true when this create c ID groups default response has a 2xx status code
func (o *CreateCIDGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create c ID groups default response has a 3xx status code
func (o *CreateCIDGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create c ID groups default response has a 4xx status code
func (o *CreateCIDGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create c ID groups default response has a 5xx status code
func (o *CreateCIDGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create c ID groups default response a status code equal to that given
func (o *CreateCIDGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create c ID groups default response
func (o *CreateCIDGroupsDefault) Code() int {
	return o._statusCode
}

func (o *CreateCIDGroupsDefault) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIDGroups default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCIDGroupsDefault) String() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-groups/v1][%d] createCIDGroups default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCIDGroupsDefault) GetPayload() *models.DomainCIDGroupsResponseV1 {
	return o.Payload
}

func (o *CreateCIDGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainCIDGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
