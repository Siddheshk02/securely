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

		filepath := args[0]

		err := lib.Admin(filepath)
		if err != nil {
			log.Fatal("Error while creating the Shares for the file!", err.Error())

		} else {
			fmt.Println("Shares created Successfully!")
		}

	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
