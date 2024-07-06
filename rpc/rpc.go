package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("did not find separator")
	}

	contentLenBytes := header[len("Content-Length: "):]
	contentLen, err := strconv.Atoi(string(contentLenBytes))
	if err != nil {
		return "", nil, err
	}

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLen], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLen], nil
}

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	contentLenBytes := header[len("Content-Length: "):]
	contentLen, err := strconv.Atoi(string(contentLenBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLen {
		return 0, nil, nil
	}

	totalLen := len(header) + 4 + contentLen
	return totalLen, data[:totalLen], nil
}
