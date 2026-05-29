package cmd

import (
	"expense_tracker/internal"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var uDescription string
var uAmount float64

var updateExpenseCmd = &cobra.Command{
	Use:     "update",
	Short:   "update the description or amount of existing expenses",
	Long:    "by providing the ID of an expense, you can update the description or the amount of that expense",
	Aliases: []string{"edit"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(args)
		id, err := uuid.Parse(args[0])
		if err != nil {
			return fmt.Errorf("failed to parse strign into uuid %v", err)
		}
		if err := internal.UpdateExpense(id, &uDescription, &uAmount); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	updateExpenseCmd.Flags().StringVarP(&uDescription, "description", "", "", "")
	updateExpenseCmd.Flags().Float64VarP(&uAmount, "amount", "", 0, "")
	rootCmd.AddCommand(updateExpenseCmd)
}
