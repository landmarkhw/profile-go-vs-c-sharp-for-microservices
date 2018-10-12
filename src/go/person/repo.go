package person

import (
	"database/sql"

	"github.com/landmarkhw/profile-go-vs-c-sharp-for-microservices/database"
)

// Get the person with the given id.
func Get(id int64) *Person {
	var firstName, lastName string

	err := pfdb.WithDb(func(db *sql.DB) {
		sql := `
			SELECT Id, FirstName, LastName
			FROM Person
			WHERE Id = $1`

		err := db.QueryRow(sql, id).Scan(&id, &firstName, &lastName)
		if err != nil {
			panic(err)
		}
	})

	switch err {
	case sql.ErrNoRows:
	case nil:
		return &Person{ID: id, FirstName: firstName, LastName: lastName}
	default:
		panic(err)
	}

	return nil
}

// Save a person into the database.
func Save(person Person) bool {
	id := 0
	pfdb.WithDb(func(db *sql.DB) {
		sql := `
			INSERT INTO Person (FirstName, LastName)
			VALUES ($1, $2)
			RETURNING Id`
		err := db.QueryRow(sql, "Test", "User").Scan(&id)
		if err != nil {
			panic(err)
		}
	})

	// If the resulting Id is valid, then the object
	// was successfully inserted
	return id > 0
}
