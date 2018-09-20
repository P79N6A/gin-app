package main

import (
	"net"
	"log"
	"encoding/binary"
	"bufio"
	"io"
	"fmt"
)

const (
	MAGIC_REQ uint8 = 0x80
	MAGIC_RES uint8 = 0x81
)

const (
	OP_GET       uint8 = 0x00
	OP_SET       uint8 = 0x01
	OP_ADD       uint8 = 0x02
	OP_REPLACE   uint8 = 0x03
	OP_DELETE    uint8 = 0x04
	OP_INCREMENT uint8 = 0x05
	OP_DECREMENT uint8 = 0x06
	OP_FLUSH     uint8 = 0x08
	OP_NOOP      uint8 = 0x0a
	OP_VERSION   uint8 = 0x0b
	OP_GETK      uint8 = 0x0c
	OP_APPEND    uint8 = 0x0e
	OP_PREPEND   uint8 = 0x0f
)

type request_header struct {
	magic    uint8
	opcode   uint8
	keylen   uint16
	extlen   uint8
	datatype uint8
	status   uint16
	bodylen  uint32
	opaque   uint32
	cas      uint64
}
type response_header struct {
	magic    uint8
	opcode   uint8
	keylen   uint16
	extlen   uint8
	datatype uint8
	status   uint16
	bodylen  uint32
	opaque   uint32
	cas      uint64
}

type response struct {
	header   *response_header
	bodyByte []byte
	body     interface{}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:11211")
	fmt.Println("000",conn,err)
	if err != nil {
		log.Fatalf("connect memcached fail: %v", err)
	}

	key := "foo"
	header := &request_header{
		magic:    MAGIC_REQ,
		opcode:   OP_GET,
		keylen:   uint16(len(key)),
		extlen:   0x00,
		datatype: 0x00,
		status:   0x00,
		bodylen:  uint32(len(key)),
		opaque:   0x00,
		cas:      0x00,
	}
	fmt.Println("1111",header)

	in_buf := make([]byte, 24)
	in_buf[0] = byte(header.magic)
	in_buf[1] = byte(header.opaque)
	binary.BigEndian.PutUint16(in_buf[2:4], header.keylen)
	in_buf[4] = byte(header.extlen)
	in_buf[5] = byte(header.datatype)
	binary.BigEndian.PutUint16(in_buf[6:8], uint16(header.status))
	binary.BigEndian.PutUint32(in_buf[8:12], header.bodylen)
	binary.BigEndian.PutUint32(in_buf[12:16], header.opaque)
	binary.BigEndian.PutUint64(in_buf[16:24], header.cas)

	buffered := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	buffered.WriteString(key)
	fmt.Println("222",key)
	err = buffered.Flush()
	fmt.Println("333",err)
	if err != nil {
		log.Fatalf("flush fail: %v", err)
	}

	b := make([]byte, 24)
	_, err = buffered.Read(b)
	fmt.Println("444",err)
	if err != nil {
		log.Fatalf("buffer read fail: %v", err)
	}
	fmt.Println("333",b)
	rheader := &response_header{
		magic:    uint8(b[0]),
		opcode:   uint8(b[1]),
		keylen:   uint16(binary.BigEndian.Uint16(b[2:4])),
		extlen:   uint8(b[4]),
		datatype: uint8(b[5]),
		status:   uint16(binary.BigEndian.Uint16(b[6:8])),
		bodylen:  uint32(binary.BigEndian.Uint32(b[8:12])),
		opaque:   uint32(binary.BigEndian.Uint32(b[12:16])),
		cas:      uint64(binary.BigEndian.Uint64(b[16:24])),
	}

	res := &response{header: rheader}
	if rheader.bodylen > 0 {
		res.bodyByte = make([]byte, rheader.bodylen)
		io.ReadFull(buffered, res.bodyByte)
	}
	fmt.Println(string(res.bodyByte), res)

}
