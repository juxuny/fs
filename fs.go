package fs

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
	"syscall"
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

func (t *fileCleaner) Execute(keepNumberOfFile int, removeFilter FileFilter) (err error) {
	list, err := ioutil.ReadDir(t.Directory)
	if err != nil {
		return errors.Wrapf(err, "read directory failed")
	}
	validFileList := make([]os.FileInfo, 0)
	for _, item := range list {
		if item.IsDir() {
			continue
		}
		stat := item.Sys().(*syscall.Stat_t)
		createTime := time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec)
		modifiedTime := item.ModTime()
		if t.Filter(item.Name(), createTime, modifiedTime) {
			validFileList = append(validFileList, item)
		}
	}
	remainCount := len(validFileList)
	if remainCount <= keepNumberOfFile {
		return nil
	}
	for _, item := range validFileList {
		if item.IsDir() {
			continue
		}
		stat := item.Sys().(*syscall.Stat_t)
		createTime := time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec)
		modifiedTime := item.ModTime()
		if removeFilter(item.Name(), createTime, modifiedTime) {
			err = os.Remove(path.Join(t.Directory, item.Name()))
			if err != nil {
				return err
			}
			remainCount -= 1
		}
		if remainCount <= keepNumberOfFile {
			return nil
		}
	}
	return nil
}
