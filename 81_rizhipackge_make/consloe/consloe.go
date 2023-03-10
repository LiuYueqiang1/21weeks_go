package consloe

import "fmt"

type Logger struct {
}

func NewLog() Logger {
	return Logger{}
}
func (l Logger) Debug(msg string) {
	fmt.Println(msg)
}
