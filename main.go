package main

import (
	"fmt"
	"math/rand"
	"time"
)

//0 => 2 diamond
//1 => 2 club
//2 => 2 heart
//3 => 2 spade
//4 => 3
//5 => 3
//6 => 3
//7 => 3
//8 => 4 diamond
//...
//12 => 5 diamond
//...
//16 => 6 diamond
//...
//20 => 7 diamond
//...
//24 => 8 diamond
//...
//28 => 9 diamond
//...
//32 => 10 diamond
//...
//36 => J diamond
//...
//40 => Q diamond
//...
//44 => K diamond
//...
//48 => A diamond
//49 => A club
//50 => A heart
//51 => A spade

type Poker struct {
	cards [52]int
	pos   int
}

func (p *Poker) init() {
	for i := 0; i < 52; i++ {
		p.cards[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(52, func(i, j int) {
		p.cards[i], p.cards[j] = p.cards[j], p.cards[i]
	})
}

func (p *Poker) get() int {
	if p.pos < 52 {
		defer func() {
			p.pos++
		}()
		return p.cards[p.pos]
	}
	return -1
}

func match(a, b int) (int, int) {
	var p = new(Poker)
	p.init()
	return player(p, a), player(p, b)
}

func card2score(i int) int {
	switch {
	case i < 36:
		return i/4 + 2
	case i < 48:
		return 10
	case i < 52:
		return 11
	default:
		panic("Poker Error.")
	}
}

func player(p *Poker, max int) int {
	score := 0
	for score < max {
		if card := p.get(); card >= 0 {
			score += card2score(card)
			if score > 21 {
				score = 0
				break
			}
		} else {
			break
		}
	}
	return score
}

//17 is the best score
func main() {
	var a, b int
	for i := 0; i < 1000000; i++ {
		score1, score2 := match(17, 18)
		if score1 > score2 {
			a += 1
		} else if score1 < score2 {
			b += 1
		}
	}
	fmt.Println(a, b)
}
