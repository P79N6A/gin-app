package encryption

import (
	"testing"
	"crypto/md5"
	"crypto/sha1"
	"os"
	"io"
)

func TestEncrypt(t *testing.T) {
	str := "hell,bill"
	md5Inst := md5.New()
	md5Inst.Write([]byte(str))
	result := md5Inst.Sum(nil)
	t.Log(result)
	t.Log(string(result))
	t.Logf("%x\n\n", result)

	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(str))
	result = sha1Inst.Sum(nil)
	t.Logf("%x\n\n", result)

    t.Logf("%x",md5.Sum([]byte("hello,bill")))
}
func TestEncryptFile(t *testing.T) {
	file := "123.txt"
	infile, err := os.Open(file)
	if err == nil {
		var result []byte = make([]byte, 111)
		n, _ := infile.Read(result[0:])
		t.Log(string(result[0:n]))

		md5h := md5.New()
		io.Copy(md5h, infile)
		t.Logf("%x %s\n", md5h.Sum(nil), file)

		shah := sha1.New()
		io.Copy(shah, infile)
		t.Logf("%x %s\n", shah.Sum(nil), file)
	}
}
