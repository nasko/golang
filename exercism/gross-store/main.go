package main

import "fmt"

func main() {
	bill := NewBill()
	units := Units()
	ok := AddItem(bill, units, "carrot", "dozen") // 12
	fmt.Println(ok)
	fmt.Println("Current bill:", bill)
	_ = AddItem(bill, units, "carrot", "quarter_of_a_dozen") // 3
	fmt.Println("Current bill:", bill)

	_ = AddItem(bill, units, "carrot", "half_of_a_dozen") // 6
	fmt.Println("Current bill:", bill)

	_ = AddItem(bill, units, "apple", "gross") // 6
	fmt.Println("Current bill:", bill)

	_ = AddItem(bill, units, "apple", "twelve") // 6
	fmt.Println("Current bill:", bill)

	_ = RemoveItem(bill, units, "carrot", "quarter_of_a_dozen") // 3
	fmt.Println("Current bill:", bill)

	removeOk := RemoveItem(bill, units, "orange", "quarter_of_a_dozen") // 3
	fmt.Println("removeOk:", removeOk)
	fmt.Println("Current bill:", bill)

	carrotQty, carrotOk := GetItem(bill, "carrot")
	fmt.Println("carrot:", carrotQty)
	fmt.Println("carrot on bill:", carrotOk)

	orangeQty, orangeOk := GetItem(bill, "orange")
	fmt.Println("orange:", orangeQty)
	fmt.Println("orange on bill:", orangeOk)

}

func AddItem(bill map[string]int, units map[string]int, item string, unit string) bool {

	qtyToAdd, unitInUnitsOk := units[unit]
	if !unitInUnitsOk {
		return false
	}

	bill[item] += qtyToAdd

	return true
}

func RemoveItem(bill map[string]int, units map[string]int, item string, unit string) bool {

	// Return false if the given item is not in the bill
	itemBalance, itemInBillOk := bill[item]
	if !itemInBillOk {
		return false
	}

	// Return false if the given unit is not in the units map
	qtyToRemove, unitInUnitsOk := units[unit]
	if !unitInUnitsOk {
		return false
	}

	if qtyToRemove > itemBalance {
		// Return false if the new quantity would be less than 0
		return false
	} else if qtyToRemove == itemBalance {
		// If the new quantity is 0, completely remove the item from the bill then return true
		delete(bill, item)
	} else {
		// Otherwise, reduce the quantity of the item and return true
		bill[item] -= qtyToRemove
	}

	return true
}

func GetItem(bill map[string]int, item string) (int, bool) {

	// Return 0 and false if the given item is not in the bill
	// Otherwise, return the quantity of the item in the bill and true
	itemBalance, itemInBillOk := bill[item]
	return itemBalance, itemInBillOk
}

func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

func NewBill() map[string]int {
	return map[string]int{}
}
