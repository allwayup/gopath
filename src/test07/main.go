package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

//生成RSA私钥和公钥，保存到文件中
func GenerateRSAKey(bits int) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	pem.Encode(publicFile, &publicBlock)
}

//RSA加密
func RSA_Encrypt(plainText []byte, path string) []byte {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	//返回密文
	return cipherText
}

//RSA解密
func RSA_Decrypt(cipherText []byte, path string) []byte {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	//返回明文
	return plainText
}

func main() {
	//生成密钥对，保存到文件
	GenerateRSAKey(2048)
	message := []byte("MIIEowIBAAKCAQEAvu0GyfngN3IXLRPQbeUNBx0xuvBT0YuDTRpIsPcPw5yPDWk0KxVilxhH2m8DU8uRytNIbFRRAWk5iyAxzChEbtuGho4tImEmU+OTcdVGuZKmwQ4JprviF3MKavIEXyWzUHJIY4wF0yz4gINiMybBHn7D9uz3mg9stfyACjAhXEgW9bPT2GttpgQJkceQpwaT3nNcXD113ZFaSW6kgKwZBLYJX3VhTYE9tplNFGGXbFfpz7KQvMnNQHku2245QsR+ziJaZqwPpk0kqs/qsRhx90aLzX5FCkJ2uKPumFPT3MprzMaaug2iVl0Cl2O46hlLsMtnzqgOodW6xUi1bY80WQIDAQABAoIBABigL+D7Rs0//PdGd/rEsWJ6hICNIPKFISFfw4J3y2O6nMTpDd6EupuseRAWg4JaXmqfx9aZJX2eGdr1AxdHFlSKIhbW3cFycGQflVP4Y0/qN3HtIpeL4kSOBQj4QFIZZcB5jRax58puIXtJ9u+MDxqk1RfTRrhrRuVONSGbRaJGPIyPyaDvXCuMmcA948sn6AeYHnnQTopw7nos2PMHOlddnGalq0NwBvo+GnHpDicNUT/Cd1XDP7Z61SJdwM27hog0txL8QL8SQFucTXkdQZvzSOzdk8lMwkFA2ML0mrdlw8iI/5SqhJSLw9eq4GJERIWcyiCSn0g1sU5KPhqjOoECgYEA/MmJpTJcm4RF9F0sjtDd1m29lrenDBTPEvSzorLI2gMS/a8CvqxhxKTKGPU7Idhot96kETlJAKDdNNqJP3t6hu0cgGeE/MIvFaDT6K5h3jb/qIDYVqgj6Q6zdtskcL+WfsmqQktBS9gMXlXtcvfXfdjHEuH3SVyrTBHB7G3UIJECgYEAwVo33YX0vf61p4xC7+BIkGQjISLiDf9owiJEA1pA96L/VNmp6IYw0FASt/PMDmxtqBTcvJ9BwviChLvJiEFT9FBAqCBODxC6Od3HR2lTy6awb2yNrvCQYNdsqp+8pZb+tZySbG5fCA2VWMWqPayU4e0tQTpwoZGybV6HpsfIu0kCgYBAGywZBMiPd9/1tJtULIvVkUb/LdvjKHPLLttPa1+cSiNKylM736N7pv7JjYdNcgA8gO3CoHBvBFyUxsb/nmTYStFrjtUe9G/UYFDdNTwEipYTOXmjoEhbFitU/QYkwbF8vc+7uDH69fNNCSWKfmfbtlnl5AA+To3yYJ55QvEEwQKBgFLvbs0PW3Zvnd2bVU7tJlMBEOxyuQIGDxpOdlv1x64w9VKg9rdtb9y6q/zJjzqUmciiAjjKGvwVem6S2hQe6XL/RWyYRsNBio+tqH/iFvZgrodsya1DNLrFTLA3SkTA6spduZTXFt4ubWQhjS9dKpNqF6JF/e/fvegZxxfr1Bc5AoGBANMC/esbBZQBru6DVv4rO74u+Vx57VJCHr2XawEfw0kVvRW95D+LoxfOYZpJ0x/AVFA5W6ZX3KpOI038JOX7ZqNTzOg+y7iid032lrfunTNB2oFyJofzVk33DhMy/TODcBYl4lu7Gkmpz+pCnXDNcLrqFDBrBHe0PQJbjxl2xH81")
	length := len(message)
	fmt.Println(length)
	retxLen := 245
	retx := make([]byte, retxLen)
	for i := 0; i < retxLen; i++ {
		retx[i] = message[i]
	}

	//加密
	start := time.Now().UnixMilli()
	fmt.Println(start)
	cipherText := RSA_Encrypt(retx, "public.pem")
	fmt.Println(time.Now().UnixMilli() - start)
	fmt.Println("长度为: ", len(cipherText))
	//解密
	start = time.Now().UnixMilli()
	plainText := RSA_Decrypt(cipherText, "private.pem")
	fmt.Println("解密后为：", string(plainText))
	fmt.Println(time.Now().UnixMilli() - start)
}
