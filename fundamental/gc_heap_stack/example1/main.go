package main

import "fmt"

const CHG = 10

type Item struct {
	Description string
	Price       int64
}

func Value(price int64) int64 {
	total := price + price*CHG/100
	return total
}

// go run -gcflags '-m' example1.go
// go run -gcflags '-m -m' example1.go
func main() {
	item1 := Item{Description: "Monitor", Price: 100}
	item1.Price = Value(item1.Price)
	_ = item1
	tP := new(Item)
	*&tP.Description = "NoteBook"
	*&tP.Price = 100
	fmt.Println(*tP)
}
