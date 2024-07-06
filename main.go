package main

import (
	"bufio"
	"encoding/json"
	"io"
	"lilypondlsp/analysis"
	"lilypondlsp/lsp"
	"lilypondlsp/rpc"
	"log"
	"os"
)

func main() {
	logger := getLogger("/Users/antoniolesik/Code/lilypondlsp/log.txt")
	logger.Println("logger initialized")

	writer := os.Stdout

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println(err)
			continue
		}
		handleMessage(method, content, state, writer, logger)
	}
}

func handleMessage(method string, content []byte, state analysis.State, writer io.Writer, logger *log.Logger) {
	logger.Printf("received %s message", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println(err)
			return
		}
		logger.Printf("connected to %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg)
		logger.Println("initialize response sent")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocNotification
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println(err)
			return
		}
		state.OpenDocument(request.Params.TextDocItem)
		logger.Printf("opened %s:\n%s", request.Params.TextDocItem.URI, request.Params.TextDocItem.Text)
	case "textDocument/didChange":
		var request lsp.DidChangeTextDocNotification
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println(err)
			return
		}
		logger.Printf("changed %s", request.Params.TextDoc.URI)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDoc.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println(err)
			return
		}
		response := lsp.NewHoverResponse(request.ID, "masz szluga?")
		writeResponse(writer, response)
	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println(err)
			return
		}
		response := lsp.NewDefinitionResponse(request.ID,
			lsp.Location{
				URI: request.Params.TextDoc.URI,
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      0,
						Character: 0,
					},
					End: lsp.Position{
						Line:      0,
						Character: 5,
					},
				},
			},
		)
		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	response := rpc.EncodeMessage(msg)
	writer.Write([]byte(response))
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(logfile, "[lilypondlsp]", log.Ltime|log.Lshortfile)
}
