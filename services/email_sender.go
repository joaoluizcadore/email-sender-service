package services

import (
	gomail "gopkg.in/mail.v2"
)

//SendEmail - Function to send an email
func SendEmail(config Config, email Email) error {
	message := gomail.NewMessage()
	message.SetHeader("From", config.SMTP.From)
	message.SetHeader("To", email.To...)
	message.SetHeader("Subject", email.Subject)
	if email.IsHTML {
		message.SetBody("text/html", email.Body)
	} else {
		message.SetBody("text/plain", email.Body)
	}
	dialer := gomail.NewDialer(config.SMTP.Host, config.SMTP.Port, config.SMTP.From, config.SMTP.Password)
	return dialer.DialAndSend(message)
}
