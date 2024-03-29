// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"github.com/go-corelibs/x-text/language"
	"github.com/go-corelibs/x-text/message"
)

// The printer defined here will be picked up by the first print statement
// in main.go.
var printer = message.NewPrinter(language.English)
