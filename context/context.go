package context

import (
	"context"
	"fmt"
	"time"
)

func Exec() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Reservation cancelled")
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel reserved")
	}
}
