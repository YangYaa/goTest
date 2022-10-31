package main

import(
	"goTest/encryption/openssl"
)

func main(){
	data := []byte{0x01,0x02}
	key := "root"
	openssl.Des3ECBDecrypt(data,[]byte(key),openssl.ZEROS_PADDING)
}

