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
		fmt.Printf("プレイヤー:%s\nCPU:%s\n", janken[playerHand], janken[cpuHand])
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
