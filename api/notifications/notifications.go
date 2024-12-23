package notifications

import (
	"fmt"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/servlicense/servlicense/api/config"
)

// SendNotification sends an email notification with the specified subject and message
func SendNotification(subject string, message string) error {
	cfg := config.GetConfig()

	// Check if SMTP is enabled
	if !cfg.Smtp.Enabled {
		return fmt.Errorf("SMTP is disabled in the configuration")
	}

	from := cfg.Smtp.SmtpFrom
	auth := smtp.PlainAuth("", cfg.Smtp.SmtpUsername, cfg.Smtp.SmtpPassword, cfg.Smtp.SmtpHost)

	var body strings.Builder
	body.WriteString("Subject: ")
	body.WriteString(subject)
	body.WriteString("\r\n\r\n")
	body.WriteString(message)

	// Send the email
	err := smtp.SendMail(
		cfg.Smtp.SmtpHost+":"+strconv.FormatInt(int64(cfg.Smtp.SmtpPort), 10),
		auth,
		from,
		cfg.Notification.Recipients,
		[]byte(body.String()),
	)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
