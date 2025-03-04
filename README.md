# Rebis
Rebis is a simple key-value in memory cache. Don't mix Rebis with Redis. If your usage is to just store key and value, don't need lots of features of Redis, Rebis is a good choice as it is simple, fast and lightweight.

## Usage
Start the server
```bash
go run main.go
```
Rebis accepts single connection at a time. To make a connection, use `nc` or `telnet`.
```bash
nc localhost 6666
# or
telnet localhost 6666
```
Queries-

- `SET <key> <value>` - Set a key-value pair. Key must be 16 characters or less and must not contain any spaces.
- `GET <key>` - Get the value of a key
- `DELETE <key>` - Delete a key
- `RESET` - Reset the cache

Or you can use docker to run rebis
```bash
docker run -p 6666:6666 shafinhasnat/rebis:latest
```
<!-- Rebis client for different languages - [documentation](client/README.md) -->
## Build
To build rebis, run the following commands-
```bash
export VERSION=<version>
make build
make push
```
