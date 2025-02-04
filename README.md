# Golang (DDD)

A Go project structured with Domain-Driven Design (DDD) principles and leveraging Fx for dependency injection and lifecycle management. The project uses Fiber as the web framework, providing a simple, fast, and scalable HTTP server.

## Technologies

**Go**: Programming language.

**Fx**: Dependency injection and lifecycle management.

**Fiber**: Web framework for building APIs.

**PostgreSQL**: Database integration.

**UUID**: For generating unique identifiers.

**Viper**: For managing environment variables.

**Testify**: For testing utilities.

**Docker**: For containerizing the application and managing services.

## Project Structure

```
├── cmd
│   └── main.go                # Application entry point
│
├── internal
│   ├── infra                  # Infrastructure layer
│   │   ├── db                 # Database connection setup
│   │   ├── id-generator       # UUID generation logic
│   │   └── env                # environment loader
│   │
│   ├── user                   # User domain
│   │   ├── application        # Application services
│   │   ├── domain             # Domain entities and interfaces
│   │   ├── infra              # User repository implementation
│   │   └── interface          # HTTP handlers for user operations
│
├── migrations                 # Database migration scripts
│
├── docker-compose.yml         # Docker configuration file
│
└── README.md
```

## Layers Explanation

**Domain Layer**: Contains core business logic and entities.

**Application Layer**: Orchestrates domain logic for use cases.

**Infrastructure Layer**: Manages external systems like databases and UUID generation.

**Interface Layer**: Handles HTTP requests and responses.

## Running the Project

### Install dependencies

```
$ go mod tidy
```

### Start the application with Docker:

```
$ docker-compose up --build
```

This will spin up the application along with PostgreSQL and a service dedicated to running migrations using migrate.

### Run migrations manually (if needed):

```
$ docker-compose run migrate
```

## Testing

The project uses Testify for testing. Example test structure:

### Run tests with

```
$ go test ./...
```
