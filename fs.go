package fs

import (
	"time"
)

type FileSystemHelperInterface interface {
	Execute(keepNumberOfFile int, removeFilter FileFilter) (err error)
}

type FileFilter func(fileName string, createTime time.Time, modifiedTime time.Time) bool

type fileCleaner struct {
	Directory string
	Filter    FileFilter
}

func CreateFileCleaner(dir string, matchFileFilter FileFilter) FileSystemHelperInterface {
	return &fileCleaner{
		Directory: dir,
		Filter:    matchFileFilter,
	}
}
