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
	cpu	int
	ply	int
	try	int
	jdg	int
)

func main() {
	printMessage()
	//擬似乱数
	rand.Seed(time.Now().UnixNano())

	janken := [4]string{"Retry", "グー", "チョキ", "パー"}

	for try == 0 {
		fmt.Println("最初はグーじゃんけん...")
		//input(int)
		ply, err := convertInput()
		if err != nil {
			log.Println("『ちょおまwwwそういうのやめてーやww!』: ")
		}
		//CPUに擬似乱数を設定
		cpu = rand.Intn(3) + 1
		//自分-CPU
		jdg = ply - cpu
		//判定
		fmt.Printf("プレイヤー:%s\nCPU:%s\n", janken[ply], janken[cpu])

		//勝ち
		if jdg == -1 || jdg == 2 {
			fmt.Println("WIN")
		}

		//負け
		if jdg == -2 || jdg == 1 {
			fmt.Println("LOSE")
		}

		//あいこになったら
		for jdg == 0 {
			fmt.Println("あいこで…")
			ply, err := convertInput()
			if err != nil {
				log.Println("『ちょおまwwwそういうのやめてーやww!』: ")
			}
			//CPUに擬似乱数を設定
			cpu = rand.Intn(2) + 1
			//自分-CPU
			jdg = ply - cpu
			//判定
			fmt.Printf("プレイヤー:%s\nCPU:%s\n", janken[ply], janken[cpu])
		}

		//コンテニューアクション
		askContinue()
	}
	fmt.Println("遊んでくれてありがとう＼＼\\٩( 'ω' )و //／／")
}

func GetInput() (st string) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	st = sc.Text()
	st = strings.TrimSpace(st)
	return
}

func convertInput() (int, error) {
	st := GetInput()
	return strconv.Atoi(strings.TrimSpace(st))
}

func printMessage() {
	fmt.Println("じゃんけんゲームへようこそ")
	fmt.Println("じゃんけんの手は")
	fmt.Print("1:グー\n2:チョキ\n3:パー\nとなります\n")
}

func askContinue() {
	fmt.Println("続けるかい？")
	fmt.Print("0:続ける\n1:やめる\n")
	playerretry, err := convertInput()
	if err != nil {
		log.Println("『ちょおまwwwそういうのやめてーやww!』: ")
	}
	try = playerretry
}