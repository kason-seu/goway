package log

import (
	stdlog "log"
	"os"
)

var log *stdlog.Logger

type fileLog string

func (Fl fileLog) Write(data []byte) (n int, err error) {

	os.OpenFile(string(Fl), os.O_CREATE | os.O_WRONLY | os.O_APPEND,  0666)
	return 0, nil
}
