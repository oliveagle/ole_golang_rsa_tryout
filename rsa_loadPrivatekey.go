package main

import (
    // "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/gob"
    "encoding/pem"
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    // 读取 pem 密钥文件
    pemData, _ :=  ioutil.ReadFile("private.pem")
    block, _ := pem.Decode(pemData)
    if block == nil {
        fmt.Println("Error: bad key data, not PEM-encoded")
    }

    // 解析私钥对象
    privatekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        fmt.Printf("Error: bad private key: %s", err)
    }
    fmt.Println(privatekey)
    fmt.Println(&privatekey.PublicKey)

    // 私钥对象对应的公钥, 所以只要有私钥就可以直接生产公钥，而且每次公钥都一致
    var publickey *rsa.PublicKey
    publickey = &privatekey.PublicKey

    // save private and public key separately
    privatekeyfile, err := os.Create("private.key.decoded")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    // gob 编码并保存
    privatekeyencoder := gob.NewEncoder(privatekeyfile)
    privatekeyencoder.Encode(privatekey)
    privatekeyfile.Close()

    publickeyfile, err := os.Create("public.key.decoded")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    publickeyencoder := gob.NewEncoder(publickeyfile)
    publickeyencoder.Encode(publickey)
    publickeyfile.Close()

}
