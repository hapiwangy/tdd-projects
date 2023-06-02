// 宣告了package main之後代表後面的程式碼都屬於main package
// 套件管理對於go是很重要的
package main

import "testing"

// 目前的test code裡面都有兩個if，可以把重複的部分放到一起，簡化
func TestMultiplication(t *testing.T) {
	fiver := Dollar{
		amount: 5,
	}
	tenner := fiver.Times(2)
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got: [%+v]", tenner.amount)
	}
}
func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	twentyEuros := tenEuros.Times(2)
	assertEqual(t, Money{20, "EUR"}, twentyEuros)
	// if twentyEuros.amount != 20 {
	// 	t.Errorf("Expected 20, got [%+v]", twentyEuros.amount)
	// }
	// if twentyEuros.currency != "EUR" {
	// 	t.Errorf("Expected EUR, got: [%s]", twentyEuros.currency)
	// }
}

func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{amount: 5, currency: "USD"}
	tenner := fiver.Times(2)
	assertEqual(t, Money{10, "USD"}, tenner)
	// if tenner.amount != 10 {
	// 	t.Errorf("Expected 10, got [%+v]", tenner.amount)
	// }
	// if tenner.currency != "USD" {
	// 	t.Errorf("Expected USD, got: [%s]", tenner.currency)
	// }
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

type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	// 指定數值到特定的名稱
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

// 為除法編寫測試
func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}

// 定義萃取函數(相同部分提出來)
func assertEqual(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}
