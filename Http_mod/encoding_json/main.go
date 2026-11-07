package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter, r *http.Request) {
	// 1. one way
	/*
		w.Header().Set("Content-Type","application/json");
		w.WriteHeader(http.StatusOK);
		res:=map[string]any{
			"success":true,
			"msg":"JSON encode successfull",
			"Time":time.Now().UTC(),
		}

		_=json.NewEncoder(w).Encode(res) */

	// 2. another way
	res := map[string]any{
		"success": true,
		"msg":     "JSON encode successfull",
		"time":    time.Now().UTC(),
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "JSON encoding failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK);

	_,_=w.Write(jsonData)

}

func main() {
	http.HandleFunc("/ok", successHandler)

	err := http.ListenAndServe(":9008", nil)
	fmt.Println(err)
}
