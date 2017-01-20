package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func Encode(encBuf, bin []byte, e64 *base64.Encoding) []byte {
	maxEncLen := e64.EncodedLen(len(bin))
	if encBuf == nil || len(encBuf) < maxEncLen {
		encBuf = make([]byte, maxEncLen)
	}
	e64.Encode(encBuf, bin)
	return encBuf[0:]
}

func Decode(decBuf, enc []byte, e64 *base64.Encoding) []byte {
	maxDecLen := e64.DecodedLen(len(enc))
	if decBuf == nil || len(decBuf) < maxDecLen {
		decBuf = make([]byte, maxDecLen)
	}
	n, err := e64.Decode(decBuf, enc)
	_ = err
	return decBuf[0:n]
}

// 文字列を反転して返す
func reverse(s string) string {
	ans := ""
	for i := range s {
		ans += string(s[len(s)-i-1])
	}
	return string(ans)
}

func ToHex(c byte) []byte {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%02X", c)
	return b.Bytes()
}

func Base64Encode(userID string) string {
	bin := []byte(userID)
	e64 := base64.StdEncoding
	enc := Encode(nil, bin, e64)
	return string(enc)
}

// EncodeID はアクティベート処理において簡易的な暗号化をおこなうために使用されます
func EncodeID(id string) string {
	return HexEncode(Base64Encode(id))
}

// DecodeID は EncodeID によってエンコードされた文字列をデコードします
func DecodeID(encodeID string) string {
	return Base64Decode(HexDecode(encodeID))
}

// Base64Decode はBase64エンコードされた文字列をデコードします
func Base64Decode(str string) string {
	b := []byte(str)
	return string(Decode(nil, b, base64.StdEncoding))
}

// HexEncode は文字列をバイト単位で処理し、1byteを16進化させたデータをある規則で結合した結果を返します
func HexEncode(UserIDBase64 string) string {
	enc := []byte(UserIDBase64)
	var bb bytes.Buffer
	var be bytes.Buffer

	for _, c := range enc {
		hex := ToHex(c)
		bb.Write(hex[:1])
		be.Write(hex[1:])
	}
	encodedID := bb.String() + reverse(be.String())
	return encodedID
}

// HexDecode は HexEncode で取得した文字列にたいしてデコードをおこないます
func HexDecode(str string) string {
	mae := strings.ToUpper(str[:len(str)/2])
	ato := reverse(strings.ToUpper(str[len(str)/2:]))

	var bb bytes.Buffer

	for i := range mae {
		s := make([]byte, 2)
		s[0] = mae[i]
		s[1] = ato[i]
		sh, _ := hex.DecodeString(string(s))
		bb.Write(sh)
	}
	return string(bb.Bytes())
}
