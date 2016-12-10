package main

import (
	"fmt"

	"github.com/pkg/errors"
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

// Result is result after doing janken
type Result int

const (
	invalid Result = iota
	draw
	win
	lose
)

func doJanken(myHand, comHand Hand) (Result, error) {
	if err := myHand.isRight(); err != nil {
		return invalid, errors.Wrap(err, "myHand")
	}
	if err := comHand.isRight(); err != nil {
		return invalid, errors.Wrap(err, "comHand")
	}

	return myHand.janken(comHand), nil
}

func main() {
	var me, com Hand = rock, paper // 更に拡張するならばコマンドラインから受け取りたいところです。

	result, err := doJanken(me, com)
	if err != nil {
		fmt.Println(errors.Wrap(err, "error occurs").Error())
		return
	}

	fmt.Printf("me: %v, com: %v = I %v\n", me, com, result)
}
