package goMux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 여기에서는 mux 작성에 대한 내용을 추가할 예정

// 1. 핸들러 정의
// 핸들러는 http.Handler로 출력된다.
func SimpleStartHandler() http.Handler {
	mux := mux.NewRouter()

	return mux
}

// 하지만 실제 핸들러를 정의할때,
// 아래와 같이 Api에 따른 동작도 함께 정의를 한다.

// 2. 구조체 DB 정의 + 동작 정의 + 핸들러 정의
/* 정의하는 사양은 아래와 같다.
구조체 	 : Member
DB		 : MemberList
동작	 : Get
*/
// 구조체 및 구조체 리스트 정의
type Member struct {
	Idx   int
	Id    int
	Name  string
	Level int
}
type Members []Member

// 여기서는 이 리스트가 DB의 역할을 할 것이다.
var MemberList Members

func InitDB() {
	MemberList = append(MemberList, Member{Idx: 0, Id: 1, Name: "수", Level: 1})
	MemberList = append(MemberList, Member{Idx: 1, Id: 2, Name: "민", Level: 1})
}

// 핸들러 정의
func InitHandler1() http.Handler {
	// 핸들러 라우터 정의
	mux := mux.NewRouter()
	// DB내 초기값 정의
	InitDB()
	// API 동작 정의
	// 이 동작은 아래의 동작 정의를 살펴보면 된다.
	mux.HandleFunc("/members", GetMemberList).Methods("GET")

	return mux
}

// 동작 정의 : Get : DB 전체출력
// r은 요청하는 내용에 대한 값이 저장되어있다.(받는거)
// rw는 요청한 내용의 수행 결과에 대한 값들이 저장될 것이다.(보낼거)
func GetMemberList(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(MemberList)
}

// 포트 할당 및 동작 수행
func FuncMux1() {
	http.ListenAndServe(":8000", InitHandler1())
}

/* mux 동작에 대한 참고 사항
- 동작은 main.go에서 다룬다.
- Test는 크롬 앱으로 제공하는 RestApi 테스터를 이용한다.
*/

// 3. Rest API 정의
// 여기서는 대표적인 Rest API인 다음의 4가지를 정의한다.
// Read(Get) 기능은 DB에 저장된 값들을 호출하는 기능이다.
// Create(POST) 기능은 Json을 입력받아 DB에 저장하는 기능이다.
// Update(PUT) 기능은 Json을 입력받아 DB에 저장된 값을 수정하는 기능이다.
// Delete 기능은 DB의 특정 값을 삭제하는 기능이다.

// 동작 Get : DB내 "특정 데이터만" 출력 (위에서 정의한 Get은 전체 출력)
func GetMember(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                // 요청값을 저장함
	idx, _ := strconv.Atoi(vars["id"]) // 요청값 str {id}를 int로 변환
	member := MemberList[idx]
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(member)
}

// 동작 Create : Json을 입력받아 DB에 저장
func CreateMember(rw http.ResponseWriter, r *http.Request) {
	var member Member
	_ = json.NewDecoder(r.Body).Decode(&member)
	fmt.Print(member)
	MemberList = append(MemberList, member)
	rw.WriteHeader(http.StatusCreated)
}

// 동작 Update : Json을 입력받아 DB에 저장된 값 수정
func UpdateMember(rw http.ResponseWriter, r *http.Request) {
	var member Member
	_ = json.NewDecoder(r.Body).Decode(&member)
	idx := member.Idx
	MemberList[idx] = member
	rw.WriteHeader(http.StatusAccepted)
}

// 동작 Delete : DB에 저장된 특정 값 삭제
func DeleteMember(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // 요청값을 저장함
	idx, _ := strconv.Atoi(vars["id"])
	MemberList[idx] = Member{}
	rw.WriteHeader(http.StatusAccepted)
}

// 핸들러 정의
func InitHandler2() http.Handler {
	// 핸들러 라우터 정의
	mux := mux.NewRouter()
	// DB내 초기값 정의
	InitDB()
	// API 동작 정의
	// 이 동작은 아래의 동작 정의를 살펴보면 된다.
	mux.HandleFunc("/members", GetMemberList).Methods("GET")
	mux.HandleFunc("/member/{id}", GetMember).Methods("GET")
	mux.HandleFunc("/members", CreateMember).Methods("POST")
	mux.HandleFunc("/members", UpdateMember).Methods("PUT")
	mux.HandleFunc("/member/{id}", DeleteMember).Methods("DELETE")

	return mux
}

// 포트 할당 및 동작 수행
func FuncMux2() {
	http.ListenAndServe(":8000", InitHandler2())
}

/* Body 샘플
{
  "Idx" : 2,
  "Id" : 3,
  "Name" : "주",
  "Level" : 3
}

*/
