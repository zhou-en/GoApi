# GoApi
This project creates the following API endpoints using Go, GORM and Gorilla Mux
- `employees/`
- `employees/{title}/`
- `employees/{title}/disable/`
- `employees/{title}/enable/`

## Environmental Variables
```shell
export DB_NAME="<database name>"
export DB_USER="<database username>"
export DB_PASSWORD="<database password>"
```

## Running service
```shell
go build
go run main.go
```

## Examples of API Usages
- Create an employee
```shell
curl -d '{"name":"Terry","city":"Toronto","age":42,"status":false}' -H 'Content-Type: application/json' -X "POST" 127.0.0.1:8888/employees
```
- Update an employee
```shell
curl -d '{"name":"Terry2"}' -H 'Content-Type: application/json' -X "PUT" 127.0.0.1:8888/employees/Terry
```
- Retrieve an employee
```shell
curl -H 'Content-Type: application/json' -X "GET" 127.0.0.1:8888/employees/Terry
```
- Delete an employee
```shell
curl -H 'Content-Type: application/json' -X "DELETE" 127.0.0.1:8888/employees/Mark
```
