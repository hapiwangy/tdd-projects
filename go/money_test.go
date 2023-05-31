// 宣告了package main之後代表後面的程式碼都屬於main package
// 套件管理對於go是很重要的
package main

import "testing"

func TestMultiplication(t *testing.T) {
	fiver := Dollar{
		amount: 5,
	}
	tenner := fiver.Times(2)
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got: [%d]", tenner.amount)
	}
}

// 添加的程式碼讓他從red->green
// 補充Dollar以及times的相關資訊
type Dollar struct {
	amount int
}

func (d Dollar) Times(multiplier int) Dollar {
	// 指定數值到特定的名稱
	return Dollar{amount: d.amount * multiplier}
}
