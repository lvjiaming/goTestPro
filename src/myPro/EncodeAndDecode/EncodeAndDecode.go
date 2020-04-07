/**
RSA和AES加解密
 */
package EncodeAndDecode

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// AES加密(AES 加密又分为 ECB、CBC、CFB、OFB 等几种)

func EncodeAndDecodeFunc()  {
	note := "www.golang.com"
	key := "asdhaodjadjlkajdlkajdajd"
	fmt.Println("AEC加密原内容：", note)
	encodeNode := CBCEncode(note, key)
	fmt.Println("AEC加密加密内容：", encodeNode)
	decodeNode := CBCDecode(encodeNode, key)
	fmt.Println("AEC加密解密内容：", decodeNode)

	note2 := "www.cocos.com"
	fmt.Println("RSA加密原内容：", note2)
	RsaEncodeNote, _ := RSAEncode(note2)
	fmt.Println("RSA加密加密内容：", base64.StdEncoding.EncodeToString(RsaEncodeNote))
	RsaDecodeNote, _ := RSADecode(RsaEncodeNote)
	fmt.Println("RSA加密解密内容：", string(RsaDecodeNote))
}

func CBCEncode(note string, key string) string { // CBC加密法
	noteData := []byte(note) // 讲内容转成字节数组
	k := []byte(key) // 将加密的key转成字节数组
	block, _ := aes.NewCipher(k) // 分组秘钥
	blockSize := block.BlockSize() // 获取秘钥长度

	orgiData := PKCS7Padding(noteData, blockSize) // 补全码

	blockMod := cipher.NewCBCEncrypter(block, k[:blockSize]) // 加密方式
	cryted := make([]byte, len(orgiData)) // 创建数组
	blockMod.CryptBlocks(cryted, orgiData) // 加密
	return base64.StdEncoding.EncodeToString(cryted)
}
func CBCDecode(note, key string) string { // CBC解密法
	noteData, _ := base64.StdEncoding.DecodeString(note) // 内容转成字节数
	k := []byte(key) // 秘钥转成字节数
	block, _ := aes.NewCipher(k) // 分组秘钥
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	cryted := make([]byte, len(noteData))

	blockMode.CryptBlocks(cryted, noteData)
	cryted = PKCS7UnPadding(cryted) // 去除补全码
	return string(cryted)
}
// 补码
func PKCS7Padding(noteData []byte, blockSize int) []byte {
	padding := blockSize - len(noteData) % blockSize
	paddtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(noteData, paddtext...)
}
// 去码
func PKCS7UnPadding(noteData []byte) []byte {
	len := len(noteData)
	unpadding := int(noteData[len - 1])
	return noteData[:len - unpadding]
}

// RSA加密（主要用于密码的加解密）

// 私钥
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDfw1/P15GQzGGYvNwVmXIGGxea8Pb2wJcF7ZW7tmFdLSjOItn9
kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQolDOkEzNP0B8XKm+Lxy4giwwR5
LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TUGQD6QzKY1Y8xS+FoQQIDAQAB
AoGAbSNg7wHomORm0dWDzvEpwTqjl8nh2tZyksyf1I+PC6BEH8613k04UfPYFUg1
0F2rUaOfr7s6q+BwxaqPtz+NPUotMjeVrEmmYM4rrYkrnd0lRiAxmkQUBlLrCBiF
u+bluDkHXF7+TUfJm4AZAvbtR2wO5DUAOZ244FfJueYyZHECQQD+V5/WrgKkBlYy
XhioQBXff7TLCrmMlUziJcQ295kIn8n1GaKzunJkhreoMbiRe0hpIIgPYb9E57tT
/mP/MoYtAkEA4Ti6XiOXgxzV5gcB+fhJyb8PJCVkgP2wg0OQp2DKPp+5xsmRuUXv
720oExv92jv6X65x631VGjDmfJNb99wq5QJBAMSHUKrBqqizfMdOjh7z5fLc6wY5
M0a91rqoFAWlLErNrXAGbwIRf3LN5fvA76z6ZelViczY6sKDjOxKFVqL38ECQG0S
pxdOT2M9BM45GJjxyPJ+qBuOTGU391Mq1pRpCKlZe4QtPHioyTGAAMd4Z/FX2MKb
3in48c0UX5t3VjPsmY0CQQCc1jmEoB83JmTHYByvDpc8kzsD8+GmiPVrausrjj4p
y2DQpGmUic2zqCxl6qXMpBGtFEhrUbKhOiVOJbRNGvWW
-----END RSA PRIVATE KEY-----
`)
// 公钥
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDfw1/P15GQzGGYvNwVmXIGGxea
8Pb2wJcF7ZW7tmFdLSjOItn9kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQol
DOkEzNP0B8XKm+Lxy4giwwR5LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TU
GQD6QzKY1Y8xS+FoQQIDAQAB
-----END PUBLIC KEY-----
`)

func RSAEncode(note string) ([]byte, error) { //RSA加密
	noteData := []byte(note)
	block, _ := pem.Decode(publicKey) // 解密pem格式的公钥
	if block == nil {
		return nil, errors.New("publicKey is err")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) // 解析公钥
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey) // 类型断言
	return rsa.EncryptPKCS1v15(rand.Reader, pub, noteData) // 加密
}

func RSADecode(noteData []byte) ([]byte, error) { // RAS解密
	block, _ := pem.Decode(privateKey) // 解密pem格式的私钥
	if block == nil {
		return nil, errors.New("privateKey is err")
	}
	pubInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes) // 解析私钥
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, pubInterface, noteData)
}
