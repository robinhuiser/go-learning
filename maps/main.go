package main

import "fmt"

func main() {
	// var colors1 map[string]string
	// colors2 := make(map[string]string)
	colors3 := map[string]string{
		"red":   "#ff0000",
		"green": "#aabb11",
		"white": "#ffffff",
	}

	// Why does this throw an error?
	// colors1["blue"] = "#bb55cc"
	// colors2["yellow"] = "#ff33aa"
	colors3["orange"] = "#ff88dd"

	// fmt.Print("colors1:", colors1)
	// fmt.Print("colors2:", colors2)
	// fmt.Print("colors3:", colors3)

	printMap(colors3)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("The hexvalue for color", key, "is", value)
	}
}
