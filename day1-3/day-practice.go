package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)

	fmt.Println("请输入你的数字：")
	var input string
	//reader := bufio.NewReader(os.Stdin)
	for {
		//input, err := reader.ReadString('\n')
		//if err != nil {
		//	fmt.Println("输入发生错误", err)
		//	return
		//}
		//input = strings.TrimSuffix(input, "\r\n")
		//fmt.Printf("%#v", input)

		//fmt.Scanf("%s", &input)
		fmt.Scanf("%s\n", &input)
		fmt.Printf("%#v", input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("转换出错")
		}
		fmt.Println("your guess: ", guess)
		if guess > secretNumber {
			fmt.Println("大了")
		} else if guess < secretNumber {
			fmt.Println("小了")
		} else {
			fmt.Println("对了")
			break
		}
	}
}
