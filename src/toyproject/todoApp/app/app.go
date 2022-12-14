package app

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

var todoMap map[int]*Todo

// /todo.html로 리다이렉트 시키는 함수
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

func addTestTodos() {
	todoMap[1] = &Todo{1, "우유 구입 하기", false, time.Now()}
	todoMap[2] = &Todo{2, "GO 언어 공부하기", true, time.Now()}
	todoMap[3] = &Todo{3, "블록체인 공부하기", false, time.Now()}
	todoMap[4] = &Todo{4, "콩나물 사기", false, time.Now()}
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	// id 생성
	id := len(todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo
	rd.JSON(w, http.StatusOK, todo)

}

type Success struct {
	Success bool `json:"success`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("뭐지..\n")
	// id값을 뽑아내기 위함
	vars := mux.Vars(r)
	// 문자열을 숫자로 바꿔준다.
	id, _ := strconv.Atoi(vars["id"])
	// id 값이 기존 todoMap에 있는지 확인하여 있으면 지워준 후 클라이언트에게 true 값을 보내준다.
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	// id값 추출
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	// todoMap에 있는 경우 변경
	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

}
func MakeNewHandler() http.Handler {
	// todo를 인메모리에 저장하기 위한 변수선언
	todoMap = make(map[int]*Todo)
	// 테스트용 리스트 등록
	addTestTodos()

	rd = render.New()
	r := mux.NewRouter()
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("GET")
	// 인덱스 경로로 들어오면 리다이렉트시킨다.
	r.HandleFunc("/", indexHandler)

	return r
}
