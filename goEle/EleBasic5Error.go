package goEle

import (
	"fmt"
)

// 에러컨트롤을 이용하면 에러를 다른 동작으로
// 토스하기 위한 트리거로도 사용할 수 있다.
// 에러컨트롤로 대표적인 방법은 두 가지를 꼽을수 있다.
// 1) 함수 안에 에러 정의
// 2) 에러 분석을 위한 세부속성 정의

// 1) 함수 안에 에러 정의
// DivwError : 함수에 에러 반환 정의
// Golang의 일반적인 함수들은 아래와 같이
// (type, error)를 반환하도록 작성하는 것을 권장한다.
func DivwError(a, b float64) (float64, error) {
	// 오류반환 조건
	if b == 0 {
		return 0, fmt.Errorf("분모는 0이 될 수 없습니다.")
	}
	// 정상반환
	return a / b, nil
}

func SimpleError() {
	_, err := DivwError(2.0, 0.0)
	fmt.Println(err)
}

// 2) 에러 분석을 위한 세부속성 정의
// 아래는 인터페이스에 구조체를 저장하여 에러를 관리분석하는 방법이다.
// Error() 메소드를 이용한 구조체
// PwError : 에러의 세부속성이 담긴 구조체
type PwError struct {
	Len int
	Req int
}

// Error() : 실제로 출력할 오류. 인터페이스 error의 메소드다.
func (err PwError) Error() string {
	return "길이부족"
}

// 아래는 error를 반환하는 함수
func RegAccount(name, pw string) error {
	if len(pw) < 8 {
		return PwError{len(pw), 8}
	}
	return nil
}

// error가 nil이 아닌 경우를 주목.
func TestError() {
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

// 위의 인터페이스 변환에 대해서 살펴본 결과
// 일반적인 인터페이스를 출력하면 그냥 {a,b,c,...} 가 출력된다.
// 근데 인터페이스 error를 출력하면 Error() string이 출력된다.
// 하지만 둘 다, 속성(type.ID, err.ID)을 출력하면 정상적으로 출력된다.
// 무슨 차이가 있길래 이러는거지?
