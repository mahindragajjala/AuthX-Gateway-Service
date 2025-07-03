
package auth

type Authenticator interface {
    Login(email, password string) bool
}

type RealAuthenticator struct{}

func (r *RealAuthenticator) Login(email, password string) bool {
    return email == "admin@example.com" && password == "1234"
}
