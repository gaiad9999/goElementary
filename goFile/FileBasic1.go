package goFile

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// 여기에서는 파일 컨트롤에 대한 내용을 추가할 예정

var path string = "goFile/sample_a.txt"
var dirAll string = "goFile/*"
var dir string = "goFile/*.txt"

// OpenFile : 파일을 여는 동작.
// 다만 해당 함수 안에서만 열고 함수 OpenFile이 끝나면 defer file.Close()에 의해 코드가 종료된다.
func OpenFile() {
	file, _ := os.Open(path)
	defer file.Close()
}

// PrintFile : 파일의 결과를 출력한다.
func PrintFile() {
	file, _ := os.Open(path)
	defer file.Close()
	// 스캐너로 파일을 읽음
	scanner := bufio.NewScanner(file)
	// 스캐너의 결과를 출력함.
	fmt.Println(scanner) // 데이터의 비트를 반환함
	//fmt.Println(scanner.Scan()) // true와 Text결과를 반환하며 idx++ 진행
	//fmt.Println(scanner.Text()) // 현재 저장된 Text 결과를 읽음
	// 위의 동작이 셋트로 옴직인다.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// CallFileList : 해당 디렉토리에 있는 모든 파일을 호출한다.
func CallFileList() {
	// dirAll에 매칭되는 파일들을 읽음
	fileList, _ := filepath.Glob(dirAll)
	// 출력
	fmt.Println(fileList)
	for _, file := range fileList {
		fmt.Println(file)
	}
}

// MatchFileList : 해당 디렉토리에 있는 특정 파일을 호출한다.
// 사실상 위의 함수와 경로지정빼고 완전히 동일하다!
func MatchFileList() {
	// dir에 매칭되는 파일들을 읽음
	fileList, _ := filepath.Glob(dir)
	// 출력
	fmt.Println(fileList)
	for _, file := range fileList {
		fmt.Println(file)
	}
}

func FuncFile1() {
	PrintFile()
	CallFileList()
	MatchFileList()
}

// 아래의 코드를 V3가 멀웨어로 감지해버린다.
// **물론 V3 끄면 정상출력된다.
// 악성코드 정보명 : Trojan/Win32.Generic.C4156681
// https://www.ahnlab.com/kr/site/securityinfo/asec/asecCodeView.do?virusSeq=36432&tabGubun=1

// PrintFileFromList : List에 있는 모든 파일을 로그 출력
func PrintFileFromList(p []string) {
	for _, name := range p {
		file, _ := os.Open(name)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

func FuncFile2() {
	fileList, _ := filepath.Glob(dir)
	PrintFileFromList(fileList)
}
