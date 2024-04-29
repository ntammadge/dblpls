package jsonrpc2_test

import (
	"fmt"
	"testing"

	jsonrpc2 "github.com/ntammadge/dblpls/pkg"
)

type SerializeTest struct {
	Testing bool
}

func TestSerializeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := jsonrpc2.SerializeMessage(SerializeTest{true})

	if actual != expected {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

type DeserializeTest struct {
	Method string
}

func TestDeserializeMessage(t *testing.T) {
	incomingContent := "{\"Method\":\"testing\"}"
	expectedLength := len(incomingContent)
	expectedMethod := "testing"
	incomingMessage := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", expectedLength, incomingContent)
	method, content, err := jsonrpc2.DeserializeMessage([]byte(incomingMessage))
	if err != nil {
		t.Fatal(err)
	}

	if len(content) != expectedLength {
		t.Fatalf("Expected: %d, Actual: %d", expectedLength, len(content))
	}

	if method != expectedMethod {
		t.Fatalf("Expected: %s, Actual: %s", expectedMethod, method)
	}
}
