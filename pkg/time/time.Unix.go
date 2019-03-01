package main

import "time"
import "fmt"

func main() {
        now := time.Now()

        secs := now.Unix()
        nanos := now.UnixNano()
        millis := nanos / 1000000

        fmt.Println(secs)   /* output: 1551398219 */
        fmt.Println(nanos)  /* output: 1551398219224093453 */
        fmt.Println(millis) /* output: 1551398303895 */

        fmt.Println(time.Unix(secs, 0))  /* output: 2019-03-01 07:59:25 +0800 CST */
        fmt.Println(time.Unix(0, nanos)) /* output: 2019-03-01 07:59:25.46520622 +0800 CST */
}
