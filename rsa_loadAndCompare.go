package main

import (
    "crypto/x509"
    "encoding/gob"
    "encoding/pem"
    "fmt"
    "io/ioutil"
    "./stringio"
)

func cmp_bytes(b1 []byte, b2 []byte) (bool, error) {
    if len(b1) != len(b2) {
        return false, fmt.Errorf("cannot compare two byte array with differnt length")
    }
    for idx, b1c := range b1 {
        if b1c != b2[idx] {
            return false, nil
        }
    }
    return true, nil
}

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


    // 把私钥文件encode之后写到stringio(或者文件)
    sio := stringio.StringIO()
    privatekeyencoder := gob.NewEncoder(sio)
    privatekeyencoder.Encode(privatekey)
    // sio.Seek(0,0)
    fmt.Println("sio: ", sio.GetValueBytes())

    // 读取私钥数据文件
    privKeyData , _ := ioutil.ReadFile("private.key")
    fmt.Println("pdata: ", privKeyData)

    // 对比stringio 和 读取到的私钥密钥文件是否相等
    // b := make([]byte, 1)
    // for idx, c := range privKeyData {
    //     sio.Seek(int64(idx),0)
    //     sio.Read(b)

    //     // fmt.Printf("pdata idx: %d, privKeyData: %v , sio: %v \n", idx, c, b[0])
    //     // fmt.Printf("pdata idx: %d, p == sio , %v == %v, %v\n", idx, c, b[0], c == b[0])
    //     if c != b[0] {
    //         fmt.Printf("Error: different.  idx: %d, p != sio , %v != %v \n", idx, c, b)
    //     }
    // }

    equal, err := cmp_bytes(privKeyData, sio.GetValueBytes())
    fmt.Println("Compare: ", equal, err )

    // var publickey *rsa.PublicKey
    // publickey = &privatekey.PublicKey

    // var privBuff bytes.Buffer
    // var pubBuff bytes.Buffer
    // fmt.Println(privatekey.Bytes)

    // publickeyencoder := gob.NewEncoder(pubBuff)
    // publickeyencoder.Encode(publickey)
    // fmt.Println(pubBuff.String())

}
