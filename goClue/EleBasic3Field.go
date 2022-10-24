package goEle

import (
	"fmt"
)

// 단순 클래스 정의
type User struct {
	Id   int
	Name string
	Age  int
	Job  string
}

// 복합 클래스 정의. 상속에 대응하는 내용이 될 수도 있다.
type member struct {
	User
	MemberLv    int
	MemberClass string
}

// 구조체 입력 함수 정의. 기존의 함수정의와 다르게 함수명 앞에 구조체 리시버가 붙음.
func LvUpMtd(m *member, a int) {
	m.MemberLv += a
	if m.MemberLv > 5 {
		m.MemberClass = "슈퍼방장"
	}
}

// 메소드 정의. 기존의 함수정의와 다르게 함수명 앞에 구조체 리시버가 붙음.
func (m *member) LvUpMtd(a int) {
	m.MemberLv += a
}
func (m *member) LvChecker() {
	if m.MemberLv > 10 {
		m.MemberClass = "슈퍼방장"
	} else if m.MemberLv <= 10 && m.MemberLv > 5 {
		m.MemberClass = "네임드"
	} else {
		if m.MemberLv < 0 {
			m.MemberLv = 0
		}
		m.MemberClass = "병아리"
	}
}

func SimpleFuncThr() {
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

// 함수의 입력값은 값을 "복사"한 뒤, 로컬 영역에 가져와 처리하는 방식을 따른다.
// 즉, 속도를 위해서는 실제 값보단, 포인터를 보내 값을 처리하도록 만드는게 좋을 수 있다.
// 참고로 *<addr>로 보내지 않아도 된다.
// 다만, 이경우는 구조체가 복사되어 전송되므로 특정 구조체를 직접 수정하는게 아니다.
// 특정 구조체를 수정하도록 만들려면 return을 쓰면 되지만 이건 확실히 속도가 느려진다.
