package setupfiles

type file struct {
	path     string
	contents string
	symlink  string
	isDir    bool
}
