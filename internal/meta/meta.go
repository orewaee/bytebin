package meta

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/orewaee/bytebin/pkg/dto"
	"os"
)

// Read reads meta from the dir if it exists and returns a pointer of type *dto.Meta.
func Read(id string) (*dto.Meta, error) {
	path := fmt.Sprintf("./meta/meta-%s.json", id)

	info, err := os.Stat(path)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		text := fmt.Sprintf("meta %s does not exist", id)
		return nil, errors.New(text)
	}

	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		text := fmt.Sprintf("meta %s is not valid", id)
		return nil, errors.New(text)
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var meta = new(dto.Meta)
	if err := json.Unmarshal(bytes, meta); err != nil {
		return nil, err
	}

	return meta, nil
}

// Write writes meta to a file. If it already exists, the file is overwritten.
func Write(meta *dto.Meta) error {
	path := fmt.Sprintf("./meta/meta-%s.json", meta.Id)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	return json.NewEncoder(file).Encode(meta)
}
