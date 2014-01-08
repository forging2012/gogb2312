package gogb2312

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func hexprint_utf8(u8 int) {
	fmt.Printf("%x %x %x\n", u8>>16, (u8>>8)&0xff, u8&0xff)
}

func hexprint_utf8string(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

func hexprint_bytes(s []byte) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

func Test_unicode2utf8(t *testing.T) {
	s := unicode2utf8(0x90ed)
	hexprint_utf8(s)
	b := make([]byte, 3)
	b[0] = byte(s >> 16)
	b[1] = byte((s >> 8) & 0xff)
	b[2] = byte(s & 0xff)
	fmt.Println(string(b))
}

func Test_RE(t *testing.T) {
	var ss = []string{
		"who  \t are u \t\t  ? ",
		"i \t \t am\t\t\t   boy.  ",
		"i \t \t am\t\t\t 90  boy.  ",
	}
	var gb = []byte("\x90\xae")
	var b = byte(98)

	for _, s := range ss {
		ret := re_space.ReplaceAllString(s, " ")
		_ = ret
		//fmt.Println(s, " - ", ret)
	}
	_ = gb
	_ = b
}

func Test_MakeData(t *testing.T) {
	//readcp936("./CP936.TXT")
}

func Test_Convert(t *testing.T) {
	buf, err := ioutil.ReadFile("./test1.txt")
	if err != nil {
		t.Error(err.Error())
	}
	hexprint_bytes(buf)
	output, cerr, il, ol := ConvertGB2312(buf)
	if cerr != nil {
		t.Error(cerr.Error())
	}
	_ = il
	_ = ol
	hexprint_utf8string("郭铁")

	hexprint_bytes(output)
	fmt.Printf("%s\n", output)
}
