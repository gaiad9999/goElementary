package goMux

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// 여기에서는 파일 컨트롤에 대한 내용을 추가할 예정
func Cache() {
	fmt.Println("testing mux")
}

func StartHandler() http.Handler {
	mux := mux.NewRouter()

	return mux
}
