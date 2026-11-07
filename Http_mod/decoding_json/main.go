package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func writeJSON(w http.ResponseWriter, status int,data any){
	w.Header().Set("Content-Type","application/json");
	w.WriteHeader(status)
	_=json.NewEncoder(w).Encode(data)
}

type testData struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func decodeHandler(w http.ResponseWriter, r *http.Request){
	if r.Method!=http.MethodPost{
		writeJSON(
			w,
			http.StatusMethodNotAllowed,
			map[string]any{
				"success":false,
				"error":"Only POST is allowed",
			},
		)
		return
	}

	defer r.Body.Close();

	var data testData;
	err:=json.NewDecoder(r.Body).Decode(&data);
	if err!=nil{
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]any{
				"success":false,
				"error":"Invalid JSON data",
			},
		)
		return
	}

	data.Name=strings.TrimSpace(data.Name);

	if data.Name==""{
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]any{
				"success":false,
				"error":"Name is required",
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		map[string]any{
			"success":true,
			"data":data,
			"timestamp":time.Now().UTC(),
		},
	)
}


func main(){
	http.HandleFunc("/decode",decodeHandler);
	
	err:=http.ListenAndServe(":9008",nil);
	fmt.Println(err);
}