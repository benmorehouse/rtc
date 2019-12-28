package main

import(
	"net/http"
	"encoding/csv"
	"github.com/spf13/cobra"
	"strconv"
)

type RTDay struct{
	Date string
	TimeSpent int
	NumberofPeople int // normally should return one since this is a personal tool! :)
	Productivity int
}

type results struct{
	TotalTimeSpent uint
	VeryUnproductive uint
	Unproductive uint
	Nuetral uint
	Productive uint
	VeryProductive uint
	NuetralThresholdReached bool
	UnproductiveThresholdReached bool
}

type mode int
const{
	Day    mode = 0
	Week   mode = 0
	Month  mode = 0
}

var todayCmd = &cobra.Command{
	Use:"",
	DisableFlagParsing:true, // all flags passed as arguments
	Short:"Display rescuetime data",
	Run: func(cmd *cobra.Command, args []string){ // args is gonna be what we pass through 
		if len(args) == 0{ // default is to just do the day
			request := GenerateRequest() // generates our request for us
			resp, err := http.Get(request)
			if err != nil{
				return
			}

			data , err := csv.NewReader(resp.Body).ReadAll()
			if err != nil{
				return
			}

			var parsedData []RTDay
			for _, day := range data[1:]{
				var scanned RTDay
				scanned.Date = day[0]
				scanned.TimeSpent, _ = strconv.Atoi(day[1])
				scanned.NumberofPeople, _ = strconv.Atoi(day[2])
				scanned.Productivity, _ = strconv.Atoi(day[3])
				parsedData = append(parsedData,scanned)
			}
		}else{
			return
		}
	}
}

// can keep data within boltdb and can use a nice CLI interface maybe with cobra? Doesnt have to use cobra though.
func (this* results) Output(input mode){
	if input == Day{

	}else if input == Week{
		return
	}else{
		return
	}
}
