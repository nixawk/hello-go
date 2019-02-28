package main

import "encoding/json"
import "os"

func main() {
        enc := json.NewEncoder(os.Stdout)
        d := map[string]int{"XX": 11, "YY": 22}
        enc.Encode(d)
}
