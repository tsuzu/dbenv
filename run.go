// Copyright (c) 2020 tsuzu
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
)

func runAction(outputPath string) {
	var output io.WriteCloser
	var err error

	if outputPath == "-" {
		output = os.Stdout
	} else {
		output, err = os.Create(outputPath)

		if err != nil {
			log.Fatalf("failed to create output file(%s): %+v", outputPath, err)
		}
		defer output.Close()
	}

	databaseURL := os.Getenv("DATABASE_URL")

	run(output, databaseURL)
}

func run(output io.Writer, databaseURL string) {
	dbURL, err := url.Parse(databaseURL)

	if err != nil {
		log.Fatalf("$DATABASE_URL is not valid: %+v", err)
	}

	if user := dbURL.User; user != nil {
		fmt.Fprintf(output, "export DATABASE_USER=%s\n", user.Username())

		if password, ok := user.Password(); ok {
			fmt.Fprintf(output, "export DATABASE_PASSWORD=%s\n", password)
		}
	}

	lcScheme := strings.ToLower(dbURL.Scheme)
	port := dbURL.Port()

	if len(port) == 0 {
		if strings.HasPrefix(lcScheme, "mysql") {
			port = "3306"
		} else if strings.HasPrefix(lcScheme, "postgre") {
			port = "5432"
		}
	}

	fmt.Fprintf(output, "export DATABASE_SCHEME=%s\n", dbURL.Scheme)
	fmt.Fprintf(output, "export DATABASE_HOST=%s\n", dbURL.Hostname())
	fmt.Fprintf(output, "export DATABASE_PORT=%s\n", port)
	fmt.Fprintf(output, "export DATABASE_DB=%s\n", strings.TrimLeft(dbURL.Path, "/"))
}
