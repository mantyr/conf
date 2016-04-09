package conf

func (c *ConfigSection) Get(key string) string {
    return c.d[key]
}

