// beta version
package tmpl

import (
	"github.com/mantyr/conf"
	"github.com/mantyr/runner"
	"html/template"
	"path/filepath"
	"strings"
	"sync"
)

var (
	subdir_sync sync.RWMutex
	subdir      string

	funcMap template.FuncMap = template.FuncMap{
		"humanInt": HumanInt,
		"isStatus": IsStatus,
		"json":     TMPL_json,

		"join":              strings.Join,
		"joinAll":           runner.Join,
		"joinInt64":         runner.JoinInt64,
		"joinMapInt64All":   runner.JoinMapInt64All,
		"joinMapInt64":      runner.JoinMapInt64,
		"joinMapInt64Bool":  runner.JoinMapInt64Bool,
		"joinMapStringAll":  runner.JoinMapStringAll,
		"joinMapString":     runner.JoinMapString,
		"joinMapStringBool": runner.JoinMapStringBool,
	}
)

func InitDir(address string) {
	SetDir(address)
}

func GetTemplate(address string) (*template.Template, error) {
	return template.New(filepath.Base(address)).Delims("{", "}").Funcs(funcMap).ParseFiles(conf.GetDirBin() + GetDir() + address)
}

func SetDir(address string) {
	subdir_sync.Lock()
	defer subdir_sync.Unlock()

	subdir = address
}

func GetDir() string {
	subdir_sync.RLock()
	defer subdir_sync.RUnlock()

	return subdir
}
