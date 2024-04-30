package storage

import (
	"github.com/orewaee/bytebin/internal/bin"
)

var bins = make(map[string]*bin.Bin)

func AddBin(id string, bin *bin.Bin) {
	bins[id] = bin
}

func RemoveBin(id string) {
	delete(bins, id)
}

func GetBin(id string) (*bin.Bin, bool) {
	b, ok := bins[id]
	if !ok {
		return nil, false
	}

	return b, true
}
