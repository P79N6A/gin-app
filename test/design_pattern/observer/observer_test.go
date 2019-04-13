/**
 * @description observer
 * @author zhangbingbing@baidu.com
 * @date 2019/3/13
 */
package observer

import (
	"testing"
	"reflect"
	"log"
	"gin-app/test/design_pattern/observer/event"
	"time"
	"encoding/json"
	"hash/crc32"
	"fmt"
	"github.com/boguslaw-wojcik/crc32a"
)

type TestEvent struct {
	Type int
}

func TestListener(t *testing.T) {
	evt := TestEvent{1}
	tt := reflect.TypeOf(evt)
	log.Println(tt, tt.Name() == "TestEvent")
	_, ok := interface{}(evt).(TestEvent)
	log.Println(ok)
}

func TestListener1(t *testing.T) {
	context := NewDefaultApplicationContext()
	evt := &event.AEvent{}
	context.Publish(evt)
	ent := &event.BEvent{}
	context.Publish(ent)
	time.Sleep(1 * time.Second)
}

func TestListener2(t *testing.T) {
	evt := &event.AEvent{}
	DefaultContext.Publish(evt)
	ent := &event.BEvent{}
	DefaultContext.Publish(ent)
	time.Sleep(1 * time.Second)
}

type base struct {
	ID int
}
type shape struct {
	base
	Width  int
	Height int
}

type AShape struct {
	shape
	Name string
}

type BShape struct {
	ID     int
	Width  int
	Height int
	Name   string
}

func TestSome(t *testing.T) {
	a := &BShape{}
	a.ID = 1
	a.Name = "aa"
	a.Width = 10
	a.Height = 20
	bt, err := json.Marshal(a)
	log.Println(string(bt), err)

	var bb AShape
	err = json.Unmarshal(bt, &bb)
	log.Println(bb, err, bb.Height, bb.Width, bb.Name)
	log.Printf("%+v\n", bb)
	bb.Width = 1
	log.Printf("%+v\n", bb)
	update(&bb)

	crc32q := crc32.MakeTable(0x82f63b78)
	fmt.Printf("%d\n", crc32.Checksum([]byte("21D682658FA29B1971F453E8B420F956|569992630852268fcf20f"), crc32q))
	fmt.Println(crc32.ChecksumIEEE([]byte("hello")))
	crc32.Checksum([]byte("21D682658FA29B1971F453E8B420F956|569992630852268fcf20f"), crc32q)
}
func update(aShape *AShape) {
	aShape.Height = 2
	aShape.ID = 2
	log.Printf("%+v\n", aShape)
}

func TestCRC(t *testing.T) {
	var expConfigs []ExpConfig
	err := json.Unmarshal([]byte(jsonConf), &expConfigs)
	if err != nil {
		log.Println(err)
		return
	}
	salt := "fcf20f"
	CUID := "21D682658FA29B1971F453E8B420F956|569992630852268"
	var platform uint32 = 2

	data := CUID + salt
	res := crc32a.Checksum([]byte(data))
	resHex := crc32a.ChecksumHex([]byte(data))
	userHash := res % 10000
	for _, expConfig := range expConfigs {
		if platform != expConfig.Platform {
			continue
		}
		userSharding := userHash / 100
		userComp := userHash % 100
		var branchIndex uint32 = 0
		var skipComp uint32 = 0
		if expConfig.Sharding[0] <= userSharding && userSharding < expConfig.Sharding[1] {
			for userComp >= skipComp {
				skipComp = skipComp + expConfig.Rates[branchIndex]
				branchIndex++
			}
		}
		choiceID := expConfig.ChoicesIDs[branchIndex-1]
		log.Printf("res:%d resHex:%s userHash:%d userSharding:%d userComp:%d branchIndex:%d choiceID:%d\n", res, resHex, userHash, userSharding, userComp, branchIndex, choiceID)
	}

}

