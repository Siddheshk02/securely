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

// func GoogleLogin(w http.ResponseWriter, r *http.Request) {
// 	temp := config.GoogleConfig()
// 	url := temp.AuthCodeURL("randomstate")

// 	//c.Status(fiber.StatusSeeOther)
// 	//open.Run(url)
// 	//http.Redirect(w, r, url, 302)
// 	/*if err := http.Redirect(w, r, url, http.StatusSeeOther); err != nil {
// 		panic(fmt.Errorf("failed to open browser for authentication %s", err.Error()))
// 	}*/
// 	http.Redirect(w, r, url, http.StatusSeeOther)
// 	//return open.Json(url)
// 	//return err
// }

var adminCmd = &cobra.Command{
	Use:     "Admin",
	Aliases: []string{"admin"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if auth.Check() {
			fmt.Println("\nSignUp / Login as Admin")

			fmt.Print("\nPress Enter to Sign up/Login using Google")
			fmt.Scanf(" ")
			errr := auth.Signup()
			if errr != nil {
				fmt.Println("Error while Signup/Login. Please Try Again.")
				return
			}
			// var comp string
			// fmt.Println("Enter The Company/Organization Name : ")
			// fmt.Scan(&comp)

			// data, err := controllers.Whoami()
			// if err != nil {
			// 	fmt.Println("Error while getting the User Data. Please Try Again.")
			// 	return
			// }
			// database.DBconn(data)

		} else {
			var temp int
			fmt.Println("\n1. Encrypt and Secure a File.")
			fmt.Println("2. See all the Encrypted Files.")
			fmt.Println("3. Who AM I ?")
			fmt.Println("4. Logout")

			fmt.Print("\nEnter any one of the above Options(for e.g. '1') : ")
			fmt.Scan(&temp)
			admindata, err := controllers.Whoami()
			if err != nil {
				fmt.Println("Error while fetching the Admin Data. Please Try Again.")
				return
			}
			switch temp {
			case 1:
				//Encrypt and Secure a File
				fmt.Println("Encrypt and Secure a File")
				fmt.Print("\nEnter the File location : ")
				var filepath string
				fmt.Scan(&filepath)

				err := lib.Admin(filepath)
				if err != nil {
					log.Fatal("Error while creating the Shares for the file!", err.Error())

				} else {
					fmt.Println("Shares created and distributed")
				}
				// err := ui.Main(func() {
				// 	filePath := ui.OpenFile(nil)
				// 	if filePath == "" {
				// 		log.Fatalln("No file selected")
				// 	}

				// 	fmt.Println("Selected file:", filePath)
				// })
				// if err != nil {
				// 	log.Fatalln(err)
				// }
			case 2:
				//See all the Enrypted Files
				fmt.Println("See all the Enrypted Files")
				err := database.ListFilesAdmin(admindata)
				if err != nil {
					log.Fatal("Error while Listing the file!", err.Error())

				}
			case 3:
				//Who AM I ?

				fmt.Println("User Data : ", string(admindata))
			case 4:
				//Logout
				controllers.Logout()
			default:
				fmt.Println("Invalid Option. Please Try Again.")

			}

		}

		//ch := 1
		// switch temp {
		// case 1:
		// 	//login/sign up controller

		// 	//GoogleLogin(cmd)
		// 	//http.Post("/google/login", cmd.GoogleLogin)

		// case 2:
		//

		// case 3:
		// 	controllers.Logout()

		// }

		// filepath := args[0]

		// err := lib.Admin(filepath)
		// if err != nil {
		// 	log.Fatal("Error while creating the Shares for the file!", err.Error())

		// } else {
		// 	fmt.Println("Shares created Successfully!")
		// }

	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
