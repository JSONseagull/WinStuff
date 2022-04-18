package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	preparedRequest.Header.Set("cookie", "csv=2; edgebucket=LOyadGKtOryMqqkQ4C; USER=eyJwcmVmcyI6eyJnbG9iYWxUaGVtZSI6IlJFRERJVCIsImNvbGxhcHNlZFRyYXlTZWN0aW9ucyI6eyJmYXZvcml0ZXMiOmZhbHNlLCJtdWx0aXMiOmZhbHNlLCJtb2RlcmF0aW5nIjpmYWxzZSwic3Vic2NyaXB0aW9ucyI6ZmFsc2UsInByb2ZpbGVzIjpmYWxzZX0sIm5pZ2h0bW9kZSI6dHJ1ZSwicnBhbkR1RGlzbWlzc2FsVGltZSI6bnVsbCwidG9wQ29udGVudERpc21pc3NhbFRpbWUiOm51bGwsInRvcENvbnRlbnRUaW1lc0Rpc21pc3NlZCI6MH19; pc=m7; reddit_session=1460843059395,2022-02-11T22:34:39,7153e87e38e679428fee2efc407bda6c57f88348; loid=0000000000in3o8mo3.2.1642335160000.Z0FBQUFBQmlCdVJfc0sxYzhsaUdrZ0xmTlhxOEYtbmVZWlYwU2ZxV21Ud3JWMDNkS095UHJ0SDhhd3ROODZrLUNnbHZHS3Vsc1h1MUVzWWlHVFBmbHhKSnl4UWVlOGxYLWZqMWdiYXE5TGc0Tnc0T1owNEViUUVCeWx5ZUkyQVVVNjBmSG9LMnh3Rnk; show_announcements=yes; recent_srs=t5_2reca,t5_2sokd,t5_2tex6,t5_2rc7j,t5_2t0lj,t5_2ty3s,t5_2t9i0,t5_2qm21,t5_2s28b,t5_2x93b; token_v2=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTAzMjAwNjEsInN1YiI6IjE0NjA4NDMwNTkzOTUtYmFwWGJYUmhlQnBCNnYzOGN2UWx2X3F4M21qek9BIiwibG9nZ2VkSW4iOnRydWUsInNjb3BlcyI6WyIqIiwiZW1haWwiLCJwaWkiXX0.2kNj2-9634Wa0oEMtWQP2aOhQOIV_CNhxx2v0gfMh_A; session_tracker=npjhkerhccidjbkfpp.0.1650288737491.Z0FBQUFBQmlYV2hpdmV5bFctTi1URmxhaUw4MWRVeDVtYy1hc2VuZ0llcnU1TG0wZVpxWjdibll6OWt3SnRoZ0Q5U0FKZ0lXWFdxanlvQTV0QjlfLXhkeE1ZVVdCdl8tN0F2Ul8weGZGYkkzUTktNXFIaXY1SDNLeW9JdmthQUx2d3I2OFlzUmxXYW8")
	res, err := tempClient.Do(preparedRequest)
	if err != nil {
		fmt.Printf("error reqesting reddit data: %v\n", err)
	}
	defer res.Body.Close()

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading out response body: %v\n", err)
	}
	fmt.Println(string(rawBody))
	raw := Posting{}

	if json.Unmarshal(rawBody, &raw); err != nil {
		fmt.Printf("error unmarshalling response body: %v\n", err)
	}

	return raw
}
