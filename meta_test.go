package meta

import "testing"

func TestText(t *testing.T) {
	text := Text()
	t.Log(text) // Output like: GOVERSION:go1.18 GOOS: GOARCH: VCSREVISION: VCSTIME: UPTIME:2022-05-14T00:21:46+08:00
}
