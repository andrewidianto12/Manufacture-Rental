package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func SendRegistrationEmail(toEmail, fullName, username string) error {
	if fullName == "" {
		fullName = username
	}

	subject := "Selamat datang di Manufacture Rental"
	body := fmt.Sprintf("Halo %s,\n\nRegistrasi akun kamu berhasil.\n\nUsername: %s\nEmail: %s\n\nSilakan login menggunakan email dan password yang sudah kamu daftarkan.\n", fullName, username, toEmail)

	if err := sendViaSMTP(toEmail, fullName, subject, body); err == nil {
		return nil
	} else {
		apiErr := sendViaAPI(toEmail, fullName, subject, body)
		if apiErr == nil {
			return nil
		}

		return fmt.Errorf("smtp failed: %v; api fallback failed: %v", err, apiErr)
	}
}

func sendViaSMTP(toEmail, fullName, subject, body string) error {
	host := os.Getenv("MAILERSEND_SMTP_HOST")
	port := os.Getenv("MAILERSEND_SMTP_PORT")
	usernameSMTP := os.Getenv("MAILERSEND_SMTP_USERNAME")
	passwordSMTP := os.Getenv("MAILERSEND_SMTP_PASSWORD")
	fromEmail := os.Getenv("MAILERSEND_FROM_EMAIL")
	fromName := os.Getenv("MAILERSEND_FROM_NAME")

	if host == "" {
		host = "smtp.mailersend.net"
	}
	if port == "" {
		port = "587"
	}

	missing := make([]string, 0)
	if strings.TrimSpace(usernameSMTP) == "" {
		missing = append(missing, "MAILERSEND_SMTP_USERNAME")
	}
	if strings.TrimSpace(passwordSMTP) == "" {
		missing = append(missing, "MAILERSEND_SMTP_PASSWORD")
	}
	if strings.TrimSpace(fromEmail) == "" {
		missing = append(missing, "MAILERSEND_FROM_EMAIL")
	}
	if strings.TrimSpace(toEmail) == "" {
		missing = append(missing, "recipient_email")
	}

	if len(missing) > 0 {
		return fmt.Errorf("mailersend configuration is incomplete: missing %s", strings.Join(missing, ", "))
	}

	if fromName == "" {
		fromName = "Manufacture Rental"
	}

	message := strings.Join([]string{
		fmt.Sprintf("From: %s <%s>", fromName, fromEmail),
		fmt.Sprintf("To: %s", toEmail),
		fmt.Sprintf("Subject: %s", subject),
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		"",
		body,
	}, "\r\n")

	auth := smtp.PlainAuth("", usernameSMTP, passwordSMTP, host)
	if err := smtp.SendMail(host+":"+port, auth, fromEmail, []string{toEmail}, []byte(message)); err != nil {
		return fmt.Errorf("mailersend smtp failed: %w", err)
	}

	return nil
}

func sendViaAPI(toEmail, fullName, subject, body string) error {
	token := os.Getenv("MAILERSEND_API_TOKEN")
	fromEmail := os.Getenv("MAILERSEND_FROM_EMAIL")
	fromName := os.Getenv("MAILERSEND_FROM_NAME")

	missing := make([]string, 0)
	if strings.TrimSpace(token) == "" {
		missing = append(missing, "MAILERSEND_API_TOKEN")
	}
	if strings.TrimSpace(fromEmail) == "" {
		missing = append(missing, "MAILERSEND_FROM_EMAIL")
	}
	if strings.TrimSpace(toEmail) == "" {
		missing = append(missing, "recipient_email")
	}

	if len(missing) > 0 {
		return fmt.Errorf("mailersend api configuration is incomplete: missing %s", strings.Join(missing, ", "))
	}

	if fromName == "" {
		fromName = "Manufacture Rental"
	}

	payload := map[string]any{
		"from": map[string]string{
			"email": fromEmail,
			"name":  fromName,
		},
		"to": []map[string]string{
			{
				"email": toEmail,
				"name":  fullName,
			},
		},
		"subject": subject,
		"text":    body,
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest(http.MethodPost, "https://api.mailersend.com/v1/email", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		responseBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("mailersend api request failed with status: %s, body: %s", resp.Status, string(responseBody))
	}

	return nil
}
