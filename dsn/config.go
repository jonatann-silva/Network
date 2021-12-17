package dsn

import (
	"time"

	"jonatann-silva/network/dsn/structs/config"
	log "jonatann-silva/network/pkg/log"
	version "jonatann-silva/network/version"
)

const (
	defaultBindAddress = "0.0.0.0"
	defaultLogLevel    = "DEBUG"
	defaultDataDir     = "/tmp/dsn"
	defaultHTTPPort    = 8080
	defaultRPCPort     = 8081
)

// Config : DSN server configuration.
type Config struct {
	// UI enabled.
	UI bool

	// DevMode enabled.
	DevMode bool

	// Version is the version of the DSN server
	Version *version.VersionInfo

	// LogLevel is the level at which the server should output logs
	LogLevel string

	//Logger.
	Logger log.Logger

	// BindAddr.
	BindAddr string

	// RPCAdvertiseAddr is the address advertised to client nodes.
	RPCAdvertiseAddr string

	// DataDir is the directory to store our state in.
	DataDir string

	// Ports.
	Ports *Ports

	// ACL
	ACL *config.ACLConfig

	// Etcd.
	Etcd *config.EtcdConfig

	// HostGCInterval is how often we perform garbage collection of hosts.
	HostGCInterval time.Duration
}

// Ports :
type Ports struct {
	HTTP int
	RPC  int
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	return &Config{
		UI:       true,
		DevMode:  false,
		Version:  version.GetVersion(),
		LogLevel: defaultLogLevel,
		BindAddr: defaultBindAddress,
		DataDir:  defaultDataDir,
		Ports: &Ports{
			HTTP: defaultHTTPPort,
			RPC:  defaultRPCPort,
		},
		ACL:            config.DefaultACLConfig(),
		Etcd:           config.DefaultEtcdConfig(),
		HostGCInterval: 5 * time.Minute,
	}
}
