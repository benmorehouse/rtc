package main

import(
	"time" // used to get current date for API call
	"strings"
)

const ApiKey string = "B63dwvvoZp5Kh47nTVcvRg3Mf1SbkBvCkrAkFbSA" // eventually will be gotten by

func GenerateRequest()(string){
	date := time.Now()
	today := strings.Fields(date.String())[0]
	var apiCall string
	apiCall += "https://www.rescuetime.com/anapi/data?key="
	apiCall += ApiKey
	apiCall += "&perspective=interval&restrict_kind=productivity&interval=hour&restrict_begin="
	apiCall += today
	apiCall += "&restrict_end="
	apiCall += today
	apiCall += "&format=csv"
	return apiCall
}


