package conf

import (
    "strconv"
    "math/big"
    "time"
)

// Examples:
//  GetInt("key")
//  GetInt("key", "section")
func (c ConfigFile) GetInt(key string, params ...string) int {
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
func (c ConfigFile) GetInt64(key string, params ...string) int64 {
    val := c.Get(key, params...)

    i, err := strconv.ParseInt(val, 10, 64)
    if err != nil {
        return 0
    }
    return i
}

// Examples:
//  GetTimeDuration("key")
//  GetTimeDuration("key", "section")
func (c ConfigFile) GetTimeDuration(key string, params ...string) time.Duration {
    return time.Duration(c.GetInt64(key, params...))
}

// Examples:
//  GetBigRat("key")
//  GetBigRat("key", "section")
func (c *ConfigFile) GetBigRat(key string, params ...string) (r *big.Rat) {
    val := c.Get(key, params...)

    r = new(big.Rat)
    r.SetString(val)
    return
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
//  GetTimeDuration("key")
//  GetTimeDuration("key", "section")
//  GetTimeDuration("key", "section", "file")
func (c *ConfigList) GetTimeDuration(key string, params ...string) time.Duration {
    return time.Duration(c.GetInt64(key, params...))
}

// Examples:
//  GetBigRat("key")
//  GetBigRat("key", "section")
//  GetBigRat("key", "section", "file")
func (c *ConfigList) GetBigRat(key string, params ...string) (r *big.Rat) {
    val := c.Get(key, params...)

    r = new(big.Rat)
    r.SetString(val)
    return
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

// Examples:
//  GetTimeDuration("key")
//  GetTimeDuration("key", "default_key")
func (c ConfigSection) GetTimeDuration(params ...string) time.Duration {
    return time.Duration(c.GetInt64(params...))
}

// Examples:
//  GetBigRat("key")
//  GetBigRat("key", "default_key")
func (c *ConfigSection) GetBigRat(params ...string) (r *big.Rat) {
    val := c.Get(params...)

    r = new(big.Rat)
    r.SetString(val)
    return
}
