package meta

import (
	"fmt"
	"runtime/debug"
	"time"
)

var meta *META

type META struct {
	// go compiler and os info
	GoVersion string
	GoOS      string
	GoArch    string

	// vcs info
	VCSRevision string
	VCSTime     string

	// app start up time
	UpTime time.Time
}

func init() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	meta := META{
		GoVersion: bi.GoVersion,
		UpTime:    time.Now(),
	}

	for _, s := range bi.Settings {
		switch s.Key {
		case "GOOS":
			meta.GoOS = s.Value
		case "GOARCH":
			meta.GoArch = s.Value
		case "vcs.revision":
			meta.VCSRevision = s.Value
		case "vcs.time":
			meta.VCSTime = s.Value
		}
	}
}

// Meta return MEAT pointer.
func Meta() *META {
	return meta
}

// Println use `fmt.Println` print to stdout.
func Println() {
	fmt.Println(Text())
}

// Text return a formatted string, or an empty string if some error occur.
func Text() string {
	return fmt.Sprintf("GOVERSION:%s GOOS:%s GOARCH:%s VCSREVISION:%s VCSTIME:%s UPTIME:%s",
		meta.GoVersion, meta.GoOS, meta.GoArch,
		meta.VCSRevision, meta.VCSTime,
		meta.UpTime.Format(time.RFC3339),
	)
}
