package utils

import (
	"bufio"
	"bytes"
	"io"
	"mime/multipart"
)

// UTF8 with BOM -> UTF8 encodingに変換するメソッド
func UTF8Encoding(file multipart.File) (*bufio.Reader, error) {
	bom := []byte{0xEF, 0xBB, 0xBF}
	reader := bufio.NewReader(file)
	peek, err := reader.Peek(3)
	if err != nil && err != io.EOF {
		return nil, err
	}
	if bytes.Equal(peek, bom) {
		reader.Discard(3)
	}
	return reader, err
}
