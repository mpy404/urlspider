package scan

import (
	"URLScan/config"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(1)
	log.SetFlags(log.Ldate | log.Llongfile | log.Ltime)
}

func bingWorker(poc string, num chan int, wg *sync.WaitGroup) {
	for i := range num {
		fmt.Println("---------------------------开始第", i/10+1, "页---------------------------")
		fmt.Println()
		url := fmt.Sprintf("https://cn.bing.com/search?q=%s&first=%d", poc, i)
		config.Fetch("bing", url)
		wg.Done()
	}
}

func BingRun(poc string) {
	time1 := time.Now()
	fmt.Println("---------------------------正在使用必应---------------------------")
	fmt.Println()
	var wg sync.WaitGroup
	num := make(chan int, 6)
	for i := 0; i < cap(num); i++ {
		go bingWorker(poc, num, &wg)
	}
	for i := 1; i <= 101; i += 10 {
		wg.Add(1)
		num <- i
	}
	wg.Wait()
	time2 := time.Now()
	fmt.Printf("总共收集%d条数据, 花费%v秒\n", config.Nums, time2.Sub(time1).Seconds())
}
