package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

func UpdateExpense(id uuid.UUID, description *string, amount *float64) error {
	file, err := ReadFileContent(filename)
	if err != nil {
		return err
	}
	var expenses []Expense
	var updatedExpense []Expense
	data, err := UnmarshalExpenses(file, expenses)
	if err != nil {
		return err
	}
	var existFlag bool = false
	for _, d := range *data {
		if d.ID == id {
			if description != nil && strings.TrimSpace(*description) != "" {
				d.Description = *description
			}
			if amount != nil && *amount > 0 {
				d.Amount = *amount
			}
			updatedExpense = append(updatedExpense, d)
			existFlag = true
		} else {
			updatedExpense = append(updatedExpense, d)
		}
	}
	if !existFlag {
		return fmt.Errorf("No expense with ID: %v available", id)
	}
	content, err := json.MarshalIndent(updatedExpense, "", " ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filename, content, 0644); err != nil {
		return err
	}
	BeautifyPrint(updatedExpense)
	return nil
}
