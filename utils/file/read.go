package file

import "os"

func Read(path string) (string, error) {
	f, err := os.Open(path)
	defer func() {
		err := f.Close()
		if err != nil {
			return
		}
	}()
	if err != nil {
		return "", err
	}
	data := make([]byte, 1024)
	_, err = f.Read(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
