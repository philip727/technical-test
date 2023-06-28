package smtp

import (
	"fmt"
	"net/smtp"
	"securigroup/tech-test/utils"
)

func SendNoreplyMail(recipient string, subject string, body string) error {
    var server = utils.GetEnvVar("SMTP_SERVER")
    var port = utils.GetEnvVar("SMTP_PORT")
    var sender = utils.GetEnvVar("SMTP_SENDER")

    message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

    auth := smtp.PlainAuth("", "", "", server)
    err := smtp.SendMail(fmt.Sprintf("%s:%s", server, port), auth, sender, []string{recipient}, []byte(message))

    if err != nil {
        return err
    }

    return nil
}
