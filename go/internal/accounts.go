package accounts

import (
	"fmt"
	"log"

	db "github.com/SupaJuke/pooe-guessing-game/go/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID  string `json:"id"`
	Pwd string `json:"password"`
}

func GetByID(id string) (User, error) {
	query := "SELECT id, password FROM Users WHERE id = ? LIMIT 1"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	row, err := stmt.Query(id)
	if err != nil {
		return User{}, err
	}
	defer row.Close()

	user := User{}
	row.Next()
	if err := row.Scan(&user.ID, &user.Pwd); err != nil {
		return User{}, err
	}

	return user, err
}

func (user User) Create() error {
	query := "INSERT INTO Users(id, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = stmt.Exec(user.ID, string(hashedPwd))
	if err != nil {
		log.Fatal(err)
		return err
	}

	logStr := fmt.Sprintf("User %s created", user.ID)
	log.Println(logStr)
	return nil
}

func (user User) Update() error {
	query := "UPDATE User SET password = ? WHERE id = ?"
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

	_, err = stmt.Exec(string(hashedPwd), user.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	logStr := fmt.Sprintf("User %s updated", user.ID)
	log.Println(logStr)
	return nil
}

func (user User) Delete() error {
	query := "DELETE FROM Users WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	logStr := fmt.Sprintf("User %s deleted", user.ID)
	log.Println(logStr)
	return nil
}

/*
	// albumsByArtist queries for albums that have the specified artist name.
	func albumsByArtist(name string) ([]Album, error) {
		// An albums slice to hold data from returned rows.
		var albums []Album

		rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
		if err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		defer rows.Close()
		// Loop through rows, using Scan to assign column data to struct fields.
		for rows.Next() {
			var alb Album
			if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
				return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
			}
			albums = append(albums, alb)
		}
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		return albums, nil
	}
*/
