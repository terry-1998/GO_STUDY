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
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secrectNum := rand.Intn(maxNum)
	//println("the secret number is", secrectNum)
	fmt.Println("please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("an error occured", err)
			return
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("an error occured", err)
			return
			continue
		}
		fmt.Println("your guess is", guess)

		if guess > secrectNum {
			fmt.Println("你太大了")
		} else if guess < secrectNum {
			fmt.Println("你太小了")
		} else {
			fmt.Println("你太聪明了")
			break
		}
	}

}
