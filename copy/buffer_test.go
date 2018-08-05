package copy

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

var sampleString = "Example string 123 456"

func CopyBuffer(r io.Reader, buf []byte) error {
	tmp := bytes.NewBuffer([]byte{})
	if _, err := io.CopyBuffer(tmp, r, buf); err != nil {
		return err
	}

	return nil
}

func BenchmarkCopyBuffer1(b *testing.B) {
	r := strings.NewReader(sampleString)
	buf := make([]byte, 1024)

	if err := CopyBuffer(r, buf); err != nil {
		b.Error(err)
	}
}
func BenchmarkCopyBuffer2(b *testing.B) {
	r := strings.NewReader(sampleString)
	buf := make([]byte, 1)

	if err := CopyBuffer(r, buf); err != nil {
		b.Error(err)
	}
}

func BenchmarkCopyBuffer3(b *testing.B) {
	r := strings.NewReader(sampleString)
	var buf []byte

	if err := CopyBuffer(r, buf); err != nil {
		b.Error(err)
	}
}
