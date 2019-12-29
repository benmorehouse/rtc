package main

import(
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"github.com/gookit/color"
	"fmt"
	"strings"
	"log"
)

var InitCmd = &cobra.Command{
	Use:"init",
	Short:"Initialize your Rescuetime CLI",
	Run: func(cmd *cobra.Command, args []string){
		color.Magenta.Println("Welcome to your very own Rescue Time CLI")
		for{
			fmt.Print("Have you created a rescuetime account yet? [Y/N]")
			var input string
			fmt.Scan(&input)
			input = strings.ToUpper(input)
			if input == "YES" || input == "Y"{
				break
			}else if input == "NO" || input == "N"{
				fmt.Print("Go to ")
				color.Magenta.Print("https://www.rescuetime.com/")
				fmt.Print(" and make an account. You must do so before continuing!")
				return
			}else{
				fmt.Println("** Not a valid input **")
				continue
			}
		}
		fmt.Println("Next go to ")
		color.Magenta.Print("https://www.rescuetime.com/apidoc")
		fmt.Println(" and create an API key. \nThis will act as your way of getting access to your account through this CLI.")
		for{
			fmt.Print("Have you created a rescuetime API Key yet? [Y/N]")
			var input string
			fmt.Scan(&input)
			input = strings.ToUpper(input)
			if input == "YES" || input == "Y"{
				break
			}else if input == "NO" || input == "N"{
				fmt.Print("Go to ")
				color.Magenta.Print("https://www.rescuetime.com/")
				fmt.Print(" and make an account. You must do so before continuing!")
				return
			}else{
				fmt.Println("** Not a valid input **")
				continue
			}
		}

		// now we have our APIKey
		db, err := bolt.Open("mainDatabase.db", 0600, nil)
		if err != nil{
			return
		}
		defer db.Close()
		_ = db.Update(func(tx *bolt.Tx) error{
			fmt.Print("Paste your API Key here:")
			var APIKey string
			fmt.Scan(&APIKey)
			bucket , err := tx.CreateBucketIfNotExists([]byte("apiAuth"))
			err = bucket.Put([]byte("apiAuth"),[]byte(APIKey)) //creates the new bucket with nothing in it 
			if bucket == nil{
				log.Fatal(err," -> Possible cause is that bolt binary isnt installed correctly")
			}
			return nil
		})
	},
}



