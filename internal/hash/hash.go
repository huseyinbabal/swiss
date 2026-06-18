package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash/crc32"
)

// MD5 returns the hex-encoded MD5 digest of the input.
func MD5(in string) string {
	sum := md5.Sum([]byte(in))
	return hex.EncodeToString(sum[:])
}

// SHA1 returns the hex-encoded SHA-1 digest of the input.
func SHA1(in string) string {
	sum := sha1.Sum([]byte(in))
	return hex.EncodeToString(sum[:])
}

// SHA256 returns the hex-encoded SHA-256 digest of the input.
func SHA256(in string) string {
	sum := sha256.Sum256([]byte(in))
	return hex.EncodeToString(sum[:])
}

// SHA512 returns the hex-encoded SHA-512 digest of the input.
func SHA512(in string) string {
	sum := sha512.Sum512([]byte(in))
	return hex.EncodeToString(sum[:])
}

// CRC32 returns the hex-encoded IEEE CRC-32 checksum of the input.
func CRC32(in string) string {
	sum := crc32.ChecksumIEEE([]byte(in))
	return fmt.Sprintf("%08x", sum)
}
