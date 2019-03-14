package web_forum

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
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

func registerAccount(username string, email string, password string) error {

	// validate inputs on the server because
	// that's what you're supposed to do or something

	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]{1,16}", username)
	if matched == false {
		return errors.New("Invalid username")
	}

	matched, _ = regexp.MatchString("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$", email)
	if matched == false {
		return errors.New("Invalid email address")
	}

	matched, _ = regexp.MatchString("^[a-zA-Z0-9]{12,512}", password)
	if matched == false {
		return errors.New("Invalid password")
	}

	// because the database is UNIQUE, an error will be raised if someone tries to create
	// an account with a username which is already present in the database

	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		username, email, password)

	if err != nil {
		return errors.New("Username exists")
	}

	return nil
}

func signupPageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {

		username := r.Form["username"][0]
		email := r.Form["email"][0]
		password := r.Form["password"][0]

		if len(username) != 0 && len(email) != 0 && len(password) != 0 {
			err := registerAccount(username, email, password)

			if err != nil {
				fmt.Fprintf(w, "error: %s", err)
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
