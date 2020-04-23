package main

import (
	"net/url"
	"strings"
	"time" // used to get current date for API call
)

const baseUrl string = "https://www.rescuetime.com/anapi/data"

func GenerateRequest(apiKey string) (string, error) {
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
