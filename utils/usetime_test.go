package utils

import (
	"fmt"
	"testing"
)

func Test_CurrentTimeMsString1(t *testing.T) {
	s := CurrentTimeMsString()
	t.Log(s)
	s1 := CurrentTimeMsString1()
	t.Log(s1)
	t1 := CurrentTimeMs()
	t1 = 1628647638000
	fmt.Println(t1)
	s2 := MsTs2Formart1(t1)
	t.Log(s2)

	t.Log(CurrentTimeMsString2())

}
