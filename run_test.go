// Copyright (c) 2020 tsuzu
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"bytes"
	"testing"
)

func testRun(t *testing.T, input, output string) {
	t.Helper()
	buf := bytes.NewBuffer(nil)

	run(buf, input)

	if buf.String() != output {
		t.Errorf("unexpected result\nexpected:\n%s\nactual:%s\n", output, buf.String())
	}
}

func TestRun(t *testing.T) {
	t.Run("Heroku ClearDB", func(t *testing.T) {
		input := "mysql://USERNAME:PASSWORD@REGIONNAME.cleardb.net/heroku_92599a6734fc1a6?reconnect=true"
		output := `export DATABASE_USER=USERNAME
export DATABASE_PASSWORD=PASSWORD
export DATABASE_SCHEME=mysql
export DATABASE_HOST=REGIONNAME.cleardb.net
export DATABASE_PORT=3306
export DATABASE_DB=heroku_92599a6734fc1a6
`
		testRun(t, input, output)
	})

	t.Run("Heroku Postgres", func(t *testing.T) {
		input := "postgres://USERNAME:PASSWORD@REGION-NAME.compute-1.amazonaws.com:5432/d2eel7vtbj70hk"
		output := `export DATABASE_USER=USERNAME
export DATABASE_PASSWORD=PASSWORD
export DATABASE_SCHEME=postgres
export DATABASE_HOST=REGION-NAME.compute-1.amazonaws.com
export DATABASE_PORT=5432
export DATABASE_DB=d2eel7vtbj70hk
`
		testRun(t, input, output)
	})

	t.Run("Postgres no port", func(t *testing.T) {
		input := "postgres://USERNAME:PASSWORD@REGION-NAME.compute-1.amazonaws.com/d2eel7vtbj70hk"
		output := `export DATABASE_USER=USERNAME
export DATABASE_PASSWORD=PASSWORD
export DATABASE_SCHEME=postgres
export DATABASE_HOST=REGION-NAME.compute-1.amazonaws.com
export DATABASE_PORT=5432
export DATABASE_DB=d2eel7vtbj70hk
`
		testRun(t, input, output)
	})

	t.Run("MySQL no port", func(t *testing.T) {
		input := "mysql://USERNAME:PASSWORD@REGIONNAME.cleardb.net/heroku_92599a6734fc1a6?reconnect=true"
		output := `export DATABASE_USER=USERNAME
export DATABASE_PASSWORD=PASSWORD
export DATABASE_SCHEME=mysql
export DATABASE_HOST=REGIONNAME.cleardb.net
export DATABASE_PORT=3306
export DATABASE_DB=heroku_92599a6734fc1a6
`
		testRun(t, input, output)
	})

}
