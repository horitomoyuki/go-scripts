package main

import (
	"fmt"
	"unsafe"
)

type controller interface {
	speedUp() int
	speedDown() int
}
type vehicle struct {
	speed       int
	enginePower int
}
type bicycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int { // ポインタレシーバーとしてvehicleの実体を受け取れるように
	v.speed += 10 * v.enginePower // メソッドの中でspeedの値を更新
	return v.speed // 現在のspeedをreturnで返す
}
func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}
func (b *bicycle) speedUp() int {
	b.speed += 3 * b.humanPower
	return b.speed
}
func (b *bicycle) speedDown() int {
	b.speed -= 1 * b.humanPower
	return b.speed
}
func speedUpAndDown(c controller) { // 引数としてcontrollerのインターフェースを受け取りspeedUpとspeedDownを実行する
	fmt.Printf("current speed: %v\n", c.speedUp()) // 受け取ったcontrollerのspeedUpとspeedDownのメソッドを実行
	fmt.Printf("current speed: %v\n", c.speedDown())
}
func (v vehicle) String() string { // Go標準パッケージのStringerインターフェースを使用
	return fmt.Sprintf("Vehicle current speed is %v (enginePower %v)", v.speed, v.enginePower)
} // vehicleの型に対してStringのメソッドを追加したのでStringerインターフェースを満たす

func main() { // vehicleとbicycleの構造の実体を定義
	v := &vehicle{0, 5} // ポインタの構造体を取得 speed:0 enginePower:5
	speedUpAndDown(v) // vehicle型はspeedUpDownのメソッドを実装しているので、controllerの型を満たすものとみなされるのでvの値を渡すことができる
	b := &bicycle{0, 5} // ポインタの構造体を取得 speed:0 humanPower:5
	speedUpAndDown(b)
	fmt.Println(v) // 渡されたvの値がStringerインターフェースを実装しているかチェックする

	var i1 interface{} // 空のインターフェースは0個のメソッドを実装するとみなす
	var i2 any // 型が予め決まらないany型
	fmt.Printf("%[1]v %[1]T %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]v %[1]T %v\n", i2, unsafe.Sizeof(i2))
	checkType(i2)
	i2 = 1
	checkType(i2)
	i2 = "hello"
	checkType(i2)

}
func checkType(i any) { // any型の値の実行時に評価するための関数
	switch i.(type) { // 渡された引数iの型を取得し、渡された型の値に応じてPrintの値を変更する
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
