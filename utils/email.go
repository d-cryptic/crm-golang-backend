package utils

import (
    "crypto/tls"
    "fmt"
    "net/smtp"
)

// SMTPConfig represents SMTP configuration
type SMTPConfig struct {
    Host     string
    Port     int
    Username string
    Password string
}

// SendEmail sends an email using SMTP with a tracking pixel
func SendEmail(smtpConfig SMTPConfig, from, to, subject, body, trackingID string) error {
    auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)

    conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", smtpConfig.Host, smtpConfig.Port), nil)
    if err != nil {
        return err
    }
    defer conn.Close()

    client, err := smtp.NewClient(conn, smtpConfig.Host)
    if err != nil {
        return err
    }
    defer client.Quit()

    if err = client.Auth(auth); err != nil {
        return err
    }

    if err = client.Mail(from); err != nil {
        return err
    }
    if err = client.Rcpt(to); err != nil {
        return err
    }

    wc, err := client.Data()
    if err != nil {
        return err
    }
    defer wc.Close()

    trackingPixel := fmt.Sprintf(`<img src="http://localhost:8080/track/open/%s" alt="" style="display:none">`, trackingID)
    msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s\n\n%s", from, to, subject, body, trackingPixel)
    _, err = wc.Write([]byte(msg))
    if err != nil {
        return err
    }

    return nil
}