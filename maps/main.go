package main

import "fmt"

func main() {
	// var colors map[string]string

	/* colors := map[string]string{
		"red":   "#ff0000",
		"green": "#50C878",
	} */

	colors := make(map[int]string)
	colors[10] = "#ff00ff"
	colors[20] = "#ffffff"

	delete(colors, 10)
	fmt.Println(colors)
}
