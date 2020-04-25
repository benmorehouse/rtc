package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// InitCmd is the cobra command which manually initializes a new instance of rtc.
// NOTE: this should also be invoked upon starting the rtc application.
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your Rescuetime CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		return initRTC()
	},
}

// initRTC wraps the initialization process of the cli
func initRTC() error {

	config, err := initConfiguration()
	if err != nil {
		log.Error(err)
		return err
	}

	setLogger(&config.LogFile)

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
			log.Error(ErrNoAccount)
			return ErrNoAccount
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
			log.Error(ErrNoKey)
			return ErrNoKey
		} else {
			fmt.Println("** Not a valid input **")
			continue
		}
	}

	fmt.Print("Paste your API Key here:")
	var APIKey string
	fmt.Scan(&APIKey)

	config.Key = APIKey

	request, err := GenerateRequest(config.Key) // generates our request for us
	if err != nil {
		log.Error(err)
	}

	_, err = http.Get(request)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(config)
	if err := setConfiguration(*config); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
