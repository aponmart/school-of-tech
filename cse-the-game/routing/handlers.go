package routing

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/labstack/echo"
	"io/ioutil"
	"io"
	"github.com/grant/CSE-The-Game/cse-the-game/schemas"
	"github.com/grant/CSE-The-Game/cse-the-game/db"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("." + STATIC_DIR + "views/")).ServeHTTP(w, r)
}

func Db(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "db")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(echo.ContentType, echo.ApplicationJSONCharsetUTF8)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(db.Todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo schemas.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set(echo.ContentType, echo.ApplicationJavaScriptCharsetUTF8)
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := db.RepoCreateTodo(todo)
	w.Header().Set(echo.ContentType, echo.ApplicationJSONCharsetUTF8)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func Static(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("." + STATIC_DIR))).ServeHTTP(w, r)
}