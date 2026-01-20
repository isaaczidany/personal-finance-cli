package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	Description string
	Amount      float64
	Category    string
	Date        time.Time
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func showMenu() {
	fmt.Println(strings.Repeat("\033[1;36m=\033[0m", 36))
	fmt.Println(strings.Repeat(" ", 6), "\033[1;36mPersonal Finance - CLI\033[0m")
	fmt.Println(strings.Repeat("\033[1;36m=\033[0m", 36))
	fmt.Println("\n\033[1;36m1 | Add Expense \n2 | List Expenses \n3 | View Total Expenses\033[0m")
	fmt.Println("\n\033[1;36m4 | Filter by Category \n5 | Filter by Date \n0 | Exit\033[0m \n")
	fmt.Println(strings.Repeat("\033[1;36m=\033[0m", 36), "\n\033[1;36mSelect an option: \033[0m")
}

func readExpense(reader *bufio.Reader) (Expense, error) {
	var expense Expense

	fmt.Print("Description: ")
	text, _ := reader.ReadString('\n')
	expense.Description = strings.TrimSpace(text)

	fmt.Print("Amount: ")
	amountstr, _ := reader.ReadString('\n')
	amountstr = strings.TrimSpace(amountstr)
	amount, err := strconv.ParseFloat(amountstr, 64)
	if err != nil {
		return Expense{}, err
	}
	expense.Amount = amount

	fmt.Print("Category: ")
	catstr, _ := reader.ReadString('\n')
	expense.Category = strings.TrimSpace(catstr)

	fmt.Print("Date (YYYY-MM-DD): ")
	dateStr, _ := reader.ReadString('\n')
	dateStr = strings.TrimSpace(dateStr)
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return Expense{}, err
	}
	expense.Date = date

	return expense, nil
}

func listExpenses(expenses []Expense, reader *bufio.Reader) {
	if len(expenses) == 0 {
		fmt.Println("[No expenses found.")
	}
	fmt.Println("\033[1;34m============== ALL EXPENSES ==============\033[0m")
	for i, expense := range expenses {
		fmt.Printf(
			"\033[1;36m%d | %s | %.2f | %s | %s \033[0m",
			i+1,
			expense.Description,
			expense.Amount,
			expense.Category,
			expense.Date.Format("02/01/2006"),
		)
		fmt.Println("\nOptions:\nR | Refresh List\nV | Back to Main Menu")
		fmt.Print("Choose an Option: ")
		text, _ := reader.ReadString('\n')
		option := strings.TrimSpace(text)
		switch option {
		case "R", "r":
			continue
		case "V", "v":
			return
		default:
			continue
		}
	}
}

func totalExpenses(expenses []Expense, reader *bufio.Reader) {
	var option string
	var total float64 = 0.0
	for _, expense := range expenses {
		total += expense.Amount
	}
	fmt.Printf("Total Expenses: %.2f \n", total)
	fmt.Println("Option: 1 | Exit ")
	fmt.Print("Enter Option: ")
	text, _ := reader.ReadString('\n')
	option = strings.TrimSpace(text)
	switch option {
	case "1":
		return
	default:
		return
	}
}

func filterCategory(expenses []Expense, reader *bufio.Reader) {
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}
	fmt.Println("Enter category to filter: ")
	text, _ := reader.ReadString('\n')
	category := strings.TrimSpace(strings.ToLower(text))
	fmt.Println("\nFiltered Expenses: \n")
	found := false
	for _, expense := range expenses {
		if strings.ToLower(expense.Category) == category {
			fmt.Printf(
				"%s | %.2f | %s | %s\n",
				expense.Description,
				expense.Amount,
				expense.Category,
				expense.Date.Format("02/01/2006"),
			)
			found = true
		}
	}
	if !found {
		fmt.Println("No expenses found for this category")
	}
	fmt.Println("Option: \n 1 | Exit \n Enter option: ")
	text2, _ := reader.ReadString('\n')
	option := strings.TrimSpace(text2)
	switch option {
	case "1":
		return
	default:
		return
	}
}

func filterDate(expenses []Expense, reader *bufio.Reader) {
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}
	fmt.Println("Enter date to filter:")
	strDate, _ := reader.ReadString('\n')
	strDate = strings.TrimSpace(strDate)
	dateFilter, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		fmt.Println("Invalid date format")
		return
	}
	found := false
	for _, expense := range expenses {
		if expense.Date.Format("2006-01-02") == dateFilter.Format("2006-01-02"){
			fmt.Printf(
				"%s | %.2f | %s | %s\n",
				expense.Description,
				expense.Amount,
				expense.Category,
				expense.Date.Format("01/02/2006"),
			)
			found = true
		}
	}
	if !found {
		fmt.Println("No expenses found for this date.")
	}
	fmt.Println("Option:\n1 | Exit\nEnter option ")
	text3, _ := reader.ReadString('\n')
	option := strings.TrimSpace(text3)
	switch option {
	case "1":
		return
	default:
		return

	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var expenses []Expense
	var running bool = true
	for running == true {
		clearTerminal()
		showMenu()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			expense, err := readExpense(reader)
			if err != nil {
				fmt.Println("Error adding expense:", err)
				break
			}
			expenses = append(expenses, expense)
		case "2":
			clearTerminal()
			listExpenses(expenses, reader)
		case "3":
			totalExpenses(expenses, reader)
		case "4":
			filterCategory(expenses, reader)
		case "5":
			filterDate(expenses, reader)
		case "0":
			running = false
		}
	}
}
