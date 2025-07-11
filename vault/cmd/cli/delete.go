package cli

import (
	"fmt"
	"github.com/lcarneli/go-playground/vault/pkg/vault"
	"github.com/spf13/cobra"
	"os"
)

var deleteCommand = &cobra.Command{
	Use:     "delete <key>",
	Short:   "Delete data and secrets",
	Aliases: []string{"del"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.New(encryptionKey, vaultPath)
		key := args[0]

		err := v.DeleteValue(key)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete value for key \"%s\": %v.\n", key, err)
			return
		}

		fmt.Printf("Success! Value deleted for key: \"%s\".\n", key)
	},
}

func init() {
	RootCommand.AddCommand(deleteCommand)
}
