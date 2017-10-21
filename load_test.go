package conf

import (
	"testing"
)

func TestLoadFileDefault(t *testing.T) {
	conf := NewConfig()
	conf.SetDefaultCatalog("./testdata")

	val := conf.Get("title")
	if val != "test conf" {
		t.Errorf("Error default section, default file, %q", val)
	}

	val = conf.Get("date")
	if val != "2016-04-02 22:03+03:00" {
		t.Errorf("Error ignore comment, %q", val)
	}

	val = conf.Get("name", "section api1")
	if val != "test_ok" {
		t.Errorf("Error revalue section, %q", val)
	}
	val = conf.Get("pass", "server_2", "storage")
	if val != "password2" {
		t.Errorf("Error ignore comment, %q", val)
	}
	val = conf.Get("comment", "server_2", "storage")
	if val != "comment / comment" {
		t.Errorf("Error ignore comment, %q", val)
	}
	val = conf.Get("comment2", "server_2", "storage", "default_value")
	if val != "comment	/ comment" {
		t.Errorf("Error ignore comment, %q", val)
	}
	val = conf.Get("default_value", "server_2", "storage", "default_ok")
	if val != "default_ok" {
		t.Errorf("Error default_value, %q", val)
	}
}

func TestInitLoadFile(t *testing.T) {
	SetDefaultFile("properties")
	//    SetDefaultCatalog("./testdata") // Set default dir ./configs/
	LoadFile("storage")

	section := GetSection("server_1", "storage")
	val := section.Get("port")
	if val != "1234_v2" {
		t.Errorf("Error first section.Get() value, %q", val)
	}
}

func TestLoadFile(t *testing.T) {
	conf := NewConfig()
	conf.SetDefaultFile("properties")
	conf.SetDefaultCatalog("./testdata") // default "./configs"

	section := conf.GetSection("server_1", "storage")
	val := section.Get("port")
	if val != "1234" {
		t.Errorf("Error first section.Get() value, %q", val)
	}

	section2 := conf.GetSection("server_1", "storage")
	val = section2.Get("port")
	if val != "1234" {
		t.Errorf("Error second section.Get() value, %q", val)
	}

	file := conf.GetFile("storage")

	/**************************************************************************/

	val = section.Get("host")
	if val != "example.com" {
		t.Errorf("Error section.Get() value, %q", val)
	}

	section_none := conf.GetSection("server_none", "storage")
	val = section_none.Get("host")
	if val != "" {
		t.Errorf("Error section.Get() value for nil section, %q", val)
	}

	section_none = conf.GetSection("server_none", "storage_none")
	val = section_none.Get("host")
	if val != "" {
		t.Errorf("Error section.Get() value for nil file, %q", val)
	}

	val = conf.Get("host", "server_1", "storage")
	if val != "example.com" {
		t.Errorf("Error ConfigList.Get(), %q", val)
	}

	val = conf.Get("host_none", "server_1", "storage")
	if val != "" {
		t.Errorf("Error ConfigList.Get(), %q", val)
	}

	val = conf.Get("host", "server_none", "storage")
	if val != "" {
		t.Errorf("Error ConfigList.Get(), %q", val)
	}

	val = conf.Get("host", "server", "storage_none")
	if val != "" {
		t.Errorf("Error ConfigList.Get(), %q", val)
	}

	/**************************************************************************/

	val_int := conf.GetInt("port", "server_1", "storage")
	if val_int != 1234 {
		t.Errorf("Error convert, int, %q", val_int)
	}

	val_int = section.GetInt("port")
	if val_int != 1234 {
		t.Errorf("Error convert, int, %q", val_int)
	}

	val_int = file.GetInt("port", "server_1")
	if val_int != 1234 {
		t.Errorf("Error convert, int, %q", val_int)
	}

	/******************/

	val_int64 := conf.GetInt64("port", "server_1", "storage")
	if val_int64 != 1234 {
		t.Errorf("Error convert, int64, %q", val_int64)
	}

	val_int64 = section.GetInt64("port")
	if val_int64 != 1234 {
		t.Errorf("Error convert, int64, %q", val_int64)
	}

	val_int64 = file.GetInt64("port", "server_1")
	if val_int64 != 1234 {
		t.Errorf("Error convert, int64, %q", val_int)
	}
}

func TestIdValue(t *testing.T) {
	SetDefaultFile("properties")
	SetDefaultCatalog("./testdata")
	LoadFile("storage")

	section := GetSection("section_id")
	val := section.Get("title")
	if val != "#id_123 .title" {
		t.Errorf("Error #id section.Get() value, %q", val)
	}
	val = section.Get("title2")
	if val != "#id_234 .title" {
		t.Errorf("Error #id section.Get() value, %q", val)
	}
}

func TestUrlParamsValue(t *testing.T) {
	SetDefaultFile("properties")
	SetDefaultCatalog("./testdata")
	LoadFile("properties")

	section := GetSection("section_id")
	val := section.Get("address1")
	if val != "http://example.com/page/details/offers.ashx?productId=5159" {
		t.Errorf("Error ?attr=value section.Get() value, %q", val)
	}
	val = section.Get("address2")
	if val != "http://example.com/page/details/offers.ashx?productId=5159#page=1" {
		t.Errorf("Error ?attr=value section.Get() value, %q", val)
	}
	val = section.Get("catalog/all/all/all/all/all/?np=200")
	if val != "true" {
		t.Errorf("Error '\\=' in key section.Get() value, %q", val)
	}
	val = section.Get("catalog/all/all/all/all/all/?np=200&test=123")
	if val != "true2" {
		t.Errorf("Error '\\=' in key section.Get() value, %q", val)
	}
}

func TestBigRat(t *testing.T) {
	SetDefaultFile("properties")
	SetDefaultCatalog("./testdata")
	LoadFile("money")

	section := GetSection("money", "money")
	val := section.GetBigRat("money1")
	if val.FloatString(2) != "12.00" {
		t.Errorf("Error Get big.Rat, %q", val.FloatString(2))
	}

	val = section.GetBigRat("money2")
	if val.FloatString(2) != "0.00" {
		t.Errorf("Error Get big.Rat, %q", val.FloatString(2))
	}

	val = section.GetBigRat("money3")
	if val.FloatString(2) != "0.00" {
		t.Errorf("Error Get big.Rat, %q", val.FloatString(2))
	}
}

func BenchmarkLoad(b *testing.B) {
	conf := NewConfig()
	conf.SetDefaultFile("properties")
	conf.SetDefaultCatalog("./testdata") // default "./configs"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		conf.GetSection("server_1", "storage")
	}
}

func BenchmarkLoadPreload(b *testing.B) {
	conf := NewConfig()
	conf.SetDefaultFile("properties")
	conf.SetDefaultCatalog("./testdata") // default "./configs"

	conf.GetSection("server_1", "storage")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		conf.GetSection("server_1", "storage")
	}
}

func BenchmarkLoadGo(b *testing.B) {
	conf := NewConfig()
	conf.SetDefaultFile("properties")
	conf.SetDefaultCatalog("./testdata") // default "./configs"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		go conf.GetSection("server_1", "storage")
	}
}

func BenchmarkLoadGoPreload(b *testing.B) {
	conf := NewConfig()
	conf.SetDefaultFile("properties")
	conf.SetDefaultCatalog("./testdata") // default "./configs"

	conf.GetSection("server_1", "storage")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		go conf.GetSection("server_1", "storage")
	}
}
