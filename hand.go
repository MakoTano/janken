package main

import "fmt"

// 最終的な出力結果が数字だと寂しいので、独自型を定義して文字変換するようにしています
//go:generate stringer -type=Hand

// Hand is a sign on janken
type Hand int

const (
	_ Hand = iota
	rock
	paper
	scissors
)

func (h Hand) isRight() error {
	switch h {
	case rock, paper, scissors:
		return nil
	}
	return fmt.Errorf("%d is not permmited", h)
}

func (h Hand) janken(com Hand) Result {
	if h == com {
		return draw
	}

	switch {
	case h == rock && com == scissors,
		h == scissors && com == paper,
		h == paper && com == rock:
		return win
	}

	return lose
}
