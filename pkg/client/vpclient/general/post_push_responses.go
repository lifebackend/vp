// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lifebackend/vp/pkg/client/vp"
)

// PostPushReader is a Reader for the PostPush structure.
type PostPushReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPushReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPushOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostPushInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /push returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPostPushOK creates a PostPushOK with default headers values
func NewPostPushOK() *PostPushOK {
	return &PostPushOK{}
}

/*PostPushOK handles this case with default header values.

Successful Response
*/
type PostPushOK struct {
	Payload *vp.PostMessageResponse
}

func (o *PostPushOK) Error() string {
	return fmt.Sprintf("[POST /push][%d] postPushOK  %+v", 200, o.Payload)
}

func (o *PostPushOK) GetPayload() *vp.PostMessageResponse {
	return o.Payload
}

func (o *PostPushOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vp.PostMessageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPushInternalServerError creates a PostPushInternalServerError with default headers values
func NewPostPushInternalServerError() *PostPushInternalServerError {
	return &PostPushInternalServerError{}
}

/*PostPushInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostPushInternalServerError struct {
	Payload *vp.ErrorMessage
}

func (o *PostPushInternalServerError) Error() string {
	return fmt.Sprintf("[POST /push][%d] postPushInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPushInternalServerError) GetPayload() *vp.ErrorMessage {
	return o.Payload
}

func (o *PostPushInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vp.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
