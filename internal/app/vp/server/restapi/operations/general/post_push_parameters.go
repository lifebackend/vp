// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/pkg/scope"
)

// NewPostPushParams creates a new PostPushParams object
// no default values defined in spec.
func NewPostPushParams() *PostPushParams {

	return &PostPushParams{}
}

// PostPushParams contains all the bound params for the post push operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostPush
type PostPushParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	// Body
	RequestBody []byte

	// Scope
	Scope *scope.Scope
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostPushParams() beforehand.
func (o *PostPushParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
