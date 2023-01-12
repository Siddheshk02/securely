package cmd

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"

	"github.com/Siddheshk02/Securely/config"
	"github.com/spf13/cobra"
)

func GoogleLogin() {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

	//c.Status(fiber.StatusSeeOther)
	open.Run(url)
	//return open.Json(url)
	//return err
}

var adminCmd = &cobra.Command{
	Use:     "Admin",
	Aliases: []string{"admin"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var temp int
		fmt.Println("\t1. Sign Up / Login as Admin")
		fmt.Println("\t2. Who AM I ?")
		fmt.Print("Enter any one of the above Options(for e.g. '1') : ")
		fmt.Scanf("%d", &temp)
		switch temp {
		case 1:
			//login/sign up controller

			GoogleLogin()

		case 2:
			//Works only if the user is live orelse please signup or login to Securely
			fmt.Println("Please Sign Up or Login to Securely to make this feature work :/")

		}

		//filepath := args[0]

		//err := lib.Admin(filepath)
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
