package rwfs

import (
	"io"
	"io/fs"
)

type FS interface {
	Open(name string) (fs.File, error)
	OpenRW(name string) (File, error)
	Truncate(name string, size int64) error
}

type File interface {
	fs.File
	io.Writer
}
