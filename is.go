package conf

// Examples:
//  Is("key")
//  Is("key", "section")
//  Is("key", "section", "file")
func (c *ConfigList) Is(key string, params ...string) bool {
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

        _, ok = config_file.d[section].d[key]
        return ok
    }
    _, ok = v.d[section].d[key]
    return ok
}

// Examples:
//  Is("key")
//  Is("key", "section")
func (c ConfigFile) Is(key string, params ...string) bool {
    section := Default_section
    if len(params) > 0 {
        section = params[0]
    }
    _, ok := c.d[section].d[key]
    return ok
}

func (c ConfigSection) Is(key string) bool {
    _, ok := c.d[key]
    return ok
}