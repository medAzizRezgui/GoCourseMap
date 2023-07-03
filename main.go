package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#FFBB00",
		"white": "#FFFFFF",
	}
	printMap(colors)

}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex ", hex, "Color", color)
	}
}
