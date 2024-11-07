package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxnum := 100
	rand.Seed(time.Now().UnixNano())
	secretnumber := rand.Intn(maxnum)
	fmt.Println("请输入猜的数字")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("输入错误")
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		guess, errL := strconv.Atoi(input)
		if errL != nil {
			fmt.Println("输入错误")
			continue
		}
		if guess == secretnumber {
			fmt.Println("猜对了")
			break
		} else if guess < secretnumber {
			fmt.Println("猜小了")
		} else {
			fmt.Println("猜大了")
		}
	}
}

///1.首先应该用rand.Intn生成一个随机数
///2.用bufio和os.Stdin输入数字
///3.用ReadString只读第一行
///4.再用TrimSuffix去除末尾的\n
///5.用strconv.Atoi转换字符串为数字
