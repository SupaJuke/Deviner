package accounts

import (
	"fmt"
	"log"

	db "github.com/SupaJuke/Deviner/go/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"Username"`
	Pwd      string `json:"password"`
}

// ------------------------- CRUD functionalities -------------------------

func Create(username string, pwd string) (User, error) {
	user := User{}
	query := "INSERT INTO Users(username, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	defer stmt.Close()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return user, err
	}

	_, err = stmt.Exec(username, string(hashedPwd))
	if err != nil {
		log.Fatal(err)
		return user, err
	}

	logStr := fmt.Sprintf("User %s created", user.Username)
	log.Println(logStr)
	user = User{username, string(hashedPwd)}
	return user, nil
}

func GetByUsername(username string) (User, error) {
	query := "SELECT username, password FROM Users WHERE username = $1 LIMIT 1"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()
	row, err := stmt.Query(username)
	if err != nil {
		return User{}, err
	}
	defer row.Close()

	user := User{}
	row.Next()
	if err := row.Scan(&user.Username, &user.Pwd); err != nil {
		return User{}, err
	}

	return user, nil
}

func (user User) Update() error {
	query := "UPDATE User SET password = ? WHERE username = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(string(hashedPwd), user.Username)
	if err != nil {
		log.Fatal(err)
		return err
	}

	logStr := fmt.Sprintf("User %s updated", user.Username)
	log.Println(logStr)
	return nil
}

func (user User) Delete() error {
	query := "DELETE FROM Users WHERE username = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username)
	if err != nil {
		log.Fatal(err)
		return err
	}

	logStr := fmt.Sprintf("User %s deleted", user.Username)
	log.Println(logStr)
	return nil
}

// ------------------------------------------------------------------------

// Compares given password with the user's hashed password.
// Returns nil on successful, error otherwise
func (user User) Authenticate(pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd))
}
