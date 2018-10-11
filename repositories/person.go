package personrepo

import (
	"database/sql"
	"fmt"

	"landmarkhw.com/models"

	// The "_" allows this package to register itself with database/sql, but doesn't expose the package to our code
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "profile-test"
	dbname   = "test"
)

// Get a connection string suitable to connect to the database
func getConnectionString() string {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname)
	return connectionString
}

// A lambda expression that takes action on the database (while the connection is open)
type dbaction func(*sql.DB)

// withDb opens a connection to the database, performs the provided action, and closes the connection
func withDb(fn dbaction) {
	connectionString := getConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fn(db)
}

// Test a connection to the database.
func Test() bool {
	withDb(func(db *sql.DB) {
		err := db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully connected!")
	})

	if r := recover(); r != nil {
		// An error occurred
		return false
	}
	return true
}

// Get the person with the given id.
func Get(id int64) *models.Person {
	var err error
	var firstName, lastName string

	withDb(func(db *sql.DB) {
		sql := `
			SELECT Id, FirstName, LastName
			FROM Person
			WHERE Id = $1`

		err = db.QueryRow(sql, id).Scan(&id, &firstName, &lastName)
	})

	switch err {
	case sql.ErrNoRows:
	case nil:
		return nil
	default:
		panic(err)
	}

	return &models.Person{ID: id, FirstName: firstName, LastName: lastName}
}

// Save a person into the database.
func Save(person models.Person) bool {
	id := 0
	withDb(func(db *sql.DB) {
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
