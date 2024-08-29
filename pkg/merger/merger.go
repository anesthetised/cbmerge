package merger

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
)

func Merge(ctx context.Context, dst io.Writer, src []Reader) error {
	zw := zip.NewWriter(dst)
	defer func() { _ = zw.Close() }()

	for _, r := range src {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		size, err := r.Size()
		if err != nil {
			return fmt.Errorf("size: %w", err)
		}

		zr, err := zip.NewReader(r, size)
		if err != nil {
			return fmt.Errorf("open zip reader: %w", err)
		}

		for _, zf := range zr.File {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			if err = zw.Copy(zf); err != nil {
				return fmt.Errorf("copy zip file %s: %w", zf.Name, err)
			}
		}

		if err = zw.Flush(); err != nil {
			return fmt.Errorf("flush zip writer: %w", err)
		}
	}

	if err := zw.Close(); err != nil {
		return fmt.Errorf("close zip writer: %w", err)
	}

	return nil
}
