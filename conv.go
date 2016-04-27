package conf

import (
    "strconv"
)

// Examples:
//  GetInt("key")
//  GetInt("key", "section")
func (c *ConfigFile) GetInt(key string, params ...string) int {
    val := c.Get(key, params...)

    i, err := strconv.Atoi(val)
    if err != nil {
        return 0
    }
    return i
}

// Examples:
//  GetInt64("key")
//  GetInt64("key", "section")
func (c *ConfigFile) GetInt64(key string, params ...string) int64 {
    val := c.Get(key, params...)

    i, err := strconv.ParseInt(val, 10, 64)
    if err != nil {
        return 0
    }
    return i
}

// Examples:
//  GetInt("key")
//  GetInt("key", "section")
//  GetInt("key", "section", "file")
func (c *ConfigList) GetInt(key string, params ...string) int {
    val := c.Get(key, params...)

    i, err := strconv.Atoi(val)
    if err != nil {
        return 0
    }
    return i
}

// Examples:
//  GetInt64("key")
//  GetInt64("key", "section")
//  GetInt64("key", "section", "file")
func (c *ConfigList) GetInt64(key string, params ...string) int64 {
    val := c.Get(key, params...)

    i, err := strconv.ParseInt(val, 10, 64)
    if err != nil {
        return 0
    }
    return i
}

// Examples:
//  GetInt("key")
//  GetInt("key", "default_key")
func (c ConfigSection) GetInt(params ...string) int {
    val := c.Get(params...)

    i, err := strconv.Atoi(val)
    if err != nil {
        return 0
    }
    return i
}

// Examples:
//  GetInt64("key")
//  GetInt64("key", "default_key")
func (c ConfigSection) GetInt64(params ...string) int64 {
    val := c.Get(params...)

    i, err := strconv.ParseInt(val, 10, 64)
    if err != nil {
        return 0
    }
    return i
}

