package conf

func (c ConfigFile) addSection(name string) {
    if _, ok := c.d[name]; !ok {
        c.d[name] = NewConfigSection()
    }
}

func (c ConfigFile) addOption(section, name, value string) {
    c.addSection(section)
    c.d[section].d[name] = value
}






func (c *ConfigFile) GetSection(key string) ConfigSection {
    return c.d[key]
}

func (c *ConfigFile) Get(key string, params ...string) string {
    section := Default_section
    if len(params) > 0 {
        section = params[0]
    }
    return c.d[section].d[key]
}
