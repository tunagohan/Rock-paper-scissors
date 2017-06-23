package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	cpuHand    int = 0
	playerHand int = 0
	result     int = 0
	janken         = [4]string{"Retry", "グー", "チョキ", "パー"}
)

func inputs() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput = scanner.Text()
	stringInput = strings.TrimSpace(stringInput)
	return
}

func converting() (int, error) {
	stringInput := inputs()
	return strconv.Atoi(strings.TrimSpace(stringInput))
}

func printmessage() {
	fmt.Println("*-----------------------*")
	fmt.Println("じゃんけんゲームへようこそ")
	fmt.Println("じゃんけんの手")
	fmt.Println("1: グーー")
	fmt.Println("2: チョキ")
	fmt.Println("3: パーー")
	fmt.Println("*-----------------------*")
}

//判定
func rpsjudge() {
	if result == -1 || result == 2 {
		fmt.Println("【勝ち】")
	}
	//負け
	if result == -2 || result == 1 {
		fmt.Println("【負け】")
	}
	return
}

func main() {
	printmessage()
	//擬似乱数
	rand.Seed(time.Now().UnixNano())
	fmt.Println("最初はグーじゃんけん...")
	fmt.Print(">>> ")
	for playerHand == 0 {
		input, err := converting()
		if err != nil {
			log.Println("不正な文字: ")
		}
		playerHand = input
	}
	//CPUに擬似乱数を設定
	cpuHand = rand.Intn(3) + 1
	//自分-CPU
	result = playerHand - cpuHand
	//結果
	fmt.Printf("Player:%s\nCPU:%s\n", janken[playerHand], janken[cpuHand])
	fmt.Println("*-----------------------*")
	for result == 0 {
		playerHand = 0
		fmt.Println("あいこで...")
		fmt.Print(">>> ")
		for playerHand == 0 {
			//input(int)
			input, err := converting()
			if err != nil {
				log.Println("不正な文字: ")

			}
			playerHand = input
		}
		//CPUに擬似乱数を設定
		cpuHand = rand.Intn(3) + 1
		//自分-CPU
		result = playerHand - cpuHand
		fmt.Printf("Player:%s\nCPU:%s\n", janken[playerHand], janken[cpuHand])
	}
	rpsjudge()

	fmt.Println("*----------------------------------------*")
	fmt.Println("遊んでくれてありがとう＼＼\\٩( 'ω' )و //／／")
	fmt.Println("*----------------------------------------*")
}
