# Bahrain CPR Validator

This project validates Bahraini Civil Personal Registration (CPR) numbers using a checksum-based validation routine.

## What it does

The service exposes a small HTTP API that accepts a CPR string and returns whether it is valid.

## Project structure

- `main.go` - HTTP server and request handlers
- `internal/validator/cpr.go` - CPR validation logic
- `internal/validator/cpr_test.go` - validation tests

## Run the application

```bash
go run .
```

The server listens on port `8080`.

## Endpoints

### Health check

```bash
curl http://localhost:8080/health
```

Example response:

```json
{"status":"ok"}
```

### Validate CPR

```bash
curl -X POST http://localhost:8080/validate \
  -H "Content-Type: application/json" \
  -d '{"cpr":"901208302"}'
```

Example response for a valid CPR:

```json
{"valid":true}
```

Example response for an invalid CPR:

```json
{"valid":false,"error":"CPR is invalid"}
```

## Validation behavior

- rejects empty input
- rejects non-numeric CPR values
- rejects invalid dates or malformed CPR strings
- validates the checksum before accepting the number

## Run tests

```bash
go test ./...
```
