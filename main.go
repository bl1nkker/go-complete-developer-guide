package main

import (
	"go-complete-developer-guide/cards"
	"go-complete-developer-guide/channels"
	"go-complete-developer-guide/helper"
	"go-complete-developer-guide/interfaces"
	"go-complete-developer-guide/maps"
	"go-complete-developer-guide/structs"
)

func main() {
	cards.Run()
	structs.Run()
	helper.Run()
	helper.RunExtra()
	maps.Run()
	interfaces.Run()
	// interfaces.RunHttp()
	// interfaces.RunAssignment1()
	// interfaces.RunAssignment2()
	channels.RunChannels()
}
