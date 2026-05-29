package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/google/uuid"
)

const (
	filename = "./expense.json"
)

type Expense struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}

func ReadFileContent(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("No expenses has been added yet.")
		} else {
			return nil, fmt.Errorf("ListExpenses: failed to read file, %v", err)
		}
	}
	if len(file) < 1 {
		return nil, fmt.Errorf("No expenses has been added yet.")
	}
	return file, nil
}

func UnmarshalExpenses(file []byte, expenses []Expense) (*[]Expense, error) {
	if err := json.Unmarshal(file, &expenses); err != nil {
		return nil, fmt.Errorf("ListExpenses: failed to unmarshal file data into expenses, %v", err)
	}
	return &expenses, nil
}

func BeautifyPrint(expenses []Expense) {
	tabWriter := tabwriter.NewWriter(os.Stdout, 5, 5, 5, ' ', tabwriter.Debug)
	fmt.Fprintln(tabWriter, "ID \t DESCRIPTION \t AMOUNT")
	for _, e := range expenses {
		fmt.Fprintf(tabWriter, "%v\t%s\t%f\n", e.ID, e.Description, e.Amount)
	}
	tabWriter.Flush()
}
