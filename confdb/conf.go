package confdb

import (
    "github.com/mantyr/conf"
    "github.com/mantyr/database/sql"
  _ "github.com/mantyr/go-sql-driver-mysql"
)

func NewDBList() (d DbList) {
    d.d = make(map[string]*sql.DB)
    return
}

func (d *DbList) ConnectDB(params ...string) {
    if len(params) == 0 {
        return
    }
    if len(params) > 1 {
        for _, section := range params {
            d.ConnectDB(section)
        }
        return
    }

    section := Default_section
    if len(params) == 1 {
        section = params[0]
    }

    db_conf := conf.GetSection("storage \""+section+"\"", "storage")

    db, err := sql.Open("mysql", db_conf.Get("username")+":"+db_conf.Get("password")+"@"+db_conf.Get("protocol", "tcp")+"("+db_conf.Get("host", "localhost")+":"+db_conf.Get("port", "3306")+")/"+db_conf.Get("db")+"?charset="+db_conf.Get("charset", "utf8"))
    if err == nil {
        d.Lock()
        d.d[section] = db
        d.Unlock()
    }
}

func (d *DbList) GetDB(section string) *sql.DB {
    d.RLock()
    db, _ := d.d[section]
    d.RUnlock()
    return db
}
