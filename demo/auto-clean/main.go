package main

import (
	"fmt"
	"github.com/juxuny/fs"
	"strings"
	"time"
)

func main() {
	helper := fs.CreateFileCleaner("./data", func(fileName string, createTime time.Time, modifiedTime time.Time) bool {
		fmt.Println(fileName, createTime.Format("2006-01-02 15:04:05"))
		if strings.Index(fileName, "001") > 0 {
			return true
		}
		return false
	})
	err := helper.Execute(1)
	if err != nil {
		panic(err)
	}
}
