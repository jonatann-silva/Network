package dsn

import (
	auth "jonatann-silva/network/dsn/auth"
	state "jonatann-silva/network/dsn/state"
	structs "jonatann-silva/network/dsn/structs"
)

// StatusService is used to check on server status
type StatusService struct {
	config      *Config
	state       state.Repository
	authHandler auth.AuthorizationHandler
}

// NewStatusService ...
func NewStatusService(config *Config, state state.Repository, authHandler auth.AuthorizationHandler) *StatusService {
	return &StatusService{
		config:      config,
		state:       state,
		authHandler: authHandler,
	}
}

// Ping is used to check for connectivity
func (s *StatusService) Ping(args structs.GenericRequest, out *structs.GenericResponse) error {
	return nil
}

// Version returns the version of the server
func (s *StatusService) Version(in structs.GenericRequest, out *structs.StatusVersionResponse) error {
	out.Version = s.config.Version.VersionNumber()
	return nil
}
