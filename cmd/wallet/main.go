package main

import (
	"context"
	"log"

	"github.com/evalphobia/logrus_sentry"
	"github.com/lifebackend/vp/pkg/scope"
	prometheusmetrics "github.com/lifebackend/vp/tools/prometheus-metrics"

	"github.com/lifebackend/vp/internal/app/wallet/config"
	"github.com/lifebackend/vp/internal/app/wallet/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetReportCaller(true)

	logger := logrus.NewEntry(logrus.StandardLogger())

	hook, err := logrus_sentry.NewAsyncSentryHook("", []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})

	if err == nil {
		hook.StacktraceConfiguration.Enable = true
		logger.Logger.Hooks.Add(hook)
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		logger.Fatalln(err)
	}

	sScope := scope.NewScope(context.Background(), logger, "", prometheusmetrics.NewMetrics())

	server, _, err := handlers.PrepareServer(sScope, cfg, "core", logger)
	if err != nil {
		logger.Fatalln(err)
	}

	if cfg.SentryEnabled {
		// DSN passing via SENTRY_ENV variable
		hook, err := logrus_sentry.NewAsyncSentryHook("", []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		})

		if err == nil {
			hook.StacktraceConfiguration.Enable = true
			logger.Logger.Hooks.Add(hook)
		}
	}

	sScope.Logger().Println("starting")

	if err := server.Serve(sScope); err != nil {
		log.Fatalln(err)
	}
}
