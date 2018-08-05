package pipe

import (
	"bytes"
	"io"
	"testing"
)

func TestPipe1(t *testing.T) {
	s := "foo123"
	r, w := io.Pipe()

	go func() {
		var b bytes.Buffer
		b.WriteString(s)
		b.WriteTo(w)
		w.Close()
	}()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(r); err != nil {
		panic(err)
	}

	if s != buf.String() {
		t.Errorf("Want %s, got %s", s, buf.String())
	}
}
