package utils

import (
	"bytes"
	"os"
)

func LoadFile(fp string) (string, error) {
	f, err := os.Open(fp)
	if err != nil {
		return "", err
	}

	st, _ := f.Stat()
	sz := st.Size()
	data := make([]byte, sz)
	_, err = f.Read(data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func FileExists(bpath string) bool {
	_, err := os.Lstat(bpath)
	return !os.IsNotExist(err)
}

func MkDir(bpath string) error {
	err := os.Mkdir(bpath, os.ModePerm)
	return err
}

func MkDirP(bpath string) error {
	return os.MkdirAll(bpath, os.ModePerm)
}

func WriteToFile(fpath, data string) error {

	// open file, setup defer
	f, err := os.Create(fpath)
	if err != nil {
		return err
	}

	defer f.Close()

	// load data into a buffer and save it
	buf := bytes.NewBufferString(data)
	f.Write(buf.Bytes())

	return nil
}
