package ShowMeBug

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	DOG  = "DOG1"
	FISH = "FISH2"
	CAT  = "CAT3"
)

func Do() {
	var dogcount uint64
	var fishcount uint64
	var catcount uint64

	dogch := make(chan string, 1) // 有缓冲的管道，管道满时主携程阻塞
	fishch := make(chan string, 1)
	catch := make(chan string, 1)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go dog(wg, dogcount, dogch, fishch)
	go fish(wg, fishcount, fishch, catch)
	go cat(wg, catcount, catch, dogch)

	dogch <- DOG
	wg.Wait()
	fmt.Println("Done!")
}

func dog(wg *sync.WaitGroup, count uint64, dogch, fishch chan string) {
	for {
		if count >= uint64(100) {
			wg.Done()
			return
		}
		str := <-dogch
		fmt.Println(str)
		atomic.AddUint64(&count, 1)
		fishch <- FISH
	}
}

func fish(wg *sync.WaitGroup, count uint64, fishch, catch chan string) {
	for {
		if count >= uint64(100) {
			wg.Done()
			return
		}
		str := <-fishch
		fmt.Println(str)
		atomic.AddUint64(&count, 1)
		catch <- CAT
	}
}

func cat(wg *sync.WaitGroup, count uint64, catch, dogch chan string) {
	for {
		if count >= uint64(100) {
			wg.Done()
			return
		}
		str := <-catch
		fmt.Println(str)
		atomic.AddUint64(&count, 1)
		dogch <- DOG
	}
}
