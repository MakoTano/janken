package main

import (
	"errors"
	"fmt"
)

// 最終的な出力結果が数字だと寂しいので、独自型を定義して文字変換するようにしています
//go:generate stringer -type=Hand
//go:generate stringer -type=Result

// Hand is a sign on janken
type Hand int

const (
	_ Hand = iota
	rock
	paper
	scissors
)

// Result is result after doing janken
type Result int

const (
	_ Result = iota
	draw
	win
	lose
)

func doJanken(me, com Hand) (Result, error) {
	if me > 3 {
		return 0, fmt.Errorf("me Hand %d is not permmited", me)
	}
	if com > 3 {
		return 0, fmt.Errorf("com Hand %d is not permmited", com)
	}

	switch me {
	case rock:
		switch com {
		case rock:
			return draw, nil
		case paper:
			return lose, nil
		case scissors:
			return win, nil
		}
	case paper:
		switch com {
		case rock:
			return win, nil
		case paper:
			return draw, nil
		case scissors:
			return lose, nil
		}
	case scissors:
		switch com {
		case rock:
			return lose, nil
		case paper:
			return win, nil
		case scissors:
			return draw, nil
		}
	}
	return 0, errors.New("something happen")
}

func main() {
	var me, com Hand = rock, paper // 更に拡張するならばコマンドラインから受け取りたいところです。

	result, err := doJanken(me, com)
	if err != nil {
		fmt.Printf("error occurs %s\n", err.Error())
	}

	fmt.Printf("me: %v, com: %v = I %v\n", me, com, result)
}
