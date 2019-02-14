package timet

import (
	"math"
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

func TestSome(t *testing.T) {
	var arr []int
	for i:=0;i<1510;i++{
		arr=append(arr,i)
	}
	// int(math.Ceil(total / float64(request.Num))
	count:=int(math.Ceil(float64(len(arr))/100))
	t.Log(count)
	for j := 0; j < count; j++ {
		start:=j*100
		end:=(j+1)*100
		if end>len(arr) {
			end=len(arr)
		}
		arr1:=arr[start:end]
		t.Log(len(arr1),arr1)
	}
}
