package logging

import (
	"log"
	"os"
)

type FLog struct {
	Logger *log.Logger
}

func NewFLog(name string) *FLog {
	out := &FLog{Logger: log.New(os.Stdout, "["+name+"] INFO: ", 0)}
	return out
}

func (w *FLog) Write(p []byte) (n int, err error) {
	w.Logger.Printf("%s", p)
	return len(p), nil
}
