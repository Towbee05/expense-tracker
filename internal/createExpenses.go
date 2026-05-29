package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func CreateExpense(description, category string, amount float64) error {
	newExpense := &Expense{
		ID:          uuid.New(),
		Description: description,
		Amount:      amount,
		Category:    category,
		CreatedAt:   time.Now(),
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, err := os.Create(filename)
			if err != nil {
				return fmt.Errorf("CreateExpense: failed to create json file: %v", err)
			}
			defer file.Close()
			_, err = file.WriteString("[]")
			if err != nil {
				return fmt.Errorf("CreateExpense: failed to write '[]' into json file: %v", err)
			}
		} else {
			return fmt.Errorf("CreateExpense: failed to read json file: %v", err)
		}
	}
	var expenses []Expense
	if len(content) < 1 {
		expenses = []Expense{*newExpense}
	} else {
		if err := json.Unmarshal(content, &expenses); err != nil {
			return fmt.Errorf("CreateExpense: failed to unmarshal expenses into existing file content: %v", err)
		}
		expenses = append(expenses, *newExpense)
	}
	updatedData, err := json.MarshalIndent(expenses, "", "\t")
	if err != nil {
		return fmt.Errorf("CreateExpense: failed to marshal expenses: %v \n", err)
	}
	if err := os.WriteFile(filename, updatedData, 0644); err != nil {
		return fmt.Errorf("CreateExpense: failed to write expenses into json file: %v \n", err)
	}
	fmt.Printf("Expense added successfully (ID: %v)\n", newExpense.ID)
	return nil
}
