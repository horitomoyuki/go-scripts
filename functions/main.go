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
func addExt(f func(file string) string, name string) { // 関数の第一引数が無名関数 第2引数でファイル名を受け取れるようにする
	fmt.Println(f(name)) // 関数の中で受け取った無名関数に拡張子なしのファイル名を渡しで実行
}
func multiply() func(int) int { // 返り値の型を無名関数にする intの引数をint型をreturnで返す無名関数
	return func(n int) int { // 返り値の型に合わせたreturnの無名関数のロジックを追加
		return n * 1000
	}
}
func countUp() func(int) int { // 関数の返り値として無名関数を返す
	count := 0 // closure 変数を定義
	return func(n int) int { // int型でnという引数を受け取り、
		count += n // nの値をcountにプラスする
		return count // returnで現在のcountの値を返す
	}
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

	i := 1 // 変数iを定義
	func(i int) { // 無名関数の場合は関数名を書かず直接引数を指定
		fmt.Println(i) // 受け取った引数をPrintlnで出力
	}(i) // 最後に括弧をつけることで即座に実行される
	f1 := func(i int) int { // 無名関数の最後の値に括弧をつけずf1に無名関数を代入している
		return i + 1 // 引数で受け取ったiに+1をしてint型で返す
	}
	fmt.Println(f1(i)) // f1関数を任意の場所で実行できる

	f2 := func(file string) string { // 引数として渡す無名関数
		return file + ".csv"
	}
	addExt(f2, "file1")

	f3 := multiply() // 関数を1度実行すると返り値として無名関数が返り、その無名関数を代入
	fmt.Println(f3(2)) // f3には無名関数が入っているので引数を渡すとこの関数を実行できる

	f4 := countUp() // countUpの実行結果を代入
	for i := 1; i <= 5; i++ {
		v := f4(2) // f4の無名関数を括弧をつけて実行している
		fmt.Printf("%v\n", v) // 実行するたびに無名関数の返り値をPrintで出力
	}
}
