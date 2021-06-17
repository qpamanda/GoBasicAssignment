package main

import (
	"fmt"
	"strconv"
)

var itemList map[string]item
var categoryName = []string{"Household", "Food", "Drinks"}

func main() {
	itemList = make(map[string]item)

	//setupItemList() // setup test data

	for {
		userChoice := printShoppingListMenu()
		iChoice, _ := strconv.ParseInt(userChoice, 10, 0)

		if iChoice == 99 { // allow user to exit the program if choice is 99
			break
		} else {
			switch iChoice {
			case 1:
				viewShoppingList()
			case 2:
				generateReport()
			case 3:
				addItems()
			case 4:
				modifyItems()
			case 5:
				deleteItems()
			default:
				fmt.Println(">>Invalid choice! Please re-enter.")
			}
		}
		fmt.Println("")
	}
}

func printShoppingListMenu() string {
	var sChoice string

	fmt.Println("SHOPPING LIST APPLICATION")
	fmt.Println("=========================")
	fmt.Println("1. View Entire Shopping List.")
	fmt.Println("2. Generate Shopping List Report.")
	fmt.Println("3. Add Items.")
	fmt.Println("4. Modify Items.")
	fmt.Println("5. Delete Items.")
	fmt.Println("Select your choice:")
	fmt.Scanln(&sChoice)
	fmt.Println("")

	return sChoice
}

func viewShoppingList() {
	fmt.Println("Shopping List Contents:")

	if len(itemList) == 0 {
		fmt.Println("The shopping list is empty. ")
	} else {
		for key, value := range itemList { // key = item name, value = item struct
			value.printItemInfo(key)
		}
	}
}

func generateReport() {
	for {
		userChoice := printReportMenu()
		iChoice, _ := strconv.ParseInt(userChoice, 10, 0)

		if iChoice == 1 {
			printTotalCostByCategory()
		} else if iChoice == 2 {
			printListOfItemByCategory()
		} else if iChoice == 3 {
			break
		} else {
			fmt.Println(">>Invalid choice! Please re-enter.")
		}

		fmt.Println("")
	}
}

func printReportMenu() string {
	var userChoice string

	fmt.Println("GENERATE REPORT")
	fmt.Println("1. Total Cost of each category.")
	fmt.Println("2. List of item by category.")
	fmt.Println("3. Main Menu.")
	fmt.Println("")
	fmt.Println("Choose your report:")
	fmt.Scanln(&userChoice)
	fmt.Println("")

	return userChoice
}
