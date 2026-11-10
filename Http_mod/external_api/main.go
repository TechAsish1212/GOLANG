package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CatFactResponse struct{
	Fact string `json:"fact"`
	Length int `json:"length"` 
}

func writeJson(w http.ResponseWriter,status int,data any){
	w.Header().Set("Content-Type","application/json");
	w.WriteHeader(status);
	_=json.NewEncoder(w).Encode(data);
}

func fetchCatFact()(CatFactResponse,error){
	url:="https://catfact.ninja/fact";
	res,err:=http.Get(url);

	if err!=nil{
		return CatFactResponse{} ,err
	}

	defer res.Body.Close();

	if res.StatusCode!=http.StatusOK{
		return CatFactResponse{},fmt.Errorf("External Api failed: %s",res.Status)
	}

	bodyBytes,err:=io.ReadAll(res.Body);
	if err!=nil{
		return CatFactResponse{},err;
	}

	var data CatFactResponse;
	err=json.Unmarshal(bodyBytes,&data); 
	if err !=nil{
		return CatFactResponse{},err;
	}

	return data,nil;
}

func externalFunc(w http.ResponseWriter,r *http.Request){
	if r.Method!=http.MethodGet{
		writeJson(w,http.StatusMethodNotAllowed,map[string]any{
			"ok":false,
			"error":"Only GET Method is allowed",
		})
		return
	}

	data,err:=fetchCatFact()
	if err!=nil{
		writeJson(w,http.StatusBadGateway,map[string]any{
			"ok":false,
			"error":"Failed to fetch data",
		})
		return
	}

	writeJson(w,http.StatusOK,map[string]any{
		"Ok":true,
		"timeStamp":time.Now().UTC(),
		"external":map[string]any{
			"source":"catfact.ninja",
			"fact":data.Fact,
			"length":data.Length,
		},
	})



}


func main(){

	http.HandleFunc("/external",externalFunc);

	err:=http.ListenAndServe(":5005",nil);
	fmt.Println(err);
}