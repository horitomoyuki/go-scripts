package main

import "fmt"

func main() {
  var a1 [3]int // 明示的な代入を省略した場合は0値で初期化 0が3つの要素で初期化
  var a2 = [3]int{10, 20, 30} // 明示的な代入
  a3 := [...]int{10, 20} // 要素数のところでスリードットをつけると自動的にカウントしてくれる
  fmt.Printf("%v %v %v\n", a1, a2, a3)
	fmt.Printf("%v %v\n", len(a3), cap(a3)) // 配列の要素数とキャパシティーを出力
	fmt.Printf("%T %T\n", a2, a3) // 型を出力する

  var s1 []int // 要素数を空にするスライス
  s2 := []int{} // 明示的な代入が必要なのでカーリーブラケットをつける
  fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
  fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
  fmt.Println(s1 == nil) // varで初期化した場合はnilとして認識される
  fmt.Println(s2 == nil) // := で初期化された場合はnilとして認識されない

	s1 = append(s1, 1, 2, 3) // append で要素の追加をするスライス
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1)) // 要素が追加され、要素数とキャパシティが3になる
	s3 := []int{4, 5, 6} // 別に定義したスライスを追加する
	s1 = append(s1, s3...) // 可変長に変換して追加することができる
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1)) // 要素が追加され、要素数とキャパシティが6になる
}
