package functions

import (
	"fmt"
	"net/smtp"
	"os"
)

func Email(email, nimi, sisu string) string {

	from := os.Getenv("emailAddr")

	password := os.Getenv("emailPass")

	//toEmail := os.Getenv("toEmail")
	toEmail := os.Getenv("TSemail")
	rain := os.Getenv("rain")

	to := []string{toEmail, rain}

	host := "smtp.gmail.com"
	port := "587"

	address := host + ":" + port

	subject := "Furnituuri kalkulaatori tagasiside kasutajalt " + nimi
	body := "\nKasutaja email: " + email + "\n" + sisu

	message := []byte("Subject: " + subject + "\r\n" + body)

	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("error:", err)
		return "Tekkis viga, vabandame. Võite saata oma tagasiside meilile tsenter@tsenter.ee"
	}
	fmt.Println("Email saadetud")
	return "Saadetud! Aitäh tagasiside eest."
}
