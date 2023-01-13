package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/Siddheshk02/securely/controllers"
	"github.com/skratchdot/open-golang/open"

	"github.com/gorilla/mux"
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
		var temp int
		fmt.Println("\n1. Sign Up / Login as Admin")
		fmt.Println("2. Who AM I ?")
		fmt.Print("\nEnter any one of the above Options(for e.g. '1') : ")
		fmt.Scanf("%d", &temp)
		switch temp {
		case 1:
			//login/sign up controller
			r := mux.NewRouter()
			r.HandleFunc("/google/login", controllers.GoogleLogin).Methods("GET")
			r.HandleFunc("/google/callback", controllers.GoogleCallback).Methods("GET")
			l, err := net.Listen("tcp", "localhost:8088")
			if err != nil {
				log.Fatal(err)
			}

			//browser.OpenURL("http://localhost:8080/google/login")

			open.Start("http://localhost:8088/google/login")
			//http.ListenAndServe(":8080", r)
			http.Serve(l, r)
			//GoogleLogin(cmd)
			//http.Post("/google/login", cmd.GoogleLogin)

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
