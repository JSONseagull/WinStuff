package main

import (
	"WinStuff/utils"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	//Using godotenv to grab cookie data(and maybe entry credentials, dont know yet)
	//While Reddit does not like this type of scraping, it makes it easier to use as you dont need to set up oauth2
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading .env file: %v\n", err)
	}
	rawData := utils.GetRedditJson("giveaways", 100, "month", "new")

	for i := 0; i < len(rawData.Data.Children); i++ {
		fmt.Println(rawData.Data.Children[i].Data.UrlOverriddenByDest)
	}
}
