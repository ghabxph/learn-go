package util

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Path(path string) string {
	p, e := filepath.Abs(path)
	catch(e)
	return p
}

func Workdir() string {
	p, e := os.Getwd()
	catch(e)
	return p
}

func ReadFileB(path string) []byte {
	b, e := ioutil.ReadFile(path)
	catch(e)
	return b
}

func ReadFileS(path string) string {
	return string(ReadFileB(path))
}

func ReadAllB(r io.Reader) []byte {
	b, e := ioutil.ReadAll(r)
	catch(e)
	return b
}

func ReadAllS(r io.Reader) string {
	return string(ReadAllB(r))
}
