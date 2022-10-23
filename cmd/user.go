package cmd

import (
	"fmt"
	"log"

	"github.com/Siddheshk02/securely/lib"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:     "User",
	Aliases: []string{"user"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		filepath := args[0]

		test := lib.User(filepath)
		if test != nil {
			log.Fatal("Error!! ", test.Error())

		} else {
			fmt.Println("File Decrypted Successfully!")
		}

	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
