# Social Media Backend API

[![Go Version](https://img.shields.io/badge/Go-1.24.1-blue.svg)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://docker.com)
[![Educational](https://img.shields.io/badge/License-Educational--Use-orange.svg)](#credits--license)

A Go-based social media backend API built with modern web development practices.

> ⚠️ **Disclaimer**: This project is based on a tutorial by [Tiago](https://youtu.be/h3fqD6IprIA?si=MyOXi4flYJrDWH0k). It was recreated as a personal learning exercise with modifications. Please refer to the original author's work for the official version.

## Why I Built This

This project was built to strengthen my skills in Go backend development, PostgreSQL, and containerized deployments. It demonstrates my ability to learn from tutorials, implement improvements, and create production-style applications for learning purposes.

## Project Overview

This is a learning project that demonstrates building a RESTful API using Go, featuring:
- User management system
- Post management system
- PostgreSQL database with Docker
- Chi router for HTTP handling
- Environment-based configuration
- Database migrations

## How This Project Differs from the Original

While following the tutorial structure, I've made several modifications and improvements:
- Enhanced error handling and logging
- Improved database connection management
- Added comprehensive documentation
- Customized the project structure for better maintainability
- Fixed several implementation details to match the User struct requirements

## Features

- **User Management**: Create, read, update, and delete users
- **Post Management**: Create and manage social media posts
- **Database**: PostgreSQL with proper migrations
- **API**: RESTful endpoints with proper HTTP status codes
- **Docker**: Containerized development environment
- **Environment Configuration**: Flexible configuration management

## Tech Stack

- **Language**: Go 1.24.1
- **Framework**: Chi router
- **Database**: PostgreSQL 16.3
- **Containerization**: Docker & Docker Compose
- **Environment**: dotenv for configuration

## Prerequisites

- Go 1.24+ installed
- Docker and Docker Compose installed
- Git

## Quick Start

### 1. Clone the Repository
```bash
git clone <your-repo-url>
cd social
```

### 2. Start the Services
```bash
docker-compose up -d
```

This will start:
- PostgreSQL database on port 5433
- Go API server on port 8081

### 3. Set Up Environment Variables
```bash
# Create a .env file with your database credentials
echo "POSTGRES_PASSWORD=your_secure_password_here" > .env
echo "DB_PASSWORD=your_secure_password_here" >> .env
```

### 4. Run Database Migrations
```bash
# The database schema will be automatically initialized
# Check the scripts/schema.sql file for the schema
```

### 5. Test the API
```bash
# Health check
curl http://localhost:8081/health

# The API will be available at http://localhost:8081
```

## Project Structure

```
social/
├── cmd/                    # Application entry points
│   ├── api/              # Main API server
│   └── migrate/          # Database migration tool
├── internal/              # Private application code
│   ├── db/               # Database connection
│   ├── env/              # Environment configuration
│   └── storage/          # Data access layer
├── scripts/               # Database scripts and migrations
├── docker-compose.yml     # Docker services configuration
└── Dockerfile            # Go application container
```

## Configuration

The application uses environment variables for configuration. Create a `.env` file in the project root:

### Required Environment Variables:
- `POSTGRES_PASSWORD`: Password for PostgreSQL database
- `DB_PASSWORD`: Password for database connections

### Optional Environment Variables:
- `ADDR`: Server address (default: :8081)
- `DB_ADDR`: Database connection string (auto-generated from DB_PASSWORD)
- `DB_MAX_OPEN_CONNS`: Maximum open database connections (default: 10)
- `DB_MAX_IDLE_CONNS`: Maximum idle database connections (default: 5)
- `DB_MAX_IDLE_TIME`: Maximum idle time for connections in seconds (default: 60)

## API Endpoints

### Health Check
- `GET /v1/healthcheck` - Health check endpoint

### User Management (Storage Layer)
- User creation, reading, updating, and deletion
- Secure password handling
- Email and username validation

### Post Management (Storage Layer)
- Post creation and management
- User-post relationships
- Content management

## Testing the API

### Health Check
```bash
curl http://localhost:8081/v1/healthcheck
```

Expected response:
```json
{
  "status": "ok",
  "message": "Everything is ok"
}
```

### Using Docker
```bash
# Start the services
docker-compose up -d

# Check logs
docker-compose logs -f app

# Test health endpoint
curl http://localhost:8081/v1/healthcheck
```

## Docker Commands

```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Rebuild and restart
docker-compose up --build -d
```

## Development

### Running Locally (without Docker)
```bash
# Set up environment variables
export DB_ADDR="host=localhost port=5433 user=postgres password=your_secure_password_here dbname=social sslmode=disable"

# Run the application
go run cmd/api/main.go
```

### Hot Reload (with Air)
```bash
# Install Air if you haven't
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

## Notes

- The database will be automatically initialized when you first run the Docker containers
- **Important**: You must create a `.env` file with your database credentials before starting
- The application waits 5 seconds for the database to be ready on startup
- All sensitive data is stored in environment variables, not in code

## Troubleshooting

### Common Issues

#### Database Connection Failed
```bash
# Check if PostgreSQL is running
docker-compose ps

# View database logs
docker-compose logs db

# Ensure .env file exists and has correct passwords
cat .env
```

#### Port Already in Use
```bash
# Check what's using port 8081
netstat -ano | findstr :8081

# Or use different ports in docker-compose.yml
```

#### Permission Denied
```bash
# On Windows, run PowerShell as Administrator
# On Linux/Mac, check file permissions
ls -la
```

### Performance Tips

- **Database Connections**: Adjust `DB_MAX_OPEN_CONNS` based on your load
- **Memory**: PostgreSQL container uses ~100MB RAM by default
- **Storage**: Database data persists in Docker volume `db-data`

## Contributing

This is a learning project, but feel free to:
- Report issues
- Suggest improvements
- Fork and experiment

## Future Enhancements

### Planned Features
- [ ] Complete CRUD API endpoints for users and posts
- [ ] JWT authentication system
- [ ] Input validation and sanitization
- [ ] Rate limiting and security headers
- [ ] Comprehensive test coverage
- [ ] API documentation with Swagger/OpenAPI

### Getting Involved
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## Credits & License

This project was built as a learning exercise by following [Tiago's YouTube tutorial](https://youtu.be/h3fqD6IprIA?si=MyOXi4flYJrDWH0k). I implemented my own version, added improvements, and fixed bugs. 

**Original Source**: The tutorial and original source code are available on [Tiago's YouTube channel](https://youtu.be/h3fqD6IprIA?si=MyOXi4flYJrDWH0k) and [GitHub repository](https://github.com/sikozonpc/GopherSocial).

**License**: This project is for educational purposes only. This repository should not be used for commercial purposes. Please respect the original tutorial author's work and refer to their repository for the official version.

---

Happy coding! 