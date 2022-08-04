package project

import (
	"example/project/project2/storage"
	"fmt"
	"net/http"
)

func Project2() {
	var stor = &storage.LocalFile{}
	router := Router{storage: stor}
	http.HandleFunc("/", router.Index)
	http.HandleFunc("/getlog", router.GetLog)
	http.HandleFunc("/addlog", router.AddLog)
	http.HandleFunc("/updatelog", router.UpdateLog)
	fmt.Println("start server: http://127.0.0.1:9000")
	http.ListenAndServe(":9000", nil)
}
