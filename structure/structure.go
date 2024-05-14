package main

import "fmt"

type Task struct {
	Title    string
	Estimate int
}

func main() {
	task1 := Task{
		Title:    "Learn Golang",
		Estimate: 3,
	}
	task1.Title = "Learning Go"
	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)

	// receiver
	task1.extendEstimate()                                    // task1の実体に対してextendEstimateメソッドを実行
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate) // Estimateの値は3(変わらない)

	(&task1).extendEstimatePointer()                          // ポインタに対して実行するので&をつけてアドレスを取得する(&は省略も可)
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate) // Estimateの値は13(初期値3に対して+10)
}

func (task Task) extendEstimate() { // 特定の型から生成された型の実体をメソッドに追加（レシーバー）
	task.Estimate += 10
}
func (taskp *Task) extendEstimatePointer() { // ポインタを受け取る（ポインタレシーバー）
	taskp.Estimate += 10
}
