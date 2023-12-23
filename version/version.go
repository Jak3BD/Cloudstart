package version

import "runtime"

var Version = "UNKNOWN"
var OS = runtime.GOOS
var Arch = runtime.GOARCH
