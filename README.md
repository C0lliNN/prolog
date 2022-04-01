# Prolog
A Logging server.

## How to test manually
1. Make init
2. Start the server running `gencert`
2. Start the server running `start_everything`
3. Append a new log record using `go run cmd/client/main.go append 'Hello World!'`
4. Read the inserted record using `go run cmd/client/main.go read 1`

OBS: Make sure you have created the directory `/var/prolog` with read/write before execute `make start`