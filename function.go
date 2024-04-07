package main

import (
	"errors"
  "fmt"
	"os"
  "strings"
)

func funcDefer() {
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("hello world")
}
func trimExtension(files ...string) []string { // 可変長で受け取るテキスト一覧を関数内でテキストのスライスとして扱える
	out := make([]string, 0, len(files)) // 引数で受け取ったファイル名から拡張子を取り除くスライスをmake関数で定義
	for _, f := range files { // インデックスは使用しないので_で受け取る
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}
	return out
}
func fileChecker(name string) (string, error) { // チェックするファイル名をString型 返り値にString型とエラー型
	f, err := os.Open(name) // 標準パッケージのOSOPENを使う
	if err != nil { // エラーの場合
		return "", errors.New("file not found") // エラー出力をして関数を抜ける
	}
	defer f.Close() // 開いておいたリソースを閉じる
	return name, nil // 成功している場合はリターンで受け取ったファイル名とエラーの値は無いのでnilを返す
}

func main() {
	funcDefer()
	files := []string{"file1.csv", "file2.csv", "file3.csv"} // String型のスライスをつくり3つのファイル名で初期化
	fmt.Println(trimExtension(files...)) // filesのStringスライスを可変長引数の形に変換
	name, err := fileChecker("file.txt") // 返り値は2つなのでnameとerrという名前で受け取れるように
	if err != nil { // 呼び出した側でもエラーチェックをする
		fmt.Println(err)
		return // returnでmain関数を抜ける
	}
	fmt.Println(name)
}
