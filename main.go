package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{ // this is a global variable dont do this! Put it in main instead
}

func main(){
	rootCmd.AddCommand(todayCommand)
	rootCmd.Execute()
}
