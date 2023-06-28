package smtp

import (
	"fmt"
	"net/smtp"
)

const server = "localhost"
const port = 1025
const sender = "noreply@SecuriGroup.co.uk"

func SendNoreplyMail(recipient string, subject string, body string) error {
    message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

    auth := smtp.PlainAuth("", "", "", server)
    err := smtp.SendMail(fmt.Sprintf("%s:%d", server, port), auth, sender, []string{recipient}, []byte(message))

    if err != nil {
        return err
    }

    return nil
}
