package main

import "encoding/json"
import "fmt"

func main() {
        var dat map[string]interface{}

        str := `{"XX": 11, "YY": 22}`
        if err := json.Unmarshal([]byte(str), &dat); err != nil {
                panic(err)
        }
        fmt.Println(dat)
}