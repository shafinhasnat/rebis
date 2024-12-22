# Rebis
Rebis is a simple key-value in memory cache. Don't mix Rebis with Redis.

## Usage
Start the server
```
go run main.go
```
Rebis accepts single connection at a time. To make a connection, use `nc`
```
nc localhost 6666
```
Commands
```
SET <key> <value>
GET <key>
DELETE <key>
```
