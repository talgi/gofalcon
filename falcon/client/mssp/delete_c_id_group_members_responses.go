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

// DeleteCIDGroupMembersReader is a Reader for the DeleteCIDGroupMembers structure.
type DeleteCIDGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCIDGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCIDGroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewDeleteCIDGroupMembersMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCIDGroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCIDGroupMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCIDGroupMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteCIDGroupMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCIDGroupMembersOK creates a DeleteCIDGroupMembersOK with default headers values
func NewDeleteCIDGroupMembersOK() *DeleteCIDGroupMembersOK {
	return &DeleteCIDGroupMembersOK{}
}

/* DeleteCIDGroupMembersOK describes a response with status code 200, with default header values.

OK
*/
type DeleteCIDGroupMembersOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupMembersResponseV1
}

func (o *DeleteCIDGroupMembersOK) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/cid-group-members/v1][%d] deleteCIdGroupMembersOK  %+v", 200, o.Payload)
}
func (o *DeleteCIDGroupMembersOK) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *DeleteCIDGroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCIDGroupMembersMultiStatus creates a DeleteCIDGroupMembersMultiStatus with default headers values
func NewDeleteCIDGroupMembersMultiStatus() *DeleteCIDGroupMembersMultiStatus {
	return &DeleteCIDGroupMembersMultiStatus{}
}

/* DeleteCIDGroupMembersMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type DeleteCIDGroupMembersMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupMembersResponseV1
}

func (o *DeleteCIDGroupMembersMultiStatus) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/cid-group-members/v1][%d] deleteCIdGroupMembersMultiStatus  %+v", 207, o.Payload)
}
func (o *DeleteCIDGroupMembersMultiStatus) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *DeleteCIDGroupMembersMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCIDGroupMembersBadRequest creates a DeleteCIDGroupMembersBadRequest with default headers values
func NewDeleteCIDGroupMembersBadRequest() *DeleteCIDGroupMembersBadRequest {
	return &DeleteCIDGroupMembersBadRequest{}
}

/* DeleteCIDGroupMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteCIDGroupMembersBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *DeleteCIDGroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/cid-group-members/v1][%d] deleteCIdGroupMembersBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteCIDGroupMembersBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *DeleteCIDGroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCIDGroupMembersForbidden creates a DeleteCIDGroupMembersForbidden with default headers values
func NewDeleteCIDGroupMembersForbidden() *DeleteCIDGroupMembersForbidden {
	return &DeleteCIDGroupMembersForbidden{}
}

/* DeleteCIDGroupMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteCIDGroupMembersForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *DeleteCIDGroupMembersForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/cid-group-members/v1][%d] deleteCIdGroupMembersForbidden  %+v", 403, o.Payload)
}
func (o *DeleteCIDGroupMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *DeleteCIDGroupMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCIDGroupMembersTooManyRequests creates a DeleteCIDGroupMembersTooManyRequests with default headers values
func NewDeleteCIDGroupMembersTooManyRequests() *DeleteCIDGroupMembersTooManyRequests {
	return &DeleteCIDGroupMembersTooManyRequests{}
}

/* DeleteCIDGroupMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteCIDGroupMembersTooManyRequests struct {

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

func (o *DeleteCIDGroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/cid-group-members/v1][%d] deleteCIdGroupMembersTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteCIDGroupMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteCIDGroupMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCIDGroupMembersDefault creates a DeleteCIDGroupMembersDefault with default headers values
func NewDeleteCIDGroupMembersDefault(code int) *DeleteCIDGroupMembersDefault {
	return &DeleteCIDGroupMembersDefault{
		_statusCode: code,
	}
}

/* DeleteCIDGroupMembersDefault describes a response with status code -1, with default header values.

OK
*/
type DeleteCIDGroupMembersDefault struct {
	_statusCode int

	Payload *models.DomainCIDGroupMembersResponseV1
}

// Code gets the status code for the delete c ID group members default response
func (o *DeleteCIDGroupMembersDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCIDGroupMembersDefault) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/cid-group-members/v1][%d] deleteCIDGroupMembers default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteCIDGroupMembersDefault) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *DeleteCIDGroupMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
