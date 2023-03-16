package accounts

import (
	db "github.com/SupaJuke/pooe-guessing-game/go/internal/db"
)

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

func GetById(id string) (User, error) {
	query := "SELECT id, password FROM Users WHERE id = ? LIMIT 1"
	stmt, err := db.Db.Prepare(query)
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
	if err := row.Scan(&user.Id, &user.Password); err != nil {
		return User{}, err
	}

	return user, err
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
