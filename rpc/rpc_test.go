package rpc_test

import (
	"lilypondlsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected:\n%s\n\nActual:\n%s\n", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incoming := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incoming))
	contentLen := len(content)
	if err != nil {
		t.Fatal(err)
	}
	if method != "hi" {
		t.Fatalf("Expected hi, got %s", method)
	}
	if contentLen != 15 {
		t.Fatalf("Expected 15, got %d", contentLen)
	}
}
