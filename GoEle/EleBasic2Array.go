package goEle

import "fmt"

// SimpleCallVarIntList : 배열 선언 간단소개
func SimpleCallVarIntList(a [2]int) (int, int) {
	var b [2]int = a
	// b := [2]int{12, 14}  //이와 같이 리스트는 [n]type{value}의 형태로 정의된다.
	var DefinedVariableInt1 int = b[0] // 변수선언방식 1번 var <name> <type> = <value>
	DefinedVariableInt2 := b[1]

	return DefinedVariableInt1, DefinedVariableInt2
}

// SimpleCallVarIntLongList : 배열 선언 또 다른버전
func SimpleCallVarIntLongList(a []int) []int {
	b := []int{} // 이렇게 배열 크기 선언 안한걸 슬라이스(slice)라 부른다.

	for idx, x := range a { // range를 이용한 for문 소개. 안 쓸 값은 _로 두면 됨
		b = append(b, x+1) // 슬라이스는 append를 이용해 값을 첨가할 수 있다.
		if idx > 3 {
			break
		}
	}

	return b
}

// SimplePointer : 포인터 선언 간단정리
func SimplePointer() {
	var a int = 51
	var p *int = &a // *<type> : 포인터 선언방식. &<var> : 변수의 주소값.

	fmt.Println("Pointer testing...", *p) // *<addr> : 주소에 저장된 값 출력
}

func SimpleFuncTwo() {
	listA := [2]int{12, 14}
	a, b := SimpleCallVarIntList(listA)
	fmt.Println("a =", a, "b =", b)

	listB := []int{10, 12, 14, 16, 18, 20}
	listC := SimpleCallVarIntLongList(listB)
	fmt.Println("c = [", listC[0], listC[1], listC[2], "]")

	listD := []int{
		10,
		12,
		14,
	} // 주의. 리스트 입력방법은 listB와 listD 처럼 두 가지 방식이 존재한다.
	fmt.Println("d = [", listD[0], listD[1], listD[2], "]")

	SimplePointer()
}
