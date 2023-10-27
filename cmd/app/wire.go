//go:build wireinject
// +build wireinject

//go:generate wire
package app

import (
	"panachainy/traefik-auto-self-sign/cmd/config"
	"panachainy/traefik-auto-self-sign/cmd/log"

	"github.com/google/wire"
)

func Wire() (*Application, error) {
	wire.Build(
		Provide,
		config.ProviderSet,
		log.ProviderSet,
	)

	return &Application{}, nil
}
