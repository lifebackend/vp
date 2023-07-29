// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/lifebackend/vp/internal/vperror"
	"github.com/sirupsen/logrus"

	"github.com/lifebackend/vp/internal/app/vp/server/models"
)

// GetAppUpdateVersionOKCode is the HTTP code returned for type GetAppUpdateVersionOK
const GetAppUpdateVersionOKCode int = 200

/*GetAppUpdateVersionOK Successful Response

swagger:response getAppUpdateVersionOK
*/
type GetAppUpdateVersionOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetAppUpdateVersionOKFunc is a type the create the response func
type NewGetAppUpdateVersionOKFunc func() *GetAppUpdateVersionOK

// NewGetAppUpdateVersionOK creates GetAppUpdateVersionOK with default headers values
func NewGetAppUpdateVersionOK() *GetAppUpdateVersionOK {

	return &GetAppUpdateVersionOK{}
}

// WithPayload adds the payload to the get app update version o k response
func (o *GetAppUpdateVersionOK) WithPayload(payload io.ReadCloser) *GetAppUpdateVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app update version o k response
func (o *GetAppUpdateVersionOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppUpdateVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		logrus.Panic(err) // let the recovery middleware deal with this
	}
}

// GetAppUpdateVersionInternalServerErrorCode is the HTTP code returned for type GetAppUpdateVersionInternalServerError
const GetAppUpdateVersionInternalServerErrorCode int = 500

/*GetAppUpdateVersionInternalServerError Internal Server Error

swagger:response getAppUpdateVersionInternalServerError
*/
type GetAppUpdateVersionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetAppUpdateVersionInternalServerErrorFunc is a type the create the response func
type NewGetAppUpdateVersionInternalServerErrorFunc func() *GetAppUpdateVersionInternalServerError

// NewGetAppUpdateVersionInternalServerError creates GetAppUpdateVersionInternalServerError with default headers values
func NewGetAppUpdateVersionInternalServerError() *GetAppUpdateVersionInternalServerError {

	return &GetAppUpdateVersionInternalServerError{}
}

// WithPayload adds the payload to the get app update version internal server error response
func (o *GetAppUpdateVersionInternalServerError) WithPayload(payload *models.ErrorMessage) *GetAppUpdateVersionInternalServerError {
	o.Payload = payload
	return o
}

// WithErr adds the Error payload with a default code to the get app update version internal server error response
func (o *GetAppUpdateVersionInternalServerError) FromErr(err error) *GetAppUpdateVersionInternalServerError {
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

// WithError adds the Error payload to the get app update version internal server error response
func (o *GetAppUpdateVersionInternalServerError) FromMessage(gaemblaErr *vperror.AppMessage) *GetAppUpdateVersionInternalServerError {
	o.Payload = &models.ErrorMessage{
		Attributes: gaemblaErr.Attributes,
		Code:       gaemblaErr.Code,
		Message:    gaemblaErr.Message,
	}
	return o
}

// SetPayload sets the payload to the get app update version internal server error response
func (o *GetAppUpdateVersionInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppUpdateVersionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			logrus.Panic(err) // let the recovery middleware deal with this
		}
	}
}
