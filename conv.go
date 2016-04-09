package conf

import (
    "strconv"
)

func (c *ConfigFile) GetInt(key string, params ...string) int {
    val := c.Get(key, params...)

    i, err := strconv.Atoi(val)
    if err != nil {
        return 0
    }
    return i
}

func (c *ConfigFile) GetInt64(key string, params ...string) int64 {
    val := c.Get(key, params...)

    i, err := strconv.ParseInt(val, 10, 64)
    if err != nil {
        return 0
    }
    return i
}

func (c *ConfigList) GetInt(key string, params ...string) int {
    val := c.Get(key, params...)

    i, err := strconv.Atoi(val)
    if err != nil {
        return 0
    }
    return i
}

func (c *ConfigList) GetInt64(key string, params ...string) int64 {
    val := c.Get(key, params...)

    i, err := strconv.ParseInt(val, 10, 64)
    if err != nil {
        return 0
    }
    return i
}

func (c *ConfigSection) GetInt(key string) int {
    val := c.Get(key)

    i, err := strconv.Atoi(val)
    if err != nil {
        return 0
    }
    return i
}

func (c *ConfigSection) GetInt64(key string) int64 {
    val := c.Get(key)

    i, err := strconv.ParseInt(val, 10, 64)
    if err != nil {
        return 0
    }
    return i
}

