package main

import (
	"fmt"
)

func equalChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	equal1 := ch1 == ch2
	equal2 := ch1 == ch1

	fmt.Println(equal1)
	fmt.Println(equal2)
}

func main() {
	equalChannels()

	ch3 := make(chan int)
	ch4 := ch3
	ch5 := make(chan int)

	fmt.Println(ch3 == ch4) // true, так как ch3 и ch4 ссылаются на один и тот же канал
	fmt.Println(ch3 == ch5) // false, так как ch3 и ch5 ссылаются на разные каналы

}
