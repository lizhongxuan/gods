package redis

import (
	"testing"
	"bytes"
	"fmt"
)

func TestWriterCommand(t *testing.T)  {
	var buf bytes.Buffer
	writer := NewRESPWriter(&buf)
	writer.WriteCommand("GET", "foo")
	fmt.Println(string(buf.Bytes()))
}
