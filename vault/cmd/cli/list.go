package cli

import (
	"fmt"
	"github.com/lcarneli/go-playground/vault/pkg/vault"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List data and secrets",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.New(encryptionKey, vaultPath)

		values, err := v.GetValues()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list values: %v.\n", err)
			return
		}

		fmt.Printf("%-9s %s %-9s\n", strings.Repeat("=", 9), "Data", strings.Repeat("=", 9))
		fmt.Printf("%-18s %s\n", "Key", "Value")
		fmt.Printf("%-18s %s\n", strings.Repeat("-", 3), strings.Repeat("-", 5))
		for key, value := range values {
			fmt.Printf("%-18s %s\n", key, value)
		}
	},
}

func init() {
	RootCommand.AddCommand(listCommand)
}
