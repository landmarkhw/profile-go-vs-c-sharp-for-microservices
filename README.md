# Profiling Comparison, C# vs Golang

# Getting Started

1. [Ensure you have Docker installed.](https://docs.docker.com/install/)
1. Open a command prompt and run `docker-compose up`.  This will:
    * Download and run PostgreSQL
        * Setup a default database with some tables inside Postgres
        * Insert some default data
    * Download and run pgadmin at http://localhost:8080
        * You can connect using `profile-postgres` as the hostname, `profiler` as the username, `profile-test` as the password.
    * Run the go app at http://localhost:5000

# Running the profiling tests

1. [Ensure you have Go installed.](https://golang.org/doc/install)
1. [Install Bombardier](https://github.com/codesenberg/bombardier) - `go get -u github.com/codesenberg/bombardier`
1. The following Bombardier commands give interesting results:

<br/>`bombardier -c 100 -n 100000 -l http://localhost:5000/person/1` - Get a `Person` from the db
<br/>`bombardier -c 10 -n 10000 -l http://localhost:5000/person/1` - Get a `Person` from the db
<br/>`bombardier -c 1000 -n 100000 -l http://localhost:5000/test/json/simple` - Get a simple JSON object
<br/>`bombardier -c 1000 -n 100000 -l http://localhost:5000/test/json/complex` - Get a complex JSON object
<br/>`bombardier -c 100 -n 10000 -l --method=PUT --body-file=./test-person.json --header="Content-Type: application/json" http://localhost:5000/person` - Save a person to the db
