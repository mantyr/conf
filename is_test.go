package conf

import (
    "testing"
)

func TestIs(t *testing.T) {
    conf := NewConfig()
    conf.SetDefaultCatalog("./testdata")

    ok := conf.Is("title")
    if !ok {
        t.Errorf("Error result Is(), %t", ok)
    }

    ok = conf.Is("title_none")
    if ok {
        t.Errorf("Error result Is(), %t", ok)
    }

    ok = conf.Is("title_none")
    if ok {
        t.Errorf("Error result Is(), %t", ok)
    }
}
