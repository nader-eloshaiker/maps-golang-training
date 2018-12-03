package main

import "fmt"

func main() {
	//var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	delete(colors, "white")

	//colors := map[string]string{
	//	"red": "#ff0000",
	//	"green": "#00ff00",
	//	"blue": "0000ff",
	//}


	fmt.Println(colors)
}
