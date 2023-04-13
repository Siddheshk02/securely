package cmd

import (
	"fmt"
	"log"

	"github.com/Siddheshk02/securely/auth"
	"github.com/Siddheshk02/securely/controllers"
	"github.com/Siddheshk02/securely/database"
	"github.com/Siddheshk02/securely/lib"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:     "User",
	Aliases: []string{"user"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		if auth.UserCheck() {
			fmt.Println("\nSignUp / Login as User")

			fmt.Print("\nPress Enter to Sign up/Login ")
			fmt.Scanf(" ")
			errr := auth.UserSignup()
			if errr != nil {
				fmt.Println("Error while Signup/Login. Please Try Again.")
				return
			}

		} else {
			var temp int
			fmt.Println("\n1. Decrypt a File.")
			fmt.Println("2. Who AM I ?")
			fmt.Println("3. Logout")

			userd, err := controllers.WhoamiUser()
			if err != nil {
				fmt.Println("Error while fetching the User Data. Please Try Again.")
				return
			}

			fmt.Print("\nEnter any one of the above Options(for e.g. '1') : ")
			fmt.Scan(&temp)
			switch temp {
			case 1:
				//Decrypt a File
				fmt.Println("Decrypt a File")
				//fmt.Print("Enter the Index of the File to be Decrypted : ")
				//list all availble allocated files of the admin to the user
				// //select the file pass its name
				// fmt.Println(string(userd))
				file, adm := database.ShowFiles(userd)
				// var filepath string
				// fmt.Scan(&filepath)
				fmt.Println(file, adm)
				test := lib.User(file, adm, userd)
				if test != nil {
					//log.Fatal("Error!! ", test.Error())
					log.Fatal("Error!! ", test)

				} else {
					fmt.Println("File Decrypted Successfully!")
				}
			case 2:
				//Who AM I ?
				userdata, err := controllers.WhoamiUser()
				if err != nil {
					fmt.Println("Error while fetching the User Data. Please Try Again.")
					return
				}
				fmt.Println("User Data : ", string(userdata))
			case 3:
				//Logout
				controllers.UserLogout()
			default:
				fmt.Println("Invalid Option. Please Try Again.")

			}

		}

		// filepath := args[0]

		// test := lib.User(filepath)
		// if test != nil {
		// 	log.Fatal("Error!! ", test.Error())

		// } else {
		// 	fmt.Println("File Decrypted Successfully!")
		// }

	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
