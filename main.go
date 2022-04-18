package main

import (
	"WinStuff/utils"
	"fmt"
)

func main() {
	rawData := utils.GetRedditJson("giveaways", 100, "month", "new")

	for i := 0; i < len(rawData.Data.Children); i++ {
		fmt.Println(rawData.Data.Children[i].Data.UrlOverriddenByDest)
	}
}
