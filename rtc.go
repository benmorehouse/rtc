package main

import (
	"encoding/csv"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type RTDay struct {
	Date           string
	TimeSpent      int
	NumberofPeople int // normally should return one since this is a personal tool! :)
	Productivity   int
}

type results struct {
	TotalTimeSpent               uint
	VeryProductive               uint
	Productive                   uint
	Nuetral                      uint
	Unproductive                 uint
	VeryUnproductive             uint
	NuetralThreshold             uint
	NuetralThresholdReached      bool
	UnproductiveThreshold        uint
	UnproductiveThresholdReached bool
}

// enums used to tell which mode to be in
type mode int

const (
	Day   mode = 0
	Week  mode = 1
	Month mode = 2
)

var rootCmd = &cobra.Command{
	Use:                "",
	DisableFlagParsing: true, // all flags passed as arguments
	Short:              "Display rescuetime data",
	Run: func(cmd *cobra.Command, args []string) { // args is gonna be what we pass through
		conf, err := getCurrentConfiguration()
		if err != nil {
			InternalFatalError(ConfigFileInit)
		}

		setLogger(&conf.LogFile)
		if len(args) == 0 || args[0] == "d" { // default is to just do the day
			getDaily(*conf)
		} else if args[0] == "w" {
			log.Trace("User has request this weeks information")
			return
			//bufferResults.Ouput(Week)
		} else {
			log.Trace("User has request this month's information")
			return
			//bufferResults.Ouput(Week)
		}

	},
}

// getDaily will generate the output based on just the day's work
func getDaily(conf Config) {

	log.Info(conf)
	log.Trace("User has requested today's information")
	request, err := GenerateRequest(conf.Key) // generates our request for us
	if err != nil {
		log.Error(err)
		InternalFatalError(URLParseFailed)
	}

	log.Info(request)
	resp, err := http.Get(request)
	if err != nil {
		log.Error(err)
		InternalFatalError(GetRequestFailed)
	}

	data, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		log.Error(err)
		InternalFatalError(CSVParseFailed)
	}

	bufferResults := results{
		TotalTimeSpent:               0,
		VeryUnproductive:             0,
		Unproductive:                 0,
		Nuetral:                      0,
		Productive:                   0,
		VeryProductive:               0,
		NuetralThreshold:             10,
		NuetralThresholdReached:      false,
		UnproductiveThreshold:        10,
		UnproductiveThresholdReached: false,
	}

	for _, day := range data[1:] {
		timeSpentBuffer, _ := strconv.Atoi(day[1])
		timeSpent := uint(timeSpentBuffer)
		productivity, _ := strconv.Atoi(day[3])

		bufferResults.TotalTimeSpent += uint(timeSpent)

		switch productivity {
		case -2:
			bufferResults.VeryUnproductive += timeSpent
		case -1:
			bufferResults.Unproductive += timeSpent
		case 0:
			bufferResults.Nuetral += timeSpent
		case 1:
			bufferResults.Productive += timeSpent
		case 2:
			bufferResults.VeryProductive += timeSpent
		}
	}

	if float64(bufferResults.Nuetral)/float64(bufferResults.TotalTimeSpent) > float64(bufferResults.NuetralThreshold) {
		bufferResults.NuetralThresholdReached = true
	}

	if float64(bufferResults.Unproductive+bufferResults.VeryUnproductive)/float64(bufferResults.TotalTimeSpent) > float64(bufferResults.UnproductiveThreshold) {
		bufferResults.NuetralThresholdReached = true
	}

	bufferResults.Output(Day)
	log.Trace("Successfully return rtc data for the day thus far")
}
