package sockettest

import (
	"testing"
	"net"
)

func TestSocket(t *testing.T) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "10.64.49.20:8003")
	t.Log(tcpAddr.Network())
	t.Log(tcpAddr.String())
	t.Log(tcpAddr.IP, tcpAddr.Port, tcpAddr.Zone)
}
