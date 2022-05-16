package utils

import (
	"bufio"
	"io"
	"os"
)

func ReadFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	strs := []string{}
	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		str := string(line)
		strs = append(strs, str)
		if err != nil {
			panic(err)
		}
	}
	f.Close()
	return strs
}
