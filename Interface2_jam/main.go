package main

import "fmt"

type SpoonOfJam interface {
	String() string
}
type Jam interface {
	GetOneSpoon() SpoonOfJam
}
type Bread struct {
	val string
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon() //Jam의 인터페이스를 받았다-> "Jam Interface를 구현한것이다"
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread" + b.val
}

type StrawberryJam struct {
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{} //38~41번으로
}

type OrangeJam struct {
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

type KayaJam struct {
}

func (j *KayaJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfKayaJam{}

}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ orange"
}

type SpoonOfKayaJam struct {
}

func (s *SpoonOfKayaJam) String() string {
	return "+kaya"
}

func main() {
	fmt.Println("잼토스트 만들기")
	//1. 빵 두개를 꺼낸다.
	bread := &Bread{}

	//jam := &StrawberryJam{} //스트로베리잼이 오도록 할 때
	//jam := &OrangeJam{}
	jam := &KayaJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
