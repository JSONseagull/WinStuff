package main

import (
	"WinStuff/utils"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading .env file: %v\n", err)
	}
	rawData := utils.GetRedditJson("giveaways", 100, "month", "new")

	for i := 0; i < len(rawData.Data.Children); i++ {
		fmt.Println(rawData.Data.Children[i].Data.UrlOverriddenByDest)
	}
}
