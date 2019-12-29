# Introduction

Feel like you are wasting too much time not working on your computer?

Fix that with Rescue Time!

Rescue Time is chrome extension that is used to keep track of what exactly 
you do on your computer on any given day. Go here if you dont have it downloaded yet:

	https://www.rescuetime.com/download

After you have it downloaded and your settings created just how you like, you can use
this CLI to interact with your daily spenditures of time.

# Installation 

Firstly, make a directory in where you are comfortable running this from. Then run

	git init
	git pull https://github.com/benmorehouse/RTC.git

Check to ensure you are running the latest version of Go. To do so, run 

	go version

Then, you need to grab the following three packages using the get command:

	go get "github.com/boltdb/bolt"
	go get "github.com/spf13/cobra"
	go get "github.com/gookit/color"

Finally, build the binary and run

	./rtc init


