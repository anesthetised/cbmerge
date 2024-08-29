package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/anesthetised/cbmerge/pkg/merger"
)

func run(ctx context.Context, config *Config) error {
	var err error

	if err = config.Validate(); err != nil {
		return fmt.Errorf("validate config: %w", err)
	}

	dst, err := os.OpenFile(config.OutputFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}

	var deleteFile bool

	defer func() {
		_ = dst.Close()

		if deleteFile {
			_ = os.Remove(config.OutputFile)
		}
	}()

	matches, err := filepath.Glob(fmt.Sprintf("%s/*.cbz", config.InputDir))
	if err != nil {
		deleteFile = true
		return fmt.Errorf("list input files with cbz extension: %w", err)
	}

	var inputs []*merger.FileReader
	defer closeSlice(inputs)

	for _, filename := range matches {
		log.Println("opening file", filename)

		f, err := os.Open(filename)
		if err != nil {
			deleteFile = true
			return fmt.Errorf("open input file: %w", err)
		}

		inputs = append(inputs, merger.NewFileReader(f))
	}

	if err = merger.Merge(ctx, dst, convertFileReaderSlice(inputs)); err != nil {
		deleteFile = true
		return fmt.Errorf("merge: %w", err)
	}

	return nil
}

func closeSlice[T io.Closer](closers []T) {
	for _, c := range closers {
		_ = c.Close()
	}
}

func convertFileReaderSlice(frs []*merger.FileReader) []merger.Reader {
	rs := make([]merger.Reader, len(frs))

	for i, fr := range frs {
		rs[i] = fr
	}

	return rs
}

func main() {
	var config Config

	flag.StringVar(&config.InputDir, "src", ".", "input directory")
	flag.StringVar(&config.OutputFile, "dst", "", "output cbz file")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := run(ctx, &config)
	if errors.Is(err, context.Canceled) {
		log.Println("canceled")
		return
	}
	if err != nil {
		log.Println(fmt.Errorf("failed to %w", err))
		return
	}

	log.Printf("done merging %s/*.cbz to %s",
		strings.TrimRight(config.InputDir, string(filepath.Separator)),
		config.OutputFile,
	)
}
