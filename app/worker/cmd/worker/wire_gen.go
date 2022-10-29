// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"harvest/app/worker/internal/biz"
	"harvest/app/worker/internal/conf"
	"harvest/app/worker/internal/data"
	"harvest/app/worker/internal/server"
	"harvest/app/worker/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	equipRepo := data.NewEquipRepo(dataData, logger)
	equipUseCase := biz.NewEquipUseCase(equipRepo, logger)
	equipmentService := service.NewEquipmentService(equipUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, equipmentService, logger)
	httpServer := server.NewHTTPServer(confServer, equipmentService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
