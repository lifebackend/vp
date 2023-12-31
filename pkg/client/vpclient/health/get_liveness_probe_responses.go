// Code generated by go-swagger; DO NOT EDIT.

package health

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

// GetLivenessProbeReader is a Reader for the GetLivenessProbe structure.
type GetLivenessProbeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLivenessProbeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLivenessProbeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetLivenessProbeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /_livenessProbe returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetLivenessProbeOK creates a GetLivenessProbeOK with default headers values
func NewGetLivenessProbeOK() *GetLivenessProbeOK {
	return &GetLivenessProbeOK{}
}

/*GetLivenessProbeOK handles this case with default header values.

Successful Response
*/
type GetLivenessProbeOK struct {
	Payload *vp.LivenessProbe
}

func (o *GetLivenessProbeOK) Error() string {
	return fmt.Sprintf("[GET /_livenessProbe][%d] getLivenessProbeOK  %+v", 200, o.Payload)
}

func (o *GetLivenessProbeOK) GetPayload() *vp.LivenessProbe {
	return o.Payload
}

func (o *GetLivenessProbeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vp.LivenessProbe)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLivenessProbeInternalServerError creates a GetLivenessProbeInternalServerError with default headers values
func NewGetLivenessProbeInternalServerError() *GetLivenessProbeInternalServerError {
	return &GetLivenessProbeInternalServerError{}
}

/*GetLivenessProbeInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetLivenessProbeInternalServerError struct {
}

func (o *GetLivenessProbeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /_livenessProbe][%d] getLivenessProbeInternalServerError ", 500)
}

func (o *GetLivenessProbeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
