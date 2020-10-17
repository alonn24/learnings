module chat-server

go 1.15


replace chat-server/server => ./server

replace chat-server/protocol => ./protocol

replace chat-server/tui => ./tui

require (
	chat-server/protocol v0.0.0-00010101000000-000000000000 // indirect
	chat-server/server v0.0.0-00010101000000-000000000000 // indirect
	chat-server/tui v0.0.0-00010101000000-000000000000 // indirect
	github.com/marcusolsson/tui-go v0.4.0 // indirect
)
