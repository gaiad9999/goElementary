package goEle

import (
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

func TestEleBasic1(t *testing.T) {
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
