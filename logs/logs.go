package logs

import (
	"github.com/op/go-logging"
)

var (
	logFormat = "%{color}[%{level:.4s}] %{time:15:04:05.000000} %{id:06x} [%{shortpkg}] %{longfunc} -> %{color:reset}%{message}"
	Log       = logging.MustGetLogger("bizcraft")
)

func Start() {
	logging.SetFormatter(logging.MustStringFormatter(logFormat))
}
