package confdb

import (
    "sync"
    "github.com/mantyr/database/sql"
)

var (
    Default_section string = "default"
)

type DbList struct {
    sync.RWMutex
    d map[string]*sql.DB
}