package main

// 最終的な出力結果が数字だと寂しいので、独自型を定義して文字変換するようにしています
//go:generate stringer -type=Result

// Result is result after doing janken
type Result int

const (
	invalid Result = iota
	draw
	win
	lose
)
