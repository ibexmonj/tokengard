# Tokengard

This application sets up a simple HTTP server that returns JSON data at the `/data` endpoint.  
It requires a valid token for authentication.

## Running the Server

```bash
go run main.go
```

This starts the server at http://localhost:8080.

## Authentication

The server expects an Authorization header with a Bearer token.
```bash
curl -X GET -H "Authorization: Bearer my-secret-token" http://localhost:8080/data
```


Without a valid token, you will receive a 401 Unauthorized response.
