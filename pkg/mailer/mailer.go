package mailer

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/nabiel-syarif/playlist-api/internal/config"
)

var cfg config.MailConfig

type MailHeader struct {
	Key   string
	Value string
}

func SetMailConfig(config config.MailConfig) {
	cfg = config
}

func sendMail(from string, to []string, cc []string, headers []MailHeader, subject, message string) error {
	body := "From: " + from + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n"

	if len(headers) != 0 {
		for _, v := range headers {
			body += v.Key + ": " + v.Value + "\n"
		}
	}
	body += "\n" + message

	auth := smtp.CRAMMD5Auth(cfg.SmtpUsername, cfg.SmtpPassword)
	smtpAddr := fmt.Sprintf("%s:%d", cfg.SmtpHost, cfg.SmtpPort)

	log.Printf("SMTP connecting to addr : %v\n", smtpAddr)
	err := smtp.SendMail(smtpAddr, auth, from, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func SendAddedAsCollaboratorEmailNotification(playlistName, from, fromName, to, toName string) error {
	body := fmt.Sprintf("Hello %s.\nYou are invited as a collaborator in playlist '%s' by %s.", toName, playlistName, fromName)
	return sendMail(from, []string{to}, nil, []MailHeader{
		{Key: "Content-Type", Value: "text/plain"},
	}, "You are now collaborator in playlist : "+playlistName, body)
}
