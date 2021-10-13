package main

import "fmt"

//메뉴를 선택
type Choose interface {
	String() string //외부기능1
}

//메뉴 인터페이스
type Menu interface {
	GetOneMenu() Choose //외부기능2
}

//접시객체 생성
type Dish struct {
	val string
}

func (d *Dish) addMenu(menu Menu) {
	choose := menu.GetOneMenu()
	d.val += choose.String()
}

func (d *Dish) String() string {
	return d.val
}

//스시
type Sushi struct {
}

func (m *Sushi) GetOneMenu() Choose {
	return &MenuSushi{}
}

//대구탕
type CodfishSoup struct {
}

func (m *CodfishSoup) GetOneMenu() Choose {
	return &menuCodfishSoup{}
}

//모듬회
type Sashimi struct {
}

func (m *Sashimi) GetOneMenu() Choose {
	return &menuSashimi{}
}

//return
type MenuSushi struct {
}

func (m *MenuSushi) String() string {
	return " 부산 스시"
}

type menuCodfishSoup struct {
}

func (m *menuCodfishSoup) String() string {
	return " +대구탕"
}

type menuSashimi struct {
}

func (m *menuSashimi) String() string {
	return " +모듬회 세트"
}

//main함수
func main() {
	fmt.Println("부산 센탑 goblock 해커톤 특별메뉴")
	dish := &Dish{}
	menu1 := &MenuSushi{}
	menu2 := &menuCodfishSoup{}
	menu3 := &menuSashimi{}
	dish.addMenu(menu1)
	dish.addMenu(menu2)
	dish.addMenu(menu3)
	fmt.Println(dish)
	fmt.Println("주문이 완료되었습니다")
}
