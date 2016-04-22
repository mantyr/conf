package conf

var config_list ConfigList

func init() {
    config_list = NewConfig()
}

func SetDefaultFile(name string) {
    config_list.SetDefaultFile(name)
}

func SetDefaultCatalog(name string) {
    config_list.SetDefaultCatalog(name)
}


// Examples:
//  GetSection("key")
//  GetSection("key", "file")
func GetSection(key string, params ...string) ConfigSection {
    return config_list.GetSection(key, params...)
}

func LoadFile(file_name string) ConfigFile {
    return config_list.LoadFile(file_name)
}

func GetFile(file_name string) ConfigFile {
    return config_list.GetFile(file_name)
}

func IsFile(file_name string) bool {
    return config_list.IsFile(file_name)
}


// Examples:
//  Get("key")
//  Get("key", "section")
//  Get("key", "section", "file")
func Get(key string, params ...string) string {
    return config_list.Get(key, params...)
}

func GetInt(key string, params ...string) int {
    return config_list.GetInt(key, params...)
}

func GetInt64(key string, params ...string) int64 {
    return config_list.GetInt64(key, params...)
}
