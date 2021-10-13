package main

import "fmt"

//behavior => decoupling하는 이유
//object에 상태+기능 => 확장불가

type Dish interface {
	String() string //외부기능1
}
type SashimiSlice interface {
	GetOneDish() Dish //외부기능2
}

//접시객체 생성
type Sashimi struct {
	val string
}

//behavior
func (d *Sashimi) PutSlice(sashimiSlice SashimiSlice) {
	dish := sashimiSlice.GetOneDish()
	d.val += dish.String()
}

func (d *Sashimi) String() string {
	return "Dish =" + d.val + " +(sauce) soy sauce"
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
	return &PlateOfFlatfish{}
}

//return
type PlateOfSalmon struct {
}

func (c *PlateOfSalmon) String() string {
	return " Salmon"
}

type PlateOfEels struct {
}

func (c *PlateOfEels) String() string {
	return " +Eels"
}

type PlateOfFlatfish struct {
}

func (c *PlateOfFlatfish) String() string {
	return " +Flatfish"
}

//main함수
func main() {
	fmt.Println("부산 광안리 모듬회")
	sashimi := &Sashimi{}
	slice := &CatchSalmon{}
	slice2 := &CatchEels{}
	slice3 := &CatchFlatfish{}
	sashimi.PutSlice(slice)
	sashimi.PutSlice(slice2)
	sashimi.PutSlice(slice3)
	fmt.Println(sashimi)
	fmt.Println("주문이 완료되었습니다")
}
