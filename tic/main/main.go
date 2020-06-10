package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/odom11/playground_micro/api"
	"github.com/odom11/playground_micro/tic"
)

func main() {
	registry := etcdv3.NewRegistry()
	service := micro.NewService(
		micro.Name("tic"),
		micro.Version("latest"),
		micro.Registry(registry),
	)

	toc := micro.NewService()
	toc.Init()

	service.Init()
	api.RegisterTicHandler(service.Server(), tic.New(api.NewTocService("toc", toc.Client())))
}