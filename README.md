# gin-crud-api

Containerized Go + Gin CRUD API for managing books.

## Quick Run 

1. **Install Golang**
   ```
   https://go.dev/doc/install
   ```
2. **Install Gin**
   ```
   https://go.dev/doc/install
   ```
3. **Run Commad**
   ```
   go run main.go
   ```
## Docker Test

1. **Run a docker container**:


```
  docker run -it -p 8080:8080 --name my-ubuntu-vps-go -v /app/go-project ubuntu:latest /bin/bash
```


2. **Install Required Packages**:

   ```bash
   apt update
   apt install -y golang
   go get github.com/gin-gonic/gin
   apt install -y nano
   ```
3. **Project Setup**:
   ```bash
   touch main.go
   go mod init go-project
   nano main.go // past the code of main.go
   cat main.go // check the code has been saved or not

   go run main.go
   
   ```

## Test Endpoint

   - List books: `GET http://localhost:8080/books`
   - Create book: `POST http://localhost:8080/books`
     ```json
     { "title": "My Book", "author": "Me", "year": 2025 }
     ```
   - Get book: `GET http://localhost:8080/books/1`
   - Update book: `PUT http://localhost:8080/books/1`
     ```json
     { "title": "Updated", "author": "Me", "year": 2026 }
     ```
   - Delete book: `DELETE http://localhost:8080/books/1`

## License

MIT © Siratul Islam
