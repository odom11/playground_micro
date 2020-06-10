package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/odom11/playground_micro/api"
	"github.com/odom11/playground_micro/toc"
)

func main() {

	registry := etcdv3.NewRegistry()
	service := micro.NewService(
		micro.Name("tic"),
		micro.Version("latest"),
		micro.Registry(registry),
	)

	tic := micro.NewService()
	tic.Init()

	service.Init()
	api.RegisterTocHandler(service.Server(), toc.New(api.NewTicService("toc", tic.Client())))
}
