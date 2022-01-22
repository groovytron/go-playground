# Go Playground

## Run the API

### Requirements

You first need to have the following softwares installed:

- `make`
- `docker` and `docker-compose`
- [`air`](https://github.com/cosmtrek/air) (it must be in your `$PATH`)

### Start the API

Do the following to run the API:

1. Copy the `.env.example` file into `.env` (only the first time you want to start the API)
1. Run `docker-compose up` in one terminal (to start the database)
2. Run `make run` in another terminal (to start the API)

The API will be available on <http://localhost:9000>.

PostgreSQL is available on `localhost:5432`.

[Adminer](https://www.adminer.org/) is available on <http://localhost:8080> to allow you to access the database and insert data.

## Useful links

- [Write a Simple REST API in Golang](https://dev.to/lucasnevespereira/write-a-rest-api-in-golang-following-best-practices-pe9)
