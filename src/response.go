package httpcontext

import (
    "net/http"
)

type Responser interface{
    Send(w http.ResponseWriter)
}

type Response struct {
	body string
}

func NewResponse(body string) *Response {
    return &Response{body: body}
}

func (this *Response) Send(w http.ResponseWriter) {
    this.sendHeader(w)
    this.sendContent(w)
}

func (this *Response) sendHeader(w http.ResponseWriter) {
   // w.WriteHeader(1) 
}

func (this *Response) sendContent(w http.ResponseWriter) {
    w.Write([]byte (this.body))
}

type RedirectResponse struct {
	url  string
	code int
}

func NewRedirectResponse(url string, code int) *RedirectResponse {
	return &RedirectResponse{url: url, code: code}
}

func (this *RedirectResponse) Send(w http.ResponseWriter) {
    w.Header().Set("Location", this.url)
    w.WriteHeader(this.code)
}

