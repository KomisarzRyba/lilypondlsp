package lsp

type TextDocItem struct {
	TextDocIdentifier
	LangID  string `json:"languageId"`
	Version int    `json:"version"`
	Text    string `json:"text"`
}

type TextDocIdentifier struct {
	URI string `json:"uri"`
}

type VersionedTextDocIdentifier struct {
	TextDocIdentifier
	Version int `json:"version"`
}

type TextDocPositionParams struct {
	TextDoc  TextDocIdentifier `json:"textDocument"`
	Position Position          `json:"position"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type Location struct {
	URI   string `json:"uri"`
	Range Range  `json:"range"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
