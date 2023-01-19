package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the AddItem() function")
	value, unitExisted := units[unit]
	if !unitExisted {
		return false
	}

	_, itemExisted := bill[item]
	if itemExisted {
		bill[item] += value
	} else {
		bill[item] = value
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExisted := bill[item]
	quantity, unitExisted := units[unit]
	newQuantity := bill[item] - quantity
	if !itemExisted || !unitExisted || newQuantity < 0 {
		return false
	} else if newQuantity == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = newQuantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, itemExisted := bill[item]
	return quantity, itemExisted
}
