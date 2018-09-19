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
		[]file{{path: "foo.txt", contents: ""}},
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
		[]file{
			{path: "foo.txt", contents: "foo\n\nfoo\n"},
			{path: "bar.txt", contents: "bar\n\n  bar\n bar\n"},
		},
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
		[]file{
			{path: "foo.txt", contents: "foo\n\t\n\nfoo\n"},
			{path: "bar.txt", contents: "\tbar\n\n\t\tbar\nbar\n"},
		},
	},
	{
		"symlink",
		`
foo.txt -> bar.txt
bar.txt
	foobar`,
		[]file{
			{path: "foo.txt", contents: "foobar", symlink: "bar.txt"},
			{path: "bar.txt", contents: "foobar"},
		},
	},
	{
		"directories",
		`
foo/bar/foo.txt -> ../bar.txt

foo/bar.txt
	foobar

foo/baz/

foo/qux -> baz
`,
		[]file{
			{path: "foo/bar/foo.txt", contents: "foobar\n", symlink: "foo/bar.txt"},
			{path: "foo/bar.txt", contents: "foobar\n"},
			{path: "foo/baz/", isDir: true},
			{path: "foo/qux/", isDir: true},
		},
	},
}

func TestCreate(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.title, func(t *testing.T) {
			dir, err := CreateTemp("setupfiles", testCase.source)
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(dir)
			for _, f := range testCase.files {
				path := filepath.Join(dir, f.path)
				if f.isDir {
					fi, err := os.Stat(path)
					if err != nil {
						t.Fatal(err)
					}
					if !fi.IsDir() {
						t.Errorf("%s should be a directory but got %s", path, fi.Mode())
					}
					continue
				}
				cnt, err := ioutil.ReadFile(path)
				if err != nil {
					t.Fatal(err)
				}
				if f.contents != string(cnt) {
					t.Errorf("contents of %s should be %q but got %q", path, f.contents, string(cnt))
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
					expected := filepath.Join(dir, f.symlink)
					if got != expected {
						t.Errorf("%s should be a symlink to %s but %s", f.path, expected, got)
					}
				}
			}
		})
	}
}
