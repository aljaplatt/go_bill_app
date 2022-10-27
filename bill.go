package main

import "fmt"

// new type, name, type = struct = structure
// blueprint for any bill object
// struct is the go equivalent of a class/ data type blueprint
type bill struct {
	// property type
	name string
	// order items - map type, keys strings, values floats 
	items map[string]float64
	tip float64
}

//* function to generate new bills object 
// takes in name as string, return a bill type/object
func newBill(name string) bill {
	// b = bill object
	b := bill{
		// define initial values for the object properties
		// name property is set to the name argument passed into the function
		name: name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip: 0,
	}
	return b
}

// format the bill
func (b bill) format() string {
	fs := "Your Bill: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// total 
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

// Your Bill: 
// pie:                      ...$5.99 
// cake:                     ...$3.99 
// total:                    ...$9.98