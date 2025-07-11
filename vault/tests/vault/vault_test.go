package vault

import (
	"errors"
	"github.com/lcarneli/go-playground/vault/pkg/vault"
	"os"
	"testing"
)

func createTempVault(t *testing.T) (*vault.Vault, func()) {
	t.Helper()

	key := "super-secret-key"
	path := "/tmp/vault_test.dat"

	v := vault.New(key, path)

	cleanup := func() {
		_ = os.Remove(path)
	}

	return v, cleanup
}

func TestSetAndGetValue(t *testing.T) {
	v, cleanup := createTempVault(t)
	defer cleanup()

	key := "foo"
	value := "bar"

	err := v.SetValue(key, value)
	if err != nil {
		t.Errorf("got %q, expected %q", err, "nil")
	}

	val, err := v.GetValue(key)
	if err != nil {
		t.Errorf("got %q, expected %q", err, "nil")
	}

	if val != value {
		t.Errorf("got %q, expected %q", val, value)
	}
}

func TestGetValues(t *testing.T) {
	v, cleanup := createTempVault(t)
	defer cleanup()

	keyA := "foo"
	keyB := "bar"
	valueA := "bar"
	valueB := "foo"

	err := v.SetValue(keyA, valueA)
	if err != nil {
		t.Errorf("got %q, expected %q", err, "nil")
	}

	err = v.SetValue(keyB, valueB)
	if err != nil {
		t.Errorf("got %q, expected %q", err, "nil")
	}

	values, err := v.GetValues()
	if err != nil {
		t.Errorf("got %q, expected %q", err, "nil")
	}

	if len(values) != 2 {
		t.Errorf("got %d, expected %d", len(values), 2)
	}

	if values[keyA] != valueA {
		t.Errorf("got %q, expected %q", values[keyA], valueA)
	}

	if values[keyB] != valueB {
		t.Errorf("got %q, expected %q", values[keyB], valueB)
	}
}

func TestDeleteValue(t *testing.T) {
	v, cleanup := createTempVault(t)
	defer cleanup()

	key := "foo"
	value := "bar"

	err := v.SetValue(key, value)
	if err != nil {
		t.Errorf("got %q, expected %q", err, "nil")
	}

	err = v.DeleteValue(key)
	if err != nil {
		t.Errorf("got %q, expected %q", err, "nil")
	}

	_, err = v.GetValue(key)
	if !errors.Is(err, vault.ErrKeyNotFound) {
		t.Errorf("got %q, expected %q", err, vault.ErrKeyNotFound)
	}
}
