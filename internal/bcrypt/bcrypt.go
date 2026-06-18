package bcrypt

import (
	gobcrypt "golang.org/x/crypto/bcrypt"
)

// Hash returns a bcrypt hash of the given password.
func Hash(password string) (string, error) {
	h, err := gobcrypt.GenerateFromPassword([]byte(password), gobcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(h), nil
}

// Verify compares a bcrypt hash against a password. It returns "match" or
// "no match". An error is only returned when the hash is malformed.
func Verify(hash, password string) (string, error) {
	err := gobcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return "match", nil
	}
	if err == gobcrypt.ErrMismatchedHashAndPassword {
		return "no match", nil
	}
	return "", err
}
