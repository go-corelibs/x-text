// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

// NOTE: The command line tool already prefixes with "gotext:".
var (
	wrap = func(err error, msg string) error {
		if err == nil {
			return nil
		}
		return fmt.Errorf("%s: %v", msg, err)
	}
	errorf = fmt.Errorf
)
