package file

import "os"

func Write(path, c string) error {
	f, err := os.Create(path)
	defer func() {
		err = f.Close()
		if err != nil {
			return
		}
	}()
	if err != nil {
		return err
	}
	data := []byte(c)
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
