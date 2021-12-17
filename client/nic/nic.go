package nic

import (
	structs "jonatann-silva/network/dsn/structs"
)

// NetworkInterfaceController provides network configuration capabilities.
type NetworkInterfaceController interface {
	Interfaces() ([]*structs.Interface, error)
	CreateInterface(iface *structs.Interface) error
	UpdateInterface(iface *structs.Interface) error
	DeleteInterfaceByAlias(s string) error
	DeleteInterfaceByName(s string) error
	DeleteAllInterfaces() error
}

type PrivateKeyStore interface {
	KeyByID(id string) (*PrivateKey, error)
	UpsertKey(key *PrivateKey) error
	DeleteKey(id string) error
}

type PrivateKey struct {
	ID        string
	Key       string
	CreatedAt int64
}
