package readerer

import (
	"bufio"
	"os"
	"testing"
)

func TestFromReaderAt(t *testing.T) {
	tests := []struct {
		offset int64
		lines  []string
		err    error
	}{
		{
			0,
			[]string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
				"incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis",
				"nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
				"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu",
				"fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in",
				"culpa qui officia deserunt mollit anim id est laborum."},
			nil,
		},
		{
			335,
			[]string{"Excepteur sint occaecat cupidatat non proident, sunt in",
				"culpa qui officia deserunt mollit anim id est laborum."},
			nil,
		},
	}

	f, err := os.Open("testdata/lorem_ipsum.txt")
	if err != nil {
		t.Fatal(err)
	}

	for idx, test := range tests {
		r := FromReaderAt(f, test.offset)
		sc := bufio.NewScanner(r)
		for i := 0; sc.Scan(); i++ {
			if test.lines[i] != sc.Text() {
				t.Errorf("[%d] expected %s, but got %s", idx, test.lines[i], sc.Text())
			}
		}
		if test.err != sc.Err() {
			t.Errorf("[%d] expected %v, but got %v", idx, test.err, sc.Err())
		}
	}
}
