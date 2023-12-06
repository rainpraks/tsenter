package functions

import (
	"fmt"
	"net/smtp"
	"os"
)

func Email(email, nimi, sisu string) {

	from := os.Getenv("emailAddr")

	password := os.Getenv("emailPass")

	//toEmail := os.Getenv("toEmail")
	toEmail := "lukas.haavel@gmail.com"

	to := []string{toEmail}

	host := "smtp.gmail.com"
	port := "587"

	address := host + ":" + port

	subject := "Tagasiside kasutajalt " + nimi
	body := "Kasutaja email: " + email + "\n" + sisu

	message := []byte(subject + "\r\n" + body)

	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("saadetud")
}
