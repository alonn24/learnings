package server

import (
	"io"
	"log"

	"chat-server/protocol"
)

func broadcastClients(from *client, clients []*client, message string) error {
	messageCommand := protocol.MessageCommand{
		Message: message,
		Name:    from.name,
	}
	for _, client := range clients {
		client.writer.Write(messageCommand)
	}

	return nil
}

func (s *TcpChatServer) serveClient(client *client) {
	cmdReader := protocol.NewCommandReader(client.conn)

	for {
		cmd, err := cmdReader.Read()
		if err != nil {
			log.Printf("Read error: %v", err)
		}

		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.SendCommand:
				go broadcastClients(client, s.clients, v.Message)

			case protocol.NameCommand:
				client.name = v.Name
				go broadcastClients(client, s.clients, "connected")
			}
		}
		if err == io.EOF {
			break
		}
	}
}
