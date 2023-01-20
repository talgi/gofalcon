// Code generated by go-swagger; DO NOT EDIT.

package malquery

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

// PostMalQueryExactSearchV1Reader is a Reader for the PostMalQueryExactSearchV1 structure.
type PostMalQueryExactSearchV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMalQueryExactSearchV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostMalQueryExactSearchV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostMalQueryExactSearchV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostMalQueryExactSearchV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostMalQueryExactSearchV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostMalQueryExactSearchV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostMalQueryExactSearchV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostMalQueryExactSearchV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMalQueryExactSearchV1OK creates a PostMalQueryExactSearchV1OK with default headers values
func NewPostMalQueryExactSearchV1OK() *PostMalQueryExactSearchV1OK {
	return &PostMalQueryExactSearchV1OK{}
}

/*
PostMalQueryExactSearchV1OK describes a response with status code 200, with default header values.

OK
*/
type PostMalQueryExactSearchV1OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this post mal query exact search v1 o k response has a 2xx status code
func (o *PostMalQueryExactSearchV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post mal query exact search v1 o k response has a 3xx status code
func (o *PostMalQueryExactSearchV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query exact search v1 o k response has a 4xx status code
func (o *PostMalQueryExactSearchV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post mal query exact search v1 o k response has a 5xx status code
func (o *PostMalQueryExactSearchV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query exact search v1 o k response a status code equal to that given
func (o *PostMalQueryExactSearchV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post mal query exact search v1 o k response
func (o *PostMalQueryExactSearchV1OK) Code() int {
	return 200
}

func (o *PostMalQueryExactSearchV1OK) Error() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1OK  %+v", 200, o.Payload)
}

func (o *PostMalQueryExactSearchV1OK) String() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1OK  %+v", 200, o.Payload)
}

func (o *PostMalQueryExactSearchV1OK) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryExactSearchV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMalQueryExactSearchV1BadRequest creates a PostMalQueryExactSearchV1BadRequest with default headers values
func NewPostMalQueryExactSearchV1BadRequest() *PostMalQueryExactSearchV1BadRequest {
	return &PostMalQueryExactSearchV1BadRequest{}
}

/*
PostMalQueryExactSearchV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostMalQueryExactSearchV1BadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this post mal query exact search v1 bad request response has a 2xx status code
func (o *PostMalQueryExactSearchV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query exact search v1 bad request response has a 3xx status code
func (o *PostMalQueryExactSearchV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query exact search v1 bad request response has a 4xx status code
func (o *PostMalQueryExactSearchV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post mal query exact search v1 bad request response has a 5xx status code
func (o *PostMalQueryExactSearchV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query exact search v1 bad request response a status code equal to that given
func (o *PostMalQueryExactSearchV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post mal query exact search v1 bad request response
func (o *PostMalQueryExactSearchV1BadRequest) Code() int {
	return 400
}

func (o *PostMalQueryExactSearchV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1BadRequest  %+v", 400, o.Payload)
}

func (o *PostMalQueryExactSearchV1BadRequest) String() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1BadRequest  %+v", 400, o.Payload)
}

func (o *PostMalQueryExactSearchV1BadRequest) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryExactSearchV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMalQueryExactSearchV1Unauthorized creates a PostMalQueryExactSearchV1Unauthorized with default headers values
func NewPostMalQueryExactSearchV1Unauthorized() *PostMalQueryExactSearchV1Unauthorized {
	return &PostMalQueryExactSearchV1Unauthorized{}
}

/*
PostMalQueryExactSearchV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostMalQueryExactSearchV1Unauthorized struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this post mal query exact search v1 unauthorized response has a 2xx status code
func (o *PostMalQueryExactSearchV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query exact search v1 unauthorized response has a 3xx status code
func (o *PostMalQueryExactSearchV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query exact search v1 unauthorized response has a 4xx status code
func (o *PostMalQueryExactSearchV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post mal query exact search v1 unauthorized response has a 5xx status code
func (o *PostMalQueryExactSearchV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query exact search v1 unauthorized response a status code equal to that given
func (o *PostMalQueryExactSearchV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post mal query exact search v1 unauthorized response
func (o *PostMalQueryExactSearchV1Unauthorized) Code() int {
	return 401
}

func (o *PostMalQueryExactSearchV1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1Unauthorized  %+v", 401, o.Payload)
}

func (o *PostMalQueryExactSearchV1Unauthorized) String() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1Unauthorized  %+v", 401, o.Payload)
}

func (o *PostMalQueryExactSearchV1Unauthorized) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *PostMalQueryExactSearchV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostMalQueryExactSearchV1Forbidden creates a PostMalQueryExactSearchV1Forbidden with default headers values
func NewPostMalQueryExactSearchV1Forbidden() *PostMalQueryExactSearchV1Forbidden {
	return &PostMalQueryExactSearchV1Forbidden{}
}

/*
PostMalQueryExactSearchV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostMalQueryExactSearchV1Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this post mal query exact search v1 forbidden response has a 2xx status code
func (o *PostMalQueryExactSearchV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query exact search v1 forbidden response has a 3xx status code
func (o *PostMalQueryExactSearchV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query exact search v1 forbidden response has a 4xx status code
func (o *PostMalQueryExactSearchV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post mal query exact search v1 forbidden response has a 5xx status code
func (o *PostMalQueryExactSearchV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query exact search v1 forbidden response a status code equal to that given
func (o *PostMalQueryExactSearchV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post mal query exact search v1 forbidden response
func (o *PostMalQueryExactSearchV1Forbidden) Code() int {
	return 403
}

func (o *PostMalQueryExactSearchV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1Forbidden  %+v", 403, o.Payload)
}

func (o *PostMalQueryExactSearchV1Forbidden) String() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1Forbidden  %+v", 403, o.Payload)
}

func (o *PostMalQueryExactSearchV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *PostMalQueryExactSearchV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostMalQueryExactSearchV1TooManyRequests creates a PostMalQueryExactSearchV1TooManyRequests with default headers values
func NewPostMalQueryExactSearchV1TooManyRequests() *PostMalQueryExactSearchV1TooManyRequests {
	return &PostMalQueryExactSearchV1TooManyRequests{}
}

/*
PostMalQueryExactSearchV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PostMalQueryExactSearchV1TooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this post mal query exact search v1 too many requests response has a 2xx status code
func (o *PostMalQueryExactSearchV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query exact search v1 too many requests response has a 3xx status code
func (o *PostMalQueryExactSearchV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query exact search v1 too many requests response has a 4xx status code
func (o *PostMalQueryExactSearchV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post mal query exact search v1 too many requests response has a 5xx status code
func (o *PostMalQueryExactSearchV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query exact search v1 too many requests response a status code equal to that given
func (o *PostMalQueryExactSearchV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the post mal query exact search v1 too many requests response
func (o *PostMalQueryExactSearchV1TooManyRequests) Code() int {
	return 429
}

func (o *PostMalQueryExactSearchV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *PostMalQueryExactSearchV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *PostMalQueryExactSearchV1TooManyRequests) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryExactSearchV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMalQueryExactSearchV1InternalServerError creates a PostMalQueryExactSearchV1InternalServerError with default headers values
func NewPostMalQueryExactSearchV1InternalServerError() *PostMalQueryExactSearchV1InternalServerError {
	return &PostMalQueryExactSearchV1InternalServerError{}
}

/*
PostMalQueryExactSearchV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostMalQueryExactSearchV1InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this post mal query exact search v1 internal server error response has a 2xx status code
func (o *PostMalQueryExactSearchV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query exact search v1 internal server error response has a 3xx status code
func (o *PostMalQueryExactSearchV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query exact search v1 internal server error response has a 4xx status code
func (o *PostMalQueryExactSearchV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post mal query exact search v1 internal server error response has a 5xx status code
func (o *PostMalQueryExactSearchV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post mal query exact search v1 internal server error response a status code equal to that given
func (o *PostMalQueryExactSearchV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post mal query exact search v1 internal server error response
func (o *PostMalQueryExactSearchV1InternalServerError) Code() int {
	return 500
}

func (o *PostMalQueryExactSearchV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PostMalQueryExactSearchV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] postMalQueryExactSearchV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PostMalQueryExactSearchV1InternalServerError) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryExactSearchV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMalQueryExactSearchV1Default creates a PostMalQueryExactSearchV1Default with default headers values
func NewPostMalQueryExactSearchV1Default(code int) *PostMalQueryExactSearchV1Default {
	return &PostMalQueryExactSearchV1Default{
		_statusCode: code,
	}
}

/*
PostMalQueryExactSearchV1Default describes a response with status code -1, with default header values.

OK
*/
type PostMalQueryExactSearchV1Default struct {
	_statusCode int

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this post mal query exact search v1 default response has a 2xx status code
func (o *PostMalQueryExactSearchV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post mal query exact search v1 default response has a 3xx status code
func (o *PostMalQueryExactSearchV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post mal query exact search v1 default response has a 4xx status code
func (o *PostMalQueryExactSearchV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post mal query exact search v1 default response has a 5xx status code
func (o *PostMalQueryExactSearchV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post mal query exact search v1 default response a status code equal to that given
func (o *PostMalQueryExactSearchV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post mal query exact search v1 default response
func (o *PostMalQueryExactSearchV1Default) Code() int {
	return o._statusCode
}

func (o *PostMalQueryExactSearchV1Default) Error() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] PostMalQueryExactSearchV1 default  %+v", o._statusCode, o.Payload)
}

func (o *PostMalQueryExactSearchV1Default) String() string {
	return fmt.Sprintf("[POST /malquery/queries/exact-search/v1][%d] PostMalQueryExactSearchV1 default  %+v", o._statusCode, o.Payload)
}

func (o *PostMalQueryExactSearchV1Default) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryExactSearchV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
