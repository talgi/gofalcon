// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// CreateNetworkLocationsReader is a Reader for the CreateNetworkLocations structure.
type CreateNetworkLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkLocationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNetworkLocationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateNetworkLocationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateNetworkLocationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /fwmgr/entities/network-locations/v1] create-network-locations", response, response.Code())
	}
}

// NewCreateNetworkLocationsCreated creates a CreateNetworkLocationsCreated with default headers values
func NewCreateNetworkLocationsCreated() *CreateNetworkLocationsCreated {
	return &CreateNetworkLocationsCreated{}
}

/*
CreateNetworkLocationsCreated describes a response with status code 201, with default header values.

Created
*/
type CreateNetworkLocationsCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrAPINetworkLocationsResponse
}

// IsSuccess returns true when this create network locations created response has a 2xx status code
func (o *CreateNetworkLocationsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network locations created response has a 3xx status code
func (o *CreateNetworkLocationsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network locations created response has a 4xx status code
func (o *CreateNetworkLocationsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network locations created response has a 5xx status code
func (o *CreateNetworkLocationsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network locations created response a status code equal to that given
func (o *CreateNetworkLocationsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network locations created response
func (o *CreateNetworkLocationsCreated) Code() int {
	return 201
}

func (o *CreateNetworkLocationsCreated) Error() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations/v1][%d] createNetworkLocationsCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkLocationsCreated) String() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations/v1][%d] createNetworkLocationsCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkLocationsCreated) GetPayload() *models.FwmgrAPINetworkLocationsResponse {
	return o.Payload
}

func (o *CreateNetworkLocationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrAPINetworkLocationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNetworkLocationsBadRequest creates a CreateNetworkLocationsBadRequest with default headers values
func NewCreateNetworkLocationsBadRequest() *CreateNetworkLocationsBadRequest {
	return &CreateNetworkLocationsBadRequest{}
}

/*
CreateNetworkLocationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateNetworkLocationsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaspecResponseFields
}

// IsSuccess returns true when this create network locations bad request response has a 2xx status code
func (o *CreateNetworkLocationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create network locations bad request response has a 3xx status code
func (o *CreateNetworkLocationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network locations bad request response has a 4xx status code
func (o *CreateNetworkLocationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create network locations bad request response has a 5xx status code
func (o *CreateNetworkLocationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create network locations bad request response a status code equal to that given
func (o *CreateNetworkLocationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create network locations bad request response
func (o *CreateNetworkLocationsBadRequest) Code() int {
	return 400
}

func (o *CreateNetworkLocationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations/v1][%d] createNetworkLocationsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNetworkLocationsBadRequest) String() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations/v1][%d] createNetworkLocationsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNetworkLocationsBadRequest) GetPayload() *models.FwmgrMsaspecResponseFields {
	return o.Payload
}

func (o *CreateNetworkLocationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrMsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNetworkLocationsForbidden creates a CreateNetworkLocationsForbidden with default headers values
func NewCreateNetworkLocationsForbidden() *CreateNetworkLocationsForbidden {
	return &CreateNetworkLocationsForbidden{}
}

/*
CreateNetworkLocationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateNetworkLocationsForbidden struct {

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

// IsSuccess returns true when this create network locations forbidden response has a 2xx status code
func (o *CreateNetworkLocationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create network locations forbidden response has a 3xx status code
func (o *CreateNetworkLocationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network locations forbidden response has a 4xx status code
func (o *CreateNetworkLocationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create network locations forbidden response has a 5xx status code
func (o *CreateNetworkLocationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create network locations forbidden response a status code equal to that given
func (o *CreateNetworkLocationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create network locations forbidden response
func (o *CreateNetworkLocationsForbidden) Code() int {
	return 403
}

func (o *CreateNetworkLocationsForbidden) Error() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations/v1][%d] createNetworkLocationsForbidden  %+v", 403, o.Payload)
}

func (o *CreateNetworkLocationsForbidden) String() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations/v1][%d] createNetworkLocationsForbidden  %+v", 403, o.Payload)
}

func (o *CreateNetworkLocationsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateNetworkLocationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateNetworkLocationsTooManyRequests creates a CreateNetworkLocationsTooManyRequests with default headers values
func NewCreateNetworkLocationsTooManyRequests() *CreateNetworkLocationsTooManyRequests {
	return &CreateNetworkLocationsTooManyRequests{}
}

/*
CreateNetworkLocationsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateNetworkLocationsTooManyRequests struct {

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

// IsSuccess returns true when this create network locations too many requests response has a 2xx status code
func (o *CreateNetworkLocationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create network locations too many requests response has a 3xx status code
func (o *CreateNetworkLocationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network locations too many requests response has a 4xx status code
func (o *CreateNetworkLocationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create network locations too many requests response has a 5xx status code
func (o *CreateNetworkLocationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create network locations too many requests response a status code equal to that given
func (o *CreateNetworkLocationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create network locations too many requests response
func (o *CreateNetworkLocationsTooManyRequests) Code() int {
	return 429
}

func (o *CreateNetworkLocationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations/v1][%d] createNetworkLocationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateNetworkLocationsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations/v1][%d] createNetworkLocationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateNetworkLocationsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateNetworkLocationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
