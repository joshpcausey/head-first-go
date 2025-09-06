package main

import (
	"fmt"
	"head-first-go/hello"

	"github.com/enescakir/emoji"
)

func main() {
	hello.Hello()

	fmt.Println(emoji.WavingHand.Tone())
}
