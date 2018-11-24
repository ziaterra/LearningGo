package main

import (
	"fmt"
)

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + ": " + hex)
	}
}
func main() {
	//
	colors := map[string]string{
		"red":    "#ff0000",
		"orange": "#fg0000",
	}

	var colors2 map[string]string
	colors3 := make(map[string]string)

	printMap(colors)
	printMap(colors2)
	printMap(colors3)
	// fmt.Println(colors)
	// fmt.Println(colors["red"])
}
