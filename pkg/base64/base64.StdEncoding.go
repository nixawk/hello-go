package main

import "encoding/base64"
import "fmt"

func base64_demo(data string) {
        sEnc := base64.StdEncoding.EncodeToString([]byte(data))
        fmt.Println(sEnc)

        sDec, _ := base64.StdEncoding.DecodeString(sEnc)
        fmt.Println(string(sDec))

        uEnc := base64.URLEncoding.EncodeToString([]byte(data))
        fmt.Println(uEnc)

        uDec, _ := base64.URLEncoding.DecodeString(uEnc)
        fmt.Println(string(uDec))
}

func main() {
        base64_demo("https://golang.org/pkg/encoding/base64/#StdEncoding")
        base64_demo("ABCabc123!?$*&()'-=@~")
}
