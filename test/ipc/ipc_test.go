package ipc

import "testing"

type EchoServer struct {
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}
func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"200", method + params}
}
func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	response1, err1 := client1.Call("GET ", "From Client1")
	response2, err2 := client2.Call("GET ", "From Client2")

	if err1 != nil || err2 != nil {
		t.Error("IpcClient.Call failed. response1:", response1, "response2:", response2)
	}
	t.Log(response1)
	t.Log(response2)
	client1.Close()
	client2.Close()
}
