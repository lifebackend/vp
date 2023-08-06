// Code generated by go-swagger; DO NOT EDIT.

package health

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

func NewGetLivenessProbeResponses() *GetLivenessProbeResponses {
	return &GetLivenessProbeResponses{
		GetLivenessProbeOK:                  NewGetLivenessProbeOK,
		GetLivenessProbeInternalServerError: NewGetLivenessProbeInternalServerError,
	}
}

type GetLivenessProbeResponses struct {
	GetLivenessProbeOK                  NewGetLivenessProbeOKFunc
	GetLivenessProbeInternalServerError NewGetLivenessProbeInternalServerErrorFunc
}

// GetLivenessProbeHandlerFunc turns a function with the right signature into a get liveness probe handler
type GetLivenessProbeHandlerFunc func(*GetLivenessProbeParams, *GetLivenessProbeResponses) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLivenessProbeHandlerFunc) Handle(params *GetLivenessProbeParams, respond *GetLivenessProbeResponses) middleware.Responder {
	return fn(params, respond)
}

// GetLivenessProbeHandler interface for that can handle valid get liveness probe params
type GetLivenessProbeHandler interface {
	Handle(*GetLivenessProbeParams, *GetLivenessProbeResponses) middleware.Responder
}

// NewGetLivenessProbe creates a new http.Handler for the get liveness probe operation
func NewGetLivenessProbe(ctx *middleware.Context, handler GetLivenessProbeHandler) *GetLivenessProbe {
	return &GetLivenessProbe{Context: ctx, Handler: handler}
}

/*
GetLivenessProbe swagger:route GET /_livenessProbe health getLivenessProbe

# Liveness Probe

Liveness Probe
*/
type GetLivenessProbe struct {
	Context *middleware.Context
	Handler GetLivenessProbeHandler
}

func (o *GetLivenessProbe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLivenessProbeParams()

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
	Params.Scope.SetResource("GET /_livenessProbe")

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

	responses := NewGetLivenessProbeResponses()

	started := time.Now()

	res := o.Handler.Handle(Params, responses) // actually handle the request

	if metrics != nil {
		metrics.AddAPIResponseDuration("GET /_livenessProbe", time.Now().Sub(started))
	}

	o.Context.Respond(rw, r, route.Produces, route, res)

}
