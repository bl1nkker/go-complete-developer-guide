package maps

import "fmt"

func GetColors() map[string]string{
	// Declaration 1
	colors1 := map[string]string{
		"white": "#FFF",
		"black": "#000",
		"red": "00FF",
	}

	// Declaration 2
	// var colors2 map[string]string
	// colors2["white"] = "#FFF"

	// Declaration 3
	// colors3 := make(map[string]string)
	// colors3["white"] = "#FFF"
	return colors1
}

func DeleteItem(m map[string]string, key string){
	delete(m, key)
}

func PrintColor(m map[string]string){
	fmt.Println("Printing colors for:", &m)
	for color, hex := range m{
		fmt.Printf("The color is %s. And the hex is %s\n", color, hex)
	}
}

func Run() {
	colors := GetColors()
	PrintColor(colors)
	DeleteItem(colors, "white")
	PrintColor(colors)
}