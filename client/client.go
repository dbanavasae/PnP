package main

import (
	"fmt"
	"log"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"github.com/micro/go-grpc"
	"github.com/PnP/invoke"
	"github.com/micro/go-micro"
	"github.com/micro/cli"
	"github.com/micro/go-micro/transport"
	proto "github.com/PnP/pnp-proto"
)

func main() {
	var pnpServer string
	var pnpOpType string
	var serverCert string

	service := grpc.NewService(
		micro.Flags(
			cli.StringFlag{
				Name : "pnp_server",
				Value: "PnPService",
				Usage: "PnP server name registered to registry",
			},
			cli.StringFlag{
				Name : "pnp_op_type",
				Usage: "Specifies pnp operation type, supported values are" +
					"installPackages, deployPlatform",
			},
			cli.StringFlag{
				Name : "server_cert_file",
				Value: "/certs/server.crt",
				Usage: "Path of server certificate file",
			},
		),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			pnpServer = c.String("pnp_server")
			pnpOpType = c.String("pnp_op_type")
			serverCert = c.String("server_cert_file")
			if pnpOpType == "" {
				log.Fatalf("PnP operation type not specified, supported values are" +
					"installPackages, deploySDPMaster, deploySDPSatellite")
			}
		}),
	)

	caCert, err := ioutil.ReadFile(serverCert)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	service.Init(
		micro.Transport(transport.NewTransport(transport.Secure(true))),
		grpc.WithTLS(tlsConfig),
	)

	pnpClient := proto.PnPServiceClient(pnpServer, service.Client())

	switch pnpOpType {
	case "installPackages":
		{
			fmt.Println("Initializing package installation..")
			invoke.InstallPackages(pnpClient)
		}
	case "deploySDPMaster":
		{
			fmt.Println("Requesting deployment of SDP Master from PnP server..")
			invoke.DeployPlatform(pnpClient, "master")
		}
	case "deploySDPSatellite":
		{
			fmt.Println("Requesting deployment of SDP Satellite from PnP server..")
			invoke.DeployPlatform(pnpClient, "satellite")
		}
	}
}
