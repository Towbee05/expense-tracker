package internal

import (
	"encoding/json"
	"fmt"
)

var monthMap = map[int]string{
	1:  "January",
	2:  "Febuary",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func SummarizeExpense() error {
	file, err := ReadFileContent(filename)
	if err != nil {
		return err
	}
	var expenses []Expense
	if err := json.Unmarshal(file, &expenses); err != nil {
		return err
	}
	var summary float64 = 0
	for _, e := range expenses {
		summary = summary + e.Amount
	}
	fmt.Printf("Total expenses: $%v \n ", summary)
	return nil
}

func SummarizeExpenseMonth(month int) error {
	file, err := ReadFileContent(filename)
	if err != nil {
		return err
	}
	var expenses []Expense
	if err := json.Unmarshal(file, &expenses); err != nil {
		return err
	}
	var summary float64 = 0
	for _, e := range expenses {
		if month == int(e.CreatedAt.Month()) {
			summary = summary + e.Amount
		} else {
			continue
		}
	}
	fmt.Printf("Total expenses for %v: $%v \n ", monthMap[month], summary)
	return nil
}
