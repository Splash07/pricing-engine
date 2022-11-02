package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	HttpServerPort string `default:"8080" envconfig:"HTTP_SERVER_PORT"`
	Env            string `default:"stg" envconfig:"ENV"`
	PProf          bool   `default:"true" envconfig:"PPROF"`
}

var config AppConfig

func init() {
	if err := envconfig.Process("ct-core-gds-delivery", &config); err != nil {
		log.Fatal(err.Error())
	}
}

func Load() AppConfig {
	return config
}
