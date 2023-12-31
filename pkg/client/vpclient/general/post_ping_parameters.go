// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/lifebackend/vp/pkg/client/vp"
)

// NewPostPingParams creates a new PostPingParams object
// with the default values initialized.
func NewPostPingParams() *PostPingParams {
	var ()
	return &PostPingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPingParamsWithTimeout creates a new PostPingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPingParamsWithTimeout(timeout time.Duration) *PostPingParams {
	var ()
	return &PostPingParams{

		timeout: timeout,
	}
}

// NewPostPingParamsWithContext creates a new PostPingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPingParamsWithContext(ctx context.Context) *PostPingParams {
	var ()
	return &PostPingParams{

		Context: ctx,
	}
}

// NewPostPingParamsWithHTTPClient creates a new PostPingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPingParamsWithHTTPClient(client *http.Client) *PostPingParams {
	var ()
	return &PostPingParams{
		HTTPClient: client,
	}
}

/*PostPingParams contains all the parameters to send to the API endpoint
for the post ping operation typically these are written to a http.Request
*/
type PostPingParams struct {

	/*Body*/
	Body *vp.PingMessage

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ping params
func (o *PostPingParams) WithTimeout(timeout time.Duration) *PostPingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ping params
func (o *PostPingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ping params
func (o *PostPingParams) WithContext(ctx context.Context) *PostPingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ping params
func (o *PostPingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ping params
func (o *PostPingParams) WithHTTPClient(client *http.Client) *PostPingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ping params
func (o *PostPingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post ping params
func (o *PostPingParams) WithBody(body *vp.PingMessage) *PostPingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post ping params
func (o *PostPingParams) SetBody(body *vp.PingMessage) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
