package internal

import (
	"encoding/json"
	"fmt"
)

func ListExpenses() error {
	file, err := ReadFileContent(filename)
	if err != nil {
		return err
	}
	var expenses []Expense
	if err := json.Unmarshal(file, &expenses); err != nil {
		return err
	}
	if len(expenses) < 1 {
		return fmt.Errorf("No expense created yet.")
	}
	BeautifyPrint(expenses)
	return nil
}

func ListExpensesByCategory(category string) error {
	file, err := ReadFileContent(filename)
	if err != nil {
		return err
	}
	var expenses []Expense
	var filteredExpensees []Expense
	if err := json.Unmarshal(file, &expenses); err != nil {
		return err
	}
	if len(expenses) < 1 {
		return fmt.Errorf("No expense created yet.")
	}
	for _, e := range expenses {
		if e.Category == category {
			filteredExpensees = append(filteredExpensees, Expense{
				ID:          e.ID,
				Description: e.Description,
				Category:    e.Category,
				Amount:      e.Amount,
				CreatedAt:   e.CreatedAt,
			})
		} else {
			continue
		}
	}
	BeautifyPrint(filteredExpensees)
	return nil
}
