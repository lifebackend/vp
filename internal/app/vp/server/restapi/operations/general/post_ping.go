// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/pkg/scope"
	prometheusmetrics "github.com/lifebackend/vp/tools/prometheus-metrics"
	"github.com/sirupsen/logrus"
)

func NewPostPingResponses() *PostPingResponses {
	return &PostPingResponses{
		PostPingOK:                  NewPostPingOK,
		PostPingInternalServerError: NewPostPingInternalServerError,
	}
}

type PostPingResponses struct {
	PostPingOK                  NewPostPingOKFunc
	PostPingInternalServerError NewPostPingInternalServerErrorFunc
}

// PostPingHandlerFunc turns a function with the right signature into a post ping handler
type PostPingHandlerFunc func(*PostPingParams, *PostPingResponses) middleware.Responder

// Handle executing the request and returning a response
func (fn PostPingHandlerFunc) Handle(params *PostPingParams, respond *PostPingResponses) middleware.Responder {
	return fn(params, respond)
}

// PostPingHandler interface for that can handle valid post ping params
type PostPingHandler interface {
	Handle(*PostPingParams, *PostPingResponses) middleware.Responder
}

// NewPostPing creates a new http.Handler for the post ping operation
func NewPostPing(ctx *middleware.Context, handler PostPingHandler) *PostPing {
	return &PostPing{Context: ctx, Handler: handler}
}

/*PostPing swagger:route POST /ping general postPing

PostPing post ping API

*/
type PostPing struct {
	Context *middleware.Context
	Handler PostPingHandler
}

func (o *PostPing) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostPingParams()

	if err := o.Context.BindValidRequest(r, route, Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	// pass predefined values from middleware
	requestCtx := r.Context()
	logger := requestCtx.Value("logger").(*logrus.Entry)
	xCorrelationID := requestCtx.Value("xCorrelationID").(string)

	// pass body
	Params.RequestBody = requestCtx.Value("body").([]byte)

	metrics := requestCtx.Value("metrics").(*prometheusmetrics.Metrics)

	// prepare scope
	Params.Scope = scope.NewScope(requestCtx, logger, xCorrelationID, metrics)
	Params.Scope.SetResource("POST /ping")

	defer func() {
		if rec := recover(); rec != nil {
			Params.Scope.Logger().Errorf("%s: %s", rec, debug.Stack())

			xCorrelationID, ok := requestCtx.Value("xCorrelationID").(string)
			if ok {
				rw.Header().Add("X-Correlation-Id", xCorrelationID)
			}

			rw.WriteHeader(http.StatusInternalServerError)

			o.Context.Respond(rw, r, route.Produces, route, json.RawMessage([]byte(`{"code":"panic","message":""}`)))
		}
	}()

	responses := NewPostPingResponses()

	started := time.Now()

	res := o.Handler.Handle(Params, responses) // actually handle the request

	if metrics != nil {
		metrics.AddAPIResponseDuration("POST /ping", time.Now().Sub(started))
	}

	o.Context.Respond(rw, r, route.Produces, route, res)

}
