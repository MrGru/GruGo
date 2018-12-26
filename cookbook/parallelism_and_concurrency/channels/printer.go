package channels

import (
	"context"
	"fmt"
	"time"
)

func Printer(ctx context.Context, ch chan string) {
	t := time.Tick(200 * time.Microsecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("printer done.")
		case res := <-ch:
			fmt.Println(res)
		case <-t:
			fmt.Println("tock")
		}
	}
}
