package conf

import (
    "github.com/mantyr/runner"
)

func SetDirBin(dir string) string {
    var err error
    if dir == "" {
        dir, err = runner.GetDirBin()
    }

    if dir == "" || err != nil {
        DirBin = "."
    } else {
        DirBin = dir
    }

    SetDefaultCatalog(DirBin+"/configs")
    return DirBin
}

func GetDirBin() string {
    return DirBin
}