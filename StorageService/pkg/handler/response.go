package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	rw gin.ResponseWriter
}

func (resp *Response) SetStatusOk() {
	resp.rw.WriteHeader(http.StatusOK)
}

func (resp *Response) SetStatusUnauthorized() {
	resp.rw.WriteHeader(http.StatusUnauthorized)
}

func (resp *Response) SetStatusBadRequest() {
	resp.rw.WriteHeader(http.StatusBadRequest)
}

func (resp *Response) SetData(data []byte) {
	resp.rw.Header().Add("Content-Type", string(data))
}
