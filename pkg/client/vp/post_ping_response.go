// Code generated by go-swagger; DO NOT EDIT.

package vp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostPingResponse PostPingResponse
//
// swagger:model PostPingResponse
type PostPingResponse struct {
	SuccessResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostPingResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SuccessResponse
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SuccessResponse = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostPingResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SuccessResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post ping response
func (m *PostPingResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this post ping response based on the context it is used
func (m *PostPingResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SuccessResponse
	if err := m.SuccessResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
