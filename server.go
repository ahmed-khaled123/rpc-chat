package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"sync"
	"time"
)

type ChatMessage struct {
	Author    string
	Text      string
	Timestamp time.Time
}

type ChatServer struct {
	mu      sync.Mutex
	history []ChatMessage
}

func (s *ChatServer) SendMessage(msg ChatMessage, reply *[]ChatMessage) error {
	if msg.Author == "" {
		return errors.New("author is required")
	}
	if msg.Text == "" {
		return errors.New("text is required")
	}
	if msg.Timestamp.IsZero() {
		msg.Timestamp = time.Now()
	}

	s.mu.Lock()
	s.history = append(s.history, msg)
	h := make([]ChatMessage, len(s.history))
	copy(h, s.history)
	s.mu.Unlock()

	*reply = h
	return nil
}

func (s *ChatServer) GetHistory(_ struct{}, reply *[]ChatMessage) error {
	s.mu.Lock()
	h := make([]ChatMessage, len(s.history))
	copy(h, s.history)
	s.mu.Unlock()

	*reply = h
	return nil
}

func main() {
	port := os.Getenv("CHAT_PORT")
	if port == "" {
		port = "1234"
	}

	chat := new(ChatServer)
	if err := rpc.Register(chat); err != nil {
		log.Fatalf("register error: %v", err)
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	fmt.Printf("Chat server is currently running on port %s...\n", port)

	go func() {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			if sc.Text() == "exit" {
				fmt.Println("Shutting down server...")
				l.Close()
				return
			}
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			log.Printf("accept error: %v", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
