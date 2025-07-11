package vault

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lcarneli/go-playground/vault/pkg/crypto"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	ErrKeyNotFound      = errors.New("key not found")
	ErrFailedToLoadFile = errors.New("failed to load file")
	ErrFailedToSaveFile = errors.New("failed to save file")
)

type Vault struct {
	encryptionKey string
	filepath      string
	data          map[string]string
	mutex         sync.RWMutex
}

func New(encryptionKey, filepath string) *Vault {
	return &Vault{
		encryptionKey: encryptionKey,
		filepath:      filepath,
		data:          make(map[string]string),
	}
}

func (v *Vault) loadFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	defer file.Close()

	var sb strings.Builder
	_, err = io.Copy(&sb, file)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrFailedToLoadFile, err)
	}

	decrypted, err := crypto.Decrypt(v.encryptionKey, sb.String())
	if err != nil {
		return fmt.Errorf("%w: %s", ErrFailedToLoadFile, err)
	}

	reader := strings.NewReader(decrypted)
	dec := json.NewDecoder(reader)
	if err := dec.Decode(&v.data); err != nil {
		return fmt.Errorf("%w: %s", ErrFailedToLoadFile, err)
	}

	return nil
}

func (v *Vault) saveFile(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrFailedToSaveFile, err)
	}
	defer file.Close()

	var sb strings.Builder
	enc := json.NewEncoder(&sb)
	if err := enc.Encode(v.data); err != nil {
		return fmt.Errorf("%w: %s", ErrFailedToSaveFile, err)
	}

	encrypted, err := crypto.Encrypt(v.encryptionKey, sb.String())
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(file, encrypted)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrFailedToSaveFile, err)
	}

	return nil
}

func (v *Vault) GetValue(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	if err := v.loadFile(v.filepath); err != nil {
		return "", err
	}

	value, ok := v.data[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}

func (v *Vault) GetValues() (map[string]string, error) {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	if err := v.loadFile(v.filepath); err != nil {
		return nil, err
	}

	return v.data, nil
}

func (v *Vault) SetValue(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	if err := v.loadFile(v.filepath); err != nil {
		return err
	}

	v.data[key] = value

	if err := v.saveFile(v.filepath); err != nil {
		return err
	}

	return nil
}

func (v *Vault) DeleteValue(key string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	if err := v.loadFile(v.filepath); err != nil {
		return err
	}

	delete(v.data, key)

	if err := v.saveFile(v.filepath); err != nil {
		return err
	}

	return nil
}
