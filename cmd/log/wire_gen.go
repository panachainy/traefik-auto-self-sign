// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package log

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

// Injectors from wire.go:

func Wire() (*logrus.Logger, error) {
	logger := Provide()
	return logger, nil
}

// wire.go:

var ProviderSet = wire.NewSet(
	Provide,
)