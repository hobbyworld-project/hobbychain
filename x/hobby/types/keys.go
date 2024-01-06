package types

const (
	// ModuleName defines the module name
	ModuleName = "hobby"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hobby"
)

const (
	PrivateDataKeyMaxSize   = 128
	PrivateDataValueMaxSize = 1024
)

var (
	PrivateDataPrefix = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
