package config

import (
	"github.com/caarlos0/env/v8"
)

// nolint:maligned
type Config struct {
	ENV          string `env:"ENV" envDefault:""`
	ImageTag     string `env:"IMAGE_TAG" envDefault:""`
	Version      string `env:"CI_COMMIT_TAG" envDefault:""`
	Port         int    `env:"VP_SERVICE_PORT" envDefault:"9000"`
	ExternalPort int    `env:"VP_EXTERNAL_PORT" envDefault:"9000"`
	ExternalHost string `env:"VP_EXTERNAL_HOST"`

	KafkaBrokers   string `env:"KAFKA_BROKERS" envDefault:"kafka:9090"`
	KafkaCoreTopic string `env:"KAFKA_CORE_EVENTS_TOPIC" envDefault:"core.events"`

	StatRealTimeDataTopic string `env:"STAT_REAL_TIME_DATA_TOPIC" envDefault:"stat.realtime"`
}

func ReadConfig() (*Config, error) {
	//nolint:exhaustivestruct
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
