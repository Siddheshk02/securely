package mail

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(name string, email string, admin string, check int) error {

	from := mail.NewEmail("Securely", "noreply@securelee.tech")
	if check == 1 {
		subject := "Welcome to Securely"

		to := mail.NewEmail(name, email)

		plainTextContent := "Hey " + name + ", Welcome to Securely! We're glad you're here."

		htmlContent := "<strong>Hey " + name + "</strong>, <br> Welcome to Securely! <br> We're glad you're here. <br> To get more Information about the file, Log-in to Securely. <br> If you have any queries regarding Securely, feel free to reach out to us at enquiry@securelee.tech. We're here with you every step of the way."

		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

		client := sendgrid.NewSendClient("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")

		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
		}

	} else {
		subject := "[New Share Alloted]"

		to := mail.NewEmail(name, email)

		plainTextContent := "Hey " + name + ", New Share is alloted to you by " + admin + " To get more Information about the file, Log-in to Securely. <br> If you have any queries regarding Securely, feel free to reach out to us at enquiry@securelee.tech. We're here with you every step of the way."

		htmlContent := "Hey " + name + ", <br> <br> New Share is alloted to you by " + admin + "<br> <br> To get more Information about the file, Log-in to Securely. <br> <br> If you have any queries regarding Securely, feel free to reach out to us at <a> enquiry@securelee.tech </a>. We're here with you every step of the way."

		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

		client := sendgrid.NewSendClient("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")

		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
		}
	}

	return nil
}
