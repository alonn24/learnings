package protocol

import (
	"errors"
	"fmt"
	"io"
)

type CommandWriter struct {
	writer io.Writer
}

func NewCommandWriter(writer io.Writer) *CommandWriter {
	return &CommandWriter{
		writer: writer,
	}
}

func (w *CommandWriter) writeString(msg string) error {
	_, err := w.writer.Write([]byte(msg))
	return err
}

func getMessageString(commandName string, message string) string {
	return fmt.Sprintf("%v%v\n", commandName, message)
}

func (w *CommandWriter) Write(command interface{}) (err error) {
	message := ""
	switch v := command.(type) {
	case SendCommand:
		message = getMessageString(sendCommaneName, v.Message)
	case MessageCommand:
		str := fmt.Sprintf("%v %v", v.Name, v.Message)
		message = getMessageString(messageCommaneName, str)
	case NameCommand:
		message = getMessageString(nameCommaneName, v.Name)
	default:
		err = errors.New("Unknown command")
	}
	if message != "" {
		err = w.writeString(message)
	}
	return err
}
