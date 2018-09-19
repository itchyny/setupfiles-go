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
		[]file{{"foo.txt", "", ""}},
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
		[]file{{"foo.txt", "foo\n\nfoo\n", ""}, {"bar.txt", "bar\n\n  bar\n bar\n", ""}},
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
		[]file{{"foo.txt", "foo\n\t\n\nfoo\n", ""}, {"bar.txt", "\tbar\n\n\t\tbar\nbar\n", ""}},
	},
	{
		"symlink",
		`
foo.txt -> bar.txt
bar.txt
	foobar`,
		[]file{{"foo.txt", "foobar", "bar.txt"}, {"bar.txt", "foobar", ""}},
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
				fi, err := os.Lstat(path)
				if err != nil {
					t.Fatal(err)
				}
				if fi.Mode()&os.ModeSymlink == 0 {
					if f.symlink != "" {
						t.Errorf("%s should be a symlink but %s", f.path, fi.Mode())
					}
				} else {
					got, err := os.Readlink(path)
					if err != nil {
						t.Fatal(err)
					}
					expected := filepath.Join(filepath.Dir(path), f.symlink)
					if got != expected {
						t.Errorf("%s should be a symlink to %s but %s", f.path, expected, got)
					}
				}
			}
		})
	}
}
