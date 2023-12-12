package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/tonkeeper/tongo/ton"
)

var ElectorAccountID = ton.MustParseAccountID("Ef8zMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzM0vF")

var Config = struct {
	Dev         string `env:"DEV"`
	Port        int    `env:"PORT" envDefault:"8888"`
	MetricsPort int    `env:"METRICS_PORT" envDefault:"9010"`
}{}

func LoadConfig() {
	if err := env.Parse(&Config); err != nil {
		log.Fatalf("❗️failed to parse config: %v\n", err)
	}
}
