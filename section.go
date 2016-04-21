package conf

// Examples:
//  Get("key")
//  Get("key", "default_key")
func (c *ConfigSection) Get(params ...string) string {
    if len(params) == 0 {
        return ""
    }

    val, ok := c.d[params[0]]
    if (!ok || val == "") && len(params) > 1 {
        return params[1]
    }
    return val
}

