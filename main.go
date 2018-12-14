package main

import (
	"net/http"

	"k8s.io/kubernetes/pkg/kubelet/server"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/" , HomeHandler)
	server := server.NewKubeletAuth()
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world"))
}