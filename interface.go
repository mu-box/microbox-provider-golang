package provider

import (
)

type BackendAdaptor interface{
	Meta() Metadata
	Catalog() ([]ServerOption, error)
	Verify(Credentials) (bool, error)
	AddKey(Credentials, KeyOrder) (Key, error)
	ListKeys(Credentials) ([]Key, error)
	ShowKey(Credentials, string) (Key, error)
	DeleteKey(Credentials, string) error
	AddServer(Credentials, ServerOrder) (Server, error)
	ListServers(Credentials) ([]Server, error)
	ShowServer(Credentials, string) (Server, error)
	DeleteServer(Credentials, string) error
	RebootServer(Credentials, string) error
	RestartServer(Credentials, string) error
}