package hash

import "testing"

func TestHashes(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string) string
		in   string
		want string
	}{
		{"md5", MD5, "abc", "900150983cd24fb0d6963f7d28e17f72"},
		{"sha1", SHA1, "abc", "a9993e364706816aba3e25717850c26c9cd0d89d"},
		{"sha256", SHA256, "abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{"sha512", SHA512, "abc", "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f"},
		{"crc32", CRC32, "abc", "352441c2"},
		{"md5 empty", MD5, "", "d41d8cd98f00b204e9800998ecf8427e"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn(tt.in); got != tt.want {
				t.Errorf("%s(%q) = %q, want %q", tt.name, tt.in, got, tt.want)
			}
		})
	}
}

func TestCRC32Error(t *testing.T) {
	// CRC32 should never collide for these distinct, simple inputs.
	if CRC32("abc") == CRC32("abd") {
		t.Errorf("CRC32 unexpectedly equal for distinct inputs")
	}
}
