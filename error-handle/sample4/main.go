package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "os.Open(path:%s) failed", path)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	return config, errors.Wrapf(err, "svc.ReadFile(path:%s) failed", filepath.Join(home, ".settings.xml"))
}

func main() {
	_, err := ReadConfig()
	if err != nil {
		fmt.Println(errors.Cause(err))
		fmt.Println(err.Error())
		//fmt.Printf("%+v\n", err)
	}
}
