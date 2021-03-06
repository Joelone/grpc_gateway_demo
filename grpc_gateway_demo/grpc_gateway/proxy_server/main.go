package main

import (
	"context"
	"flag"

	"github.com/golang/glog"
	"github.com/npuichigo/grpc_gateway_demo/grpc_gateway/gateway"
)

var (
	address = flag.String("address", "0.0.0.0:8080", "http port to listen on for proxy server")
	endpoint = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service")
	network = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
)

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	opts := gateway.Options{
		Addr: *address,
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
	}
	if err := gateway.Run(ctx, opts); err != nil {
		glog.Fatal(err)
	}
}
