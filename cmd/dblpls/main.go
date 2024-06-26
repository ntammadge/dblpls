package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	jsonrpc2 "github.com/ntammadge/dblpls/pkg"
	"github.com/ntammadge/dblpls/pkg/lsp"
)

func main() {
	logger := getLogger("log.txt") // Creates the file in the vscode %localappdata% folder
	logger.Print("Started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(jsonrpc2.Split)

	for scanner.Scan() {
		data := scanner.Bytes()

		method, content, err := jsonrpc2.DeserializeMessage(data)
		if err != nil {
			logger.Printf("Error deserializing message: %s", err.Error())
			logger.Printf("Raw message: %s", data)
			continue
		}

		if method == "" {
			// Messages without a method are expected to be error messages, but that may not be every message.
			// Adding basic logging to get information for better understanding
			logger.Printf("Received message with no method: %s", content)
			continue
		}

		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("Received message with method %s", method)

	switch method {
	case "initialize":
		var initReq lsp.InitializeRequest
		if err := json.Unmarshal(content, &initReq); err != nil {
			// TODO: unable to initialize
			break
		}
		logger.Printf("Connecting to client %s %s",
			initReq.Params.ClientInfo.Name,
			initReq.Params.ClientInfo.Version)
		os.Stdout.Write([]byte(jsonrpc2.SerializeMessage(lsp.NewInitializeResponse(initReq.Id)))) // TODO: come up with something more elegant
	case "initialized":
		logger.Println("Connected to client")
	case "textDocument/didOpen":
		var req lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("Couldn't parse open notification: %s", err)
			break
		}
		logger.Printf("Opened: %s", req.Params.TextDocument.Uri)
	}
}

func getLogger(fileName string) *log.Logger {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0664)
	if err != nil {
		panic("bad log file name")
	}

	return log.New(logFile, "[dblpls]", log.Ldate|log.Ltime|log.Lshortfile)
}
