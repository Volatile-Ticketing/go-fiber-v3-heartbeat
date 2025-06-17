# Go Fiber v3 Heartbeat

A robust health check endpoint implementation using Go Fiber v3 framework.

## Features

- Detailed server health status with code and message
- Comprehensive uptime information (days, hours, minutes, seconds)
- Human-readable memory statistics (B, KB, MB, GB)
- Detailed system information
- Runtime metrics

## Installation

```bash
go get github.com/mpro4/go-fiber-v3-heartbeat
```

## Usage

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v3"
    "github.com/mpro4/go-fiber-v3-heartbeat/routes"
)

func main() {
    app := fiber.New()
    app.Get("/heartbeat", routes.HeartbeatHandler)
    log.Fatal(app.Listen(":3000"))
}
```

## API Response

The `/heartbeat` endpoint returns a JSON response with the following structure:

```json
{
    "status": {
        "code": "healthy",
        "message": "Server is running normally"
    },
    "timestamp": "2024-03-14T12:00:00Z",
    "uptime": {
        "raw": "1h2m3s",
        "days": 0,
        "hours": 1,
        "minutes": 2,
        "seconds": 3
    },
    "memory": {
        "alloc": "1.23 MB",
        "totalAlloc": "7.65 MB",
        "sys": "9.87 MB",
        "numGC": 42,
        "heapAlloc": "1.23 MB",
        "heapSys": "2.34 MB"
    },
    "system": {
        "goroutines": 10,
        "numCPU": 8,
        "goVersion": "go1.21.0",
        "os": "darwin",
        "arch": "amd64"
    }
}
```

### Response Fields

#### Status
- `code`: Current health status code
- `message`: Human-readable status message

#### Uptime
- `raw`: Raw uptime string
- `days`: Number of days
- `hours`: Number of hours
- `minutes`: Number of minutes
- `seconds`: Number of seconds

#### Memory
- `alloc`: Currently allocated memory
- `totalAlloc`: Total allocated memory
- `sys`: System memory
- `numGC`: Number of garbage collections
- `heapAlloc`: Heap allocation
- `heapSys`: Heap system memory

#### System
- `goroutines`: Number of active goroutines
- `numCPU`: Number of CPU cores
- `goVersion`: Go runtime version
- `os`: Operating system
- `arch`: System architecture

## License

MIT 