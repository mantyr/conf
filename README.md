# Golang conf (ini file)

[![Build Status](https://travis-ci.org/mantyr/conf.svg?branch=master)](https://travis-ci.org/mantyr/conf) [![GoDoc](https://godoc.org/github.com/mantyr/conf?status.png)](http://godoc.org/github.com/mantyr/conf) [![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

This stable version

## Installation

    $ go get github.com/mantyr/conf

## Example

```GO
package main

import (
    "github.com/mantyr/conf"
)

func init() {
    conf.SetDefaultCatalog("./testdata")
    conf.SetDefaultFile("properties")

    // preload files, not necessary
    conf.LoadFile("properties")
    conf.LoadFile("storage")
}

func main() {
    val := conf.Get("key")
    val  = conf.Get("key", "default")
    val  = conf.Get("key", "default", "properties")

    val_int := conf.GetInt64("key")
    val_int  = conf.GetInt64("key", "default")
    val_int  = conf.GetInt64("key", "default", "properties")

    section := conf.GetSection("section_name")
    section  = conf.GetSection("section_name", "properties")

    val       = section.Get("key")
    val_int   = section.GetInt("key")
    val_int64 = section.GetInt64("key")

    ini_file := conf.GetFile("properties")

    ...
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
