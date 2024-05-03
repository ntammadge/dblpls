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
		}
		// TODO: handle client error messages from our invalid replies
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

		os.Stdout.Write([]byte(jsonrpc2.SerializeMessage(lsp.NewInitializeResponse(initReq.Id))))
	case "initialized":
		logger.Println("Connected to client")
	}
}

func getLogger(fileName string) *log.Logger {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0664)
	if err != nil {
		panic("bad log file name")
	}

	return log.New(logFile, "[dblpls]", log.Ldate|log.Ltime|log.Lshortfile)
}
