package cmd

import (
	"doem/cmd/list"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doem",
	Short: "Todos in cli",
	Long:  "Todos in cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(list.GetCommand())
}
