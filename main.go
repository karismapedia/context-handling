package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// case #1: normal context
	fmt.Println("case #1")
	ctx := context.Background()

	// give normal context a timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// err after timeout assignment
	fmt.Println("after timeout assignment")
	isErr(ctx)

	// err after 500ms
	fmt.Println("after 500ms")
	time.Sleep(500 * time.Millisecond)
	isErr(ctx)

	// err after timeout passed
	fmt.Println("after timeout passed")
	time.Sleep(time.Second)
	isErr(ctx)

	// case #2: normal context
	fmt.Println("case #2")
	ctx2 := context.Background()

	// create child context with timeout
	childCtx2, cancel2 := context.WithTimeout(ctx2, 1*time.Second)
	defer cancel2()

	// err after timeout assignment
	fmt.Println("after timeout assignment")
	isErr(ctx2)
	isErr(childCtx2)

	// err after 500ms
	fmt.Println("after 500ms")
	time.Sleep(500 * time.Millisecond)
	isErr(ctx2)
	isErr(childCtx2)

	// err after timeout passed
	fmt.Println("after timeout passed")
	time.Sleep(time.Second)
	isErr(ctx2)
	isErr(childCtx2)

	// case #3: normal context
	fmt.Println("case #3")
	ctx3 := context.Background()

	insideFunc(ctx3, false)

	// after func
	isErr(ctx3)

	// case #4: normal context
	fmt.Println("case #4")
	ctx4 := context.Background()

	insideFunc(ctx4, true)

	// after func
	isErr(ctx4)

	// case #5: normal context
	fmt.Println("case #5")
	ctx5, _ := context.WithTimeout(context.Background(), time.Millisecond*100)

	insideFunc(ctx5, true)

	// after func
	isErr(ctx5)
}

func isErr(ctx context.Context) {
	fmt.Println(ctx.Err())
}

func insideFunc(ctx context.Context, withCancel bool) {
	// create child context with timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	if withCancel {
		defer cancel()
	}

	// err after timeout assignment
	fmt.Println("after timeout assignment")
	isErr(ctx)

	// err after 500ms
	fmt.Println("after 500ms")
	time.Sleep(500 * time.Millisecond)
	isErr(ctx)

	// err after timeout passed
	fmt.Println("after timeout passed")
	time.Sleep(time.Second)
	isErr(ctx)
}
