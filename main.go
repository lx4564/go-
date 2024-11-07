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
	fmt.Println("请输入数字")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("错误输入")
		return
	}
	input = strings.TrimSuffix(input, "\n")
	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("错误输入")
		return
	}
	switch {
	case guess < secretnumber:
		fmt.Println("小了")
	case guess > secretnumber:
		fmt.Println("大了")
	case guess == secretnumber:
		fmt.Println("你猜对了")
	}
}
