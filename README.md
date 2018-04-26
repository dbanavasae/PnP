# PnP

## Prerequisites

1. go lang: version > 1.9
2. Set env variable GOPATH
3. Run `$ go get "github.com/micro/go-micro"`
4. Run `$ go get "github.com/micro/go-grpc"`
5. Running instance of Consul

Note: To run PnP server and client you should be a root user

## Running PnP Server

`$ go run server.go --registry_address=<consul_ip> --server_name=<pnp_server_name> --package_file_path=<path/of/packageInfo.json>`

e.g.: 
`$ go run server.go --package_file_path "/../config/packageInfo.json" --registry_address "192.168.50.129" --server_name "NewPnPService"`

`packageInfo.json` recides in config directory.

## Running PnP client

`$ go run client.go --registry_address=<consul_ip> --pnp_server=<pnp server name registered to consul>`

e.g.: 
`$ go run client.go --registry_address="192.168.50.129" --pnp_server="NewPnPService"`
