package main

import "fmt"

//1.
type Bread struct {
	val string
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num) //list를 가진 Bread num개를 만들고 breads라고 정의
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

//2.
type StrawberryJam struct { //객체이름  StrawberryJam
	opened bool //opened라는 이름을가진 bool속성
}

func OpenStrawberryJam(jam *StrawberryJam) { //위 StrawberryJam 객체를 받아 정의한 함수 (참)
	jam.opened = true
}

//3.
type SpoonOfStrawberry struct { //한스푼의 잼 객체
}

func GetOneSpoon(_ *StrawberryJam) *SpoonOfStrawberry { //이 동작은 SpoonOfStrawberry객체활용
	return &SpoonOfStrawberry{}
}

//4.
func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) { //이 동작은 SpoonOfStrawberry객체활용
	bread.val += "+strawberry jam" //bread객체에도 추가해주기
}

//5.
type Sandwich struct {
	val string
}

func MakeSandwich(breads []*Bread) *Sandwich {
	sandwich := &Sandwich{}
	for i := 0; i < len(breads); i++ {
		sandwich.val += breads[i].val + "+"
	}
	return sandwich
}

//main

func main() {
	breads := GetBreads(2) //객체가 참(true)임을 가져옴 “bread ”
	jam := &StrawberryJam{}
	OpenStrawberryJam(jam) //여기서 잼을 가져오는데 &스트로베리잼 포인터로 가 객체의 주소를?

	spoon := GetOneSpoon(jam)       // SpoonOfStrawberry객체를 활용한 spoon을 쓰기 위해
	PutJamOnBread(breads[0], spoon) //위에서 정의한 spoon쓰기 하면 “+strawberry jam”출력

	sandwich := MakeSandwich(breads) //원래 sandwich는 Sandwich{}의 포인터변수를 받음(string) “+”

	fmt.Println(sandwich.val)
}
