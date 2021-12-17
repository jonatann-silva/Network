package dsn

import (
	"context"
	"time"

	auth "jonatann-silva/network/dsn/auth"
	state "jonatann-silva/network/dsn/state"
	structs "jonatann-silva/network/dsn/structs"
	log "jonatann-silva/network/pkg/log"
	uuid "jonatann-silva/network/pkg/uuid"
)

const (
	// aclBootstrapResetFileName is the name of the file in the data dir containing the reset index.
	aclBootstrapResetFileName = "acl-bootstrap-reset"
)

// ACLService :
type ACLService struct {
	config      *Config
	logger      log.Logger
	state       state.Repository
	authHandler auth.AuthorizationHandler
}

// NewACLService :
func NewACLService(config *Config, logger log.Logger, state state.Repository, authHandler auth.AuthorizationHandler) *ACLService {
	return &ACLService{
		config:      config,
		logger:      logger,
		state:       state,
		authHandler: authHandler,
	}
}

// BootstrapACL :
func (s *ACLService) BootstrapACL(args *structs.ACLBootstrapRequest, out *structs.ACLTokenUpsertResponse) error {

	if !s.config.ACL.Enabled {
		return structs.ErrACLDisabled
	}

	ctx := context.TODO()

	if !s.config.ACL.Enabled {
		return structs.ErrACLDisabled
	}

	if s.isBootstrapped(ctx) {
		return structs.ErrACLAlreadyBootstrapped
	}

	t := &structs.ACLToken{
		ID:        uuid.Generate(),
		Name:      "Root Token",
		Secret:    uuid.Generate(),
		Type:      structs.ACLTokenTypeManagement,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.state.UpsertACLToken(ctx, t)
	if err != nil {
		return err
	}

	state := s.aclStateLazy()
	state.RootTokenID = t.ID
	state.RootTokenResetIndex++

	// Update ACL state
	err = s.state.ACLSetState(ctx, state)
	if err != nil {
		return err
	}

	out.ACLToken = t

	return nil
}

// ResolveToken :
func (s *ACLService) ResolveToken(args *structs.ResolveACLTokenRequest, out *structs.ResolveACLTokenResponse) error {

	if !s.config.ACL.Enabled {
		return structs.ErrACLDisabled
	}

	t, err := s.state.ACLTokenBySecret(context.TODO(), args.Secret)
	if err != nil {
		return err
	}

	if t == nil {
		t = AnonymousACLToken
	}

	out.ACLToken = t

	return nil
}

func (s *ACLService) isBootstrapped(ctx context.Context) bool {
	return s.aclStateLazy().RootTokenID != ""
}

// ACLStateLazy implements lazy state persistence
func (s *ACLService) aclStateLazy() *structs.ACLState {
	ctx := context.TODO()
	aclState, err := s.state.ACLState(ctx)
	if err != nil {
		aclState = &structs.ACLState{}
		s.state.ACLSetState(ctx, aclState)
	}
	return aclState
}
