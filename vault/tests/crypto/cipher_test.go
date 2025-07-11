package crypto

import (
	"github.com/lcarneli/go-playground/vault/pkg/crypto"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	key := "super-secret-key"
	plaintext := "this is a top secret message"

	encrypted, err := crypto.Encrypt(key, plaintext)
	if err != nil {
		t.Errorf("got %q, expected %q", err, "nil")
	}

	decrypted, err := crypto.Decrypt(key, encrypted)
	if err != nil {
		t.Errorf("got %q, expected %q", err, "nil")
	}

	if decrypted != plaintext {
		t.Errorf("got %q, expected %q", decrypted, plaintext)
	}
}

func TestDecryptWithWrongKey(t *testing.T) {
	key := "correct-key"
	wrongKey := "wrong-key"
	plaintext := "this is a top secret message"

	encrypted, err := crypto.Encrypt(key, plaintext)
	if err != nil {
		t.Errorf("Failed to encrypt: %v", err)
	}

	_, err = crypto.Decrypt(wrongKey, encrypted)
	if err == nil {
		t.Errorf("got %q, expected %q", "nil", "invalid key error")
	}
}

func TestDecryptWithInvalidHex(t *testing.T) {
	key := "any-key"
	badHex := "ZZZ123"

	_, err := crypto.Decrypt(key, badHex)
	if err == nil {
		t.Errorf("got %q, expected %q", "nil", "invalid hex error")
	}
}
