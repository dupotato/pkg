package logger

import "testing"

func Test_zerolog(t *testing.T) {
	var lc LConfig
	lc.Type = ZEROLOG

	NewLogger(lc)
	Debugf("111%s", "ok")
}

func Test_syslog(t *testing.T) {
	var lc LConfig
	lc.Type = SYSLOG
	NewLogger(lc)
	Debugf("%s", "test syslog")
}
