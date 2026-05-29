package cmd

import (
	"expense_tracker/internal"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var description string
var amount float64

var createExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "add an expense",
	Long:  "add an expense in a temporary file storage that exist in this project path. Provide a description and amount of expenses to get started.",
	Aliases: []string{
		"create",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if strings.TrimSpace(description) == "" {
			return fmt.Errorf("createExpenseCmd: please provide a description for the expense")
		}
		if amount < 0 {
			return fmt.Errorf("createExpenseCmd: please provide a valid amount: amount must be greater than zero")
		}
		if err := internal.CreateExpense(description, amount); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	createExpenseCmd.Flags().StringVarP(&description, "description", "D", "", "")
	createExpenseCmd.Flags().Float64VarP(&amount, "amount", "A", 0, "")
	createExpenseCmd.MarkFlagsRequiredTogether("description", "amount")
	rootCmd.AddCommand(createExpenseCmd)
}
