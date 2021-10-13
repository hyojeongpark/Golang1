package main

import (
	"fmt"
	"time"
)

type SpoonOfTopping interface {
	String() string
}
type Topping interface {
	GetOneSpoon() SpoonOfTopping
}
type Bread struct {
	val string
}

//behavior 행동ex> 바른다, 얹는다
//외부공개기능으로 쓰이기 때문에 중요! (내부는 숨기고)
func (b *Bread) PutTopping(topping Topping) { //인터페이스이름
	spoon := topping.GetOneSpoon()
	b.val += spoon.String()
}
func (b *Bread) String() string {
	return "bread" + b.val //뒤에 요소들 붙이기
}

type Tomato struct {
}

func (t *Tomato) GetOneSpoon() SpoonOfTopping {
	return &SpoonOfTomatoTopping{}
}

type Lettuce struct {
}

func (t *Lettuce) GetOneSpoon() SpoonOfTopping {
	return &SpoonOfLettuceTopping{}
}

type Meat struct {
}

func (t *Meat) GetOneSpoon() SpoonOfTopping {
	return &SpoonOfMeatTopping{}
}

// return 내용
type SpoonOfTomatoTopping struct {
}

func (t *SpoonOfTomatoTopping) String() string {
	return "+tomato"
}

type SpoonOfLettuceTopping struct {
}

func (t *SpoonOfLettuceTopping) String() string {
	return "+lettuce"
}

type SpoonOfMeatTopping struct {
}

func (t *SpoonOfMeatTopping) String() string {
	return "+meat"
}

func main() {
	fmt.Printf("주문시간: ")
	p := fmt.Println
	now := time.Now()
	p(now)
	//rand.Seed(p)
	//fmt.Println("주문번호: ", rand.Int())

	fmt.Println("부산 맥도날드덕 Special menu! ")
	bread := &Bread{}

	topping := &Tomato{}
	topping1 := &Lettuce{}
	topping2 := &Meat{}

	bread.PutTopping(topping)
	bread.PutTopping(topping1)
	bread.PutTopping(topping2)
	fmt.Println(bread)
	fmt.Println("주문이 완료되었습니다")

}
