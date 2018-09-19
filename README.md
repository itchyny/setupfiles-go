# setupfiles-go
Create files and directories easily for tests in Go.

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
defer os.Remove(dir)
```

## Bug Tracker
Report bug at [Issues・itchyny/setupfiles-go - GitHub](https://github.com/itchyny/setupfiles-go/issues).

## Author
itchyny (https://github.com/itchyny)

## License
This software is released under the MIT License, see LICENSE.
