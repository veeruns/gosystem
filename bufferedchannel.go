package gosystem

import (
	"fmt"
)

func Bufferdchannel() {
	buffch := make(chan int, 2)
	buffch <- 3
	buffch <- 2
	buffch <- 4
	fmt.Println(<-buffch)
	fmt.Println(<-buffch)
	fmt.Println(<-buffch)

}
