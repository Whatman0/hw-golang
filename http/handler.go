package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type RandomHandler struct{}

func NewRandomHandler(router *http.ServeMux) {
	handler := &RandomHandler{}
	router.HandleFunc("/random", handler.Random())
}

func (handler *RandomHandler) Random() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		rand.NewSource(time.Now().UnixNano())
		x := rand.Intn(7)
		str := strconv.Itoa(x)
		w.Write([]byte(str))
	}
}
