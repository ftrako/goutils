package main

import (
	"github.com/ftrako/logger"
	"goutils/fileutil"
)

func main()  {
	fileutil.ReadTextFile("a.txt", func(text string) {
		logger.Debug(text)
	})
}
