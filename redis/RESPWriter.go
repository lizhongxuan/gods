package redis

import (
	"bufio"
	"io"
	"strconv"
)

var (
	arrayPrefix  = []byte{'*'}
	stringPrefix = []byte{'$'}
	lineEnding   = []byte{'\r', '\n'}
)

type RESPWriter struct {
	*bufio.Writer
}

func NewRESPWriter(writer io.Writer) *RESPWriter {
	return &RESPWriter{
		Writer: bufio.NewWriter(writer),
	}
}

func (w *RESPWriter) WriteCommand(args ...string) (err error) {
	// 首先写入数组的标志和数组的数量
	w.Write(arrayPrefix)
	w.WriteString(strconv.Itoa(len(args)))
	w.Write(lineEnding)
	// 写入批量字符串
	for _, arg := range args {
		w.Write(stringPrefix)
		w.WriteString(strconv.Itoa(len(arg)))
		w.Write(lineEnding)
		w.WriteString(arg)
		w.Write(lineEnding)
	}
	return w.Flush()
}
