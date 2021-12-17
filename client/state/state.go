package state

import (
	"jonatann-silva/network/client/nic"
	"jonatann-silva/network/dsn/structs"
)

// Repository :
type Repository interface {
	Name() string

	// Client state
	Interfaces() ([]*structs.Interface, error)
	UpsertInterface(*structs.Interface) error
	DeleteInterfaces(id []string) error

	// Key store
	KeyByID(id string) (*nic.PrivateKey, error)
	UpsertKey(key *nic.PrivateKey) error
	DeleteKey(id string) error
}
