package tmath

import (
	"testing"
	"math"
	rand2 "math/rand"
	"time"
	"crypto/rand"
	"math/big"
	"os"
	"strings"
	"bytes"
	"os/exec"
	"log"
	"fmt"
	"encoding/json"
)

func TestMath(t *testing.T) {
	t.Log(math.Modf(3.14))
	t.Log(math.Mod(4, 3))
	t.Log(math.Remainder(14, 3), (14-math.Trunc(14/3))*3)

	ran := rand2.New(rand2.NewSource(99))
	t.Log(ran.Int(),rand2.Int(),ran.Intn(100))

	rand2.Seed(100)
	t.Log(rand2.Int())
	t.Log(rand2.Int31())
	t.Log(rand2.Intn(100))

	rand2.Seed(time.Now().UnixNano())
	t.Log(rand2.Intn(100))

	res,_:=rand.Int(rand.Reader,big.NewInt(100))
	t.Log(res)

}

func TestOs(t *testing.T) {
	t.Log(os.Hostname())
	t.Log(os.Getpagesize())
}

func TestCmd(t *testing.T) {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())


	cmd = exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)



}

func TestCmd1(t *testing.T) {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
