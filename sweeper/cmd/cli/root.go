package cli

import (
	"context"
	"fmt"
	"github.com/lcarneli/go-playground/sweeper/pkg/cleaner"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	days       int
	noProgress bool
	paths      []string
)

var (
	ErrNoPathsProvided = fmt.Errorf("no paths provided")
)

type result struct {
	path    string
	deleted int
	err     error
}

var rootCommand = cobra.Command{
	Use:   "sweeper",
	Short: "A sweeper to clean up old files.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		defer stop()

		if len(paths) == 0 {
			return ErrNoPathsProvided
		}

		opts := cleaner.Options{
			Days:       days,
			NoProgress: noProgress,
		}
		clr := cleaner.New(&opts)

		var wg sync.WaitGroup
		progress := mpb.New(mpb.WithContext(ctx), mpb.WithWaitGroup(&wg))

		results := make(chan result, len(paths))

		start := time.Now()
		for _, path := range paths {
			wg.Add(1)
			go func() {
				defer wg.Done()

				fmt.Printf("Starting cleanup in %q (files older than %d days)\n\n", path, days)

				deleted, err := clr.Clean(ctx, path, progress)
				results <- result{
					path:    path,
					deleted: deleted,
					err:     err,
				}
			}()
		}

		progress.Wait()
		close(results)

		elapsed := time.Since(start)
		if !opts.NoProgress {
			fmt.Println()
		}
		fmt.Println("================ Summary ================")
		for res := range results {
			if res.err != nil {
				fmt.Printf("%s — Error: %v\n", res.path, res.err)
			} else {
				fmt.Printf("%s — Success: %d files deleted\n", res.path, res.deleted)
			}
		}
		fmt.Println("=========================================")
		fmt.Printf("\n✅  All cleanups completed in %s.\n", elapsed.Round(time.Millisecond).String())

		return nil
	},
}

func init() {
	rootCommand.Flags().IntVarP(&days, "days", "d", 30, "Number of days to keep files.")
	rootCommand.Flags().BoolVar(&noProgress, "no-progress", false, "Disable the progress bar display.")
	rootCommand.Flags().StringSliceVarP(&paths, "paths", "p", []string{}, "Paths to clean.")
}

func Execute() error {
	return rootCommand.Execute()
}
