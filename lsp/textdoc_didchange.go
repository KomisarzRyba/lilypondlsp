package lsp

type DidChangeTextDocNotification struct {
	Notification
	Params DidChangeTextDocParams `json:"params"`
}

type DidChangeTextDocParams struct {
	TextDoc        VersionedTextDocIdentifier  `json:"textDocument"`
	ContentChanges []TextDocContentChangeEvent `json:"contentChanges"`
}

type TextDocContentChangeEvent struct {
	Text string `json:"text"`
}
