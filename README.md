# Gofiber Repository Pattern

## Tech Stack
- [Golang](https://github.com/golang/go)
- [MySQL (Database)](https://github.com/mysql/mysql-server)

## Framework
- [GoFiber (HTTP Framework)](https://github.com/gofiber/fiber)
- [Viper (Configuration)](https://github.com/spf13/viper)
- [Go Playground Validator (Validation)](https://github.com/go-playground/validator)
- [Golang Migration](https://github.com/golang-migrate/migrate)

## Development Tools
- [air (hot reload)](https://github.com/cosmtrek/air)

## How to install
```bash
go get .
```

## Run with air (hot reload)
```bash
air
```

## Create Migration
```bash
migrate create -ext sql -dir db/migrations create_table_xxx
```

## Run Migration
```bash
migrate -database "mysql://root:password@tcp(localhost:3306)/golang_migration?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
```

## Architecture
![alt text](https://raw.githubusercontent.com/prioarief/gofiber-repository-pattern/120f16f8333800c067d5aa23797c566eaca2e0b4/diagram.svg)



