package utils

import (
	"strings"
	"time"
)

const (
	TS_SEC_FORMAT       = "2006-01-02 15:04:05"     //.999
	TS_MS_FORMAT        = "2006-01-02 15:04:05.999" //
	TS_MS_FORMAT1       = "2006-01-02 15:04:05 999"
	TS_SEC_FORMAT_SHORT = "20060102150405"
	TS_MS_FORMAT2       = "20060102150405.999"
)

//Time function

//CurrentTimeMs return MillSecond now time
func CurrentTimeMs() int64 {
	return time.Now().UnixNano() / 1e6
}

//MsTs2TimeFmt change timestamp ts to time'string
//return string time by FORMAT
func MsTs2TimeFmt(ts int64) string {
	return time.Unix(ts/1000, ts%1000*1000000).Format(TS_MS_FORMAT)
}

func MsTs2ShortTime(ts int64) string {
	return time.Unix(ts/1000, ts%1000*1000000).Format(TS_SEC_FORMAT_SHORT)
}

// TS_MS-FORMAT1
func MsTs2Formart1(ts int64) string {
	s := MsTs2TimeFmt(ts)
	if strings.Contains(s, ".") == true {
		return strings.Replace(s, ".", " ", 1)
	} else {
		return s + " 000"
	}
}

func CurrentTimeMsString() string {
	return time.Now().Format(TS_MS_FORMAT)
}

func CurrentTimeMsString1() string {
	s := CurrentTimeMsString()
	if strings.Contains(s, ".") == true {
		return strings.Replace(s, ".", " ", 1)
	} else {
		return s + " 000"
	}
	//return time.Now().Format(TS_MS_FORMAT1)
}

func CurrentTimeMsString2() string {

	s := time.Now().Format(TS_MS_FORMAT2)

	if strings.Contains(s, ".") == true {
		return strings.Replace(s, ".", " ", 1)
	} else {
		return s + " 000"
	}
}
