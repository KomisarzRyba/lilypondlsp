package lsp

type DidOpenTextDocNotification struct {
	Notification
	Params DidOpenTextDocParams `json:"params"`
}

type DidOpenTextDocParams struct {
	TextDocItem TextDocItem `json:"textDocument"`
}
