package lsp

type HoverRequest struct {
	Request
	Params HoverParams `json:"params"`
}

type HoverParams struct {
	TextDocPositionParams
}

type HoverResponse struct {
	Response
	Result *HoverResult `json:"result"`
}

type HoverResult struct {
	Contents string `json:"contents"`
}

func NewHoverResponse(id int, contents string) HoverResponse {
	return HoverResponse{
		Response: Response{
			Message: Message{
				RPC: "2.0",
			},
			ID: &id,
		},
		Result: &HoverResult{
			Contents: contents,
		},
	}
}
