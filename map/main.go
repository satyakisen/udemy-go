package main

import "fmt"

func main() {
	colors := map[string]string{
		"white": "#ffffff",
		"black": "#000000",
	}
	printColor(colors)
}

func printColor(d map[string]string) {
	for color, hex := range d {
		fmt.Println(color, hex)
	}
}
