package utils

import (
	"io/ioutil"
	"os"
)

type fn func()

func GetConsoleOutput(f fn) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out);
}