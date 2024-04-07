package main

import "fmt"

func main() {
    var num int = 10 // 変数の宣言
    var ptr *int // ポインタの宣言

    ptr = &num // numのアドレスをptrに割り当てる
		fmt.Printf("memory: %p\n", ptr) // メモリアドレスの出力

    // ポインタを介してnumの値にアクセスする
    fmt.Println("numの値:", num)
    fmt.Println("ptrの値:", *ptr) // *ptrはポインタが指す値を取得する

    // ポインタを介してnumの値を変更する
    *ptr = 20
    fmt.Println("numの値:", num) // numの値が変更される
    fmt.Println("ptrの値:", *ptr)
}
