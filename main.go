package main

import "fmt"

func main(){
	// Pass a customer name into newBill, which returns a bill object, stored as myBill - a bill instance
	myBill := newBill("Alex")

	fmt.Println(myBill.format()) // {Alex map[] 0}
}

// Your Bill: 
// pie:                      ...$5.99 
// cake:                     ...$3.99 
// total:                    ...$9.98