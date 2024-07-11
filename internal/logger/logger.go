package logger

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func init() {
	if err := createDir("logs"); err != nil {
		panic("cannot create logs directory")
	}
}

func createDir(name string) error {
	if err := os.Mkdir(name, os.ModePerm); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	return nil
}

func New(dir string) (*zerolog.Logger, error) {
	if err := createDir("logs/" + dir); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(
		fmt.Sprintf("logs/%s.log", time.Now().Format("02012006-150405MST")),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)

	if err != nil {
		return nil, err
	}

	writer := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}, file,
	)

	logger := zerolog.New(writer).With().Timestamp().Logger()

	return &logger, nil
}
