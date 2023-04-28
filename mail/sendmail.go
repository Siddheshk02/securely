package mail

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(name string, email string) error {

	err := godotenv.Load(".env")
	from := mail.NewEmail("Securely", "noreply@securelee.tech")
	key := os.Getenv("Key")

	subject := "Welcome to Securely"

	to := mail.NewEmail(name, email)

	plainTextContent := "Hey " + name + ", Welcome to Securely! We're glad you're here."

	htmlContent := "<strong>Hey " + name + "</strong>,<br> <p> Welcome to Securely! </p> <p> We're glad you're here.<p> If you have any queries regarding Securely, feel free to reach out to us at enquiry@securelee.tech . </p> <p> We're here with you every step of the way.</p> <br> <p> Best Regards,</p> <p>Securely Team</p>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(key)

	response, err := client.Send(message)
	_ = response
	if err != nil {
		log.Println(err)
	}

	return nil
}

func UserMail(name string, email string, admin string) error {

	err := godotenv.Load(".env")

	from := mail.NewEmail("Securely", "noreply@securelee.tech")
	key := os.Getenv("Key")
	subject := "[New Share Alloted]"

	//fmt.Println(name, email)
	to := mail.NewEmail(string(name), string(email))

	htmlContent := "<b>Hey " + name + "</b>, <p> New Share is alloted to you by " + admin + "</p> <p> To get more Information about the file, Log-in to Securely. </p> <p> If you have any queries regarding Securely, feel free to reach out to us at <a href=url>enquiry@securelee.tech</a>. We're here with you every step of the way.</p> <br> <p>Best Regards,</p><p>Securely Team</p>"

	plainTextContent := "Hey " + name + ", New Share is alloted to you by " + admin + ". To get more Information about the file, Log-in to Securely. If you have any queries regarding Securely, feel free to reach out to us at enquiry@securelee.tech . We're here with you every step of the way. Best Regards, Securely Team"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(key)

	response, err := client.Send(message)
	_ = response
	if err != nil {
		log.Println(err)
	}

	return nil
}

func MailAdmin(admin string, email string, filename string, name string) {

	err := godotenv.Load(".env")

	from := mail.NewEmail("Securely", "noreply@securelee.tech")
	key := os.Getenv("Key")

	subject := "[File Downloaded!]"

	//fmt.Println(name, email)
	to := mail.NewEmail(admin, email)

	t := time.Now()
	htmlContent := "<b>Hey " + admin + "</b>, <p> <b>File : " + filename + "</b> is decrypted by </p><p> <b> User : " + name + " </b> at </p> <p> <b>Time : " + t.Format("2006-01-02 15:04:05") + "</b></p> <p> To get more Information about the file, Log-in to Securely. </p> <p> If you have any queries regarding Securely, feel free to reach out to us at <a href=url>enquiry@securelee.tech</a>. We're here with you every step of the way.</p> <br> <p>Best Regards,</p><p>Securely Team</p>"

	plainTextContent := "Hey " + admin + ", New Share is alloted to you by " + admin + ". To get more Information about the file, Log-in to Securely. If you have any queries regarding Securely, feel free to reach out to us at enquiry@securelee.tech . We're here with you every step of the way. Best Regards, Securely Team"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(key)

	response, err := client.Send(message)
	_ = response
	if err != nil {
		log.Println(err)
	}
}
