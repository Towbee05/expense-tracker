package cmd

import (
	"expense_tracker/internal"

	"github.com/spf13/cobra"
)

var month int

var summarizeCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summarize your general expenses",
	Long:  "Summarize your general expenses. You can summarize by providing the month for the expense.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("month") {
			if err := internal.SummarizeExpenseMonth(month); err != nil {
				return err
			}
		} else {
			if err := internal.SummarizeExpense(); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	summarizeCmd.Flags().IntVarP(&month, "month", "", 0, "")
	rootCmd.AddCommand(summarizeCmd)
}
