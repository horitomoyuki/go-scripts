package main

import "fmt"

func main() {
	i := 1
	ui := uint16(2)
	fmt.Println(i)
	fmt.Printf("i: %v %T\n", i, i) // デフォルトの型(int型)とデータ型フォーマットを指定
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui) //インデックスを指定

	f := 1.23456 // float64 浮動小数展
	s := "hello" // string
	b := true // bool
	fmt.Printf("f: %[1]v %[1]T\n", f)
	fmt.Printf("s: %[1]v %[1]T\n", s)
	fmt.Printf("b: %[1]v %[1]T\n", b)

	pi, title := 3.14, "Go" // 複数の型を定義
	fmt.Printf("pi: %v title: %v\n", pi, title)

	x := 10 // int
	y := 1.23 // float64
	z := float64(x) + y // xとyで型が違うので型変換が必要
	fmt.Println(z)
}
