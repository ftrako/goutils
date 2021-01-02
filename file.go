package goutils

import (
	"bufio"
	"io"
	"os"
)

// 文件类

// 逐行读取文本文件
func ReadTextFile(file string, fn func(text string)) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// 逐行解析文本
	br := bufio.NewReader(f)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if fn != nil {
			fn(string(line))
		}
	}
	return nil
}
