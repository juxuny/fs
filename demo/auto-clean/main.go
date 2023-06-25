package main

import (
	"fmt"
	"github.com/juxuny/fs"
	"strings"
	"time"
)

func main() {
	helper := fs.CreateFileCleaner("./data", func(fileName string, createTime time.Time, modifiedTime time.Time) bool {
		fmt.Println(fileName)
		if strings.Index(fileName, "yaml") > 0 {
			return true
		}
		return false
	})
	err := helper.Execute(1, func(fileName string, createTime time.Time, modifiedTime time.Time) bool {
		return strings.Index(fileName, "001") > 0
	})
	if err != nil {
		panic(err)
	}
}
