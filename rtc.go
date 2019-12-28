package main

import(
	"github.com/spf13/cobra"
	"net/http"
	"encoding/csv"
	"strconv"
//	"log"
)

type RTDay struct{
	Date string
	TimeSpent int
	NumberofPeople int // normally should return one since this is a personal tool! :)
	Productivity int
}

type results struct{
	TotalTimeSpent uint
	VeryProductive uint
	Productive uint
	Nuetral uint
	Unproductive uint
	VeryUnproductive uint
	NuetralThreshold uint
	NuetralThresholdReached bool
	UnproductiveThreshold uint
	UnproductiveThresholdReached bool
}

type mode int
const(
	Day    mode = 0
	Week   mode = 1
	Month  mode = 2
)

var rootCmd = &cobra.Command{
	Use:"",
	DisableFlagParsing:true, // all flags passed as arguments
	Short:"Display rescuetime data",
	Run: func(cmd *cobra.Command, args []string){ // args is gonna be what we pass through 
		if len(args) == 0 || args[0] == "d"{ // default is to just do the day
			request := GenerateRequest() // generates our request for us
			resp, err := http.Get(request)
			if err != nil{
				return
			}

			data , err := csv.NewReader(resp.Body).ReadAll()
			if err != nil{
				return
			}

			bufferResults := results{
				TotalTimeSpent: 0,
				VeryUnproductive: 0,
				Unproductive: 0,
				Nuetral: 0,
				Productive: 0,
				VeryProductive: 0,
				NuetralThreshold:10,
				NuetralThresholdReached: false,
				UnproductiveThreshold:10,
				UnproductiveThresholdReached: false,
			}

			for _, day := range data[1:]{
				timeSpentBuffer, _ := strconv.Atoi(day[1])
				timeSpent := uint(timeSpentBuffer)
				productivity, _ := strconv.Atoi(day[3])

				bufferResults.TotalTimeSpent += uint(timeSpent)

				switch productivity{
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

			if float64(bufferResults.Nuetral) / float64(bufferResults.TotalTimeSpent) < float64(bufferResults.NuetralThreshold){
				bufferResults.NuetralThresholdReached = true
			}

			if float64(bufferResults.Unproductive + bufferResults.VeryUnproductive) / float64(bufferResults.TotalTimeSpent) < float64(bufferResults.UnproductiveThreshold){
				bufferResults.NuetralThresholdReached = true
			}

			bufferResults.Output(Day)
		}else if args[0] == "w" {
			return
			//bufferResults.Ouput(Week)
		}else{
			return
			//bufferResults.Ouput(Week)
		}
	},
}

// can keep data within boltdb and can use a nice CLI interface maybe with cobra? Doesnt have to use cobra though.

