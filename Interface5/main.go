package main

import (
	"fmt"
	"time"
)

//String이라는 외부공개메서드
type SetMenu interface {
	String() string
}

//오로지 관계만 선언
type Ingredient interface { //동작이름
	GetOneIngredient() SetMenu //behavior: GetOneIngredient라는 외부공개메서드
}

//메뉴셋팅 인터페이스
// type MakeMenu interface {
// 	Setting() string
// }

//빵객체 생성
type Bread struct {
	val string
}

//behavior 행동ex> 바른다, 얹는다
//외부공개기능으로 쓰이기 때문에 중요! (내부는 숨기고)
func (b *Bread) GetIngredient(ingredient Ingredient) { //인터페이스이름
	slice := ingredient.GetOneIngredient()
	b.val += slice.String()
}
func (b *Bread) String() string { //string이라는 Bread메소드 추가
	return "bread" + b.val + " +bread" //Bread의 메소드를 추가해준다
} //1. 아래 코드 중 "return값수정" 과 연관되어있다

type Tomato struct {
}

func (t *Tomato) GetOneIngredient() SetMenu {
	return &SliceOfTomato{}
}

type Lettuce struct {
}

func (t *Lettuce) GetOneIngredient() SetMenu {
	return &SliceOfLettuce{}
}

type Meat struct {
}

func (t *Meat) GetOneIngredient() SetMenu {
	return &SliceOfMeat{}
}

type Cheese struct {
}

func (t *Cheese) GetOneIngredient() SetMenu {
	return &SliceOfCheese{}
}

// return 내용
type SliceOfTomato struct {
}

func (t *SliceOfTomato) String() string {
	return " +tomato" //1. return값수정
}

type SliceOfLettuce struct {
}

func (t *SliceOfLettuce) String() string {
	return " +lettuce" //1. return값수정
}

type SliceOfMeat struct {
}

func (t *SliceOfMeat) String() string {
	return " +meat" //1. return값수정
}

type SliceOfCheese struct {
}

func (t *SliceOfCheese) String() string {
	return " +cheese" //1. return값수정
}
func main() {
	fmt.Printf("주문시간: ")
	p := fmt.Println
	now := time.Now()
	p(now)
	//rand.Seed(p)
	//fmt.Println("주문번호: ", rand.Int())

	fmt.Println("부산 맥도날드덕 Special menu! ")
	bread := &Bread{} //메모리상 주소

	patties := &Tomato{}
	patties1 := &Lettuce{}
	patties2 := &Meat{}
	patties3 := &Cheese{}
	bread.GetIngredient(patties)
	bread.GetIngredient(patties1)
	bread.GetIngredient(patties2)
	bread.GetIngredient(patties3)
	fmt.Println(bread)
	fmt.Println("주문이 완료되었습니다")

}
