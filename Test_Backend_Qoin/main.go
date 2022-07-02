package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var player int
	var dice int
	var score int
	var simpan string
	var dadu int
	var round int
	var pertambahan = 0
	var arrscore []int
	var arrdice []int
	fmt.Println("Input the number of Player :")
	fmt.Scan(&player)
	fmt.Println("Input the number of dice :")
	fmt.Scan(&dice)
	dadu = dice
	fmt.Println("<===================================>")
	fmt.Println("Game Start")
	for i := 0; i < player; i++ {
		simpan = ""
		score = 0
		if round > 0 {
			dice = arrdice[i]
		} else {
			dice = dadu
		}
		for j := 0; j < dice; j++ {
			var acak = rand.Intn(7-1) + 1
			if acak == 6 {
				score += 1
				dice -= 1
			}
			if acak == 1 {
				pertambahan += 1
			}
			simpan = simpan + " " + fmt.Sprint(acak)
		}

		fmt.Println("player#", i+1, score, simpan)
		arrscore = append(arrscore, score)
		arrdice = append(arrdice, dice)
		fmt.Println(arrscore)
	}
	round++

}


