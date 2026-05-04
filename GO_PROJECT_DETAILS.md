# Go Project Details

A modern Go/Gin microservice API that provides access to university and school data.

## Features

- RESTful API endpoints for querying school data
- Structured logging using Go 1.21+ `log/slog`
- Comprehensive error handling with proper error wrapping
- Input validation for GUID parameters
- Health check endpoint
- Request ID tracking for debugging
- Environment variable configuration
- In-memory data caching for improved performance
- Graceful shutdown handling
- CORS support for cross-origin requests
- Comprehensive test suite with table-driven tests
- Thread-safe data access with mutex protection

## Prerequisites

- Go >= 1.21
- `data.json` file in the project root

## Installation

1. Install dependencies:

```bash
make deps
```

Or manually:

```bash
go mod download
go mod tidy
```

## Configuration

The application can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `3000` | Port number for the server to listen on |
| `DATA_FILE_PATH` | `./data.json` | Path to the JSON data file |
| `LOG_LEVEL` | `info` | Logging level (debug, info, warn, error) |
| `JSON_LOG` | `false` | Enable JSON log output format |
| `READ_TIMEOUT` | `10s` | HTTP read timeout |
| `WRITE_TIMEOUT` | `10s` | HTTP write timeout |

Create a `.env` file (optional) or set environment variables:

```bash
PORT=3000
DATA_FILE_PATH=./data.json
LOG_LEVEL=info
JSON_LOG=false
READ_TIMEOUT=10s
WRITE_TIMEOUT=10s
```

## Running the Application

### Development Mode

```bash
make run
```

Or manually:

```bash
go run ./cmd/server
```

### Production Mode

Build the binary first:

```bash
make build
```

Then run:

```bash
./bin/server
```

The server will start on the configured port (default: 3000).

## API Endpoints

### Health Check

|Route|Description|Status Code|
|-----|-----------|-----------|
|**GET** `/health`|Returns the health status of the service.|`200 OK`|

**Response:**

```json
{
  "status": "healthy"
}
```

### Get All Data

|Route|Description|Status Code|
|-----|-----------|-----------|
|**GET** `/`|Returns all school/university data.|`200 OK`|

**Response:**

```json
[
  {
    "guid": "05024756-765e-41a9-89d7-1407436d9a58",
    "school": "Iowa State University",
    "mascot": "Cy the Cardinal",
    "nickname": "Cyclones",
    "location": "Ames, IA, USA",
    "latlong": "42.026111,-93.648333",
    "ncaa": "Division I",
    "conference": "Big 12 Conference"
  },
  ...
]
```

### Get Item by GUID

|Route|Description|Status Code|
|-----|-----------|-----------|
|**GET** `/:guid`|Returns a single school/university item by GUID. Parameters: `guid` (path parameter) - Valid GUID format `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`|`200 OK`, `404 Not Found`, `400 Bad Request`|

**Parameters:**

- `guid` (path parameter) - Valid GUID in the format `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`

**Response (Success):**

```json
{
  "guid": "05024756-765e-41a9-89d7-1407436d9a58",
  "school": "Iowa State University",
  "mascot": "Cy the Cardinal",
  "nickname": "Cyclones",
  "location": "Ames, IA, USA",
  "latlong": "42.026111,-93.648333",
  "ncaa": "Division I",
  "conference": "Big 12 Conference"
}
```

**Response (Not Found):**

```json
{
  "error": "Data not found",
  "request_id": "uuid-here"
}
```

**Response (Invalid GUID Format):**

```json
{
  "error": "Invalid GUID format",
  "request_id": "uuid-here"
}
```

## Error Responses

All error responses follow a consistent format:

```json
{
  "error": "Error description",
  "request_id": "unique-request-id"
}
```

All responses include an `X-Request-ID` header for request tracking.

## Testing

### Run Tests

```bash
make test
```

Or manually:

```bash
go test -v ./...
```

### Run Tests with Race Detection

```bash
make test-verbose
```

Or manually:

```bash
go test -v -race ./...
```

### Test Coverage

```bash
make test-coverage
```

This generates a coverage report in HTML format (`coverage.html`).

The test suite includes:

- Health check endpoint tests
- Data retrieval tests
- Error handling tests (404, 400)
- Input validation tests
- GUID format validation tests
- Table-driven tests for multiple scenarios
- Integration tests (requires `data.json`)
- Mock service for unit testing

Test results can be integrated with CI/CD pipelines.

## Code Quality

### Formatting

Format code:

```bash
make fmt
```

