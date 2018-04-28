package main

import (
	"log"
	"time"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/PnP/config"
	handler "github.com/PnP/handlers"
	proto "github.com/PnP/pnp-proto"
	"github.com/micro/cli"
)

func main() {
	service := grpc.NewService(
		micro.Name("PnPServer"),
		micro.Flags(
			cli.StringFlag{
				Name : "package_file",
				Value: "/config/packageInfo.json",
				Usage: "Path of packageInfo.json file",
			},
			cli.StringFlag{
				Name : "sdp_deploy_file",
				Value: "/config/platform-config.json",
				Usage: "Path of sdp platform deploy config json file",
			},
		),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*5),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			config.PackageFilePath = c.String("package_file")
			config.PlatformDeployFile = c.String("sdp_deploy_file")
		}),
	)

	proto.RegisterPnPHandler(service.Server(), new(handler.PnPService))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
