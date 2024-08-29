package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

var (
	ErrEmptyInput  = errors.New("input directory is empty")
	ErrEmptyOutput = errors.New("output file is empty")
)

type Config struct {
	InputDir   string
	OutputFile string
}

func (c Config) Validate() error {
	if c.InputDir == "" {
		return ErrEmptyInput
	}

	if c.OutputFile == "" {
		return ErrEmptyOutput
	}

	if ext := strings.TrimLeft(filepath.Ext(c.OutputFile), "."); ext != "cbz" {
		return fmt.Errorf("unexpected output file extension: %s", ext)
	}

	return nil
}