Or manually:

```bash
go fmt ./...
```

### Linting

Check code style:

```bash
make lint
```

This runs `golangci-lint` if available, or falls back to `go vet`.

Or manually:

```bash
go vet ./...
```

### Clean Build Artifacts

```bash
make clean
```

This removes:
- `bin/` directory (build artifacts)
- `coverage.out` and `coverage.html` (test coverage files)

## Project Structure

```bash
.
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go        # Configuration management
│   ├── handler/
│   │   ├── handler.go       # HTTP handlers
│   │   └── handler_test.go  # Handler tests
│   ├── model/
│   │   └── model.go         # Data models
│   ├── service/
│   │   └── service.go       # Business logic (data loading/caching)
│   └── middleware/
│       └── middleware.go    # HTTP middleware
├── pkg/
│   └── logger/
│       └── logger.go         # Structured logging wrapper
├── data.json                 # Data file
├── go.mod                    # Go module definition
├── go.sum                    # Go module checksums
├── Makefile                  # Build and test commands
└── README.md                 # Project README
```

## Code Quality Features

- **Structured Logging**: Uses Go 1.21+ `log/slog` for structured logging with JSON output support
- **Error Handling**: Proper error wrapping with context (Go 1.13+ error wrapping)
- **Input Validation**: GUID format validation to prevent invalid requests
- **Separation of Concerns**: Clear separation between handlers, services, models, and middleware
- **Thread Safety**: In-memory cache with mutex protection for concurrent access
- **Graceful Shutdown**: Handles SIGTERM/SIGINT for clean shutdown
- **Request Tracking**: Unique request IDs for debugging and tracing
- **Configuration Management**: Environment variable support with validation
- **Data Caching**: In-memory caching eliminates disk I/O on every request
- **Middleware Stack**: Recovery, logging, CORS, and request ID middleware
- **Service Interface**: Interface-based design for easy testing and mocking

## Development

### Code Style

The project follows standard Go conventions:

- Use `gofmt` for formatting (run `make fmt`)
- Follow Go naming conventions
- Add Go doc comments for exported functions/types
- Keep functions focused and small
- Use interfaces for dependency injection

### Adding New Endpoints

1. Add handler function to `internal/handler/handler.go`
2. Add route to `cmd/server/main.go` in `setupRouter()`
3. Add tests to `internal/handler/handler_test.go`
4. Update this documentation with API details

### Running Individual Commands

```bash
# Download dependencies
make deps

# Build the application
make build

# Run the application
make run

# Format code
make fmt

# Lint code
make lint

# Run tests
make test

# Run tests with coverage
make test-coverage

# Clean build artifacts
make clean

# Show help
make help
```

## Troubleshooting

### Port Already in Use

If you get an error that the port is already in use:

1. Change the `PORT` in your environment variables
2. Or stop the process using the port:
   ```bash
   lsof -ti:3000 | xargs kill
   ```

### Data File Not Found

Ensure `data.json` exists in the project root or update `DATA_FILE_PATH` in your environment variables.

### Tests Failing

Make sure:

1. Dependencies are installed: `make deps`
2. The server can start successfully
3. `data.json` exists for integration tests
4. Go version is 1.21 or higher

### Import Errors

If you encounter import errors:

1. Ensure you're in the project root directory
2. Run `go mod tidy` to update dependencies
3. Check that the module name in `go.mod` matches your repository path (if you've copied the code)
4. Verify Go can find the modules: `go list -m all`

### Build Errors

If you encounter build errors:

1. Ensure Go 1.21+ is installed: `go version`
2. Clean and rebuild: `make clean && make build`
3. Check for syntax errors: `go build ./...`

### Module Name

When copying this code to your own repository, update the module name in `go.mod`:

1. Change `module api` to your repository path (e.g., `module github.com/username/my-repo`)
2. Run `go mod tidy` to update imports automatically

## Performance Considerations

- **Data Caching**: Data is loaded once at startup and cached in memory, eliminating disk I/O on every request
- **Thread Safety**: Mutex protection ensures safe concurrent access to cached data
- **Connection Timeouts**: Configurable read/write timeouts prevent resource exhaustion
- **Graceful Shutdown**: Allows in-flight requests to complete before shutting down

## Security Features

- **Input Validation**: GUID format validation prevents malformed requests
- **Error Messages**: Generic error messages prevent information leakage
- **CORS Support**: Configurable CORS headers for cross-origin requests
- **Request Timeouts**: Prevents resource exhaustion from slow clients
