// Copyright (c) 2020 tsuzu
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import "fmt"

// These variables are set in build step
var (
	Version  = "unknown"
	Revision = "unknown"
)

func versionAction() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Git Revision: %s\n", Revision)
}
