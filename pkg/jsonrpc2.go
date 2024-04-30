package jsonrpc2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func SerializeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		// TODO: don't panic
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

var contentLengthIndicator = "Content-Length: "

type Message struct {
	Method string `json:"method"`
}

func DeserializeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return "", nil, errors.New("did not find expected content separator")
	}

	contentLengthBytes := header[len(contentLengthIndicator):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	message := Message{}
	err = json.Unmarshal(content[:contentLength], &message)
	if err != nil {
		return "", nil, err
	}

	return message.Method, content[:contentLength], nil
}

// Defines the input stream split operation for the main program loop.
// Takes input data, looks for the break between the "header" and content,
// and advances the data read position after receiving the specified number of
// bytes of data from the client.
func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	// TODO: can part of this be abstracted to unify operations between input reading and deserialization?
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return 0, nil, nil
	}

	contentLengthBytes := header[len("Content-Length: "):] // TODO: len(header) or something, probably
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}
