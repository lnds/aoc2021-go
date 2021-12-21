package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lnds/aoc2021-go/shared"
)

func main() {
	lines := shared.ReadLines(os.Args[1])
	pos1, _ := strconv.Atoi(strings.TrimPrefix(lines[0], "Player 1 starting position: "))
	pos2, _ := strconv.Atoi(strings.TrimPrefix(lines[1], "Player 2 starting position: "))

	w, l, s, r := game(pos1, pos2, 1000, advance3)
	fmt.Printf("winner = %d, looser = %d, rolls = %d\n", w, l, r)
	fmt.Printf("result = %d\n", r*s)

	wins := countWins(0, int(pos1), 0, int(pos2), 0)
	fmt.Printf("wins = %v\n", wins)
	if wins[0] > wins[1] {
		fmt.Printf("win 1 in  %d universes\n", wins[0])
	} else {
		fmt.Printf("win 2 in  %d universes\n", wins[1])

	}
}

func play(pos, score, roll int) (int, int) {
	newPos := ((pos - 1 + roll) % 10) + 1
	newScore := score + newPos
	return newPos, newScore
}

type args struct{ player, pos0, score0, pos1, score1 int }

var memo map[args][2]int = make(map[args][2]int)

func countWins(player, pos0, score0, pos1, score1 int) [2]int {
	if score0 >= 21 {
		return [2]int{1, 0}
	} else if score1 >= 21 {
		return [2]int{0, 1}
	}

	wins, ok := memo[args{player, pos0, score0, pos1, score1}]
	if ok {
		return wins
	}

	wins = [2]int{0, 0}
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				if player == 0 {
					newPos, newScore := play(pos0, score0, int(a+b+c))
					w := countWins(1, newPos, newScore, pos1, score1)
					wins[0] += w[0]
					wins[1] += w[1]
				} else {
					newPos, newScore := play(pos1, score1, int(a+b+c))
					w := countWins(0, pos0, score0, newPos, newScore)
					wins[0] += w[0]
					wins[1] += w[1]
				}
			}
		}
	}
	memo[args{player, pos0, score0, pos1, score1}] = wins
	return wins
}

func game(pos1, pos2, limit int, advance3 func(int, int, int, int) (int, int, int)) (winner, looser, lowScore, rolls int) {
	score1 := 0
	score2 := 0
	dice := 0
	for {
		pos1, dice, score1 = advance3(pos1, dice, score1, 1)
		rolls += 3
		if score1 >= limit {
			break
		}
		pos2, dice, score2 = advance3(pos2, dice, score2, 2)
		rolls += 3
		if score2 >= limit {
			break
		}
	}
	if score1 > score2 {
		winner = 1
		looser = 2
		lowScore = score2
	} else {
		winner = 2
		looser = 1
		lowScore = score1
	}
	return
}

func roll(dice int) int {
	if dice == 100 {
		return 1
	}
	return dice + 1
}

func advance3(pos, dice, score, _player int) (int, int, int) {
	dice = roll(dice)
	sd := dice
	dice = roll(dice)
	sd += dice
	dice = roll(dice)
	sd += dice

	pos = ((pos - 1 + sd) % 10) + 1
	score += pos
	return pos, dice, score
}
