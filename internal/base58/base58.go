package base58

import (
	"fmt"
	"math/big"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var radix = big.NewInt(58)

// Encode encodes the raw bytes of in using Bitcoin base58. It never returns an
// error but keeps the signature for consistency.
func Encode(in string) (string, error) {
	data := []byte(in)

	// Count leading zero bytes -> leading '1's.
	zeros := 0
	for zeros < len(data) && data[zeros] == 0 {
		zeros++
	}

	num := new(big.Int).SetBytes(data)
	mod := new(big.Int)
	var out []byte
	for num.Sign() > 0 {
		num.DivMod(num, radix, mod)
		out = append(out, alphabet[mod.Int64()])
	}

	for i := 0; i < zeros; i++ {
		out = append(out, alphabet[0])
	}

	// Reverse.
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return string(out), nil
}

// Decode decodes a Bitcoin base58 string back into its raw bytes. It returns an
// error when the input contains an invalid character.
func Decode(in string) (string, error) {
	// Count leading '1's -> leading zero bytes.
	zeros := 0
	for zeros < len(in) && in[zeros] == alphabet[0] {
		zeros++
	}

	num := big.NewInt(0)
	for i := 0; i < len(in); i++ {
		idx := indexOf(in[i])
		if idx < 0 {
			return "", fmt.Errorf("invalid base58 character %q", in[i])
		}
		num.Mul(num, radix)
		num.Add(num, big.NewInt(int64(idx)))
	}

	decoded := num.Bytes()
	out := make([]byte, zeros+len(decoded))
	copy(out[zeros:], decoded)

	return string(out), nil
}

func indexOf(c byte) int {
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] == c {
			return i
		}
	}
	return -1
}
