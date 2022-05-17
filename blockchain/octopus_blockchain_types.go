package blockchain

//定义hash和地址长度byte
const (
	// hash长度
	HashLength = 32
	// 地址值长度
	AddressLength = 20
)

//定义hash字节类型
type Hash [HashLength]byte

//定义地址字节类型
type Address [AddressLength]byte
