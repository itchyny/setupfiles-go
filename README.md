# setupfiles-go
Create files and directories easily for tests in Go.

When we setup sample files and directories for testing, we have to create the base directory, write files and catch errors.
Doing these things is not what we have to care in testing.
Using this setupfiles package, you can pass a recipe for what the file contents and directory structure should be to the library and it creates the files and directories.

## Usage
```go
dir, err := setupfiles.CreateTemp("test-setupfiles-sample", `
bar.txt
  bar
  contents
  here

baz.txt -> foo/qux.txt

foo/qux.txt
  qux contents

dir/
`)
if err != nil {
  t.Fatal(err)
}
defer os.RemoveAll(dir)
```

## Bug Tracker
Report bug at [Issues・itchyny/setupfiles-go - GitHub](https://github.com/itchyny/setupfiles-go/issues).

## Author
itchyny (https://github.com/itchyny)

## License
This software is released under the MIT License, see LICENSE.
