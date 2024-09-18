package pkg

import (
	"fmt"
	"sync"
	"time"
)
func ThreadTest1(){
	// channel用于线程通信
	c:=make(chan string)
	go func ()  {
		c<-"!@3123"
	}()
	
	for i:=0 ;i< 5;i++{
		fmt.Printf("123+%d\n", i)
	}
	fmt.Println(<-c)

}
func ThreadTest() {
	var currentPrint int
	var count = 3
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		for count > 0 {
			mu.Lock()
			if currentPrint == 0 {
				fmt.Println("A")
				currentPrint++
			}
			mu.Unlock()
			time.Sleep(100 * time.Millisecond) // 避免空转
		}
	}()

	go func() {
		defer wg.Done()
		for count > 0 {
			mu.Lock()
			if currentPrint == 1 {
				fmt.Println("B")
				currentPrint++
			}
			mu.Unlock()
			time.Sleep(100 * time.Millisecond) // 避免空转
		}
	}()

	go func() {
		defer wg.Done()
		for count > 0 {
			mu.Lock()
			if currentPrint == 2 {
				fmt.Println("C")
				currentPrint = 0
				count--
			}
			mu.Unlock()
			time.Sleep(100 * time.Millisecond) // 避免空转
		}
	}()

	wg.Wait() // 等待所有 Goroutine 完成
}
