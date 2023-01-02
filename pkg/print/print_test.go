package print

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestBody(t *testing.T) {
	tmpdir := t.TempDir()
	filename := path.Join(tmpdir, "a.txt")

	want := "hello"

	Body(filename, want)

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		t.Error(err)
	}

	got := string(body)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
