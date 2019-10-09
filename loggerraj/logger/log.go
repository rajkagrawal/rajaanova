package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log/syslog"
	"os"
	"sync"
	"time"
)

type entry struct {
	Time time.Time
	Msg  string
}

func init() {
	DefaultLogger()
}

type logger struct {
	writ io.Writer
}

var log *logger

func (a *logger) Write(p []byte) (int, error) {
	return a.writ.Write(p)
	//fmt.Println("this is logger", string(p))
	//return 0, nil
}

func DefaultLogger() *logger {
	var writer io.Writer
	writer, err := syslog.New(syslog.LOG_NOTICE, "rajlogger")
	if err != nil {
		writer = os.Stdout
	}
	var f = func() {
		log = &logger{writ: writer}
	}
	var once sync.Once

	once.Do(f)
	return log
}

func Debug(a ...interface{}) {
	inte := make([]interface{}, 0, 1)
	inte = append(inte, time.Now())
	a = append(inte, a...)
	s := fmt.Sprint(a)
	ent := entry{Time: time.Now(), Msg: s}
	b, _ := json.Marshal(ent)
	log.Write(b)
}
