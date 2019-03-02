// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
   Package hello implements a demo golang package.

*/

package hello

import "fmt"

// Welcome prints a greeting message.
func Welcome() {
	fmt.Println("Welcome to golang")
}
