// +build !windows

package fileinfo

func getTimespec(fi os.FileInfo) (t Times) {
	t.flags = 0
	t.atime = time.Unix(0, 0)
	t.mtime = time.Unix(0, 0)
	return t
}
