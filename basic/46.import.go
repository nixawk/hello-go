package main

import . "fmt"

/*

Import declaration          Local name of Sin

import   "lib/math"         math.Sin
import m "lib/math"         m.Sin
import . "lib/math"         Sin

An import declaration declares a dependency relation between the importing and
imported package. It is illegal for a package to import itself, directly import
a package without referring to any of its exported identifiers. To import a
package solely for its side-effects (initialization). use the blank identifier as explicit package name:

        import _ "lib/math"

*/

func main() {
        Println("https://golang.org/ref/spec#Import_declarations")
}

/* https://stackoverflow.com/questions/21220077/what-does-an-underscore-in-front-of-an-import-statement-mean-in-golang */
/* https://golang.org/ref/spec#Import_declarations */
