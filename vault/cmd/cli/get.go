package cli

import (
	"fmt"
	"github.com/lcarneli/go-playground/vault/pkg/vault"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var getCommand = &cobra.Command{
	Use:   "get <key>",
	Short: "Get data and retrieves secrets",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.New(encryptionKey, vaultPath)
		key := args[0]

		value, err := v.GetValue(key)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get value for key \"%s\": %v.\n", key, err)
			return
		}

		fmt.Printf("%-9s %s %-9s\n", strings.Repeat("=", 9), "Data", strings.Repeat("=", 9))
		fmt.Printf("%-18s %s\n", "Key", "Value")
		fmt.Printf("%-18s %s\n", strings.Repeat("-", 3), strings.Repeat("-", 5))
		fmt.Printf("%-18s %s\n", key, value)
	},
}

func init() {
	RootCommand.AddCommand(getCommand)
}
