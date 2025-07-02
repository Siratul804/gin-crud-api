# gin-crud-api

Containerized Go + Gin CRUD API for managing books.

## Quickstart

1. **Build** the Docker image (replace URL if needed):

   ```bash
   docker build \
     --build-arg REPO_URL=https://github.com/your-username/gin-crud-api.git \
     -t gin-crud-api:latest .
   ```

2. **Run** the container:

   ```bash
   docker run -d --name gin-api -p 8080:8080 gin-crud-api:latest
   ```

3. **Test** endpoints:
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

MIT © Your Name
