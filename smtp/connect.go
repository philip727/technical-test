package smtp

import (
	"fmt"
	"net/smtp"
)

func SendNoreplyMail(recipient string, subject string, body string) error {
    var server = "localhost"
    var port = 1025
    var sender = "noreply@SecuriGroup.co.uk"

    message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

    auth := smtp.PlainAuth("", "", "", server)
    err := smtp.SendMail(fmt.Sprintf("%s:%d", server, port), auth, sender, []string{recipient}, []byte(message))

    if err != nil {
        return err
    }

    return nil
}
