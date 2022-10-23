package goEle

import (
	"fmt"
)

// go인터페이스는 다음과 같은 상황에서 다양한 기법과 병행하여 사용가능함.
// 인터페이스는 다음과 같은 특징이 있다.
// 1. func(input) 이란 함수가 있을때,
//    인터페이스는 input에 여러 타입의 값을 넣을 수 있도록 만들어준다.
// 2. 인터페이스로 정의된 입력값의 종류는 정의된 메소드를 공유하는 타입들이다.

// 사용법 1) 서로다른 구조체를 입력가능한 함수 제작
// 아래의 경우는 패키지a,b 두 종류가 있음.
// 구조체도 메소드도 완전히 다른 패키지임.

// 패키지 a {KrMember, NameCalling()}
type KrMember struct {
	Id   int
	Name string
	Age  int
}

func (m *KrMember) NameCalling() string {
	return m.Name
}

// 패키지 b {JpMember, Whatyourname()}
type JpMember struct {
	Id   int
	Name [2]string
}

func (m *JpMember) Whatyourname() [2]string {
	return m.Name
}

// 이하는 위의 패키지 a,b를 이용하여 인터페이스를 정의한다.
// 인터페이스로 동작처리
type Member interface {
	Whois() string
}

func (m *KrMember) Whois() string {
	return m.NameCalling() // KrMember 동작 인터페이스
}

func (m *JpMember) Whois() string {
	return m.Whatyourname()[1] // JpMember 동작 인터페이스
}

// 인터페이스의 메소드
func CallName(m Member) {
	s := m.Whois()
	fmt.Println("Name =", s)
}

// 이 방법의 장점은 패키지가 언제 추가되더라도 기존의 패키지를 수정하지 않고,
// 인터페이스만 추가하면 CallName을 그대로 사용할 수 있다는 점이다.
func SimpleInterfaceOne() {
	// 서로다른 구조체를 CallName함수에 입력가능하다.
	member1 := JpMember{1, [2]string{"이구치", "유카"}}
	CallName(&member1)
	member2 := KrMember{1, "지연", 14}
	CallName(&member2)
}

// 사용법 2) 아무값이나 입력가능한 변수와 함수 선언
// 근데 이건 예기치못한 오류의 위험성이 있어서 추천되지는 않는 사용법임.
// 다만 이것도 가능함을 소개한다.

func f(i interface{}) {
	fmt.Print(i, ", ")
	fmt.Printf("타입 : %T\n", i)
}

func SimpleInterfaceTwo() {
	var a interface{} = "a string"
	var b interface{} = 14
	var c int = 31

	f(a)
	f(b)
	f(c)
}
