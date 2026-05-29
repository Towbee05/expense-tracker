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
