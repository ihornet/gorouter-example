package main

import (
	"fmt"
	"github.com/ihornet/gorouter"
	"net/http"
)

func main() {

	router := gorouter.New()

	router.GET("/teachers/list", func(resp http.ResponseWriter, req *http.Request, params *gorouter.Param) {
		resp.Write([]byte("/teachers/list"))
	})

	router.GET("/teachers/:id/profile", func(resp http.ResponseWriter, req *http.Request, params *gorouter.Param) {
		resp.Write([]byte(fmt.Sprintf("%s = %s", "id", params.GetValue("id"))))
	})

	router.GET("/teachers/:id/profile/:id", func(resp http.ResponseWriter, req *http.Request, params *gorouter.Param) {
		resp.Write([]byte(fmt.Sprintf("id1 = %s; id2 = %s", params.Values[0], params.Values[1])))
	})

	http.ListenAndServe(":3001", router)

}
