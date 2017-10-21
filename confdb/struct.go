package confdb

import (
	"github.com/mantyr/database/sql"
	"sync"
)

var (
	Default_section string = "default"
)

type DbList struct {
	sync.RWMutex
	d map[string]*sql.DB
}
