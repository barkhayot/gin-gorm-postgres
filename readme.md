# Go CRUD Server using Gin + Gorm + PostgeSQL

This is a simple CRUD (Create, Read, Update, Delete) server implemented in Go, using the Gin web framework and Gorm ORM with PostgreSQL as the database.

  

## Setup

Before running the server, make sure to set up the required environment variables and initialize the database.

  

### Environment Variables

Create a .env file in the project root and set the following variables:
  
``` 
DB_USER=your_postgres_user

DB_PASSWORD=your_postgres_password

DB_NAME=your_database_name

DB_HOST=localhost

DB_PORT=5432 
```

### Initialize Database

Run the following commands to apply the database migrations:

```
go run main.go migrate
```

### Run the Server

To start the server, run the following command:
```
go run main.go
```

The server will start at http://localhost:8080 by default.

  


### API Endpoints
- `GET /posts`: Get all posts.
- `GET /posts/:id`: Get post by ID.
- `POST /posts`: Create a new post.
- `PUT /posts/:id`: Update post by ID
- `DELETE /posts/:id`: Delete post by ID

  

Dependencies

- <b>Gin</b> - Web framework
- <b>Gorm</b> - ORM library
- <b>PostgreSQL</b> - Database
