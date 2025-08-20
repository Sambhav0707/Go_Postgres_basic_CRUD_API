# Go PostgreSQL CRUD API

A basic CRUD (Create, Read, Update, Delete) API built with Go (Golang) that connects to PostgreSQL database using `database/sql` and `lib/pq` driver.

## 🚀 Features

- **RESTful API** with standard CRUD operations
- **PostgreSQL Integration** using `database/sql` and `lib/pq`
- **Environment Configuration** using `.env` files
- **Gorilla Mux Router** for HTTP routing
- **JSON Response Format** for all API endpoints
- **Stock Management System** as an example implementation

## 📋 Prerequisites

- Go 1.24.5 or higher
- PostgreSQL database
- Git

## 🛠️ Installation & Setup

### 1. Clone the Repository
```bash
git clone <your-repository-url>
cd go_postgres_project
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Database Setup

Create a PostgreSQL database and run the following SQL to create the stocks table:

```sql
CREATE TABLE stocks (
    stockid SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price BIGINT NOT NULL,
    company VARCHAR(255) NOT NULL
);
```

### 4. Environment Configuration

Create a `.env` file in the root directory:

```env
POSTGRES=postgres://username:password@localhost:5432/database_name?sslmode=disable
```

**Replace the following:**
- `username`: Your PostgreSQL username
- `password`: Your PostgreSQL password
- `localhost:5432`: Your PostgreSQL host and port
- `database_name`: Your database name

### 5. Run the Application
```bash
go run main.go
```

The server will start on `http://localhost:8081`

## 📁 Project Structure

```
go_postgres_project/
├── main.go              # Application entry point
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
├── .env                 # Environment variables (create this)
├── middleware/
│   └── handler.go       # HTTP handlers and database operations
├── models/
│   └── models.go        # Data structures and models
└── router/
    └── router.go        # HTTP routing configuration
```

## 🔌 API Endpoints

### Stock Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/stocks` | Get all stocks |
| `GET` | `/api/stocks/{id}` | Get stock by ID |
| `POST` | `/api/newstock` | Create a new stock |
| `PUT` | `/api/stocks/{id}` | Update stock by ID |
| `DELETE` | `/api/deletestock/{id}` | Delete stock by ID |

## 📝 API Usage Examples

### 1. Create a New Stock
```bash
curl -X POST http://localhost:8081/api/newstock \
  -H "Content-Type: application/json" \
  -d '{
    "name": "AAPL",
    "price": 150,
    "company": "Apple Inc."
  }'
```

**Response:**
```json
{
  "id": 1,
  "message": "Stock created successfully"
}
```

### 2. Get All Stocks
```bash
curl -X GET http://localhost:8081/api/stocks
```

**Response:**
```json
[
  {
    "stockid": 1,
    "name": "AAPL",
    "price": 150,
    "company": "Apple Inc."
  },
  {
    "stockid": 2,
    "name": "GOOGL",
    "price": 2800,
    "company": "Alphabet Inc."
  }
]
```

### 3. Get Stock by ID
```bash
curl -X GET http://localhost:8081/api/stocks/1
```

**Response:**
```json
{
  "stockid": 1,
  "name": "AAPL",
  "price": 150,
  "company": "Apple Inc."
}
```

### 4. Update Stock
```bash
curl -X PUT http://localhost:8081/api/stocks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "AAPL",
    "price": 155,
    "company": "Apple Inc."
  }'
```

**Response:**
```json
{
  "id": 1,
  "message": "Stock updated successfully. Total rows/record affected 1"
}
```

### 5. Delete Stock
```bash
curl -X DELETE http://localhost:8081/api/deletestock/1
```

**Response:**
```json
{
  "id": 1,
  "message": "Stock deleted successfully. Total rows/record affected 1"
}
```

## 🏗️ Architecture

### Key Components

1. **Models** (`models/models.go`)
   - Defines the `Stock` struct with JSON tags
   - Represents the database schema

2. **Router** (`router/router.go`)
   - Uses Gorilla Mux for HTTP routing
   - Defines RESTful endpoints with proper HTTP methods

3. **Middleware** (`middleware/handler.go`)
   - Contains all HTTP handlers
   - Database connection management
   - CRUD operations implementation

4. **Main** (`main.go`)
   - Application entry point
   - Server initialization and startup

### Database Connection

The application uses:
- `database/sql` package for database operations
- `lib/pq` driver for PostgreSQL connectivity
- Environment variables for database configuration
- Connection pooling and proper resource management

## 🔧 Dependencies

- `github.com/gorilla/mux` - HTTP router and URL matcher
- `github.com/lib/pq` - PostgreSQL driver
- `github.com/joho/godotenv` - Environment variable loader

## 🚨 Error Handling

The application includes comprehensive error handling for:
- Database connection failures
- Invalid JSON requests
- Missing or invalid parameters
- Database query errors

## 🔒 Security Considerations

- Use environment variables for sensitive database credentials
- Implement proper input validation in production
- Consider adding authentication and authorization
- Use HTTPS in production environments

## 🧪 Testing

To test the API endpoints, you can use:
- cURL commands (as shown in examples)
- Postman or similar API testing tools
- Automated testing frameworks like `httptest`

## 📈 Future Enhancements

- Add authentication and authorization
- Implement request validation
- Add logging and monitoring
- Create automated tests
- Add pagination for large datasets
- Implement caching mechanisms

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 👨‍💻 Author

Created with ❤️ using Go and PostgreSQL

---

**Note:** Make sure to replace the database connection string in the `.env` file with your actual PostgreSQL credentials before running the application.
