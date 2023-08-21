package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	colors["white"] = "#FFFFFF"

	delete(colors, "red")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Value for", key, "is", value)
	}
}
