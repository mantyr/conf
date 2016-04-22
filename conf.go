package conf

import (
    "strings"
)

// Examples:
//  GetSection("key")
//  GetSection("key", "file")
func (c *ConfigList) GetSection(key string, params ...string) ConfigSection {
    file_name := c.default_file
    if len(params) > 0 {
        file_name = params[0]
    }

    c.RLock()
    v, ok := c.d[file_name]
    c.RUnlock()

    if !ok {
        config_file := Load(c.get_file_address(file_name))

        c.Lock()
        c.d[file_name] = config_file
        c.Unlock()

        return config_file.d[key]
    }
    return v.d[key]
}

func (c *ConfigList) get_file_address(file_name string) string {
    return c.default_dir+"/"+strings.Replace(file_name, ".", "/", -1)+".ini"
}

// Preload config file
func (c *ConfigList) LoadFile(file_name string) ConfigFile {
    c.RLock()
    if v, ok := c.d[file_name]; ok {
        c.RUnlock()
        return v
    }
    c.RUnlock()

    config_file := Load(c.get_file_address(file_name))
    c.Lock()
    c.d[file_name] = config_file
    c.Unlock()

    return config_file
}

func (c *ConfigList) GetFile(file_name string) ConfigFile {
    return c.LoadFile(file_name)
}

func (c *ConfigList) IsFile(file_name string) bool {
    c.RLock()
    if v, ok := c.d[file_name]; !ok || v.Error != nil {
        return false
    }
    return true
}

// Examples:
//  Get("key")
//  Get("key", "section")
//  Get("key", "section", "file")
func (c *ConfigList) Get(key string, params ...string) string {
    section   := Default_section
    file_name := c.default_file

    if len(params) > 0 {
        section = params[0]
    }
    if len(params) > 1 {
        file_name = params[1]
    }

    c.RLock()
    v, ok := c.d[file_name]
    c.RUnlock()

    if !ok {
        // загружаем конфиг
        config_file := Load(c.get_file_address(file_name))

        c.Lock()
        c.d[file_name] = config_file
        c.Unlock()

        return config_file.d[section].d[key]
    }
    return v.d[section].d[key]
}


func (c *ConfigList) SetDefaultFile(name string) {
    c.default_file = name
}

func (c *ConfigList) SetDefaultCatalog(name string) {
    c.default_dir = name
}

func NewConfig() (c ConfigList) {
    c = ConfigList{}
    c.d = make(map[string]ConfigFile)

    c.SetDefaultFile("properties")
    c.SetDefaultCatalog(DirBin+"/configs")
    return c
}

func NewConfigFile() (c ConfigFile) {
    c = ConfigFile{}
    c.d = make(map[string]ConfigSection)
    return c
}

func NewConfigSection() (c ConfigSection) {
    c = ConfigSection{}
    c.d = make(map[string]string)
    return c
}

