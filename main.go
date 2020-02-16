// Copyright (c) 2020 tsuzu
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import "flag"

var (
	version = flag.Bool("version", false, "Show version")
	help    = flag.Bool("help", false, "Show usage")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()

		return
	}

	if *version {
		versionAction()

		return
	}

	arg := ".env"

	if len(flag.Args()) >= 1 {
		arg = flag.Arg(0)
	}

	runAction(arg)
}
