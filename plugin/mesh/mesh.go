package main

import (
	"net/rpc"

	log "jonatann-silva/network/pkg/log"
)

// TODO: Implementation of the dsn.Plugin interface, with services exposed through RPC
// Whenever a new interface is created (event InterfaceCreated), this plugin interacts with the
// DSN RPC API in order to Link this interface with every other publicly exposed interface in
// the same network, thus creating a mesh overlay.

// MeshPlugin :
type MeshPlugin struct {
	logger    log.Logger
	rpcServer *rpc.Server
}

// Config :
type Config struct {
}

// NewMeshPlugin : Creates a new mesh plugin object parameterized according to the provided configurations.
func NewMeshPlugin(config *Config) (*MeshPlugin, error) {
	p := &MeshPlugin{}
	return p, nil
}

func main() {
	_, err := NewMeshPlugin(&Config{})
	if err != nil {
		panic(err)
	}
}
