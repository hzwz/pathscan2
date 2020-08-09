package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
	"github.com/kirinlabs/HttpRequest"
	"github.com/schollz/progressbar/v3"
	"github.com/EDDYCJY/fake-useragent"
)

var (
	url string
	t   int
)

func produce(ch chan string) {
	fileName := "dicc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)

	for {
		line, err := rd.ReadString('\n')
		line = strings.TrimSpace(line)

		if err != nil || io.EOF == err {
			continue
		}
		ch <- line

	}

}

func use(ch chan string, url string, bar *progressbar.ProgressBar) {
	var path string
	req := HttpRequest.NewRequest()
	var random string
	random = browser.Random()

	req.SetTimeout(8)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	for {
		random = browser.Random()
		
	    req.SetHeaders(map[string]string{
	    	"User-Agent": random,})
		path = <-ch
		res, err := req.Get(url + path)
		if err != nil {
			continue
		}
		body, err := res.Body()
		CheckStatusCode(url+path, res.StatusCode(), string(body))
		bar.Add(1)
		fmt.Print("\r")

	}

}

func CheckStatusCode(url string, status int, body string) {
	if !strings.Contains(body, "无法访问系统资源") {

		switch status {
		case 200:
			color.Green.Print(url + "        " + strconv.Itoa(status) + "                                                            \n")
		case 300:
			color.Blue.Print(url + "        " + strconv.Itoa(status) + "                                                            \n")
		case 301:
			color.Blue.Print(url + "        " + strconv.Itoa(status) + "                                                            \n")
		case 302:
			color.Blue.Print(url + "        " + strconv.Itoa(status) + "                                                            \n")
		case 400:
			color.Yellow.Print(url + "        " + strconv.Itoa(status) + "                                                            \n")
		case 401:
			color.Yellow.Print(url + "        " + strconv.Itoa(status) + "                                                            \n")
		case 402:
			color.Yellow.Print(url + "        " + strconv.Itoa(status) + "                                                            \n")
		case 403:
			color.Yellow.Print(url + "        " + strconv.Itoa(status) + "                                                            \n")
		case 500:
			color.Red.Print(url + "        " + strconv.Itoa(status) + "                                                            \n")

		}

	}

}

func main() {

	flag.StringVar(&url, "u", "", "website")
	flag.IntVar(&t, "t", 30, "the num of threads")

	flag.Parse()
	charray := []byte(url)
	if string(charray[len(charray)-1]) != "/" {
		url = url + "/"
	}
	var wg sync.WaitGroup

	// ch := make(chan int, 10)
	ch := make(chan string)
	go produce(ch)
	time.Sleep(3 * time.Second)
	//fmt.Println(num)
	bar := progressbar.Default(231435)
	wg.Add(t)

	for i := 1; i <= t; i++ {
		go func() {
			use(ch, url, bar)

		}()
	}

	wg.Wait()
}
