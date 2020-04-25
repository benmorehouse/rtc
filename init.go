package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

type config struct {
	LogFile string `yaml:"LogFile"`
	Verbose bool   `yaml:"Verbose"`
	Key     string `yaml:"Key"`
}

const configFileName string = ".rtc.yaml"

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your Rescuetime CLI",
	Run: func(cmd *cobra.Command, args []string) {

		color.LightBlue.Println("Welcome to your very own Rescue Time CLI")
		for {
			fmt.Print("Have you created a rescuetime account yet? [Y/N]")
			var input string
			fmt.Scan(&input)
			input = strings.ToUpper(input)
			if input == "YES" || input == "Y" {
				break
			} else if input == "NO" || input == "N" {
				fmt.Print("Go to ")
				color.Magenta.Print("https://www.rescuetime.com/")
				fmt.Print(" and make an account. You must do so before continuing!")
				return
			} else {
				fmt.Println("** Not a valid input **")
				continue
			}
		}

		fmt.Print("Next go to ")
		color.LightBlue.Print("https://www.rescuetime.com/apidoc")
		fmt.Println(" and create an API key. \nThis will act as your way of getting access to your account through this CLI.")

		var input string
		for {
			fmt.Print("Have you created a rescuetime API Key yet? [Y/N]")
			fmt.Scan(&input)
			input = strings.ToUpper(input)
			if input == "YES" || input == "Y" {
				break
			} else if input == "NO" || input == "N" {
				fmt.Print("Go to ")
				color.Magenta.Print("https://www.rescuetime.com/")
				fmt.Print(" and make an account. You must do so before continuing!")
				return
			} else {
				fmt.Println("** Not a valid input **")
				continue
			}
		}

	},
}

// getCurrentConfiguration will return the current configuration
func getCurrentConfiguration() (*config, error) {

	conf, err := os.Open(configFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
	}
}

// setLogger will set the logging for logrus, and will then take place through
// the rest of the application.
// NOTE: default file: rtc.log
func setLogger() {

	log.SetOutput(ioutil.Discard)
	log.SetReportCaller(true)
}
