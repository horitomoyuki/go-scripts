package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err01 := errors.New("something wrong") // 標準パッケージerrorsのNew関数を呼び出す
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01) // %[1]pでerr01のポインタの値を表示
	fmt.Println(err01)

	// エラーのラップ
	err0 := fmt.Errorf("add info: %w", errors.New("original error")) // Errorfのフォーマットパッケージ%wで付加情報を付与
	fmt.Printf("%[1]p %[1]T %[1]v\n", err0)
	
	// Unwrap
	fmt.Println(errors.Unwrap(err0))

	// fileChecker
	file := "dummy.txt"
	err3 := fileChecker(file) // fileCheckerの引数に渡し、返り値をerr3にする
	if err3 != nil {
		if errors.Is(err3, os.ErrNotExist) { // OSOPENを実行しファイルが存在するかどうかをerrors.Isでチェック
			fmt.Printf("%v file not found\n", file) // ファイルが存在しない場合
		} else {
			fmt.Println("unknown error")
		}
	}
}

func fileChecker(name string) error {
	f, err := os.Open(name) // 引数で渡されたファイル名に対してOSOPENを実行
	if err != nil { // エラーが発生した場合はラップしてリターンで返す
		return fmt.Errorf("in checker: %w", err)
	}
	defer f.Close() // 関数が終了したらファイルをクローズ(メモリを解放させる)
	return nil // 成功した場合はnilを返す
}
