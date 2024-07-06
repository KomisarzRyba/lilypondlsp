package lsp

type Message struct {
	RPC string `json:"jsonrpc"`
}

type Request struct {
	Message
	ID     int    `json:"id"`
	Method string `json:"method"`
}

type Response struct {
	Message
	ID *int `json:"id,omitempty"`
}

type Notification struct {
	Message
	Method string `json:"method"`
}
