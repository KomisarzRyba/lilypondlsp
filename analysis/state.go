package analysis

import "lilypondlsp/lsp"

type State struct {
	Documents map[string]lsp.TextDocItem
}

func NewState() State {
	return State{
		Documents: make(map[string]lsp.TextDocItem),
	}
}

func (s *State) OpenDocument(doc lsp.TextDocItem) {
	s.Documents[doc.URI] = doc
}

func (s *State) UpdateDocument(uri, content string) {
	if doc, ok := s.Documents[uri]; ok {
		doc.Text = content
		s.Documents[uri] = doc
	}
}
