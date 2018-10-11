package timet

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Log(time.Now().Unix())
	t.Log(time.Now().Date())
	t.Log(time.Now().Format("2006-01-02"))
	tt,_:=time.Parse("2006-01-02 15:04:05","2013-08-11 11:18:46")
	t.Log(tt,tt.Format("2006-01-02 15:04:05"))
	t.Log(time.Unix(1538038032,0).Format("2006-01-02 15:04:05"))
}
