# Websocket chatroom api 



## How to use


### 1. Install dependencies

```bash
go mod tidy
```

### 2. Run server

```bash
go run main.go
```

### 3. Create a room

```bash
curl -X GET http://localhost:8080/'
```

```json
{
 "created_room": "9ae71801-abbb-4759-92fb-88714900811b"
}
```

### 3. Open client

Use a websocket client to connect to `ws://localhost:8080/ws/{room_id}` and send a message to the server. The server will broadcast the message to all connected clients.

## Database

Create mongo database

```bash
    docker run -d --name mongodb -p 27017:27017 mongo
```
