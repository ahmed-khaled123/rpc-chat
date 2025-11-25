# RPC Chat Server (Dockerized)

This project is a **Go RPC Chat Server** with Docker, allowing you to run the server inside a container and connect clients for messaging.

---

## 📁 File Structure

rpc-chat/
 ├── server.go
 ├── client.go
 ├── Dockerfile
 └── README.md

> **Note:** Dockerfile must be in the same folder as `server.go`.

---

## 1️⃣ Build Docker Image

Open a terminal inside the project folder and run:

```bash
docker build -t ahmedelbanna123/rpc-chat-server:new .
```

This will create a Docker image ready to run the server.

---

## 2️⃣ Run Server Inside Docker

```bash
docker run --rm -p 1234:1234 --name rpc-chat-server ahmedelbanna123/rpc-chat-server:new
```

You should see:

Chat server is currently running on port 1234...

> **Note:** The server inside Docker does not accept direct terminal input; messages are exchanged via the client.

---

## 3️⃣ Run the Client

Open a new terminal and run:

```bash
go run client.go


- You will be prompted to enter your name:
Enter your name please:

- After entering your name, you can send and receive messages through the server.

> You can open multiple terminals to run multiple clients for testing multi-user chat.

---

## 4️⃣ Docker Hub Login

If you want to push the image to Docker Hub, log in first:

```bash
docker login -u ahmedelbanna123
```

- Use your **Access Token** as the password.

---

## 5️⃣ Push Image to Docker Hub

```bash
docker push ahmedelbanna123/rpc-chat-server:new
```

After pushing, anyone can pull and run the server directly.

---

## 6️⃣ Pull Image from Docker Hub

```bash
docker pull ahmedelbanna123/rpc-chat-server:new
```

Useful for running the server on another machine without rebuilding.

---

## 7️⃣ Run Server from Docker Hub

```bash
docker run --rm -p 1234:1234 ahmedelbanna123/rpc-chat-server:new
```

The server will run directly on port 1234.

---

## 8️⃣ Multi-Client Example

1. Open a terminal for each client.
2. Run each client:

```bash
go run client.go

3. Each client will send and receive messages through the server running in Docker.


## 🔗 Links

- **Docker Hub Repository:** [https://hub.docker.com/repository/docker/ahmedelbanna123/rpc-chat-server/general]

---

## ✅ Notes

- The Docker image contains only the server (`server.go`).
- The client (`client.go`) runs from your machine or any terminal connected to the same network/port.
- Make sure port 1234 is open for client-server communication.

