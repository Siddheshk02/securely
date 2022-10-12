package cmd

import (
	"fmt"
	"log"

	"github.com/Siddheshk02/securely/lib"
	"github.com/spf13/cobra"
)

var adminCmd = &cobra.Command{
	Use:     "Admin",
	Aliases: []string{"admin"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		//if len(args) != 2 {
		//	fmt.Println("Please provide a filepath")
		//	return
		//}
		filepath := args[0]
		//fmt.Println(filepath)
		//file := args[2]

		test := lib.Admin(filepath)
		if test != nil {
			log.Fatal("Error while creating the Shares for the file!", test.Error())

		} else {
			fmt.Println("Shares created Successfully!")
		}

	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
