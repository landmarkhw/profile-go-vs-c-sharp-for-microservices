# Profiling Comparison, C# vs Golang

# Getting Started

1. [Ensure you have Docker installed.](https://docs.docker.com/install/)
1. Open a command prompt and run `docker-compose up`.  This will download PostgreSQL, run it in a container, and setup some default data.

# Running the profiling tests

1. [Ensure you have Go installed.](https://golang.org/doc/install)
1. [Install Bombardier](https://github.com/codesenberg/bombardier) - `go get -u github.com/codesenberg/bombardier`
1. The following Bombardier commands give interesting results:

`bombardier -c 100 -n 100000 http://localhost:5000/person/1` - Get a `Person` from the db, 100 users, 100000 calls
`bombardier -c 10 -n 10000 http://localhost:5000/person/1` - Get a `Person` from the db, 10 users, 10000 calls
`bombardier -c 1000 -n 100000 http://localhost:5000/test/json/simple` - Get a simple JSON object, 1000 users, 100000 calls
`bombardier -c 1000 -n 100000 http://localhost:5000/test/json/complex` - Get a complex JSON object, 1000 users, 100000 calls
