package server

import (
	"log"
	"net"

	"chat-server/protocol"
)

func (s *TcpChatServer) Listen(listener net.Listener) error {
	s.listener = listener

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Print(err)
			return err
		}
		// handle connection
		log.Printf("Accepting connection from %v, total clients: %v", conn.RemoteAddr().String(), len(s.clients)+1)
		client := s.createClient(conn)
		defer s.removeClient(client)
		go s.serveClient(client)
	}
	return nil
}

func (s *TcpChatServer) Close() {
	log.Print("Closing server")
	for _, client := range s.clients {
		client.writer.Write("Closing connection")
	}
	s.listener.Close()
}

func (s *TcpChatServer) createClient(conn net.Conn) *client {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	client := &client{
		conn:   conn,
		writer: protocol.NewCommandWriter(conn),
	}
	s.clients = append(s.clients, client)
	return client
}

func (s *TcpChatServer) removeClient(client *client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// remove the connections from clients array
	for i, check := range s.clients {
		if check == client {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
		}
	}

	log.Printf("Closing connection from %v", client.conn.RemoteAddr().String())
	client.conn.Close()
}
