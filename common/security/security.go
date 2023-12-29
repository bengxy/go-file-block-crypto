package security

import (
	"fmt"
	"github.com/deatil/go-cryptobin/cryptobin/crypto"
)

var key = "dfertf12dfertf12" // length = 16

func Encrypt(data []byte) string {

	cypt := crypto.FromBytes(data).
		SetKey(key).
		Aes().
		ECB().
		PKCS7Padding().
		Encrypt().
		ToBase64String()
	fmt.Println("加密", "from=", data, ",to=", cypt)
	return cypt
}

func Decrypt(data string) []byte {
	cyptde := crypto.
		FromBase64String(data).
		SetKey(key).
		Aes().
		ECB().
		PKCS7Padding().
		Decrypt().
		ToBytes()
	fmt.Println("解密", string(cyptde))
	return cyptde
}
