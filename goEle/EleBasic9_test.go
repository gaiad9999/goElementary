package goEle

import (
	"fmt"
	"math"
	"testing"
)

/*
아래는 EleBasic1에서 다룬 내용을 임시로 test처럼 만든 것이다.
하지만 test 함수는 별도의 파일로 분리해서 관리하는 것이 권장되므로
이와같이 패키지 파일 내부에 작성하는 건 추천되지 않는다.
따라서 이 파일에서는 이와 같은 파일들의 테스트 버전을 전부 만들고자 한다.
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

*/

/*

테스트 작성에는 3가지 권장규칙이 있다.
1) 파일명을 _test.go 로 끝나도록 작성한다.
2) testing 패키지를 import한다.
3) 테스크 코드는 func Test--(t *testing.T) 로 정의한다.

*/

// 테스트 코드의 동작은 go test로 가능하다.

func TestEleBasic1(t *testing.T) {
	fmt.Println("Test EleBasic1...")
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

func TestEleBasic2(t *testing.T) {
	fmt.Println("Test EleBasic2...")
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

func TestEleBasic3(t *testing.T) {
	fmt.Println("Test EleBasic3...")
	user1 := User{1, "하나", 14, "가수"}
	fmt.Println(user1)

	member1 := member{
		User{2, "나루", 12, "사장"},
		1,
		"운영자",
	}
	fmt.Println(member1)
	fmt.Println(member1.User.Name) // >>이름 출력

	LvUpMtd(&member1, 5) // 구조체 함수 실행. 이경우 "포인터"를 입력함
	member1.LvUpMtd(5)   // 메소드 실행  // 오버로딩
	member1.LvChecker()
	fmt.Println(member1.MemberLv, member1.MemberClass)

}

func TestEleBasic4Inter1(t *testing.T) {
	fmt.Println("Test EleBasic4-1...")
	// 서로다른 구조체를 CallName함수에 입력가능하다.
	member1 := JpMember{1, [2]string{"이구치", "유카"}}
	CallName(&member1)
	member2 := KrMember{1, "지연", 14}
	CallName(&member2)
}

func TestEleBasic4Inter2(t *testing.T) {
	fmt.Println("Test EleBasic4-2...")
	var a interface{} = "a string"
	var b interface{} = 14
	var c int = 31

	f(a)
	f(b)
	f(c)
}

func TestEleBasic5Error1(t *testing.T) {
	fmt.Println("Test EleBasic5-1...")
	_, err := DivwError(2.0, 0.0)
	fmt.Println(err)
}

func TestEleBasic5Error2(t *testing.T) {
	fmt.Println("Test EleBasic5-2...")
	err := RegAccount("myID", "mypw")
	if err != nil {
		// 여기 나오는 구문은 (구조체, bool) := 인터페이스.(구조체타입)
		// 인터페이스를 구조체로 변환시 변환 가능한지 여부를 체크하여 bool로 표시
		if errInfo, ok := err.(PwError); ok {
			fmt.Println(err.(PwError), errInfo.Len, ok)
			fmt.Printf("%v Len:%d Req:%d\n",
				errInfo, errInfo.Len, errInfo.Req)
		}
	} else {
		fmt.Println("가입완료")
	}
}

/* 테스트코드 돌려본 결과
- go test는 모든 테스트코드의 실행결과를 출력해준다.
- 테스트 코드 하나하나의 결과는 위에서 제공하는 run test를 이용하여 동작가능함
- go test -run <func> 으로 테스트 코드 하나만 돌릴수도 있음
  go test -run Ele 라고쳐도 Ele로 시작하는 모든 테스트 동작함
  go test -run EleBasic5 라고 치면 딱 EleBasic5-1,-2가 동작한다!
- 테스트 코드의 결과가 필요한 경우엔 어떻게?

*/

/* 벤치마크 코드의 경우
위의 테스트와 동일하게 Benchmark~~(b *testing.B)로 작성한 뒤,
go test -bench .
로 동작하면 된다.

아래의 벤치만 돌리는 방법
go test -run Sqrt -bench .

*/

func sqrt1(n int) int {
	return n * n
}

func sqrt2(n float64) float64 {
	return math.Pow(n, 2)
}

func BenchmarkSqrt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sqrt1(99)
	}
}

func BenchmarkSqrt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sqrt2(99)
	}
}
