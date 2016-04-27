package conf

// Examples:
//  Get("key")
//  Get("key", "default_key")
func (c ConfigSection) Get(params ...string) string {
    if len(params) == 0 {
        return ""
    }

    val, ok := c.d[params[0]]
    if !ok || val == "" {
        if len(params) > 1 {
            return params[1]
        }
        return ""
    }
    return val
}

// No safe. Do not change the values if you do not believe that non-concurrent access
func (c ConfigSection) GetItems() map[string]string {
    return c.d
}