package conf

import (
	"sync"
)

const Default_section string = "default"

var (
	DirBin string = "."
)

type ConfigList struct {
	sync.RWMutex
	d map[string]ConfigFile

	default_dir  string
	default_file string
}

type ConfigFile struct {
	Error error
	d     map[string]ConfigSection
}

type ConfigSection struct {
	d map[string]string
}
