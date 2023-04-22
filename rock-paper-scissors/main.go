package main

import (
	"fmt"
	"math/rand"
)

// 模擬剪刀石頭布多人遊戲，計算平手的機率

type condition = int // 情況分類

const (
	no_tie          condition = iota // 不是平手
	tie_by_three                     // 有三個以上的數字，是平手
	tie_by_the_same                  // 全部只有一種數字，是平手
)

func main() {
	fmt.Printf("count, tie_by_three, tie_by_the_same\n")
	for count := 2; count <= 10; count++ {
		tieMap := map[int]int{}
		measure := 1000000
		for i := 0; i < measure; i++ {
			results := make([]int, 0, count)
			for i := 0; i < count; i++ {
				results = append(results, rng())
			}
			isTie := checkTie(results)
			tieMap[isTie]++
		}
		fmt.Printf("%d,%f,%f\n", count, float64(tieMap[tie_by_three])/float64(measure), float64(tieMap[tie_by_the_same])/float64(measure))
	}
}

func rng() int {
	return rand.Intn(3)
}

func checkTie(results []int) condition {
	m := map[int]bool{}
	for _, v := range results {
		m[v] = true
		if len(m) == 3 {
			return tie_by_three // 有三個以上的數字，是平手
		}
	}
	if len(m) == 1 {
		return tie_by_the_same // 全部只有一種數字，是平手
	}
	return no_tie // 不是平手
}
