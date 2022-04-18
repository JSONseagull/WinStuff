package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func GetRedditJson(sub string, limit int, timeframe string, sort string) Posting {
	tempClient := http.Client{Timeout: 10 * time.Second}

	baseUrl := fmt.Sprintf("https://www.reddit.com/r/%v/%v.json?limit=%v&t=%v", sub, sort, limit, timeframe)

	preparedRequest, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		fmt.Printf("error assembling request: %v\n", err)
	}

	preparedRequest.Header.Set("User-Agent", "not golang")
	preparedRequest.Header.Set("cookie", os.Getenv("cookie"))
	res, err := tempClient.Do(preparedRequest)
	if err != nil {
		fmt.Printf("error reqesting reddit data: %v\n", err)
	}
	defer res.Body.Close()

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading out response body: %v\n", err)
	}

	raw := Posting{}

	if json.Unmarshal(rawBody, &raw); err != nil {
		fmt.Printf("error unmarshalling response body: %v\n", err)
	}

	return raw
}
