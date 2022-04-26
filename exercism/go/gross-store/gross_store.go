package gross

// Units stores the Gross Store unit measurements.
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

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	qtyToAdd, unitInUnitsOk := units[unit]
	if !unitInUnitsOk {
		return false
	}

	bill[item] += qtyToAdd

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
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

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemBalance, itemInBillOk := bill[item]
	return itemBalance, itemInBillOk
}
