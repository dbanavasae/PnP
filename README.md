# PnP

## Prerequisites

1. go lang: version > 1.9
2. Set env variable GOPATH
3. Run `$ go get "github.com/micro/go-micro"`
4. Run `$ go get "github.com/micro/go-grpc"`
5. Running instance of Consul.
6. Generate the server certificate and key file: 
   
   6.1. Go to `'../PnP/util/'` folder, and run the `GenerateTLSCertificate.go`. This generates the `server.crt` & `server.key` files in `../PnP/certs` folder.
    
   6.2. Transfer these files to Client machine in the folder: `'../PnP/certs'`.

Note: To run PnP server and client you should be a root user

## Running PnP Server

`$ go run server.go --registry_address=<consul_ip> --server_name=<pnp_server_name> --package_file=<path/of/packageInfo.json> --pnp_op_type=<operation_type> --cert_file "../certs/server.crt" --key_file "../certs/server.key"`

e.g.: 
`$ go run server.go --package_file "/../config/packageInfo.json" --registry_address "192.168.50.129" --server_name "NewPnPService" --pnp_op_type="installPackages" --server_cert_file "../certs/server.crt"`

`packageInfo.json` recides in config directory.

## Running PnP client

`$ go run client.go --registry_address=<consul_ip> --pnp_server=<pnp server name registered to consul> --pnp_op_type=<operation_type> --server_cert_file <path_of_server_cert>`

e.g.: 
`$ go run client.go --registry_address="192.168.50.129" --pnp_server="NewPnPService" --pnp_op_type="installPackages" --server_cert_file "../certs/server.crt"`
