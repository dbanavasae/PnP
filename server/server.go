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

/*func init() {
	flag.StringVar(&consulAddress, "consulIP", "127.0.0.1",
		"Consul server's IP address at which gRPC server")
	flag.StringVar(&pnpServerName, "PnPServerName", "pnpService",
		"PnP server's registered service name in Consul")
	flag.Parse()
}*/
func main() {
	//os.Setenv("MICRO_REGISTRY_ADDRESS", "192.168.50.129")

	service := grpc.NewService(
		micro.Name("PnPServer"),
		micro.Flags(
			cli.StringFlag{
				Name : "PackageFilePath",
				Value: "/config/packageInfo.json",
				Usage: "Path of packageInfo.json file",
			},
		),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			config.PackageFilePath = c.String("PackageFilePath")
		}),
	)

	proto.RegisterPnPHandler(service.Server(), new(handler.PnPService))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
