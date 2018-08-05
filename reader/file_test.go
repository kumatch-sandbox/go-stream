package reader

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	f, err := os.Open("./foo.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(f); err != nil {
		panic(err)
	}

	if buf.String() != "foo123" {
		t.Errorf("want foo123, got %s", buf.String())
	}

	r := bytes.NewReader(buf.Bytes())
	buf2 := make([]byte, 6)
	r.Read(buf2)

	if string(buf2) != "foo123" {
		t.Errorf("want foo123, got %s", buf2)
	}
}

func foo(r io.Reader) {}
