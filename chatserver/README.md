# Chat Server Example

This is a minimal customer service chat server inspired by the `go-fly` project. It provides simple HTTP endpoints to send and fetch messages.

## Building and Running

```
# from the chatserver directory
go run main.go
```

The server listens on port `8080` and exposes the following endpoints:

- `POST /api/message` - Send a message with JSON body `{ "sender": "name", "content": "text" }`.
- `GET /api/messages` - Fetch all messages as JSON.

This implementation keeps messages in memory and does not support authentication or persistence.
