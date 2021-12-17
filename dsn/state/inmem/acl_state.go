package inmem

import (
	"context"
	"errors"
	"fmt"

	structs "jonatann-silva/network/dsn/structs"
)

// ACLState :
func (r *StateRepository) ACLState(ctx context.Context) (*structs.ACLState, error) {
	key := aclStateKey()
	if v, found := r.kv.Get(key); found {
		return v.(*structs.ACLState), nil
	}
	return nil, errors.New("not found")
}

// ACLSetState :
func (r *StateRepository) ACLSetState(ctx context.Context, s *structs.ACLState) error {
	key := aclStateKey()
	r.kv.Set(key, s)
	return nil
}

func aclStateKey() string {
	return fmt.Sprintf("%s/%s/global/%s", defaultPrefix, "acl", "state")
}
