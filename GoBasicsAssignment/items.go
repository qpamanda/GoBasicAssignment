package main

import (
	"fmt"
	"strconv"
)

type item struct {
	categoryId int
	quantity   int
	unitCost   float64
}

func addItems() {
	var itemName, catName, quantity, unitCost string

	fmt.Println("Add Items.")

	fmt.Println("What is the name of your item?")
	fmt.Scanln(&itemName)

	if itemName != "" {
		fmt.Println("What category does it belong to?")
		fmt.Scanln(&catName)
		catId, _ := validateCategory(catName, true)

		fmt.Println("How many units are there?")
		fmt.Scanln(&quantity)
		quantityValue, _ := validateUnits(quantity, true)

		fmt.Println("How much does it cost per unit?")
		fmt.Scanln(&unitCost)
		unitCostValue, _ := validateUnitCost(unitCost, true)

		addNewItem(itemName, catId, quantityValue, unitCostValue)
		fmt.Println(">>Added", itemName)
	} else {
		fmt.Println(">>Please enter a valid item name.")
	}
}

func addNewItem(itemName string, catId int, quantity int, unitCost float64) {
	addedItem := item{catId, quantity, unitCost}
	itemList[itemName] = addedItem
}

func modifyItems() {
	var inputItemName, newItemName, inputCatName, catName, inputQuantity, quantityValue, inputUnitCost, unitCostValue string
	var catId, newCatId, quantity, newQuantity int
	var unitCost, newUnitCost float64

	fmt.Println("Modify Items.")

	if len(itemList) == 0 {
		fmt.Println("The shopping list is empty. Nothing to modify.")
	} else {
		// -- Enter item name and validate if item exists --
		fmt.Println("Which item would you wish to modify?")
		fmt.Scanln(&inputItemName)

		if inputItemName != "" {
			itemName := validateItem(inputItemName)

			if itemName != "" {
				fmt.Println("Enter new Name. Enter for no change.")
				fmt.Scanln(&newItemName)

				selectedItem := itemList[itemName] // this is the existing item struct

				if newItemName == "" {
					newItemName = itemName
				}

				// -- Enter category name and validate if category exists --
				fmt.Println("Enter new Category. Enter for no change.")
				fmt.Scanln(&inputCatName)

				catId = selectedItem.categoryId // Set to category id of existing item
				newCatId = catId
				if inputCatName != "" {
					newCatId, catName = validateCategory(inputCatName, false)
					if catName == "" { // User press enter for no change
						newCatId = catId
					}
				}

				// -- Enter quantity and validate that quantity is a valid number --
				fmt.Println("Enter new Quantity. Enter for no change.")
				fmt.Scanln(&inputQuantity)

				quantity = selectedItem.quantity // Set to quantity of existing item
				newQuantity = quantity
				if inputQuantity != "" {
					newQuantity, quantityValue = validateUnits(inputQuantity, false)
					if quantityValue == "" { // User press enter for no change
						newQuantity = quantity
					}
				}

				// -- Enter unit cost and validate that unit cost is a valid number --
				fmt.Println("Enter new Unit Cost. Enter for no change.")
				fmt.Scanln(&inputUnitCost)

				unitCost = selectedItem.unitCost
				newUnitCost = unitCost
				if inputUnitCost != "" {
					newUnitCost, unitCostValue = validateUnitCost(inputUnitCost, false)
					if unitCostValue == "" {
						newUnitCost = unitCost
					}
				}

				if catId == newCatId {
					fmt.Println("No changes to category made")
				}

				if quantity == newQuantity {
					fmt.Println("No changes to quantity made")
				}

				if unitCost == newUnitCost {
					fmt.Println("No changes to unit cost made")
				}

				if itemName == newItemName {
					fmt.Println("No changes to item name made")
				}

				modifyExistingItem(itemName, newItemName, newCatId, newQuantity, newUnitCost)
				fmt.Println(">>Updated", itemName)
			}
		} else {
			fmt.Println(">>Please enter a valid item name.")
		}
	}
}

func modifyExistingItem(itemName string, newItemName string, newCatId int, newQuantity int, newUnitCost float64) {
	deleteExistingItem(itemName)
	addNewItem(newItemName, newCatId, newQuantity, newUnitCost)
}

