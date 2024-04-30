package main

import (
	"bufio"
	"log"
	"os"

	jsonrpc2 "github.com/ntammadge/dblpls/pkg"
)

func main() {
	logger := getLogger("log.txt") // Creates the file in the vscode %localappdata% folder
	logger.Print("Started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(jsonrpc2.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(fileName string) *log.Logger {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0664)
	if err != nil {
		panic("bad log file name")
	}

	return log.New(logFile, "[dblpls]", log.Ldate|log.Ltime|log.Lshortfile)
}
