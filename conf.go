package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	LogFile string `yaml:"LogFile"`
	Key     string `yaml:"Key"`
}

const (
	configFileName     = ".rtc.yaml"
	defaultLoggingFile = ".rtc.log"
)

// getCurrentConfiguration will return the current configuration
// if nothing is found, then it will return a default configuration
func getCurrentConfiguration() (*Config, error) {

	_, err := os.Stat(configFileName)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = initConfiguration()
			if err != nil {
				return nil, err
			}
		}
	}

	buffer, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// setConfiguration will take in a certain configuration that is edited and the marshal it
// into the .yml file that we have built
func setConfiguration(conf Config) error {

	fmt.Println("Writing the configuration file")
	fmt.Println(conf)
	confFile, err := os.OpenFile(configFileName, os.O_RDWR, 0740)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Error(err)
			return err
		}

		confFile, err = os.Create(configFileName)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	buffer, err := yaml.Marshal(conf)
	if err != nil {
		log.Error(err)
		return err
	}

	if _, err := confFile.Write(buffer); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// initConfiguration will initialize the simple rtc configuration
// NOTE: I think that the file should be created without the user permission.
func initConfiguration() (*Config, error) {

	// if they previously had a file, then we are going to delete it.
	if _, err := os.Stat(configFileName); err == nil {
		if err := os.Remove(configFileName); err != nil {
			return nil, err
		}
	}

	config := &Config{
		Key:     "",
		LogFile: defaultLoggingFile,
	}

	if err := setConfiguration(*config); err != nil {
		log.Error(err)
		return nil, err
	}

	return config, nil
}

// setLogger will set the logging for logrus, and will then take place through
// the rest of the application.
// NOTE: this should be one of the if not the first called function when using rtc.
func setLogger(filename *string) error {

	fn := defaultLoggingFile
	if filename != nil && strings.TrimSpace(*filename) != "" {
		fn = *filename
	}

	logFile, err := os.OpenFile(fn, os.O_RDWR, 0740)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		logFile, err = os.Create(fn)
		if err != nil {
			return err
		}
	}

	log.SetOutput(logFile)
	log.SetReportCaller(true)
	return nil
}
