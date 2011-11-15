package httpcontext

import "net/http"

type Request struct {
    *http.Request
    attributes map[string]string
}

func NewRequest(request *http.Request) *Request {
    this := &Request{}
    this.Request = request
    this.attributes = make(map[string]string)
    return this
}

func (this *Request) GetAttributes() map[string]string{
        return this.attributes
}
