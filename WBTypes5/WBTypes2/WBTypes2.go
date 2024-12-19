package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeoutCause(context.Background(), time.Second, errors.New("timeodffdfdut"))
	defer cancel() // show difference

	<-ctx.Done()

	fmt.Println(ctx.Err())
	fmt.Println(context.Cause(ctx))

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//
	//<-ctx.Done()
	//
	//fmt.Println(ctx.Err())          // Вывод: context deadline exceeded
	//fmt.Println(context.Cause(ctx)) // Вывод: context deadline exceeded
}
