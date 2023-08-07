// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// GetRuleTypesReader is a Reader for the GetRuleTypes structure.
type GetRuleTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuleTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuleTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRuleTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRuleTypesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRuleTypesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /ioarules/entities/rule-types/v1] get-rule-types", response, response.Code())
	}
}

// NewGetRuleTypesOK creates a GetRuleTypesOK with default headers values
func NewGetRuleTypesOK() *GetRuleTypesOK {
	return &GetRuleTypesOK{}
}

/*
GetRuleTypesOK describes a response with status code 200, with default header values.

OK
*/
type GetRuleTypesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIRuleTypesResponse
}

// IsSuccess returns true when this get rule types o k response has a 2xx status code
func (o *GetRuleTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get rule types o k response has a 3xx status code
func (o *GetRuleTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule types o k response has a 4xx status code
func (o *GetRuleTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rule types o k response has a 5xx status code
func (o *GetRuleTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule types o k response a status code equal to that given
func (o *GetRuleTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get rule types o k response
func (o *GetRuleTypesOK) Code() int {
	return 200
}

func (o *GetRuleTypesOK) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-types/v1][%d] getRuleTypesOK  %+v", 200, o.Payload)
}

func (o *GetRuleTypesOK) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-types/v1][%d] getRuleTypesOK  %+v", 200, o.Payload)
}

func (o *GetRuleTypesOK) GetPayload() *models.APIRuleTypesResponse {
	return o.Payload
}

func (o *GetRuleTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIRuleTypesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleTypesForbidden creates a GetRuleTypesForbidden with default headers values
func NewGetRuleTypesForbidden() *GetRuleTypesForbidden {
	return &GetRuleTypesForbidden{}
}

/*
GetRuleTypesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRuleTypesForbidden struct {

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

// IsSuccess returns true when this get rule types forbidden response has a 2xx status code
func (o *GetRuleTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rule types forbidden response has a 3xx status code
func (o *GetRuleTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule types forbidden response has a 4xx status code
func (o *GetRuleTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rule types forbidden response has a 5xx status code
func (o *GetRuleTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule types forbidden response a status code equal to that given
func (o *GetRuleTypesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get rule types forbidden response
func (o *GetRuleTypesForbidden) Code() int {
	return 403
}

func (o *GetRuleTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-types/v1][%d] getRuleTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetRuleTypesForbidden) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-types/v1][%d] getRuleTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetRuleTypesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRuleTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRuleTypesNotFound creates a GetRuleTypesNotFound with default headers values
func NewGetRuleTypesNotFound() *GetRuleTypesNotFound {
	return &GetRuleTypesNotFound{}
}

/*
GetRuleTypesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRuleTypesNotFound struct {

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

// IsSuccess returns true when this get rule types not found response has a 2xx status code
func (o *GetRuleTypesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rule types not found response has a 3xx status code
func (o *GetRuleTypesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule types not found response has a 4xx status code
func (o *GetRuleTypesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rule types not found response has a 5xx status code
func (o *GetRuleTypesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule types not found response a status code equal to that given
func (o *GetRuleTypesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get rule types not found response
func (o *GetRuleTypesNotFound) Code() int {
	return 404
}

func (o *GetRuleTypesNotFound) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-types/v1][%d] getRuleTypesNotFound  %+v", 404, o.Payload)
}

func (o *GetRuleTypesNotFound) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-types/v1][%d] getRuleTypesNotFound  %+v", 404, o.Payload)
}

func (o *GetRuleTypesNotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRuleTypesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRuleTypesTooManyRequests creates a GetRuleTypesTooManyRequests with default headers values
func NewGetRuleTypesTooManyRequests() *GetRuleTypesTooManyRequests {
	return &GetRuleTypesTooManyRequests{}
}

/*
GetRuleTypesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRuleTypesTooManyRequests struct {

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

// IsSuccess returns true when this get rule types too many requests response has a 2xx status code
func (o *GetRuleTypesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rule types too many requests response has a 3xx status code
func (o *GetRuleTypesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule types too many requests response has a 4xx status code
func (o *GetRuleTypesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rule types too many requests response has a 5xx status code
func (o *GetRuleTypesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule types too many requests response a status code equal to that given
func (o *GetRuleTypesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get rule types too many requests response
func (o *GetRuleTypesTooManyRequests) Code() int {
	return 429
}

func (o *GetRuleTypesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-types/v1][%d] getRuleTypesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRuleTypesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-types/v1][%d] getRuleTypesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRuleTypesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRuleTypesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
