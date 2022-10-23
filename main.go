package main

import (
	"fmt"

	"github.com/gaiad9999/goElementary/goEle"
)

func main() {
	fmt.Println("Start testing")
	basicEle()
	fmt.Println("End testing")
}

func basicEle() {
	goEle.Prod()

	fmt.Println("Test 1...")
	//goEle.SimpleFuncOne() // EleBasic1.go 변수선언, if문, for문

	fmt.Println("Test 2...")
	//goEle.SimpleFuncTwo() // EleBasic2Array.go 배열선언, 배열for문, 포인터선언

	fmt.Println("Test 3...")
	//goEle.SimpleFuncThr() // EleBasic3Field.go 구조체선언

	fmt.Println("Test 4...")
	goEle.SimpleInterfaceOne() // EleBasic4Interface.go 인터페이스 다루기
	goEle.SimpleInterfaceTwo()
}
