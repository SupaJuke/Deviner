package users

import (
	"fmt"
	"log"
	"math/big"

	"crypto/rand"

	db "github.com/SupaJuke/Deviner/go/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ------------------------- CRUD functionalities -------------------------

func Create(username string, pwd string) (User, error) {
	user := User{}
	query := "INSERT INTO Users(username, password) VALUES($1, $2)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return user, err
	}
	defer stmt.Close()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return user, err
	}

	_, err = stmt.Exec(username, string(hashedPwd))
	if err != nil {
		log.Println(err)
		return user, err
	}

	user = User{Username: username, Password: string(hashedPwd)}
	logStr := fmt.Sprintf("User %s created", user.Username)
	log.Println(logStr)
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
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		return User{}, err
	}

	return user, nil
}

func (user User) UpdatePwd() error {
	query := "UPDATE User SET password = $1 WHERE username = $2"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec(string(hashedPwd), user.Username)
	if err != nil {
		log.Println(err)
		return err
	}

	logStr := fmt.Sprintf("User %s updated", user.Username)
	log.Println(logStr)
	return nil
}

func (user User) Delete() error {
	query := "DELETE FROM Users WHERE username = $1"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username)
	if err != nil {
		log.Println(err)
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
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
}

func (user User) GetCode() (string, error) {
	query := "SELECT code FROM Users WHERE username = $1"
	code := ""
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return code, err
	}
	defer stmt.Close()

	row, err := stmt.Query(user.Username)
	if err != nil {
		return code, err
	}
	defer row.Close()

	row.Next()
	if err := row.Scan(&code); err != nil {
		return code, err
	}

	return code, nil
}

func (user User) GenerateNewCode() error {
	// Randomizing a new code
	bigI, err := rand.Int(rand.Reader, big.NewInt(100000))
	if err != nil {
		log.Println("Failed to generate new code for user: ", user.Username)
		return err
	}

	// Inserting the new code to DB
	code := bigI.String()
	query := "UPDATE User SET code = $1 WHERE username = $2"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println("Error while preparing statement [GenerateNewCode]")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(code, user.Username)
	if err != nil {
		log.Println("Error while updating new code")
		return err
	}

	// Successfully updated the new code
	logStr := fmt.Sprintf("Updated user %s code to %s", user.Username, code)
	log.Println(logStr)
	return nil
}
