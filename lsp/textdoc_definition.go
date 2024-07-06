package lsp

type DefinitionRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

type DefinitionParams struct {
	TextDocPositionParams
}

type DefinitionResponse struct {
	Response
	Result Location `json:"result"`
}

func NewDefinitionResponse(id int, location Location) DefinitionResponse {
	return DefinitionResponse{
		Response: Response{
			Message: Message{
				RPC: "2.0",
			},
			ID: &id,
		},
		Result: location,
	}
}
