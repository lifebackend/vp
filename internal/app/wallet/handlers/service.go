package handlers

import (
	"context"

	kafkasenderservice "github.com/lifebackend/vp/internal/app/common/services/kafka-sender-service"
	"github.com/lifebackend/vp/internal/app/wallet/config"
	"github.com/lifebackend/vp/internal/app/wallet/server/restapi"
	"github.com/lifebackend/vp/internal/app/wallet/server/restapi/operations"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"

	"github.com/lifebackend/vp/pkg/scope"
	prometheusmetrics "github.com/lifebackend/vp/tools/prometheus-metrics"
	"github.com/lifebackend/vp/tools/run"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer   *restapi.Server
	kafkaService *kafkasenderservice.KafkaSenderService
}

type Dependencies struct{}

func PrepareServer(scope *scope.Scope, cfg *config.Config, serviceName string, logger *logrus.Entry) (*Server, *Handlers, error) {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, nil, err
	}

	logger.Info("Initializing services...")

	kafkaService, err := kafkasenderservice.NewKafkaSender(cfg)
	if err != nil {
		return nil, nil, err
	}
	//eventSendingService := eventsendingservice.NewEventSendingService(kafkaService, cfg.KafkaCoreTopic, serviceName)

	h := NewHandlers(
		cfg.ImageTag,
	)

	logger.Info("Initializing API...")
	api := operations.NewWalletAPI(swaggerSpec)

	// add API routes
	AddDefaultHandlers(api, h)

	api.UseSwaggerUI()
	api.SetDefaultProduces("application/json")
	api.SetDefaultConsumes("application/json")
	api.HTMLProducer = runtime.TextProducer()

	logger.Info("Initializing integrations")

	server := restapi.NewServerWithMiddleware(api, serviceName, logger, prometheusmetrics.NewMetrics())
	server.Port = cfg.Port
	server.EnabledListeners = []string{"http"}

	return &Server{
		httpServer:   server,
		kafkaService: kafkaService,
	}, h, nil
}

func (s *Server) Serve(scope *scope.Scope) error {
	group := run.NewNamedGroup(scope.Ctx, "core")

	//group.AddWithContextNamed("taskService", func(ctx context.Context) error {
	//	return s.taskService.Start(scope.ForkWithCtx(ctx))
	//})

	group.AddWithContextNamed("swagger", func(ctx context.Context) error {
		go func() {
			<-ctx.Done()

			// nolint:errcheck
			_ = s.httpServer.Shutdown()
		}()

		return s.httpServer.Serve()
	})

	return group.Run()
}

func (s *Server) GetHost() string {
	return s.httpServer.Host
}

func (s *Server) GetPort() int {
	return s.httpServer.Port
}
