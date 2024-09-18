package pkg

import (
	"fmt"
	"sync"
	"time"
)
func ThreadTest1(){
	// channel用于线程通信
	c:=make(chan int,3)
	go func ()  {
		for i:=0;i<3;i++{
			c<-i;
			fmt.Printf("sub goroutine sen:%d, channel's len: %d,cap:%d\n", i,len(c),cap(c))
		}
		close(c)
		time.Sleep(3*time.Second)
		// channel会使线程阻塞
		c<-566
	}()
	
	for i:=0;i<4;i++{
		fmt.Printf("%d\n", <-c)
	}

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
