package utils

import (
	"log"
)

func SendConfirmationEmail(to string) error {
	// Dummy email log
	log.Printf("📧 Sent confirmation email to: %s", to)
	return nil
}
