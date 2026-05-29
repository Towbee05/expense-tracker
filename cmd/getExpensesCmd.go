package cmd

import (
	"expense_tracker/internal"

	"github.com/spf13/cobra"
)

var listExpensesCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all saved expenses",
	Long:  "Fetch all saved expenses, from an in-house memory",
	Aliases: []string{
		"view", "get-all",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := internal.ListExpenses(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listExpensesCmd)
}
