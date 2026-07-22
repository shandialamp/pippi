# Pippi Errors

A simple error handling package that combines standard Go errors with structured logging fields from `zap`.

## Overview

The `errors` package provides a lightweight wrapper for Go errors, allowing you to attach structured logging fields (`zap.Field`) directly to error instances. This makes it easy to capture contextual information alongside errors for structured logging.

## Features

- **Error Wrapping**: Wrap standard Go errors with additional context and fields
- **Structured Logging**: Attach `zap.Field` directly to errors for structured logging
- **JSON Marshaling**: Error messages and fields are JSON serializable

## Installation

```bash
go get github.com/shandialamp/pippi/exception
```

## Usage

### Creating Errors with Fields

```go
package main

import (
	"errors"
	"github.com/shandialamp/pippi/exception"
	"go.uber.org/zap"
)

func main() {
	// Create an error with structured fields
	err := errors.New(
		errors.New("database connection failed"),
		"failed to connect to database",
		zap.String("host", "localhost"),
		zap.Int("port", 5432),
		zap.String("database", "mydb"),
	)
	
	// Use the error
	println(err.Message)  // "failed to connect to database"
	println(err.Trace)    // "database connection failed"
	// err.Fileds contains the zap fields
}
```

## API

### `Error`

```go
type Error struct {
	Message string     // Human-readable error message
	Trace   string     // Original error trace
	Fileds  []zap.Field // Structured logging fields
}
```

### `New(err error, msg string, fields ...zap.Field) Error`

Creates a new `Error` instance from a standard Go error, a message, and optional `zap.Field` fields.

**Parameters:**
- `err`: The underlying error
- `msg`: A descriptive error message
- `fields`: Variable number of `zap.Field` for structured context

**Returns:** An `Error` struct containing the message, trace, and fields.

## License

MIT
