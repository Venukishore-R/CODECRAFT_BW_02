# Persistent Storage with Database Integration

## Task Overview
This project extends the previously developed REST API by integrating persistent storage using a relational database. The API now uses **PostgreSQL** for storing user data, providing robust data management and retrieval capabilities.

---

## Features
- Integration with a **PostgreSQL** database for persistent storage.
- Usage of an **ORM** (e.g., GORM) to simplify database operations.
- Implementation of database migrations to set up and manage the `users` table schema.
- Connection pooling for efficient database access.
- Environment-specific configurations managed using `.env` files.
- Secure handling of database credentials using environment variables.

---

## User Model
The `users` table schema:
- `id` (UUID): A unique identifier for each user.
- `name` (string): The name of the user.
- `email` (string): The user's email address (validated).
- `age` (integer): The age of the user.
- `created_at` (timestamp): The timestamp when the user was created.
- `updated_at` (timestamp): The timestamp when the user was last updated.

---

## Endpoints

| Method | Endpoint           | Description               |
|--------|--------------------|---------------------------|
| POST   | `/users`           | Create a new user         |
| GET    | `/users`           | Get all users             |
| GET    | `/users/:email`    | Get details of a user     |
| PUT    | `/users/:email`    | Update details of a user  |
| DELETE | `/users/:email`    | Delete a user by ID       |

---

## Requirements
- **Golang** for API implementation.
- **PostgreSQL** as the database.
- **GORM** (or any ORM of your choice) for interacting with the database.
- **godotenv** for loading environment variables.

---

## Setup Instructions

### Prerequisites
- Install [Go](https://golang.org/dl/).
- Install [PostgreSQL](https://www.postgresql.org/).

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/persistent-storage-api.git
   cd persistent-storage-api
   ```

2. Create a `.env` file with the following variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database
   ```

3. Install Go dependencies:
   ```bash
   go mod tidy
   ```

4. Start the server:
   ```bash
   go run cmd/main.go
   ```

---

## Usage
- Use tools like **Postman** or **curl** to interact with the API.
- Example: Create a new user
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com", "age": 30}' http://localhost:3000/users
  ```

---

## Technologies Used
- **Golang**: Programming language.
- **PostgreSQL**: Database.
- **GORM**: ORM for database interactions.
- **godotenv**: Environment variable loader.

---

## Error Handling
- `404 Not Found`: When a resource is not found.
- `400 Bad Request`: For invalid or missing inputs.

---

## Contributing
Contributions are welcome! Please fork the repository and create a pull request.

---

## License
This project is licensed under the MIT License.

---

## Author
Developed by [Venukishore-R](https://github.com/Venukishore-R).
