package pfdb

import (
	"database/sql"
	"fmt"

	// The "_" allows this package to register itself with database/sql, but doesn't expose the package to our code
	_ "github.com/lib/pq"
)

const (
	host     = "profile-postgres"
	port     = 5432
	user     = "profiler"
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
func WithDb(fn dbaction) (err error) {
	connectionString := getConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	// Ensure the database gets closed when we're done
	defer db.Close()

	// Recover from any panic()s inside the dbaction function
	defer func() {
		if r := recover(); r != nil {
			err = r.(error) // typecast to error
		}
	}()

	fn(db)
	return err
}

// Test a connection to the database.
func Test() bool {
	err := WithDb(func(db *sql.DB) {
		err := db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully connected!")
	})

	if err != nil {
		// An error occurred
		return false
	}
	return true
}
