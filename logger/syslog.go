//go:build linux || darwin || freebsd || openbsd || solaris
// +build linux darwin freebsd openbsd solaris

package logger

import (
	"fmt"
	"log/syslog"
	"os"

	zz "github.com/rs/zerolog/log"
)

type syslogBackend struct {
	writer [numSeverity]*syslog.Writer
	buf    [numSeverity]chan []byte
}

var SyslogPriorityMap = map[string]syslog.Priority{
	"local0": syslog.LOG_LOCAL0,
	"local1": syslog.LOG_LOCAL1,
	"local2": syslog.LOG_LOCAL2,
	"local3": syslog.LOG_LOCAL3,
	"local4": syslog.LOG_LOCAL4,
	"local5": syslog.LOG_LOCAL5,
	"local6": syslog.LOG_LOCAL6,
	"local7": syslog.LOG_LOCAL7,
}

var pmap = []syslog.Priority{syslog.LOG_EMERG, syslog.LOG_ERR, syslog.LOG_WARNING, syslog.LOG_INFO, syslog.LOG_DEBUG}
var slog *syslogBackend

func InitSyslog(prior string, tag string) {
	var err error
	if slog, err = NewSyslogBackend(prior, tag); err != nil {
		fmt.Printf(err.Error())
	}
}

func NewSyslogBackend(priorityStr string, tag string) (*syslogBackend, error) {
	priority, ok := SyslogPriorityMap[priorityStr]
	if !ok {
		return nil, fmt.Errorf("unknown syslog priority: %s", priorityStr)
	}
	var err error
	var b syslogBackend
	for i := 0; i < numSeverity; i++ {
		b.writer[i], err = syslog.New(priority|pmap[i], tag)
		if err != nil {
			return nil, err
		}
		b.buf[i] = make(chan []byte, 1<<16)
	}
	b.log()
	return &b, nil
}

func DialSyslogBackend(network, raddr string, priority syslog.Priority, tag string) (*syslogBackend, error) {
	var err error
	var b syslogBackend
	for i := 0; i < numSeverity; i++ {
		b.writer[i], err = syslog.Dial(network, raddr, priority|pmap[i], tag+severityName[i])
		if err != nil {
			return nil, err
		}
		b.buf[i] = make(chan []byte, 1<<16)
	}
	b.log()
	return &b, nil
}

func (sl *syslogBackend) Log(s Severity, msg []byte) {
	msg1 := make([]byte, len(msg))
	copy(msg1, msg)
	switch s {
	case FATAL:
		sl.tryPutInBuf(FATAL, msg1)
	case ERROR:
		sl.tryPutInBuf(ERROR, msg1)
	case WARN:
		sl.tryPutInBuf(WARN, msg1)
	case INFO:
		sl.tryPutInBuf(INFO, msg1)
	case DEBUG:
		sl.tryPutInBuf(DEBUG, msg1)
	}
}

func (sl *syslogBackend) Close() {
	for i := 0; i < numSeverity; i++ {
		sl.writer[i].Close()
	}
}

func (sl *syslogBackend) tryPutInBuf(s Severity, msg []byte) {
	select {
	case sl.buf[s] <- msg:
	default:
		os.Stderr.Write(msg)
	}
}

func (sl *syslogBackend) log() {
	for i := 0; i < numSeverity; i++ {
		go func(index int) {
			for {
				msg := <-sl.buf[index]
				slog.writer[index].Write(msg[27:])
			}
		}(i)
	}
}

func sysDebugf(s *string, v *[]interface{}) {
	slog.Log(DEBUG, []byte(fmt.Sprintf(*s, v)))
}
func sysInfof(s *string, v *[]interface{}) {
	slog.Log(INFO, []byte(fmt.Sprintf(*s, v)))
}
func sysWarnf(s *string, v *[]interface{}) {
	slog.Log(WARN, []byte(fmt.Sprintf(*s, v)))
	logEvent(zz.Warn(), s, v)
}
func sysErrorf(s *string, v *[]interface{}) {
	logEvent(zz.Error(), s, v)
}
func sysFatalf(s *string, v *[]interface{}) {
	logEvent(zz.Fatal(), s, v)
}
