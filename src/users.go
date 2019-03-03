package web_forum

import (
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Hash struct{}

type User struct {
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

func registerAccount(username string, email string, passwordHash string) {

}

func (c *Hash) Generate(s string) (string, error) {
	saltedBytes := []byte(s)

	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

func (c *Hash) Compare(hash string, s string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
	return err == nil
}

func signupPageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {

		username := r.Form["username"][0]
		email := r.Form["email"][0]
		password := r.Form["password"][0]

		if len(username) != 0 && len(email) != 0 && len(password) != 0 {
			// hash password

			hashedPassword := password

			_, err := db.Query("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
				username, email, hashedPassword)

			if err != nil {
				fmt.Fprintf(w, "error: Username exists\n\n%s", err)
			} else {
				fmt.Fprintf(w, "Account created")
			}
		}
	} else {
		frontPage := getFrontPageData()
		tmpl, err := template.ParseFiles("template/signup.html")

		if err != nil {
			fmt.Fprint(w, "503: ", err)
		} else {
			tmpl.Execute(w, frontPage)
		}
	}
}
