package main

import(
	"github.com/gookit/color"
	"fmt"
)

func (this* results) Output(input mode){
	if this.NuetralThresholdReached{
		color.Yellow.Println("*** ALERT *** \nDESIGNATED NUETRAL THRESHOLD HAS BEEN REACHED. \nSUGGESTION: CHECK RESCUETIME ACCOUNT FOR DISCREPENCY IN CATEGORIZATION")
	}

	if this.UnproductiveThresholdReached{
		color.Yellow.Println("*** ALERT *** \nDESIGNATED UNPRODUCTIVE THRESHOLD HAS BEEN REACHED. \nSUGGESTION: CHECK RESCUETIME ACCOUNT FOR DISCREPENCY IN CATEGORIZATION")
	}

	if input == Day{
		this.OutputHelper()
	}else if input == Week{
		return
	}else{
		return
	}
}

func (this* results) OutputHelper(){
	percentages := [5]float64{
		(float64(this.VeryProductive)   / float64(this.TotalTimeSpent))*100,
		(float64(this.Productive)       / float64(this.TotalTimeSpent))*100,
		(float64(this.Nuetral)          / float64(this.TotalTimeSpent))*100,
		(float64(this.Unproductive)     / float64(this.TotalTimeSpent))*100,
		(float64(this.VeryUnproductive) / float64(this.TotalTimeSpent))*100,
	}
	color.Cyan.Println("Time logged today:",this.TotalTimeSpent/3600," hours, ",this.TotalTimeSpent%60," minutes")
	for i:=0;i<5;i++{
		switch i{
		case 0:
			color.Blue.Print("Very Productive    ")
		case 1:
			color.LightBlue.Print("Productive         ")
		case 2:
			color.Gray.Print("Nuetral            ")
		case 3:
			color.LightRed.Print("Unproductive       ")
		case 4:
			color.Red.Print("Very Unproductive  ")
		}
		fmt.Print("[")
		for j:=0.0;j<50;j++{
			if j > percentages[i]/2{
				fmt.Print("-")
			}else{
				fmt.Print("#")
			}
		}

		fmt.Print("]")
		fmt.Printf("%9.2f",percentages[i])
		fmt.Println("%")
	}
}

