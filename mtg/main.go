package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 模擬魔法風雲會連續回合有地牌可使用的回合

type cardType = int

const (
	cardType_land cardType = iota
	cardType_normal
)

type deck []int

func newDeck(land int) deck {
	d := deck{}
	for i := 0; i < land; i++ {
		d = append(d, cardType_land)
	}
	for i := land; i < 60; i++ {
		d = append(d, cardType_normal)
	}
	return d
}
func (d *deck) len() int {
	return len(*d)
}

func (d *deck) shuffle() {
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}
func (d *deck) draw(n int) []int {
	ans := (*d)[:n]
	*d = (*d)[n:]
	return ans
}

func removeElement(s []int, target int) []int {
	for i, v := range s {
		if v == target {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func isHaveland(hand []int) bool {
	for _, v := range hand {
		if v == cardType_land {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())

	landNum := 24
	measureNum := 1000000

	stat := map[int]int{}

	for i := 0; i < measureNum; i++ {
		deck := newDeck(landNum)
		deck.shuffle()
		hand := deck.draw(7)
		round := 0
		for isHaveland(hand) && len(deck) > 0 {
			hand = removeElement(hand, cardType_land)
			round++
			hand = append(hand, deck.draw(1)...)
		}
		stat[round]++
	}

	for i := 0; i <= landNum; i++ {
		fmt.Printf("%d,%f\n", i, float64(stat[i])/float64(measureNum))
	}
}