func TestCRCIOS(t *testing.T) {

	expConf := map[string]interface{}{
		"salt": "fcf20f",
		//"salt":          "cfabcd",
		"exp_id":        1472,
		"platform":      1,
		"version_range": []int{0, 14},
		"sharding":      []int{0, 100},
		"choices_ids":   []int{3437, 3438},
		"rates":         []int{50, 50},
	}
	CUID := "D0FD9077732331EBACDE81EF4176E7685815BC16DFGLHFJORGC"
	//CUID := "55418458E8AE10630E881A86A9B2A0E536EB8140COMHNEMOKNC"
	data := CUID + expConf["salt"].(string)
	log.Println(data)
	res := crc32.ChecksumIEEE([]byte(data))

	log.Println(res)
	log.Println("/100", ((res%10000+10000)%10000)/100)
	log.Println("%100", ((res%10000+10000)%10000)%100)
	userSharding := ((res%10000 + 10000) % 10000) / 100 //57
	userComp := ((res%10000 + 10000) % 10000) % 100     //79

	rates := []int{50, 50}
	sharding := []uint32{0, 100}

	log.Println("userSharding:", userSharding)
	log.Println("userComp:", userComp)
	if sharding[0] <= userSharding && userSharding < sharding[1] {
		branchIndex := 0
		skipComp := rates[0]
		for int(userComp) >= skipComp {
			skipComp = skipComp + rates[branchIndex]
			branchIndex++
		}
		log.Println("index:", branchIndex)
	}
}

func TestXXX(t *testing.T) {
	salt := "fcf20f"
	cuids := []string{
		"D0FD9077732331EBACDE81EF4176E7685815BC16DFGLHFJORGC",
		"21D682658FA29B1971F453E8B420F956|569992630852268",
		"F7663B3DE215D9B83212D15BC4355561|57B49E0400000A",
		"F81D111D254C023A35D3DB2936CD332A|5A2F0A2500001A",
		"F78635A0E1C00806B670AC6B6A8ED92C|85939415300099",
		"F5FDD466F033A078010D0ABE7C4C2D4F|DB1EA65700000A",
		"F80E1BA9BFEFD3AA2F5B06BA4F0E4A31|59348920200099",
		"F5C097B9C20711F4F984EB087CCE8C09|FB25BE8400000A",
		"F81D191A8CB04F6A36AAFF145A4A62C5|CCE9210500001A",
		"F58694D60AF15B4D35B36AFAF50E1775|8127E19500000A",
		"F810C695096BC566139BCD277F571D6C|23261916500099",
		"F5D32CFA72EC7DB093E6BA84BF79CCCE|CBACF06500001A",
		"F81CB319B8EF3F6E7F0934AEEB42EE11|046ABC0500001A",
		"F816CF4F89FDF83CBE3AB24574B2E9BF|9DEED585500099",
		"F789BE477F481CFB491B20B4DEB414B7|10586827000099",
		"F815568849583C5BE6C879945B9D84FA|25988301600099",
		"F5D917676EF82A901EFE0995382C3085|CD88C00200032A",
		"F5058A98A5EA1C07E4BBA81871FE6CAA|226C96B500001A",
		"F4D4C14F253671A5F3A190ABEA2E0079|46454127000099",
		"F62887DAC966F005FA7A0623605D7B6A|AC04749400000A",
		"F781E1F358405BD1D709982F77064761|758D1A0500001A",
		"F76F30CA5D8DA0CB12A7C7F3174508AC|209980F400000A",
		"552CF79160593B90EAE3B22D124F22B8E89771A56FMHGPMGABG",
		"552C125EE853EBA23DA0BC10D349177E35B714290FMFRDJTLOH",
		"55499B01902219051795DBDD775C48B8E3372E5F6OMMGEGIGRB",
		"5549269E92A1513F2E0669A709219583FA8D09DA2OHHGESPOBK",
		"55488BF4FE08000767AB9DD858021BA023760DB68FMMGESFCGO",
		"55484F4A68BA396E6AED3B1227013DD4E6B36B83AFCNFIHKLAO",
		"55483516545F9967B484341E61FED097CE86592C4FHMNEFQDNH",
		"5547AD4F36CCB2F9E3F62C0DA63C3E9E76CCECDC4ORHDHCPKII",
		"5546C0386A031D1F5BFA59DF8445FBDE794F3D238ORQCEMLITD",
		"5546376872E7FCA625FBEBDB252AFF156341AB8B6OMMGPMHIST",
		"5545A578E72D329DFF6E4D63CE4EDEA844E96EBC8OMMGECIKIA",
		"55455EE0042A13C565B1F2FBC531ACE9D09F488D6FHNGAOHPQT",
		"554549009807CC5EAA5448A0FD1B92CB058020A5AFCBQAMRQPQ",
		"55446F7EDCFF3AC756881E3E7F6153441E6970D11OHFHHTQEHJ",
		"554442D1344C0A92F26AC6A67C0EEC9A8ED41E69DFMDBICDFMT",
		"55442B8D951E28940DD9B37538312B9D795D6E471FMHGPHDQMB",
		"55426651F90AC33D9C87393CB4E599288574F91D2FMHGPMOJGH",
		"55425FBE377089C1972008F617054C112FD08F76FFHEALTEABL",
		"55418458E8AE10630E881A86A9B2A0E536EB8140COMHNEMOKNC",
		"55409A7078E1C9FB3583BF61CDE52CD27762C3B4CFHPBBAPNJG",
	}
	//str := "21D682658FA29B1971F453E8B420F956|569992630852268fcf20f"
	//hash := md5.New()
	//io.WriteString(hash, str)
	//data := hash.Sum(nil)
	//tmp := hex.EncodeToString(data)
	//log.Println(crc32.ChecksumIEEE([]byte(tmp)))
	//
	//crc32q := crc32.MakeTable(0x82f63b78)
	//log.Printf("%d\n", crc32.Checksum([]byte(str), crc32q))
	//
	//crc32q1 := crc32.MakeTable(0xeb31d82e)
	//log.Printf("%d\n", crc32.Checksum([]byte(str), crc32q1))
	//
	//log.Println(crc32a.Checksum([]byte(str)))
	//log.Println(crc32a.ChecksumHex([]byte(str)))

	for _, cuid := range cuids {
		intstr := crc32a.Checksum([]byte(cuid + salt))
		hexstr := crc32a.ChecksumHex([]byte(cuid + salt))
		//intstr := crc32.Checksum([]byte(cuid+salt), crc32q)
		//hexstr := fmt.Sprintf("%u", intstr)
		userHash := intstr % 10000
		userSharding := float64(userHash) / 100
		userComp := userHash % 100
		fmt.Printf("cuid:%s hex:%s int:%d userHash:%d userSharding:%v userComp:%d\n", cuid, hexstr, intstr, userHash, userSharding, userComp)
	}

	fmt.Println(crc32a.Checksum([]byte("hello")))

	var conf []ExpConfig
	_ = json.Unmarshal([]byte(jsonConf), &conf)
	fmt.Printf("%+v\n", conf)
}

