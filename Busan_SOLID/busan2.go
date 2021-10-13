package main

import "fmt"

//메서드 2개 : ReturnMenuName() => string(메뉴이름) 리턴, GetOneDish() => Dish 구조체 주소 리턴.
//각각의 메서드를 위한 인터페이스 : ReturnMenuName() => Dish, GetOneDish() => Put 인터페이스

//ReturnSting(각 메뉴의 string 값을 리턴) 위한 인터페이스
type Dish interface {
	ReturnMenuName() string //외부기능1
}

//GetOneDish 메소드(각 메뉴의 주소값을 리턴)를 위한 인터페이스
type Put interface {
	GetOneDish() Dish //외부기능2
}

//메뉴객체 생성
type Menu struct {
	val string
}

//Cook 메소드
func (d *Menu) Cook(put Put) { //put인터페이스를 받는다
	//각 메뉴의 주소값을 담는다.
	dish := put.GetOneDish()
	//각 메뉴이름(string 값)을 메뉴에 더하여 준다.
	d.val += dish.ReturnMenuName()
}

//부산스시
type Sushi struct {
}

func (c *Sushi) GetOneDish() Dish {
	return &GetSushi{}
}

type GetSushi struct {
}

func (c *GetSushi) ReturnMenuName() string {
	return "★ 부산스시 ★"
}

//대구탕
type CodfishSoup struct {
}

func (c *CodfishSoup) GetOneDish() Dish {
	return &GetCodfish{}
}

type GetCodfish struct {
}

func (c *GetCodfish) ReturnMenuName() string {
	return "\n★ 대구탕 ★"
}

//모듬회 세트
type Sashimi struct {
}

type GetSashimi struct {
}

func (c *Sashimi) GetOneDish() Dish {
	return &GetSashimi{}
}

func (c *GetSashimi) ReturnMenuName() string {
	return "\n★ 모듬회 세트 ★"
}

//main함수
func main() {
	fmt.Println("부산 goBlock 해커톤 특별메뉴")
	menu := &Menu{}
	sushi := &Sushi{}
	codfishSoup := &CodfishSoup{}
	sashimi := &Sashimi{}

	//cook 함수 : put 구조체를 매개변수를 받고, 전달받은 메뉴의 주소값에 담긴 값(메뉴명)을
	//menu(메뉴 string 값)에 더하여 준다.
	menu.Cook(sushi)
	menu.Cook(codfishSoup)
	menu.Cook(sashimi)

	fmt.Println(menu)
	fmt.Println("요리가 완료되었습니다")
}
