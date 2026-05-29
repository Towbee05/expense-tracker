# Expense Tracker

A CLI expense tracker built with Go and [Cobra](https://github.com/spf13/cobra). Stores expenses in a local JSON file.

## Project Structure

```
expense_tracker/
├── main.go              # Entry point
├── cmd/                 # Cobra command definitions
│   ├── root.go
│   ├── createExpensesCmd.go
│   ├── deleteExpensesCmd.go
│   ├── getExpensesCmd.go
│   ├── summarizeExpensesCmd.go
│   └── updateExpensesCmd.go
├── internal/            # Business logic
│   ├── helper.go        # Expense struct, file I/O, table formatting
│   ├── createExpenses.go
│   ├── deleteExpenses.go
│   ├── getExpenses.go
│   ├── summarizeExpenses.go
│   └── updateExpenses.go
├── expense.json         # Data store
├── go.mod
└── go.sum
```

## Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `add` | `create` | Add an expense (`-D` description, `-A` amount) |
| `delete` | `drop` | Delete an expense by UUID |
| `list` | `view, get-all` | List all expenses in a table |
| `summary` | — | Show total expenses; `--month N` for monthly filter |
| `update` | `edit` | Update description/amount of an expense by UUID |

## Usage

```bash
# Build
go build -o expense_tracker

# Add an expense
./expense_tracker add -D "Lunch" -A 15.50

# List all expenses
./expense_tracker list

# Delete an expense
./expense_tracker delete <uuid>

# Update an expense
./expense_tracker update <uuid> --description "New desc" --amount 20

# Get total expenses
./expense_tracker summary

# Get expenses for a specific month
./expense_tracker summary --month 5
```

- Project submitted at [roadmap for backend project](https://roadmap.sh/projects/expense-tracker)