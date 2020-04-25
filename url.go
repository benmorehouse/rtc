package main

import (
	"net/url"
	"strings"
	"time" // used to get current date for API call
)

// GenerateRequest will return the url query for a day
func GenerateRequest(apiKey string) (string, error) {

	baseUrl := "https://www.rescuetime.com/anapi/daily_summary_feed"
	date := time.Now()
	today := strings.Fields(date.String())[0]

	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("key", apiKey)
	q.Set("perspective", "interval")
	q.Set("restrict_kind", "productivity")
	q.Set("interval", "hour")
	q.Set("restrict_begin", today)
	q.Set("restrict_end", today)
	q.Set("format", "csv")

	u.RawQuery = q.Encode()

	return u.String(), nil
}
