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

func baiduWorker(poc string, num chan int, wg *sync.WaitGroup) {
	for i := range num {
		fmt.Println("---------------------------第", i/10+1, "页---------------------------")
		fmt.Println()
		url := fmt.Sprintf("http://www.baidu.com/s?wd=%s&pn=%d", poc, i)
		config.Fetch("baidu", url)
		wg.Done()
	}
}

func BaiduRun(poc string) {
	fmt.Println("---------------------------正在使用百度扫描---------------------------")
	fmt.Println()
	time1 := time.Now()
	var wg sync.WaitGroup
	num := make(chan int, 10)
	for i := 0; i < cap(num); i++ {
		go baiduWorker(poc, num, &wg)
	}
	for i := 0; i <= 750; i += 10 {
		wg.Add(1)
		num <- i
	}
	wg.Wait()
	time2 := time.Now()
	fmt.Printf("总共收集%d条数据, 花费%v秒\n", config.Nums, time2.Sub(time1).Seconds())
}
