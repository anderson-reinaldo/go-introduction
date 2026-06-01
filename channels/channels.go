package channels

import (
	"fmt"
	"time"
)

// Thread 1
func Exemplo1() {
	//Chanel entre as duas threads, faz com que elas se comuniquem T1 <-> T2
	hello := make(chan string)

	//Thread 2
	go func() {
		//Envia a mensagem para o canal
		hello <- "Hello World"
	}()

	//Exibe a mensagem
	msg := <-hello
	println(msg)

}

func Exemplo2() {

	forever := make(chan string)

	go func() {
		x := true
		for {
			if x == true {
				continue
			}
		}
	}()

	fmt.Println("Aguardando para sempre...")
	<-forever

}

func Exemplo4() {
	hello := make(chan string)

	go func() {
		hello <- "Hello World"
	}()

	select {
	case msg := <-hello:
		fmt.Println(msg)

	default:
		fmt.Println("No message received")
	}

}

func Exemplo5() {
	queue := make(chan int)

	go func() {
		i := 0
		for {
			queue <- i
			i++
		}
	}()

	for msg := range queue {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func worker(workerID int, msg chan int) {
	for i := range msg {
		fmt.Printf("Worker %d received %d\n", workerID, i)
		time.Sleep(time.Second)
	}
}

func Exemplo6() {
	msg := make(chan int)

	go worker(1, msg)
	go worker(2, msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}
}
