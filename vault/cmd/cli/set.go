package cli

import (
	"fmt"
	"github.com/lcarneli/go-playground/vault/pkg/vault"
	"github.com/spf13/cobra"
	"os"
)

var setCommand = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set data and secrets",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.New(encryptionKey, vaultPath)
		key, value := args[0], args[1]

		err := v.SetValue(key, value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to set value for key \"%s\": %v.\n", key, err)
			return
		}

		fmt.Printf("Success! Value set for key: \"%s\".\n", key)
	},
}

func init() {
	RootCommand.AddCommand(setCommand)
}
