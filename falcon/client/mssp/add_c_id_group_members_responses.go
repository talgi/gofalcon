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

// AddCIDGroupMembersReader is a Reader for the AddCIDGroupMembers structure.
type AddCIDGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCIDGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddCIDGroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewAddCIDGroupMembersMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddCIDGroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddCIDGroupMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddCIDGroupMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddCIDGroupMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddCIDGroupMembersOK creates a AddCIDGroupMembersOK with default headers values
func NewAddCIDGroupMembersOK() *AddCIDGroupMembersOK {
	return &AddCIDGroupMembersOK{}
}

/* AddCIDGroupMembersOK describes a response with status code 200, with default header values.

OK
*/
type AddCIDGroupMembersOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupMembersResponseV1
}

func (o *AddCIDGroupMembersOK) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-group-members/v1][%d] addCIdGroupMembersOK  %+v", 200, o.Payload)
}
func (o *AddCIDGroupMembersOK) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *AddCIDGroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCIDGroupMembersMultiStatus creates a AddCIDGroupMembersMultiStatus with default headers values
func NewAddCIDGroupMembersMultiStatus() *AddCIDGroupMembersMultiStatus {
	return &AddCIDGroupMembersMultiStatus{}
}

/* AddCIDGroupMembersMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type AddCIDGroupMembersMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupMembersResponseV1
}

func (o *AddCIDGroupMembersMultiStatus) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-group-members/v1][%d] addCIdGroupMembersMultiStatus  %+v", 207, o.Payload)
}
func (o *AddCIDGroupMembersMultiStatus) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *AddCIDGroupMembersMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCIDGroupMembersBadRequest creates a AddCIDGroupMembersBadRequest with default headers values
func NewAddCIDGroupMembersBadRequest() *AddCIDGroupMembersBadRequest {
	return &AddCIDGroupMembersBadRequest{}
}

/* AddCIDGroupMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddCIDGroupMembersBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *AddCIDGroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-group-members/v1][%d] addCIdGroupMembersBadRequest  %+v", 400, o.Payload)
}
func (o *AddCIDGroupMembersBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *AddCIDGroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddCIDGroupMembersForbidden creates a AddCIDGroupMembersForbidden with default headers values
func NewAddCIDGroupMembersForbidden() *AddCIDGroupMembersForbidden {
	return &AddCIDGroupMembersForbidden{}
}

/* AddCIDGroupMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddCIDGroupMembersForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *AddCIDGroupMembersForbidden) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-group-members/v1][%d] addCIdGroupMembersForbidden  %+v", 403, o.Payload)
}
func (o *AddCIDGroupMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *AddCIDGroupMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddCIDGroupMembersTooManyRequests creates a AddCIDGroupMembersTooManyRequests with default headers values
func NewAddCIDGroupMembersTooManyRequests() *AddCIDGroupMembersTooManyRequests {
	return &AddCIDGroupMembersTooManyRequests{}
}

/* AddCIDGroupMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddCIDGroupMembersTooManyRequests struct {

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

func (o *AddCIDGroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-group-members/v1][%d] addCIdGroupMembersTooManyRequests  %+v", 429, o.Payload)
}
func (o *AddCIDGroupMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AddCIDGroupMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddCIDGroupMembersDefault creates a AddCIDGroupMembersDefault with default headers values
func NewAddCIDGroupMembersDefault(code int) *AddCIDGroupMembersDefault {
	return &AddCIDGroupMembersDefault{
		_statusCode: code,
	}
}

/* AddCIDGroupMembersDefault describes a response with status code -1, with default header values.

OK
*/
type AddCIDGroupMembersDefault struct {
	_statusCode int

	Payload *models.DomainCIDGroupMembersResponseV1
}

// Code gets the status code for the add c ID group members default response
func (o *AddCIDGroupMembersDefault) Code() int {
	return o._statusCode
}

func (o *AddCIDGroupMembersDefault) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/cid-group-members/v1][%d] addCIDGroupMembers default  %+v", o._statusCode, o.Payload)
}
func (o *AddCIDGroupMembersDefault) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *AddCIDGroupMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
