package routines

import (
	"fmt"
	"time"
)

func Hello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
		time.Sleep(time.Millisecond * 100)
	}
}

func World() {
	for i := 0; i < 10; i++ {
		fmt.Println("World")
		time.Sleep(time.Millisecond * 100)
	}
}

func Exec() {
	go Hello()
	go World()
	time.Sleep(time.Second)
}
