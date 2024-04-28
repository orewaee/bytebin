package storage

import (
	"github.com/orewaee/bytebin/internal/bin"
)

var bins = make(map[string]*bin.Bin)

func AddBin(id string, bin *bin.Bin) {
	bins[id] = bin
}

func GetBin(id string) (*bin.Bin, bool) {
	b, ok := bins[id]
	return b, ok
}
