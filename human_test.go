package main

import (
	"testing"
)

func TestHumanSize(t *testing.T) {
	type args struct {
		size int64
	}
	tests := []struct {
		name string
		args uint64
		want string
	}{
		{"0", 0, "0"},
		{"123", 123, "123"},
		{"1kb", 1024, "1kb"},
		{"1025", 1025, "1025"},
		{"2mb", 1024 * 1024 * 2, "2mb"},
		{"2gb", 1024 * 1024 * 1024 * 2, "2gb"},
		{"2048gb", 1024 * 1024 * 1024 * 1024 * 2, "2048gb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HumanSize(tt.args); got != tt.want {
				t.Errorf("HumanSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
