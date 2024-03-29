package protocol

var (
	messageCommaneName = "MESSAGE "
	sendCommaneName    = "SEND "
	nameCommaneName    = "NAME "
)

// SendCommand is used for sending new message from client
type SendCommand struct {
	Message string
}

// NameCommand is used for setting client display name
type NameCommand struct {
	Name string
}

// MessageCommand is used for notifying new messages
type MessageCommand struct {
	Name    string
	Message string
}
