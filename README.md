# Sunday Project

A learning project built with Go and MongoDB to practice backend
development, HTTP handlers, and session-based authentication.

## Features

- User registration
- Password hashing with bcrypt
- Login with a username and password
- Server-side sessions stored in MongoDB
- Cookie-based authentication
- A protected profile endpoint
- HTML pages rendered with Go templates

## Requirements

- A Go version compatible with go.mod
- MongoDB
- mongosh for database setup

## Configuration

The application reads configuration from environment variables.
If a variable is not set or is empty, the application uses its default value.

| Variable | Description | Default |
|----------|-------------|---------|
| MONGO_URI | MongoDB connection URI | mongodb://localhost:27017 |
| DATABASE_NAME | MongoDB database name | sundayProject |
| PORT | HTTP server port, without a colon | 9999 |

For example, in Bash or Zsh:

```bash
export MONGO_URI="mongodb://localhost:27017"
export DATABASE_NAME="sundayProject"
export PORT="9999"
```

Run the application from the same terminal where these variables were set.

## Database Setup

Make sure MongoDB is running, then open the MongoDB Shell:

```bash
mongosh
```

Select the application database:

```javascript
use sundayProject
```

If you use a different DATABASE_NAME, select that database instead.

Create a unique index on the login field in the accounts collection:

```javascript
db.accounts.createIndex({ login: 1 }, { unique: true })
```

This index prevents multiple accounts from using the same login.

The index is stored in MongoDB and does not need to be recreated
every time the application starts. It must be configured for each
new database used by the application.

To inspect existing indexes:

```javascript
db.accounts.getIndexes()
```

Check that the login index contains:

```javascript
key: { login: 1 }
unique: true
```

If the collection already contains duplicate login values,
resolve them before creating the unique index.

## Running the Application

Run the following commands from the project root,
where go.mod is located:

```bash
go mod download
go run ./cmd/server
```

By default, the application is available at:

http://localhost:9999

If you change PORT, use the corresponding port in your browser.

## Routes

| Method | Path | Description |
|--------|------|-------------|
| GET | / | Home page |
| GET | /register | Registration page |
| POST | /register | Create an account |
| GET | /login | Login page |
| POST | /login | Authenticate and create a session |
| GET | /profile | Protected profile endpoint |

## Project Structure

```text
.
├── auth/          # Session token generation
├── cmd/server/    # Application entry point
├── database/      # MongoDB connection
├── handlers/      # HTTP handlers and authentication middleware
├── models/        # Account and session models
├── templates/     # HTML templates
├── go.mod
├── go.sum
└── README.md
```

## Development Notes

This project is intended for learning and local development.
It is not production-ready.

The session cookie currently uses Secure: false for local HTTP
development. HTTPS deployments should use Secure: true.