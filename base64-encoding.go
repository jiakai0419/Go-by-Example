package main

import "fmt"
import b64 "encoding/base64"

func main() {
	P := fmt.Println

	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	P(sEnc)
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	P(string(sDec))
	P()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	P(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	P(string(uDec))
}