func deleteItems() {
	var inputItemName string
	fmt.Println("Delete Items.")

	// -- Enter item name and validate if item exists --
	fmt.Println("Enter item name to delete:")
	fmt.Scanln(&inputItemName)

	_, bFound := itemList[inputItemName]

	if bFound {
		deleteExistingItem(inputItemName)
		fmt.Println(">>Deleted", inputItemName)
	} else {
		fmt.Println(">>Item not found. Nothing to delete!")
	}
}

func deleteExistingItem(itemName string) {
	delete(itemList, itemName)
}

func validateItem(itemName string) string {
	_, bFound := itemList[itemName]

	if bFound {
		catId := itemList[itemName].categoryId
		catName := categoryName[catId]
		quantity := itemList[itemName].quantity
		unitCost := itemList[itemName].unitCost
		fmt.Println("Current item name is", itemName, "- Category is", catName, "- Quantity is", quantity, "- Unit Cost", unitCost)
	} else {
		itemName = "" // clear itemName
		fmt.Println(">>Please enter a valid item name.")
		fmt.Scanln(&itemName)

		if itemName != "" {
			itemName = validateItem(itemName)
		}
	}

	return itemName
}

func validateCategory(inputCatName string, bAddItem bool) (int, string) {
	var retCatId int
	var inputValue string

	bFound := false
	for catId, catName := range categoryName {
		if catName == inputCatName {
			bFound = true
			retCatId = catId
			inputValue = inputCatName
		}
	}

	if !bFound {
		fmt.Println(">>Please enter a valid category from:", categoryName)
		fmt.Scanln(&inputValue)

		if inputValue != "" {
			retCatId, inputValue = validateCategory(inputValue, bAddItem)
		} else { // check if calling function is AddItems. If so, continue to check for valid category input
			if bAddItem {
				retCatId, inputValue = validateCategory(inputValue, bAddItem)
			}
		}
	}

	return retCatId, inputValue
}

func validateUnits(inputQuantity string, bAddItem bool) (int, string) {
	var inputValue string
	retValue, err := strconv.Atoi(inputQuantity)

	if err != nil {
		fmt.Println(">>Please enter a valid quantity.")
		fmt.Scanln(&inputValue)
		fmt.Println("User input: ", inputValue)

		if inputValue != "" {
			retValue, inputValue = validateUnits(inputValue, bAddItem)
		} else { // check if calling function is AddItems. If so, continue to check for valid category input
			if bAddItem {
				retValue, inputValue = validateUnits(inputValue, bAddItem)
			}
		}
	} else {
		inputValue = inputQuantity
	}
	return retValue, inputValue
}

func validateUnitCost(inputUnitCost string, bAddItem bool) (float64, string) {
	var inputValue string
	retValue, err := strconv.ParseFloat(inputUnitCost, 64)

	if err != nil {
		fmt.Println(">>Please enter valid cost per unit.")
		fmt.Scanln(&inputValue)
		fmt.Println("User input: ", inputValue)

		if inputValue != "" {
			retValue, inputValue = validateUnitCost(inputValue, bAddItem)
		} else { // check if calling function is AddItems. If so, continue to check for valid category input
			if bAddItem {
				retValue, inputValue = validateUnitCost(inputValue, bAddItem)
			}
		}
	} else {
		inputValue = inputUnitCost
	}
	return retValue, inputValue
}

func printTotalCostByCategory() {
	fmt.Println("Total Cost By Category.")

	for i := range categoryName {
		totalCost := 0.0
		for _, value := range itemList { // key = item name, value = item struct
			if value.categoryId == i {
				totalCost += float64(value.quantity) * value.unitCost
			}
		}
		fmt.Printf("%s cost : %.2f\n", categoryName[i], totalCost)
	}
}

func printListOfItemByCategory() {
	fmt.Println("List By Category.")

	if len(itemList) == 0 {
		fmt.Println("The shopping list is empty. ")
	} else {
		for i := range categoryName {
			for key, value := range itemList { // key = item name, value = item struct
				if value.categoryId == i {
					value.printItemInfo(key)
				}
			}
		}
	}
}

func (iItem item) printItemInfo(itemName string) {
	catName := categoryName[iItem.categoryId]
	fmt.Println("Category:", catName, "- Item:", itemName, "Quantity:", iItem.quantity, "Unit Cost:", iItem.unitCost)
}
