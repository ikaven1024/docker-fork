package main

import (
	"fmt"
)

func HumanSize(size uint64) string {
	if size == 0 {
		return ""
	}
	units := []string{"", "kb", "mb", "gb"}
	i := 0
	var base uint64 = 1024

	for ; i < len(units)-1 && size >= base && size%base == 0; i++ {
		size /= base
	}
	return fmt.Sprintf("%d%s", size, units[i])
}
