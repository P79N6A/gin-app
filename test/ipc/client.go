package ipc

import "encoding/json"

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (response *Response, err error) {
	request := &Request{method, params}
	var b []byte
	b, err = json.Marshal(request)
	if err != nil {
		return
	}
	client.conn <- string(b)
	str := <-client.conn
	var response1 Response
	err = json.Unmarshal([]byte(str), &response1)
	response = &response1
	return
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
