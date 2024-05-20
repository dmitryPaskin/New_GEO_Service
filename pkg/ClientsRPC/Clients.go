package ClientsRPC

import (
	"GeoServiseAppDate/internal/service"
	"fmt"
	"os"
)

type FactoryRPC interface {
	service.Service
}

func GetRPCProtocol() (FactoryRPC, error) {
	protocol := os.Getenv("RPC_PROTOCOL")

	switch protocol {
	case "ServersRPC":
		return &RPC{}, nil
	case "JSON-ServersRPC":
		return &JSONRPC{}, nil
		//	case "gRPC":
		//		return
	}
	return nil, fmt.Errorf("%s", "protocol not found")
}
