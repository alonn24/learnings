package protocol

import (
	"bufio"
	"errors"
	"io"
	"log"
)

type CommandReader struct {
	reader *bufio.Reader
}

func NewCommandReader(reader io.Reader) *CommandReader {
	return &CommandReader{
		reader: bufio.NewReader(reader),
	}
}

func withoutLast(s string) string {
	return s[:len(s)-1]
}

func firstError(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *CommandReader) Read() (interface{}, error) {
	commandName, err := r.reader.ReadString(' ')
	if err != nil {
		return nil, err
	}

	switch commandName {
	case messageCommaneName:
		user, err1 := r.reader.ReadString(' ')
		message, err2 := r.reader.ReadString('\n')
		if firstError(err1, err2) != nil {
			return nil, firstError(err1, err2)
		}

		return MessageCommand{
			withoutLast(user),
			withoutLast(message),
		}, nil

	case sendCommaneName:
		message, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}
		return SendCommand{withoutLast(message)}, nil

	case nameCommaneName:
		name, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}
		return NameCommand{withoutLast(name)}, nil

	default:
		log.Printf("Unknown command: %v", commandName)
	}

	return nil, errors.New("Unknown command")
}

func (r *CommandReader) ReadAll() ([]interface{}, error) {
	commands := []interface{}{}

	for {
		command, err := r.Read()

		if command != nil {
			commands = append(commands, command)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return commands, err
		}
	}

	return commands, nil
}
