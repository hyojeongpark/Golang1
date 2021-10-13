package main

import "fmt"

type Bread struct {
	val      string
	location Location
}
type Location struct {
	name string
	tel  int
}

func (b Bread) FindBread() {
	fmt.Println(b.location)
}

//추가된 부분
func (b Bread) NewStore(name string, tel int) {
	b.location.name = name
	b.location.tel = tel
}
func NewStore(b Bread, name string, tel int) {
	b.location.name = name //이건그냥 name으로 간단히 쓰려고 이렇게 쓰는건가???
	b.location.tel = tel
}

func FindBread(b Bread) {
	fmt.Println(b.location)
}

func main() {
	var b Bread
	b.location.name = "광안리"
	b.location.tel = 511234567

	b.FindBread()
	FindBread(b) //이것도 b.FindBread()와 같은 출력기능

	b.NewStore("해운대", 514557555) //이 주황색 NewStore는 이름은 같지만 값만 복사된 다른 주소!
	b.FindBread()
}
