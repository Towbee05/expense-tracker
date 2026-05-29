package internal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func DeleteExpense(id uuid.UUID) error {
	file, err := ReadFileContent(filename)
	if err != nil {
		return err
	}
	if len(file) < 1 {
		return fmt.Errorf("No expense created yet.")
	}
	var expenses []Expense
	var updatedExpenses []Expense
	if err := json.Unmarshal(file, &expenses); err != nil {
		return err
	}
	for _, e := range expenses {
		if e.ID == id {
			continue
		}
		newExpense := &Expense{
			ID:          e.ID,
			Description: e.Description,
			Amount:      e.Amount,
		}
		updatedExpenses = append(updatedExpenses, *newExpense)
	}
	bytes, err := json.MarshalIndent(updatedExpenses, "", " ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filename, bytes, 0644); err != nil {
		return err
	}
	BeautifyPrint(updatedExpenses)
	return nil
}
