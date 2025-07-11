package cli

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	vaultPath     string
	encryptionKey string
)

var RootCommand = &cobra.Command{
	Use:   "vault",
	Short: "A vault to store secrets",
}

func defaultVaultPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Failed to get user home directory.")
	}

	path := filepath.Join(home, "vault/", "vault.dat")

	return path
}

func init() {
	RootCommand.PersistentFlags().StringVarP(&vaultPath, "path", "p", defaultVaultPath(), "Path to vault file")
	RootCommand.PersistentFlags().StringVarP(&encryptionKey, "key", "k", "", "Encryption key for vault file")
}
