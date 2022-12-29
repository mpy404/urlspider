安装第三方库

```
go get github.com/axgle/mahonia
go install github.com/axgle/mahonia
go get golang.org/x/text/  
go install golang.org/x/text/
go get github.com/PuerkitoBio/goquery 
go install github.com/PuerkitoBio/goquery
```

修改配置

1. 添加Cookie（在config/config.go中添加）

![](images/Snipaste_2022-12-29_16-59-12.jpg)

不然运行多次后会出现百度安全验证

![](images/Snipaste_2022-12-29_17-00-34.jpg)

2. 配置忽略的站点(在config/config.go中添加)

![](images/Snipaste_2022-12-29_17-03-11.jpg)

运行结果

![](images/Snipaste_2022-12-29_16-58-44.jpg)
![](images/Snipaste_2022-12-29_17-21-40.jpg)

每次10个请求，每次延时2秒，请求太快怕触发安全验证，想要去掉延时，在config/config.go中去掉这行

![](images/Snipaste_2022-12-29_17-24-36.jpg)
