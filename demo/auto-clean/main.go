package main

import (
	"github.com/juxuny/fs"
	"strings"
	"time"
)

func main() {
	helper := fs.CreateFileCleaner("./data", func(fileName string, createTime time.Time, modifiedTime time.Time) bool {
		if strings.Index(fileName, "001") > 0 {
			return true
		}
		return false
	})
	err := helper.Execute(1, func(fileName string, createTime time.Time, modifiedTime time.Time) bool {
		return true
	})
	if err != nil {
		panic(err)
	}
}