//var jsonConf = "{\"exp_id\":1471,\"platform\":1,\"time_window\":[1547740800,1559232000],\"version_range\":[\"0\",\"14\"],\"sharding\":[0,100],\"is_upload\":true,\"choices\":[{\"abtest_novel_recommend\":false},{\"abtest_novel_recommend\":false}],\"events\":[[],[]],\"choices_ids\":[3435,3436],\"rates\":[50,50],\"is_immediate\":true}"
var jsonConf = "[{\"exp_id\":1471,\"platform\":1,\"time_window\":[1547740800,1559232000],\"version_range\":[\"0\",\"14\"],\"sharding\":[0,100],\"is_upload\":true,\"choices\":[{\"abtest_novel_recommend\":false},{\"abtest_novel_recommend\":false}],\"events\":[[],[]],\"choices_ids\":[3435,3436],\"rates\":[50,50],\"is_immediate\":true},{\"exp_id\":1472,\"platform\":2,\"time_window\":[1547740800,1559232000],\"version_range\":[\"0\",\"14\"],\"sharding\":[0,100],\"is_upload\":true,\"choices\":[{\"abtest_novel_recommend\":false},{\"abtest_novel_recommend\":false}],\"events\":[[],[]],\"choices_ids\":[3437,3438],\"rates\":[50,50],\"is_immediate\":true}]"

type ExpConfig struct {
	ExpID        uint32 `json:"exp_id"`
	Platform     uint32 `json:"platform"`
	TimeWindow   []uint32 `json:"time_window"`
	VersionRange []string `json:"version_range"`
	Sharding     []uint32 `json:"sharding"`
	IsUpload     bool    `json:"is_upload"`
	Choices      []map[string]bool`json:"choices"`
	Events       []interface{}`json:"events"`
	ChoicesIDs   []uint32`json:"choices_ids"`
	Rates        []uint32 `json:"rates"`
	IsImmediate  bool    `json:"is_immediate"`
}
