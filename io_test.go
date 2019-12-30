package readerer

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFromReaderAt(t *testing.T) {
	tests := []struct {
		offset int64
		str    string
		err    error
	}{
		{0, `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis
nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu
fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.
`,
			nil,
		},
		{335, `Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.
`,
			nil},
		{446, "", io.EOF},
	}

	f, err := os.Open("testdata/lorem_ipsum.txt")
	if err != nil {
		t.Fatal(err)
	}

	for idx, test := range tests {
		r := FromReaderAt(f, test.offset)
		b, err := ioutil.ReadAll(r)
		if err != nil {
			t.Fatalf("[%d] unexpected error", idx)
		}
		if test.str != string(b) {
			t.Errorf("[%d] expected %s, got %s", idx, test.str, string(b))
		}
	}
}
