package errorGroup

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func ErrorGroupNotBreak() {
	var eg errgroup.Group
	for i := 0; i < 10; i++ {
		i := i
		eg.Go(func() error {
			time.Sleep(2 * time.Second)
			if i > 6 {
				return fmt.Errorf("Error occurred: %d", i)
			}
			fmt.Println("End:", i)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}

func ErrorGroupBreak() {
	eg, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 100; i++ {
		i := i
		eg.Go(func() error {
			time.Sleep(2 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("Canceled:", i)
				return nil
			default:
				if i > 90 {
					fmt.Println("Error:", i)
					return fmt.Errorf("Error: %d", i)
				}
				fmt.Println("End:", i)
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
