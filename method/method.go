package main

import "fmt"

type Account struct {
	balance int
}

//범용 인출함수
func withdrawFunc(a *Account, amount int) {
	a.balance -= amount //원래(*a).balance인데 구조체이기 때문에 생략가능
}

//Account 구조체용 메서드
func (a *Account) withdrawMethod(amount int) {
	a.balance -= amount
}

//별칭 리시버
type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}
func main() {
	//a := &Account{100} //a는 Account 변수자체
	var a *Account = &Account{100}
	var b *Account = &Account{200}
	fmt.Println(a, b)

	withdrawFunc(b, 30)
	withdrawFunc(a, 30)
	a.withdrawMethod(30)
	fmt.Printf("%d \n", a.balance)

	var c myInt = 10
	fmt.Println(c.add(30))
	fmt.Println(a, b)

}
