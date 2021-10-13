package main

import (
	"fmt"
	"time"
)

type Bread struct {
	val string
}
type SetMenu interface {
	String() string
}
type Patty interface {
	GetOnePatty() SetMenu
}

//
func (b *Bread) String() string {
	return "bread" + b.val //뒤에 요소들 붙이기
}

//behavior 행동ex> 바른다, 얹는다
//외부공개기능으로 쓰이기 때문에 중요! (내부는 숨기고)
func (b *Bread) PutPatty(patties Patty) { //인터페이스이름
	slice := patties.GetOnePatty()
	b.val += slice.String()
}

//
type Tomato struct {
}

func (t *Tomato) GetOnePatty() SetMenu {
	return &SpoonOfTomatoTopping{}
}

type Lettuce struct {
}

func (t *Lettuce) GetOnePatty() SetMenu {
	return &SpoonOfLettuceTopping{}
}

type Meat struct {
}

func (t *Meat) GetOnePatty() SetMenu {
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

	patties := &Tomato{}
	patties1 := &Lettuce{}
	patties2 := &Meat{}

	bread.PutPatty(patties)
	bread.PutPatty(patties1)
	bread.PutPatty(patties2)
	fmt.Println(bread)
	fmt.Println("주문이 완료되었습니다")

}
