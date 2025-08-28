package cleaner

import (
	"context"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
	"os"
	"path/filepath"
	"time"
)

type Options struct {
	Days       int
	NoProgress bool
}

type Cleaner struct {
	options *Options
}

func New(options *Options) *Cleaner {
	return &Cleaner{
		options: options,
	}
}

func (c *Cleaner) countFiles(path string) int {
	count := 0
	if err := filepath.WalkDir(path, func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !dir.IsDir() {
			count++
		}

		return nil
	}); err != nil {
		return 0
	}

	return count
}

func (c *Cleaner) Clean(ctx context.Context, path string, progress *mpb.Progress) (int, error) {
	var total = 0
	var bar *mpb.Bar
	if !c.options.NoProgress {
		total = c.countFiles(path)
		bar = progress.AddBar(int64(total),
			mpb.PrependDecorators(
				decor.Name(path, decor.WCSyncSpaceR),
				decor.CountersNoUnit("%d / %d", decor.WCSyncWidth),
			),
			mpb.AppendDecorators(
				decor.Percentage(),
			),
		)
	}

	duration := time.Hour * 24 * time.Duration(c.options.Days)
	deleted := 0

	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if time.Since(info.ModTime()) < duration {
			return nil
		}

		if err := os.Remove(path); err != nil {
			return err
		}

		deleted++
		if bar != nil {
			bar.Increment()
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		return nil
	}); err != nil {
		if bar != nil {
			bar.SetTotal(int64(total), true)
		}
		return 0, err
	}

	if bar != nil {
		bar.SetTotal(int64(total), true)
	}

	return deleted, nil
}
