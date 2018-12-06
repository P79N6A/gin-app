package sockettest

import (
	"math"
	"net"
	"testing"
)

func TestSocket(t *testing.T) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "10.64.49.20:8003")
	t.Log(tcpAddr.Network())
	t.Log(tcpAddr.String())
	t.Log(tcpAddr.IP, tcpAddr.Port, tcpAddr.Zone)
}
func TestSome(t *testing.T) {
	t.Log(int(math.Ceil(8 / float64(10))))
}
