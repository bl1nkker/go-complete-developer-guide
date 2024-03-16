package maps

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