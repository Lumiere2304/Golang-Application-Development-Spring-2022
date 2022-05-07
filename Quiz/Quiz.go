package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sayHello(name string) {
	defer wg.Done() // decrease WaitGroup counter by one as function is finished
	fmt.Println("Hello", name)
	time.Sleep(time.Second)
	fmt.Println("See you next time", name)
}

func main() {
	wg.Add(3) // Устанавливает счетчик (количество активных горутин)
	go sayHello("Temirbolat")
	go sayHello("Anuar")
	go sayHello("Andrew")
	wg.Wait() // блокирует дальнейшую работу кода пока счеткик не станет 0.
	fmt.Println("The programm is out")
}