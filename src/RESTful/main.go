package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func main() {
	router := httprouter.New()
	router.DELETE("/", deleteUser)
}
