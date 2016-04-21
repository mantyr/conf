package confdb

import (
    "github.com/mantyr/database/sql"
)

var db_list DbList

func init() {
    db_list = NewDBList()
}

func ConnectDB(params ...string) {
    db_list.ConnectDB(params...)
}

func GetDB(section string) *sql.DB {
    return db_list.GetDB(section)
}
