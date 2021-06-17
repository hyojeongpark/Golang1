package main

import "fmt"

func main() {
	//3번예제 - break 라고만 써있으면 for문 하나만 빠져나옴
	//break L1이 있기 때문에 L1을 빠져나온다
	i := 0
L1:
	for {
		for {
			if i == 0 {
				break L1
			}
		}
	}
	fmt.Println("L1끝")
	//2번예제
	// var f [10]float32 = [10]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for i := 0; i < 20; i++ {
	// 	if i == 10 {
	// 		break
	// 	}
	// 	fmt.Printf("%d는 %2.4f\n", i, f[i]) //20까지하면 컴파일러에서 에러남 but if-break문을 적어주면 괜찮음
	// }

	//1번예제
	// j := [5]int{1, 2, 3, 4, 5}     //var j= [5] int 도 가능
	// name := [3]byte{'a', 'b', 'c'} //작은따옴표 유의
	// fmt.Println(j, name)
	// for i := 0; i < 3; i++ {
	// 	fmt.Printf("%c", name[i]) //i의 name을 넣기위함
	// }
}
