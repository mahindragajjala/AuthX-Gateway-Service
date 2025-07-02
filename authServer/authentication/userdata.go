package controllers

import (
	"authserver/utils"
	"regexp"
)

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func IsEmailExists(email string) (bool, error) {
	// Assume `db` is a *sql.DB
	var exists bool
	err := utils.Db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=?)", email).Scan(&exists)
	return exists, err
}

/*
func insertUser(email, hashedPassword string) error {
	_, err := db.Exec("INSERT INTO users (email, password, created_at) VALUES (?, ?, NOW())", email, hashedPassword)
	return err
}

func sendConfirmationEmail(email string) error {
	// You can replace this with a proper SMTP service like Mailgun, SendGrid, etc.
	from := "noreply@yourapp.com"
	pass := "smtp-password"
	to := email
	msg := "Subject: Confirm your account\n\nClick to confirm your account."

	err := smtp.SendMail("smtp.example.com:587",
		smtp.PlainAuth("", from, pass, "smtp.example.com"),
		from, []string{to}, []byte(msg))
	return err
} */
