package main

import "fmt"

func main() {
	//var colors map[string]string
	//colors := make(map[string]string)

	//colors["white"] = "#ffffff"
	//delete(colors, "white")

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "0000ff",
		"white": "ffffff",
	}


	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("key:" , key, "value:", value)
	}
}
