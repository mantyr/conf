package conf

import (
    "runtime"
)

func MinCPU(min_cpu int) {
    cpu := runtime.NumCPU()
    if min_cpu > cpu {
        cpu = min_cpu
    }
    runtime.GOMAXPROCS(cpu)
}