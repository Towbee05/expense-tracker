package cmd

import (
	"expense_tracker/internal"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var deleteExpenseCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete an expense.",
	Long:    "Delete an existing expense using its ID.",
	Aliases: []string{"drop"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := uuid.Parse(args[0])
		if err != nil {
			return err
		}
		if err := internal.DeleteExpense(id); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteExpenseCmd)
}
