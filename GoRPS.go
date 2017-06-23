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
	cpuHand    int
	playerHand int
	retry      int
	judge      int
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
	fmt.Println("じゃんけんゲームへようこそ")
	fmt.Println("じゃんけんの手は")
	fmt.Print("1:グー\n2:チョキ\n3:パー\nとなります\n")
}

func main() {
	printmessage()
	//擬似乱数
	rand.Seed(time.Now().UnixNano())
	janken := [4]string{"Retry", "グー", "チョキ", "パー"}
	for retry == 0 {
		fmt.Println("最初はグーじゃんけん...")
		//input(int)
		playerHand, err := converting()
		if err != nil {
			log.Println("『ちょおまwwwそういうのやめてーやww!』: ")
		}
		//CPUに擬似乱数を設定
		cpuHand = rand.Intn(3) + 1
		//自分-CPU
		judge = playerHand - cpuHand
		//判定
		fmt.Printf("プレイヤー:%s\nCPU:%s\n", janken[playerHand], janken[cpuHand])
		//あいこになったら
		for judge == 0 {
			fmt.Println("あいこで…")
			playerHand, err := converting()
			if err != nil {
				log.Println("『ちょおまwwwそういうのやめてーやww!』: ")
			}
			//CPUに擬似乱数を設定
			cpuHand = rand.Intn(2) + 1
			//自分-CPU
			judge = playerHand - cpuHand
			//判定
			fmt.Printf("プレイヤー:%s\nCPU:%s\n", janken[playerHand], janken[cpuHand])
		}
		//勝ち
		if judge == -1 || judge == 2 {
			fmt.Println("WIN")
		}
		//負け
		if judge == -2 || judge == 1 {
			fmt.Println("LOSE")
		}
		//コンテニューアクション
		fmt.Println("続けるかい？")
		fmt.Print("0:続ける\n1:やめる\n")
		playerretry, err := converting()
		if err != nil {
			log.Println("『ちょおまwwwそういうのやめてーやww!』: ")
		}
		retry = playerretry
	}
	fmt.Println("遊んでくれてありがとう＼＼\\٩( 'ω' )و //／／")
}
