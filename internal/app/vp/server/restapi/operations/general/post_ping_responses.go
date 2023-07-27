// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/lifebackend/vp/internal/vperror"
	"github.com/sirupsen/logrus"

	"github.com/lifebackend/vp/internal/app/vp/server/models"
)

// PostPingOKCode is the HTTP code returned for type PostPingOK
const PostPingOKCode int = 200

/*PostPingOK Successful Response

swagger:response postPingOK
*/
type PostPingOK struct {

	/*
	  In: Body
	*/
	Payload *models.PostPingResponse `json:"body,omitempty"`
}

// NewPostPingOKFunc is a type the create the response func
type NewPostPingOKFunc func() *PostPingOK

// NewPostPingOK creates PostPingOK with default headers values
func NewPostPingOK() *PostPingOK {

	return &PostPingOK{}
}

// WithPayload adds the payload to the post ping o k response
func (o *PostPingOK) WithPayload(payload *models.PostPingResponse) *PostPingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post ping o k response
func (o *PostPingOK) SetPayload(payload *models.PostPingResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			logrus.Panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPingInternalServerErrorCode is the HTTP code returned for type PostPingInternalServerError
const PostPingInternalServerErrorCode int = 500

/*PostPingInternalServerError Internal Server Error

swagger:response postPingInternalServerError
*/
type PostPingInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewPostPingInternalServerErrorFunc is a type the create the response func
type NewPostPingInternalServerErrorFunc func() *PostPingInternalServerError

// NewPostPingInternalServerError creates PostPingInternalServerError with default headers values
func NewPostPingInternalServerError() *PostPingInternalServerError {

	return &PostPingInternalServerError{}
}

// WithPayload adds the payload to the post ping internal server error response
func (o *PostPingInternalServerError) WithPayload(payload *models.ErrorMessage) *PostPingInternalServerError {
	o.Payload = payload
	return o
}

// WithErr adds the Error payload with a default code to the post ping internal server error response
func (o *PostPingInternalServerError) FromErr(err error) *PostPingInternalServerError {
	type swaggerErr interface {
		Plain() (code string, message string, attributes map[string]string)
	}

	if swaggerErr, ok := err.(swaggerErr); ok {
		code, message, attributes := swaggerErr.Plain()

		o.Payload = &models.ErrorMessage{
			Code:       code,
			Message:    message,
			Attributes: attributes,
		}
		return o
	}

	o.Payload = &models.ErrorMessage{
		Code:       "InternalServiceError",
		Message:    err.Error(),
		Attributes: nil,
	}

	return o
}

// WithError adds the Error payload to the post ping internal server error response
func (o *PostPingInternalServerError) FromMessage(gaemblaErr *vperror.AppMessage) *PostPingInternalServerError {
	o.Payload = &models.ErrorMessage{
		Attributes: gaemblaErr.Attributes,
		Code:       gaemblaErr.Code,
		Message:    gaemblaErr.Message,
	}
	return o
}

// SetPayload sets the payload to the post ping internal server error response
func (o *PostPingInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPingInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			logrus.Panic(err) // let the recovery middleware deal with this
		}
	}
}
