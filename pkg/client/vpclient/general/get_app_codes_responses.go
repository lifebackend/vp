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

// GetAppCodesReader is a Reader for the GetAppCodes structure.
type GetAppCodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppCodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppCodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAppCodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /app-codes returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetAppCodesOK creates a GetAppCodesOK with default headers values
func NewGetAppCodesOK() *GetAppCodesOK {
	return &GetAppCodesOK{}
}

/*GetAppCodesOK handles this case with default header values.

Successful Response
*/
type GetAppCodesOK struct {
	Payload vp.GetAppMessagesResponse
}

func (o *GetAppCodesOK) Error() string {
	return fmt.Sprintf("[GET /app-codes][%d] getAppCodesOK  %+v", 200, o.Payload)
}

func (o *GetAppCodesOK) GetPayload() vp.GetAppMessagesResponse {
	return o.Payload
}

func (o *GetAppCodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppCodesInternalServerError creates a GetAppCodesInternalServerError with default headers values
func NewGetAppCodesInternalServerError() *GetAppCodesInternalServerError {
	return &GetAppCodesInternalServerError{}
}

/*GetAppCodesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAppCodesInternalServerError struct {
	Payload *vp.ErrorMessage
}

func (o *GetAppCodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /app-codes][%d] getAppCodesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppCodesInternalServerError) GetPayload() *vp.ErrorMessage {
	return o.Payload
}

func (o *GetAppCodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vp.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
