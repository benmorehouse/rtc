package main

//simply calling the Cobra CLI framework

func main(){
	rootCmd.AddCommand(InitCmd)
	rootCmd.Execute()
}
