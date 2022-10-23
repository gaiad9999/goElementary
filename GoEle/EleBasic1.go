package goEle

import (
	"fmt"
)

// SimplePrint : fmt패키지의 출력함수 Println 간단소개
func SimplePrint() {
	// fmt.Println : 값을 출력해주는 기본 함수. Print와 다르게 간격조정있음
	fmt.Println("Printing message...")
}

// SimpleCallVarInt : 변수(int) 선언 간단소개
func SimpleCallVarInt(a int, b int) (int, int) {
	var DefinedVariableInt int = a // 변수선언방식 1번 var <name> <type> = <value>
	SimpleDefinedVariableInt := b  // 변수선언방식 2번 <name> := <value>  //타입 자동지정

	return DefinedVariableInt, SimpleDefinedVariableInt
}
func SimplePrintInt(a, b int) {
	fmt.Println("a =", a, "b =", b)
}

// SimpleCallVarStr : 변수(str) 선언 간단소개
func SimpleCallVarStr(a string, b string) (string, string) {
	var DefinedVariableStr string = a // 변수선언방식 1번 var <name> <type> = <value>
	SimpleDefinedVariableStr := b     // 변수선언방식 2번 <name> := <value>  //타입 자동지정

	return DefinedVariableStr, SimpleDefinedVariableStr
}
func SimplePrintStr(a, b string) {
	fmt.Println(a, b)
}

// SimpleOperator : 연산자 간단소개
// 연산자는 입력 타입 = 출력 타입 조건을 따른다.
func SimpleOperator(a int, b int) (sum int, sub int, prod int, div int) {
	if b == 0 {
		return 0, 0, 0, 0
	}
	sum = a + b
	sub = a - b
	prod = a * b
	div = a / b // float가 아닌 것을 명심

	return
}
func SimplePrintOper(a int, b int, c int, d int) {
	fmt.Println("a + b =", a)
	fmt.Println("a - b =", b)
	fmt.Println("a * b =", c)
	fmt.Println("a / b =", d)

}

// SimpleCond : if문 간단설명
// if문 기본  : if <조건문>
// if문 상세  : if <전처리문>; <조건문>
func SimpleCond(a int) {
	if a > 30 {
		fmt.Println("입력값은 30보다 큰 수입니다")
	} else if a == 30 {
		fmt.Println("입력값은 30입니다")
	} else if a < 30 && a > 20 {
		fmt.Println("입력값은 20과 30 사이의 수입니다")
	} else {
		fmt.Println("입력값은 20보다 작은 수입니다")
	}
}

// SimpleLoop : for문 루프 간단설명
// for문 기본  : for <조건문>
// for문 상세1 : for <전처리문>; <조건문>; <후처리문>
// for문 상세2 : for ; <조건문>; <후처리문>   (;생략하면 안됨)
// for문 상세3 : for <전처리문>; <조건문>;    (;생략하면 안됨)
// for문 연동 : break
func SimpleLoop(a int) {
	for idx := 0; idx <= a; idx++ {
		fmt.Print(idx + 1)
		if idx+1 == a {
			fmt.Println()
			break
		} else {
			fmt.Print(", ")
		}
	}
}

// 위에서 다룬 내용 정리
func SimpleFuncOne() {
	SimplePrint()

	a, b := SimpleCallVarInt(30, 12)
	SimplePrintInt(a, b)
	c, d := SimpleCallVarStr("Hello", "world!")
	SimplePrintStr(c, d)

	sum, sub, prod, div := SimpleOperator(a, b)
	SimplePrintOper(sum, sub, prod, div)

	SimpleCond(sum)

	SimpleLoop(b)
}
