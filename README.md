# go-hex

A simple web API illustrating [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) or Martin's [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) in Go.

## Structure

This project is built as a Go module. All code is privy to this module under the `internal` folder.

Main directories are:
- `domain` is the core application containing domain entities and interfaces. It does not depend on any external implementations or framework as it should be.
- `rest` contains HTTP handlers to process API requests
- `storage` contains implementations of the repository interface declared in the `domain` folder. The only coded implementation is a SQLite3 in-memory database. 

```
.
├── Makefile
├── README.md
├── bin
│   ├── main-darwin-amd64
│   ├── main-linux-amd64
│   └── main-windows-amd64
├── cmd
│   └── go-hex
│       └── main.go
├── go.mod
├── go.sum
└── internal
    └── user
        ├── domain
        │   ├── model.go
        │   ├── repository.go
        │   └── service.go
        ├── rest
        │   ├── handler.go
        │   └── message.go
        └── storage
            └── sqlite
                └── user.go
```

## Libraries

This project makes use of:
- Gin
- Gorm

## Running

```
make run
curl -X POST http://localhost:8080/user -d '{"name": "tester"}'
# {"user":{"id":"f508278b-d502-4621-abb0-2c218d564b73","name":"tester","created":"2020-05-21T20:37:43.636748+02:00"}}%
curl http://localhost:8080/user/f508278b-d502-4621-abb0-2c218d564b73
# {"id":"f508278b-d502-4621-abb0-2c218d564b73","name":"tester","created":"2020-05-21T20:37:43.636748+02:00"}
```
