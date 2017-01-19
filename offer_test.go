package offer

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestStrings(t *testing.T) {

	s1 := "golang"
	s2 := "apple1"
	joint := make([]byte, 2)
	joint[0] = s1[0]
	joint[1] = s2[0]
	if string(joint) != "ga" {
		t.Error("error")
	}
	t.Log("OK")
}

func TestHexpkg(t *testing.T) {

	e := hex.EncodeToString([]byte("M"))
	d, _ := hex.DecodeString(e)
	if string(d) != "M" {
		t.Error("error")
	}
	sss, _ := hex.DecodeString("4D54497A4E413D3D")
	if string(sss) != "MTIzNA==" {
		t.Error("error")
	}
}

func TestB64(t *testing.T) {
	e64 := base64.StdEncoding
	enc := Encode(nil, []byte("1234"), e64)
	if string(enc) != "MTIzNA==" {
		t.Error("error")
	}
}
func TestOfferDecode(t *testing.T) {

	if d := DecodeID("45474433DD1EA94D"); d != "1234" {
		t.Errorf("ERR: not match\n expect[%s]\n actual[%s]", "1234", d)
	}
}

func TestOfferEncode(t *testing.T) {

	if e := EncodeID("1234"); e != "45474433DD1EA94D" {
		t.Errorf("ERR: not match\n expect[%s]\n actual[%s]", "45474433DD1EA94D", e)
	}
}
