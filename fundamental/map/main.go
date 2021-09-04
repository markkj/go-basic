package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf75",
	}
	fmt.Println(colors["red"])

	colors["white"] = "#00000"
	fmt.Println(colors)

	for key, v := range colors {
		fmt.Println(key, v)
	}
}
