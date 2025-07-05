package postgres

type User struct {
	ID           string
	Email        string
	PasswordHash string
}

func GetUserByEmail(email string) (*User, error) {
	query := `SELECT id, email, password_hash FROM users WHERE email=$1`
	row := DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
