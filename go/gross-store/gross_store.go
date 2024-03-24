package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	res := map[string]int{}
	res["quarter_of_a_dozen"] = 3
	res["half_of_a_dozen"] = 6
	res["dozen"] = 12
	res["small_gross"] = 120
	res["gross"] = 144
	res["great_gross"] = 1728
	return res
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	res := map[string]int{}
	return res
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if val, exis := units[unit]; exis {
		if _, item_exis := bill[item]; !item_exis {
			bill[item] = val
		} else {
			bill[item] += val
		}
		return true
	}
	return false
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Check if the unit exists in the units map
	val, ok := units[unit]
	if !ok {
		return false
	}

	// Check if the item exists in the bill
	quantity, exists := bill[item]
	if !exists {
		return false
	}

	// Calculate the new quantity after removing the unit
	newQuantity := quantity - val
	if newQuantity < 0 {
		return false
	}

	if newQuantity == 0 {
		// Remove the item from the bill if the new quantity is 0
		delete(bill, item)
	} else {
		// Update the quantity of the item in the bill
		bill[item] = newQuantity
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	res, exis := bill[item]
	if !exis {
		return 0, false
	}
	return res, true
}
