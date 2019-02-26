package main

import "errors"
import "math"
import "fmt"

func Sqrt(value float64) (float64, error) {
        if (value < 0) {
                return 0, errors.New("Math: negative number passed to Sqrt")
        }
        return math.Sqrt(value), nil
}

func Output(result float64, err error) {
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println(result)
        }
}

func main() {
        result, err := Sqrt(-1)
        Output(result, err)

        result, err = Sqrt(9)
        Output(result, err)
}

/* Go programming provides a pretty simple error handling framework with
 * inbuilt error interface type of the following declaration.

 * type error interface {
 *      Error() string      
 * }

 * Functions normally return error as last return value. Use errors.New to
 * construct a basic error message as following

 * func Sqrt(value float64) (float, error) {
 *        if (value < 0) {
 *                return 0, errors.New("Math: negative number passed to Sqrt")
 *        }
 *        return math.Sqrt(value)
 * }
 *
 */