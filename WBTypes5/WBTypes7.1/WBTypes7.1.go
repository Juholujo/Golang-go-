package main

import (
	"context"
	"fmt"
)

func main() {
	{
		ctx := context.WithValue(context.Background(), "key", "Kolya1")
		fmt.Println("value1 =", ctx.Value("key").(string))
		ctx = context.WithValue(ctx, "key", "Kolya2")

		fmt.Println("value2 =", ctx.Value("key").(string))
	}
	{
		type str1 string // type definition, not type alias
		type str2 string // type definition, not type alias
		const k1 str1 = "ANDREW"
		const k2 str2 = "ANDREW"

		ctx := context.WithValue(context.Background(), k1, "WB1")
		ctx = context.WithValue(ctx, k2, "WB2")

		fmt.Println("str1 =", ctx.Value(k1).(string))
		fmt.Println("str1 =", ctx.Value(k2).(string))

		//fmt.Println("k1", k1)
		//k1 = "value1"
	}
}
