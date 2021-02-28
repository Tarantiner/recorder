package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var title, link, price, uid string
var f *os.File
var err error
var reg *regexp.Regexp

func writeFile() {
	defer f.Close()
	f, err = os.OpenFile("goods.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	b := make([]byte, 1024)
	var t []byte
	for {
		n, err := f.Read(b)
		if err != nil {
			break
		}
		t = append(t, b[:n]...)
		if n < 1024{
			break
		}
	}
	if strings.Contains(string(t), uid) {
		fmt.Println("已存在")
		return
	}
	_, err = f.WriteString(fmt.Sprintf(`%s|%s|%s|%s`, title, uid, price, link) + "\r\n")
	if err != nil {
		fmt.Println("写入文件失败", err)
		return
	}
}

func isWrongLink() bool{
	reg, _ = regexp.Compile(`com/(.*?).html`)
	ret := reg.FindStringSubmatch(link)
	if len(ret) == 2{
		uid = ret[1]
		return false
	}
	return true
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(strings.Repeat("#", 80))
		fmt.Print("输入商品名>>>")
		title, err = rd.ReadString('\n')
		if err != nil {
			continue
		}
		title = strings.TrimSpace(title)
		if title == "q" {
			fmt.Println("bye!")
			break
		}
		if title == "" {
			fmt.Println("请输入商品名")
			continue
		}

		fmt.Print("输入链接>>>")
		link, err = rd.ReadString('\n')
		if err != nil {
			continue
		}
		link = strings.TrimSpace(link)
		if link == "q" {
			fmt.Println("bye!")
			break
		}
		if isWrongLink() {
			fmt.Println("请输入有效链接")
			continue
		}

		fmt.Print("输入期望价格>>>")
		price, err = rd.ReadString('\n')
		if err != nil {
			continue
		}
		price = strings.TrimSpace(price)
		if price == "q" {
			fmt.Println("bye!")
			break
		}
		_, err = strconv.ParseFloat(strings.TrimSpace(price), 64)
		if err != nil {
			fmt.Println("请输入有效数字价格")
			continue
		}
		writeFile()
	}
}
