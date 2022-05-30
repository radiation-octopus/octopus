package utils

import "encoding/hex"

// Encode将b编码为前缀为0x的十六进制字符串。
func Encode(b []byte) string {
	enc := make([]byte, len(b)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], b)
	return string(enc)
}

// BytesToHash sets b to hash.
// If b is larger than len(h), b will be cropped from the left.

func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}

// CopyBytes返回所提供字节的精确副本。
func CopyBytes(b []byte) (copiedBytes []byte) {
	if b == nil {
		return nil
	}
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)
	return
}
