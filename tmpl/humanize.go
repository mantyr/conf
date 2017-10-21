package tmpl

import (
	"github.com/dustin/go-humanize"
	"reflect"
	"strconv"
)

func HumanInt(st interface{}) string {
	Type := reflect.TypeOf(st).String()
	if Type == "string" {
		i, _ := strconv.Atoi(st.(string))
		return humanize.Comma(int64(i))
	} else if Type == "int" {
		return humanize.Comma(int64(st.(int)))
	} else if Type == "int64" {
		return humanize.Comma(st.(int64))
	} else {
		return ""
	}
}
