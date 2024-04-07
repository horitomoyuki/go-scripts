package main

import "fmt"

func funcDefer() {
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("hello world")
}
func trimExtension(files ...string) []string { // 引数で3ドットString 可変長で受け取るテキスト一覧を関数内でテキストのスライスとして扱える
	out := make([]string, 0, len(files)) // 引数で受け取ったファイル名から拡張子を取り除くスライスをmake関数で定義
	for _, f := range files { // インデックスは使用しないので_で受け取る
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}
	return out
}

func main() {
	funcDefer()
	files := []string("file1.csv", "file2.csv", "file3.csv") // String型のスライスをつくり3つのファイル名で初期化
	fmt.Println(trimExtension(files...)) // filesのStringスライスを可変長引数の形に変換
}