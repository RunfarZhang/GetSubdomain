package main

// Written by H0XM4N Runfar Zhang

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

var (
	searchurl string
	domain    string
	save      string
	suffix    string

	proxys []string = []string{"http://61.216.185.88:60808",
		"http://117.114.149.66:55443",
		"http://183.236.123.242:8080",
		"http://112.74.17.146:8118",
		"http://47.92.158.181:8118",
		"http://183.236.123.242:8080",
		"http://106.15.197.250:8001",
		"http://106.54.128.253:999",
		"http://103.37.141.69:80",
		"http://202.55.5.209:8090",
		"http://152.136.62.181:9999",
		"http://117.34.25.11:55443",
		"http://60.170.204.30:8060",
		"http://118.163.120.181:58837",
	}
	logFile *os.File
	file    *os.File
	dateStr = time.Now().Format("20060102")
)

//func (t *io.multiWriter) Write(p []byte) (n int, err error) {
//	for _, w := range t.writers {
//		n, err = w.Write(p)
//		if err != nil {
//			return
//		}
//		if n != len(p) {
//			err = io.ErrShortWrite
//			return
//		}
//	}
//	return len(p), nil
//}
//
//func MultiWriter(writers ...io.Writer) io.Writer {
//	allWriters := make([]io.Writer, 0, len(writers))
//	for _, w := range writers {
//		if mw, ok := w.(multiWriter); ok {
//			allWriters = append(allWriters, mw.writers...)
//
//		} else {
//			allWriters = append(allWriters, w)
//		}
//	}
//	return &multiWriter{allWriters}
//}

func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.SetPrefix("[-h0xm4n-] ")
}

func handleErr(err error) {
	if err != nil {
		//fmt.Println(when, err)
		log.Panicln(err, "\n-----------------------------------------------------------------")
		//os.Exit(1)
	}

}

func parseArgs() {
	//flag.StringVar(&searchurl, "u", "https://chaziyu.com/", "-u 使用的工具网站")
	flag.StringVar(&searchurl, "u", "https://site.ip138.com/", "-u 使用的工具网站")
	flag.StringVar(&domain, "t", "baidu.com", "-t 要查询的域名")
	flag.StringVar(&suffix, "s", "/domain.htm", "-s url后缀补充")
	flag.StringVar(&save, "o", "nil", `-o 将结果保存到到指定位置,
	如果不指定位置但使用了该参数, 请输入default, 将保存在默认位置: [./20220428-out.txt]
	必须对参数值添加单或双引号！`)
	flag.Parse()
}

func getResponseBytes(url string) (bytes []byte) {
	//resp, err := http.Get(url)
	client := &http.Client{
		Timeout: 5 * time.Second,
		//Transport: trans,
	}
	req, err := http.NewRequest("GET", url, nil)
	handleErr(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	resp, err := client.Do(req)
	handleErr(err)
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	handleErr(err)
	return
}

func save2file(strs string) {
	// 保存到文件
	//myself, err := user.Current()
	//handleErr(err)
	//homedir := myself.HomeDir
	//now := time.Now()
	//fmt.Println(now)

	if save == "default" {
		save = "./" + dateStr + "-out.txt"
	}
	file, err := os.OpenFile(save, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	handleErr(err)
	defer file.Close()

	write := bufio.NewWriter(file)
	//for i := 0; i < len(strs); i++ {
	write.WriteString(strs)
	write.WriteString("\n")
	//}
	write.Flush()
	//fmt.Println("密码文件保存成功！位置:", save)
}

func getMatch(reStr string, ss string) (res []string) {
	re := regexp.MustCompile(reStr)
	allString := re.FindAllStringSubmatch(ss, -1)
	for _, s := range allString {
		// 输出分组，第一个是全部，后面的为正则中的分组
		//fmt.Println(s)
		res = append(res, s...)
	}
	//fmt.Println(res)
	res = deduplicated(res)
	return
}

func deduplicated(input []string) (output []string) {
	deMap := make(map[string]bool)
	for _, de := range input {
		if !deMap[de] {
			deMap[de] = true
		}
	}
	for key := range deMap {
		output = append(output, key)
	}
	return
}

//
//func useProxy() *http.Transport {
//	rand.Seed(time.Now().UnixNano())
//	ri1 := rand.Intn(len(proxys))
//	//fmt.Println(ri)
//	proxyAddr1 := proxys[ri1]
//	proxy1, err := url.Parse(proxyAddr1)
//	handleErr(err, "proxy1, err := url.Parse(proxyAddr1)")
//	fmt.Println("proxy: ", proxy1)
//	netTranspot := &http.Transport{
//		Proxy:                 http.ProxyURL(proxy1),
//		MaxIdleConnsPerHost:   4,
//		ResponseHeaderTimeout: time.Second * time.Duration(5)}
//	return netTranspot
//}

func main() {

	parseArgs()
	log.Println("1. 用户输入参数填充完毕！\n\t\t\t\t\t( -u: ", searchurl, " )\n\t\t\t\t\t( -t: ", domain, " )\n\t\t\t\t\t( -s: ", suffix, " )\n\t\t\t\t\t( -o: ", save, " )")
	url := searchurl + domain
	if suffix != "" {
		url = searchurl + domain + suffix
	}
	log.Println("2. url拼接完成！ url:", url)
	//fmt.Println(url)
	//bytes := getResponseBytes(url, useProxy())
	bytes := getResponseBytes(url)
	resp := string(bytes)
	log.Println("3. 响应获取成功！ 长度:", len(resp))
	//fmt.Println(resp)

	reSubDomain := `([a-zA-Z0-9]+\.*[a-zA-z0-9]+\.` + domain + `)`
	//fmt.Println(reSubDomain)
	res := getMatch(reSubDomain, resp)
	log.Println("4. 内容过滤成功！ 长度:", len(res))
	//var slice []string

	for _, value := range res {
		fmt.Println(value)
	}
	fmt.Println(len(res))
	//useProxy()

	if save != "nil" {
		for _, value := range res {
			//fmt.Println(value)
			save2file(value)
			//save2file("\n")
		}
		log.Println("5. 结果保存在:", save)
		log.Println("6. 结果保存成功！")
		//save2file(dateStr)
	}
	log.Println("完成！\n-----------------------------------------------------------------")
	//log.Println("\n-----------------------------------------------------------------")
	defer logFile.Close()
}
