package goMux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// 여기에서는 좀 더 실제로 사용하기에 적합한 mux 코드 샘플을 작성한다.
// 수정내역
// 1) json 리스트를 입력받을 수 있는 샘플코드이다.
// 2) render 패키지를 이용하여 반환결과를 간단하게 작성함.
// 3) negroni 패키지를 이용하여 동작 제어관리기능을 추가함.

// 구조체 DB 정의 + 동작 정의 + 핸들러 정의
/* 정의되는 사양은 아래와 같다.
구조체 	 : user
DB       : userList
Util	 : Len, IdxReset
API리스트
메소드    URL        동작
Get     /users      DB 호출
Get     /user/{id}  DB[id] 호출
Post    /users      DB += []User
Put     /user/{id}  DB[id] = User
Delete  /user/{id}  DB[id] 삭제
*/
// 구조체 및 구조체 리스트 정의
type User struct {
	Idx   int    `json:"index"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"Lv"`
}
type Users []User

// DB 정의
var UserList Users

func (u *Users) InitUserDB() {
	*u = append(*u, User{Idx: 0, Id: 1, Name: "수", Level: 1})
	*u = append(*u, User{Idx: 0, Id: 2, Name: "민", Level: 1})
	u.IdxReset() //#추가
}

// Utils
// render 관련 정의
var rd *render.Render

type Success struct {
	Success bool `json:"success"`
	result  int  // 당연하지만 이 값은 공개가 안된다!
}

// Len : DB의 크기
func (u *Users) Len() int { //#추가
	return len(*u)
}

// IdxReset : Idx 할당
func (u Users) IdxReset() { //#추가
	for idx := 0; idx < u.Len(); idx++ {
		u[idx].Idx = idx
	}
}

// 핸들러 정의
func InitUserHandler() http.Handler {
	mux := mux.NewRouter()
	UserList.InitUserDB()
	fmt.Println(UserList)
	mux.Handle("/", http.FileServer(http.Dir("goMux"))) //#웹 샘플코드
	mux.HandleFunc("/users", GetUserList).Methods("GET")
	mux.HandleFunc("/user/{id}", GetUser).Methods("GET")
	mux.HandleFunc("/users", CreateUser).Methods("POST")
	mux.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	mux.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	return mux
}

// 동작 Get : DB 전체출력
func GetUserList(rw http.ResponseWriter, r *http.Request) {
	// 입력

	// 주요 코드

	// 출력
	rd.JSON(rw, http.StatusOK, UserList) //#수정 render
	/*
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(UserList)
	*/
}

// 동작 Get : DB내 특정 데이터 출력
func GetUser(rw http.ResponseWriter, r *http.Request) {
	// 입력
	vars := mux.Vars(r)
	idx, _ := strconv.Atoi(vars["id"])
	// 주요 코드
	user := UserList[idx]
	// 출력
	rd.JSON(rw, http.StatusOK, user) //#수정 render
	/*
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(user)
	*/
}

// 동작 Create : Json을 입력받아 DB에 저장
// 리스트를 받아 추가하도록 수정함
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	// 입력
	var users Users                            //#수정
	_ = json.NewDecoder(r.Body).Decode(&users) //#수정
	// 주요 코드
	UserList = append(UserList, users...) //#추가
	UserList.IdxReset()                   //#추가
	// 출력
	rd.JSON(rw, http.StatusCreated, Success{true, 1}) //#수정 render
	//rw.WriteHeader(http.StatusCreated)
}

// 동작 Update : Json을 입력받아 DB에 저장된 값 수정
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	// 입력
	vars := mux.Vars(r)                //#추가
	idx, _ := strconv.Atoi(vars["id"]) //#추가
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	// 주요 코드
	UserList[idx] = user
	UserList.IdxReset() //#추가
	// 출력
	rd.JSON(rw, http.StatusAccepted, Success{true, 1}) //#수정 render
	//rw.WriteHeader(http.StatusAccepted)
}

// 동작 Delete : DB에 저장된 특정 값 삭제
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	// 입력
	vars := mux.Vars(r)
	idx, _ := strconv.Atoi(vars["id"])
	// 주요 코드
	UserList = append(UserList[:idx], UserList[idx+1:]...) //#수정
	UserList.IdxReset()                                    //#추가
	// 출력
	rd.JSON(rw, http.StatusAccepted, Success{true, 1}) //#수정 render
	//rw.WriteHeader(http.StatusAccepted)
}

// 포트 할당 및 동작 수행
func FuncMux3() {
	rd = render.New()                 //#추가 render  (이거 빼먹었더니 오류 발생함)
	mux := InitUserHandler()          //#추가 negroni
	neg := negroni.Classic()          //#추가 negroni
	neg.UseHandler(mux)               //#추가 negroni
	http.ListenAndServe(":8000", neg) //#수정 negroni

	//http.ListenAndServe(":8000", InitUserHandler())
}

/* Mux 관련
실제로 많이 사용하는 Mux 패키지는 echo 패키지인듯 하다.
이 고릴라 mux 패키지는 굉장히 쉽게 다룰수 있어서 좋다는 장점이 있지만,
좀 더 세세한 백엔드 기능들을 컨트롤 하려면 결국 echo를 다룰수 있어야 할지도?
*/

/* Web 관련
Web의 핵심이 되는 코드는 다음과 같다.
mux.Handle("/", http.FileServer(http.Dir("dir")))
이 코드가 실행되는 경우, /dir/index.html이 실행된다.

a)
mux.Handle("/", http.FileServer(http.Dir("")))
로 테스트 해보니 파일 리스트가 뜬다 ㄷㄷ
근데 b를 한뒤 다시 a를 하면 b결과가 뜨는데 아마 렌더링 덮어쓰기 처리를 해버리는듯 하다.
b)
mux.Handle("/", http.FileServer(http.Dir("goMux")))
로 테스트 해보니 정상적으로 /goMux 안에있는 index.html이 열린다.
"/goMux"를 하면 오류가 발생하니 주의.
*/

/* Body 샘플
[
{
  "Id" : 3,
  "Name" : "철",
  "Level" : 3
},
{
  "Id" : 4,
  "Name" : "미미",
  "Level" : 4
},
{
  "Id" : 5,
  "Name" : "시",
  "Level" : 5
}
]
*/
