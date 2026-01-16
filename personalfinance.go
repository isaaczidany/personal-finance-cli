package main
//packages
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)
//set Expense type
type Expense struct {
	Description string
	Amount float64
	Category string
	Date time.Time
}
func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
func showMenu () {
	fmt.Println("\n", strings.Repeat("=", 36))
	fmt.Println(strings.Repeat(" ", 6), "Personal Finance - CLI")
	fmt.Println(strings.Repeat("=", 36))
	fmt.Println("  \n" + "1 | Add Expense\n" + "2 | List Expenses \n" + "3 | View Total Expenses")
	fmt.Println("\n 4 | Filter by Category \n" + "5 | Filter by Date \n" + "0 | Exit")
	fmt.Println("\n", strings.Repeat("=", 36), "\n Select an option: ")
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
func listExpenses (expenses []Expense)  {
	var option string
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
	}
	fmt.Println("\033[1;34m============== ALL EXPENSES ==============\033[0m")
		for i, expense := range expenses {
			fmt.Printf("%d | %s | %.2f | %s | %s", i+1, expense.Description, expense.Amount, expense.Category, expense.Date.Format("02/01/2006"))
			fmt.Println("\nOptions:" + "\n R | Refresh List" + "\n V | Back to Main Menu")
			fmt.Print("Choose an Option: ")
			fmt.Scanln(&option)
			switch option {
		 		case "R", "r":
			 		continue
		 		case "V", "v":
			 			return
		}


		}
	}
func main () {
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
				fmt.Println("Expense added successfully")
			case "2":
				clearTerminal()
				listExpenses(expenses)
			case "0":
				running = false
		}


	}

}
