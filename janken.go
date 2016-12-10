package main

import (
	"fmt"

	"github.com/pkg/errors"
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
