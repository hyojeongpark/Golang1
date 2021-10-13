package main

import "fmt"

//behavior => decoupling하는 이유
//object에 상태+기능 => 확장불가

type Dish interface {
	String() string //외부기능1
}
type Put interface {
	GetOneDish() Dish //외부기능2
}

//접시객체 생성
type Menu struct {
	val string
}

//behavior
func (d *Menu) Cook(put Put) {
	dish := put.GetOneDish()
	d.val += dish.String()
}

func (d *Menu) String() string {
	return d.val
}

//생선들
type CatchSalmon struct {
} //연어

func (c *CatchSalmon) GetOneDish() Dish {
	return &PlateOfSalmon{}
}

type CatchEels struct {
} //장어

func (c *CatchEels) GetOneDish() Dish {
	return &PlateOfEels{}
}

type CatchFlatfish struct {
} //광어

func (c *CatchFlatfish) GetOneDish() Dish {
	return &FlatfishSoup{}
}

//return
type PlateOfSalmon struct {
}

func (c *PlateOfSalmon) String() string {
	return "★ 부산스시 ★"
}

type PlateOfEels struct {
}

func (c *PlateOfEels) String() string {
	return "\n★ 대구탕 ★"
}

type FlatfishSoup struct {
}

func (c *FlatfishSoup) String() string {
	return "\n★ 모듬회 세트 ★"
}

//main함수
func main() {
	fmt.Println("부산 goBlock 해커톤 특별메뉴")
	menu := &Menu{}
	salmon := &CatchSalmon{}
	slice2 := &CatchEels{}
	slice3 := &CatchFlatfish{}
	menu.Cook(salmon)
	menu.Cook(slice2)
	menu.Cook(slice3)
	fmt.Println(menu)
	fmt.Println("주문이 완료되었습니다")
}
