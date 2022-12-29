package config

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"
)

var UserAgent = []string{
	"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.0) Gecko/20060728 Firefox/1.5.0 Opera 9.22 ",
	"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.0) Gecko/20060728 Firefox/1.5.0 Opera 9.24 ",
	"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.0) Gecko/20060728 Firefox/1.5.0 Opera 9.26 ",
	"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.1) Gecko/20061208 Firefox/2.0.0 Opera 9.51 ",
	"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.1) Gecko/20061208 Firefox/5.0 Opera 11.11 ",
	"Mozilla/5.0 (Windows NT 5.1; U; es-la; rv:1.8.0) Gecko/20060728 Firefox/1.5.0 Opera 9.27 ",
	"Mozilla/5.0 (Windows NT 5.1; U; fr) Opera 8.51 ",
	"Mozilla/5.0 (Windows NT 5.1; U; pl) Opera 8.54 ",
	"Mozilla/5.0 (Windows NT 5.1; U; pl; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6 Opera 11.00 ",
	"Mozilla/5.0 (Windows NT 5.1; U; ru) Opera 8.51 ",
	"Mozilla/5.0 (Windows NT 5.1; U; zh-cn; rv:1.8.1) Gecko/20061208 Firefox/2.0.0 Opera 9.50 ",
	"Opera/9.12 (Windows NT 5.0; U) ",
	"Opera/9.12 (Windows NT 5.0; U; ru) ",
	"Opera/9.12 (X11; Linux i686; U; en) (Ubuntu) ",
	"Opera/9.20 (Windows NT 5.1; U; en) ",
	"Opera/9.20(Windows NT 5.1; U; en) ",
	"Opera/9.20 (Windows NT 5.1; U; es-AR) ",
	"Opera/9.20 (Windows NT 5.1; U; es-es) ",
	"Opera/9.20 (Windows NT 5.1; U; it) ",
	"Opera/9.20 (Windows NT 5.1; U; nb) ",
	"Opera/9.20 (Windows NT 5.1; U; zh-tw) ",
	"Opera/9.20 (Windows NT 5.2; U; en) ",
	"Opera/9.20 (Windows NT 6.0; U; de) ",
	"Opera/9.20 (Windows NT 6.0; U; en) ",
	"Opera/9.20 (Windows NT 6.0; U; es-es) ",
	"Opera/9.20 (X11; Linux i586; U; en) ",
	"Opera/9.20 (X11; Linux i686; U; en) ",
	"Opera/9.23 (Windows NT 5.1; U; ja) ",
	"Opera/9.23 (Windows NT 5.1; U; pt) ",
	"Opera/9.23 (Windows NT 5.1; U; zh-cn) ",
	"Opera/9.23 (Windows NT 6.0; U; de) ",
	"Opera/9.23 (X11; Linux i686; U; en) ",
	"Opera/9.27 (Windows NT 5.1; U; ja) ",
	"Opera/9.27 (Windows NT 5.2; U; en) ",
	"Opera/9.27 (X11; Linux i686; U; en) ",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.0.13) Gecko/2009073022 Firefox/3.0.13 (.NET CLR 3.5.30729) ",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.2.28) Gecko/20120306 Firefox/3.6.28 (.NET CLR 3.5.30729; .NET4.0C) ",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; pl; rv:1.9.2.2) Gecko/20100316 Firefox/3.6.2 GTB6 (.NET CLR 3.5.30729) ",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; pt-BR; rv:1.8.0.2) Gecko/20060308 Firefox/1.5.0.2 ",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; pt-BR; rv:1.8.0.9) Gecko/20061206 Firefox/1.5.0.9 ",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; pt-BR; rv:1.8.1.15) Gecko/20080623 Firefox/2.0.0.15 ",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; pt-BR; rv:1.9.0.13) Gecko/2009073022 Firefox/3.0.13 ",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36 ",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:84.0) Gecko/20100101 Firefox/84.0 ",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15 ",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.2 Safari/605.1.15 ",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36 ",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36 ",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0 ",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36 ",
	"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko ",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36 ",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36 ",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36 ",
}                                                // 设置UA
var Nums = 0                                     // 统计获取的URL
var DisCardUrl = []string{"gov", "edu", "index"} // 配置过滤站点

// discardURL 过滤站点
func discardURL(url string) bool {

	for i := 0; i < len(DisCardUrl); i++ {
		re := regexp.MustCompile(DisCardUrl[i])
		findString := re.FindString(url)
		if len(findString) != 0 {
			return false
		}
	}
	return true
}

// setHeaders 设置请求头
func setHeaders(method string, request *http.Request) *http.Request {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(57)
	if method == "baidu" {
		request.Header.Add("User-Agent", UserAgent[n])
		request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Header.Add("Host", "www.baidu.com")
		return request
	} else if method == "bing" {
		request.Header.Add("User-Agent", UserAgent[n])
		request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Header.Add("Host", "cn.bing.com")
		request.Header.Add("Referer", "https://cn.bing.com/search?q=inurl%3aphp%3fid%3d1+%e5%85%ac%e5%8f%b8&first=10&FORM=PERE")
		return request
	}
	return request
}

// Fetch 处理数据包
func Fetch(method, url string) {
	// 打开文件
	file, err := os.OpenFile("url.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	writer := bufio.NewWriter(file)
	if err != nil {
		log.Panic("File Error : ", err)
	}
	defer file.Close()
	// 创建请求
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print("Http Error : ", err)
	}

	request = setHeaders(method, request)

	time.Sleep(time.Second * 2)
	resp, err := client.Do(request)
	if err != nil {
		log.Print("Content Error : ", err)
	}

	if resp.StatusCode != http.StatusOK {
		// 如果异常保存源码
		fmt.Println("StatusCode Error : ", resp.StatusCode)
		all, _ := io.ReadAll(resp.Body)
		_ = os.WriteFile("error.html", all, os.ModePerm)
		log.Panic(err)
	}

	defer resp.Body.Close()

	//匹配百度获取的数据
	if method == "baidu" {
		reader, _ := goquery.NewDocumentFromReader(resp.Body)
		reader.Find(".result").Each(func(i int, selection *goquery.Selection) {
			after, _ := selection.Attr("mu")
			gov := discardURL(after)
			if after != "null" && gov {
				_, _ = writer.WriteString(after + "\n")
				_ = writer.Flush()
				Nums++
			}
		})
	} else if method == "bing" {
		reader, _ := goquery.NewDocumentFromReader(resp.Body)
		reader.Find("a.sh_favicon").Each(func(i int, selection *goquery.Selection) {
			after, _ := selection.Attr("href")
			gov := discardURL(after)
			if after != "null" && gov {
				_, _ = writer.WriteString(after + "\n")
				_ = writer.Flush()
				Nums++
			}
		})
	}
}
