package file

import (
	"bufio"
	"io"
	"os"
)

func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			return
		}
	}()
	reader := bufio.NewReader(f)
	var data []string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		data = append(data, string(line))
	}
	return data, nil
}
