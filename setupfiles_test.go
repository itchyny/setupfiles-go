package setupfiles

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var testCases = []struct {
	title  string
	source string
	files  []file
}{
	{
		"empty",
		``,
		[]file{},
	},
	{
		"one file",
		`foo.txt`,
		[]file{{"foo.txt", ""}},
	},
	{
		"file and contents",
		`
foo.txt
  foo

  foo

bar.txt
  bar

    bar
   bar
`,
		[]file{{"foo.txt", "foo\n\nfoo\n"}, {"bar.txt", "bar\n\n  bar\n bar\n"}},
	},
	{
		"tab indent",
		`
foo.txt
	foo
		

	foo

bar.txt
		bar

			bar
	bar
`,
		[]file{{"foo.txt", "foo\n\t\n\nfoo\n"}, {"bar.txt", "\tbar\n\n\t\tbar\nbar\n"}},
	},
}

func TestCreate(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.title, func(t *testing.T) {
			dir, err := CreateTemp("setupfiles", testCase.source)
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(dir)
			for _, f := range testCase.files {
				path := filepath.Join(dir, f.path)
				cnt, err := ioutil.ReadFile(path)
				if err != nil {
					t.Fatal(err)
				}
				if f.contents != string(cnt) {
					t.Errorf("contents should be %q but got %q", f.contents, string(cnt))
				}
			}
		})
	}
}
