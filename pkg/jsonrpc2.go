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
