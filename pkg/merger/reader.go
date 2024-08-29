package merger

import (
	"fmt"
	"io"
	"os"
)

type Reader interface {
	io.ReaderAt
	Size() (int64, error)
}

var _ Reader = (*FileReader)(nil)

type FileReader struct {
	file *os.File
}

func NewFileReader(f *os.File) *FileReader {
	if f == nil {
		panic("merger: file reader: *os.File is nil")
	}
	return &FileReader{file: f}
}

func (r *FileReader) ReadAt(b []byte, off int64) (n int, err error) {
	return r.file.ReadAt(b, off)
}

func (r *FileReader) Size() (int64, error) {
	fi, err := r.file.Stat()
	if err != nil {
		return 0, fmt.Errorf("stat: %w", err)
	}
	return fi.Size(), nil
}

func (r *FileReader) Close() error {
	return r.file.Close()
}
